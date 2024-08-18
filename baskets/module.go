package baskets

import (
	"context"

	"eda-in-go/baskets/internal/application"
	"eda-in-go/baskets/internal/grpc"
	"eda-in-go/baskets/internal/handlers"
	"eda-in-go/baskets/internal/logging"
	"eda-in-go/baskets/internal/postgres"
	"eda-in-go/baskets/internal/rest"
	"eda-in-go/internal/ddd"
	"eda-in-go/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	baskets := postgres.NewBasketRepository("baskets.baskets", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(baskets, stores, products, orders, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogDomainEventHandlerAccess(
		application.NewOrderHandlers(orders),
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
	// order handlers is OnBasketCheckOut handler, which is triggered as soon as
	// the basket status changes to checkout. This handler extracts the contents
	// of the basket to create a new order.
	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return
}
