package ordering

import (
	"context"

	"eda-in-go/internal/ddd"
	"eda-in-go/internal/monolith"
	"eda-in-go/ordering/internal/application"
	"eda-in-go/ordering/internal/grpc"
	"eda-in-go/ordering/internal/handlers"
	"eda-in-go/ordering/internal/logging"
	"eda-in-go/ordering/internal/postgres"
	"eda-in-go/ordering/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	orders := postgres.NewOrderRepository("ordering.orders", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, invoices, shopping, domainDispatcher)
	app = logging.NewApplication(app, mono.Logger())
	// set application handlers
	notificationHandlers := logging.LogDomainEventHandlerAccess(
		application.NewNotificationHandlers(notifications),
		mono.Logger(),
	)
	invoicesHandlers := logging.LogDomainEventHandlerAccess(
		application.NewInvoiceHandlers(invoices),
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterNotifcationHandlers(notificationHandlers, domainDispatcher)
	handlers.RegisterInvoiceHandlers(invoicesHandlers, domainDispatcher)

	return nil
}
