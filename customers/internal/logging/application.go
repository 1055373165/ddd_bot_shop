package logging

import (
	"context"
	"eda-in-go/customers/internal/application"
	"eda-in-go/customers/internal/domain"

	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) RegisterCustomer(ctx context.Context, register application.RegisterCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.RegisterCustomer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Customers.RegisterCustomer") }()
	return a.App.RegisterCustomer(ctx, register)
}

func (a Application) AuthorizeCustomer(ctx context.Context, authorize application.AuthorizeCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.AuthorizeCustomer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Customers.AuthorizeCustomer") }()
	return a.App.AuthorizeCustomer(ctx, authorize)
}

func (a Application) GetCustomer(ctx context.Context, get application.GetCustomer) (customer *domain.Customer, err error) {
	a.logger.Info().Msg("--> Customers.GetCustomer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Customers.GetCustomer") }()
	return a.App.GetCustomer(ctx, get)
}

func (a Application) EnableCustomer(ctx context.Context, enable application.EnableCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.EnableCustomer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Customers.EnableCustomer") }()
	return a.App.EnableCustomer(ctx, enable)
}

func (a Application) DisableCustomer(ctx context.Context, disable application.DisableCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.DisableCustomer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Customers.DisableCustomer") }()
	return a.App.DisableCustomer(ctx, disable)
}
