package application

import (
	"context"

	"eda-in-go/internal/ddd"
	"eda-in-go/ordering/internal/application/commands"
	"eda-in-go/ordering/internal/application/queries"
	"eda-in-go/ordering/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrder) error
		CancelOrder(ctx context.Context, cmd commands.CancelOrder) error
		ReadyOrder(ctx context.Context, cmd commands.ReadyOrder) error
		CompleteOrder(ctx context.Context, cmd commands.CompleteOrder) error
	}
	Queries interface {
		GetOrder(ctx context.Context, query queries.GetOrder) (*domain.Order, error)
		GetOrders(ctx context.Context, query queries.GetOrders) ([]*domain.Order, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateOrderHandler
		commands.CancelOrderHandler
		commands.ReadyOrderHandler
		commands.CompleteOrderHandler
	}
	appQueries struct {
		queries.GetOrderHandler
		queries.GetOrdersHandler
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository, customers domain.CustomerRepository, payments domain.PaymentRepository,
	invoices domain.InvoiceRepository, shopping domain.ShoppingRepository,
	domainPublisher ddd.EventPublisher,
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler:   commands.NewCreateOrderHandler(orders, customers, payments, shopping, domainPublisher),
			CancelOrderHandler:   commands.NewCancelOrderHandler(orders, shopping, domainPublisher),
			ReadyOrderHandler:    commands.NewReadyOrderHandler(orders, invoices, domainPublisher),
			CompleteOrderHandler: commands.NewCompleteOrderHandler(orders, domainPublisher),
		},
		appQueries: appQueries{
			GetOrderHandler:  queries.NewGetOrderHandler(orders),
			GetOrdersHandler: queries.NewGetOrdersHandler(orders),
		},
	}
}
