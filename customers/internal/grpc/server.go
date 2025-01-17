package grpc

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"eda-in-go/customers/customerspb"
	"eda-in-go/customers/internal/application"
	"eda-in-go/customers/internal/domain"
)

type server struct {
	app                                             application.App
	customerspb.UnimplementedCustomersServiceServer // 继承+重写
}

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	customerspb.RegisterCustomersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) RegisterCustomer(ctx context.Context, request *customerspb.RegisterCustomerRequest) (*customerspb.RegisterCustomerResponse, error) {
	// 122 bit security.
	id := uuid.New().String()
	err := s.app.RegisterCustomer(ctx, application.RegisterCustomer{
		ID:        id,
		Name:      request.GetName(),
		SmsNumber: request.SmsNumber,
	})
	return &customerspb.RegisterCustomerResponse{Id: id}, err
}

func (s server) AuthorizeCustomer(ctx context.Context, request *customerspb.AuthorizeCustomerRequest) (*customerspb.AuthorizeCustomerResponse, error) {
	err := s.app.AuthorizeCustomer(ctx, application.AuthorizeCustomer{
		ID: request.GetId(),
	})
	return &customerspb.AuthorizeCustomerResponse{}, err
}

func (s server) GetCustomer(ctx context.Context, request *customerspb.GetCustomerRequest) (*customerspb.GetCustomerResponse, error) {
	customer, err := s.app.GetCustomer(ctx, application.GetCustomer{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &customerspb.GetCustomerResponse{
		Customer: s.customerFromDomain(customer),
	}, nil
}

func (s server) GetCustomers(ctx context.Context, request *customerspb.GetCustomersRequest) (*customerspb.GetCustomersResponse, error) {
	domainCustomers, err := s.app.GetCustomers(ctx)
	if err != nil {
		return nil, err
	}

	pbCustomers := []*customerspb.Customer{}
	for _, domainCustomer := range domainCustomers {
		pbCustomers = append(pbCustomers, s.customerFromDomain(domainCustomer))
	}

	return &customerspb.GetCustomersResponse{
		Customers: pbCustomers,
	}, nil
}

func (s server) EnableCustomer(ctx context.Context, request *customerspb.EnableCustomerRequest) (*customerspb.EnableCustomerResponse, error) {
	err := s.app.EnableCustomer(ctx, application.EnableCustomer{ID: request.GetId()})
	return &customerspb.EnableCustomerResponse{}, err
}

func (s server) DisableCustomer(ctx context.Context, request *customerspb.DisableCustomerRequest) (*customerspb.DisableCustomerResponse, error) {
	err := s.app.DisableCustomer(ctx, application.DisableCustomer{ID: request.GetId()})
	return &customerspb.DisableCustomerResponse{}, err
}

func (s server) customerFromDomain(customer *domain.Customer) *customerspb.Customer {
	return &customerspb.Customer{
		Id:        customer.ID,
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}
