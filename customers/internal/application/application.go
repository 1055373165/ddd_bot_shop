package application

import (
	"context"

	"eda-in-go/customers/internal/domain"
	"eda-in-go/internal/ddd"

	"github.com/stackus/errors"
)

type (
	RegisterCustomer struct {
		ID        string
		Name      string
		SmsNumber string
	}

	AuthorizeCustomer struct {
		ID string
	}

	GetCustomer struct {
		ID string
	}

	GetCustomers struct {
	}

	EnableCustomer struct {
		ID string
	}

	DisableCustomer struct {
		ID string
	}

	App interface {
		RegisterCustomer(ctx context.Context, register RegisterCustomer) error
		AuthorizeCustomer(ctx context.Context, authorize AuthorizeCustomer) error
		GetCustomer(ctx context.Context, get GetCustomer) (*domain.Customer, error)
		GetCustomers(ctx context.Context, get GetCustomers) ([]*domain.Customer, error)
		EnableCustomer(ctx context.Context, enable EnableCustomer) error
		DisableCustomer(ctx context.Context, discable DisableCustomer) error
	}

	Application struct {
		customers       domain.CustomerRepository
		domainPublisher ddd.EventPublisher
	}
)

var _ App = (*Application)(nil)

func New(customers domain.CustomerRepository, domainPublisher ddd.EventPublisher) *Application {
	return &Application{
		customers:       customers,
		domainPublisher: domainPublisher,
	}
}

func (a Application) RegisterCustomer(ctx context.Context, register RegisterCustomer) error {
	customer, err := domain.RegisterCustomer(register.ID, register.Name, register.SmsNumber)
	if err != nil {
		return err
	}

	err = a.customers.Save(ctx, customer)
	if err != nil {
		return err
	}

	// publish domain event
	if err = a.domainPublisher.Publish(ctx, customer.GetEvents()...); err != nil {
		return err
	}

	return nil
}

func (a Application) AuthorizeCustomer(ctx context.Context, authorize AuthorizeCustomer) error {
	customer, err := a.customers.Find(ctx, authorize.ID)
	if err != nil {
		return err
	}

	if !customer.Enabled {
		return errors.Wrap(errors.ErrUnauthorized, "customer is not authorized")
	}

	return nil
}

func (a Application) GetCustomer(ctx context.Context, get GetCustomer) (*domain.Customer, error) {
	return a.customers.Find(ctx, get.ID)
}

func (a Application) GetCustomers(ctx context.Context, get GetCustomers) ([]*domain.Customer, error) {
	return a.customers.FindAll(ctx)
}

func (a Application) EnableCustomer(ctx context.Context, enable EnableCustomer) error {
	customer, err := a.customers.Find(ctx, enable.ID)
	if err != nil {
		return err
	}

	err = customer.Enable()
	if err != nil {
		return err
	}

	return a.customers.Update(ctx, customer)
}

func (a Application) DisableCustomer(ctx context.Context, disable DisableCustomer) error {
	customer, err := a.customers.Find(ctx, disable.ID)
	if err != nil {
		return err
	}

	err = customer.Disable()
	if err != nil {
		return err
	}

	return a.customers.Update(ctx, customer)
}
