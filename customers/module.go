package customers

import (
	"context"
	"eda-in-go/customers/internal/application"
	"eda-in-go/customers/internal/grpc"
	"eda-in-go/customers/internal/logging"
	"eda-in-go/customers/internal/postgres"
	"eda-in-go/customers/internal/rest"
	"eda-in-go/internal/monolith"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	var app application.App
	app = application.New(customers)
	// 多态
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
