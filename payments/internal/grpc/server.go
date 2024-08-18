package grpc

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"eda-in-go/payments/internal/application"
	"eda-in-go/payments/paymentspb"
)

type server struct {
	app application.App
	paymentspb.UnimplementedPaymentsServiceServer
}

var _ paymentspb.PaymentsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	paymentspb.RegisterPaymentsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) AuthorizePayment(ctx context.Context, request *paymentspb.AuthorizePaymentRequest) (*paymentspb.AuthorizePaymentResponse, error) {
	id := uuid.New().String()
	err := s.app.AuthorizePayment(ctx, application.AuthorizePayment{
		ID:         id,
		CustomerID: request.GetCustomerId(),
		Amount:     request.GetAmount(),
	})
	return &paymentspb.AuthorizePaymentResponse{Id: id}, err
}

func (s server) ConfirmPayment(ctx context.Context, request *paymentspb.ConfirmPaymentRequest) (*paymentspb.ConfirmPaymentResponse, error) {
	err := s.app.ConfirmPayment(ctx, application.ConfirmPayment{
		ID: request.GetId(),
	})
	return &paymentspb.ConfirmPaymentResponse{}, err
}

func (s server) CreateInvoice(ctx context.Context, request *paymentspb.CreateInvoiceRequest) (*paymentspb.CreateInvoiceResponse, error) {
	id := uuid.New().String()
	err := s.app.CreateInvoice(ctx, application.CreateInvoice{
		ID:      id,
		OrderID: request.GetOrderId(),
		Amount:  request.GetAmount(),
	})
	return &paymentspb.CreateInvoiceResponse{
		Id: id,
	}, err
}

func (s server) AdjustInvoice(ctx context.Context, request *paymentspb.AdjustInvoiceRequest) (*paymentspb.AdjustInvoiceResponse, error) {
	err := s.app.AdjustInvoice(ctx, application.AdjustInvoice{
		ID:     request.GetId(),
		Amount: request.GetAmount(),
	})
	return &paymentspb.AdjustInvoiceResponse{}, err
}

func (s server) PayInvoice(ctx context.Context, request *paymentspb.PayInvoiceRequest) (*paymentspb.PayInvoiceResponse, error) {
	err := s.app.PayInvoice(ctx, application.PayInvoice{
		ID: request.GetId(),
	})
	return &paymentspb.PayInvoiceResponse{}, err
}

func (s server) CancelInvoice(ctx context.Context, request *paymentspb.CancelInvoiceRequest) (*paymentspb.CancelInvoiceResponse, error) {
	err := s.app.CancelInvoice(ctx, application.CancelInvoice{
		ID: request.GetId(),
	})
	return &paymentspb.CancelInvoiceResponse{}, err
}

func (s server) GetPayments(ctx context.Context, req *paymentspb.GetPaymentsRequest) (*paymentspb.GetPaymentsResponse, error) {
	payments, err := s.app.GetPayments(ctx, application.GetPayments{})
	if err != nil {
		return &paymentspb.GetPaymentsResponse{}, err
	}

	pbPayments := make([]*paymentspb.Payment, len(payments))
	for _, payment := range payments {
		pbPayment := &paymentspb.Payment{
			Id:         payment.ID,
			CustomerId: payment.CustomerID,
			Amount:     payment.Amount,
		}
		pbPayments = append(pbPayments, pbPayment)
	}

	return &paymentspb.GetPaymentsResponse{
		Payments: pbPayments,
	}, nil
}

func (s server) GetInvoices(ctx context.Context, req *paymentspb.GetInvoicesRequest) (*paymentspb.GetInvoicesResponse, error) {
	invoices, err := s.app.GetInvoices(ctx, application.GetInvoices{})
	if err != nil {
		return &paymentspb.GetInvoicesResponse{}, err
	}

	pbInvoices := make([]*paymentspb.Invoice, len(invoices))
	for _, invoice := range invoices {
		pbInvoice := &paymentspb.Invoice{
			Id:      invoice.ID,
			OrderId: invoice.OrderID,
			Amount:  invoice.Amount,
			Status:  invoice.Status.String(),
		}
		pbInvoices = append(pbInvoices, pbInvoice)
	}

	return &paymentspb.GetInvoicesResponse{
		Invoices: pbInvoices,
	}, nil
}
