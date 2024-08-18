package application

import (
	"context"

	"github.com/stackus/errors"

	"eda-in-go/payments/internal/models"
)

type (
	AuthorizePayment struct {
		ID         string
		CustomerID string
		Amount     float64
	}

	ConfirmPayment struct {
		ID string
	}

	GetPayments struct {
	}

	CreateInvoice struct {
		ID        string
		OrderID   string
		PaymentID string
		Amount    float64
	}

	GetInvoices struct {
	}

	AdjustInvoice struct {
		ID     string
		Amount float64
	}

	PayInvoice struct {
		ID string
	}

	CancelInvoice struct {
		ID string
	}

	App interface {
		AuthorizePayment(ctx context.Context, authorize AuthorizePayment) error
		ConfirmPayment(ctx context.Context, confirm ConfirmPayment) error
		CreateInvoice(ctx context.Context, create CreateInvoice) error
		AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error
		PayInvoice(ctx context.Context, pay PayInvoice) error
		CancelInvoice(ctx context.Context, cancel CancelInvoice) error
		GetPayments(ctx context.Context, get GetPayments) ([]*models.Payment, error)
		GetInvoices(ctx context.Context, get GetInvoices) ([]*models.Invoice, error)
	}

	Application struct {
		invoices InvoiceRepository
		payments PaymentRepository
		orders   OrderRepository
	}
)

var _ App = (*Application)(nil)

func New(invoices InvoiceRepository, payments PaymentRepository, orders OrderRepository) *Application {
	return &Application{
		invoices: invoices,
		payments: payments,
		orders:   orders,
	}
}

func (a Application) GetPayments(ctx context.Context, get GetPayments) ([]*models.Payment, error) {
	payments, err := a.payments.FindAll(ctx)

	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, "payment cannot be find all")
	}

	return payments, nil
}

func (a Application) AuthorizePayment(ctx context.Context, authorize AuthorizePayment) error {
	return a.payments.Save(ctx, &models.Payment{
		ID:         authorize.ID,
		CustomerID: authorize.CustomerID,
		Amount:     authorize.Amount,
	})
}

func (a Application) ConfirmPayment(ctx context.Context, confirm ConfirmPayment) error {
	if payment, err := a.payments.Find(ctx, confirm.ID); err != nil || payment == nil {
		return errors.Wrap(errors.ErrNotFound, "payment cannot be confirmed")
	}

	return nil
}

func (a Application) CreateInvoice(ctx context.Context, create CreateInvoice) error {
	return a.invoices.Save(ctx, &models.Invoice{
		ID:      create.ID,
		OrderID: create.OrderID,
		Amount:  create.Amount,
		Status:  models.InvoicePending,
	})
}

func (a Application) GetInvoices(ctx context.Context, get GetInvoices) ([]*models.Invoice, error) {
	invoices, err := a.invoices.FindAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "invoice cannot be find all")
	}

	return invoices, nil
}

func (a Application) AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error {
	invoice, err := a.invoices.Find(ctx, adjust.ID)
	if err != nil {
		return err
	}

	invoice.Amount = adjust.Amount

	return a.invoices.Update(ctx, invoice)
}

func (a Application) PayInvoice(ctx context.Context, pay PayInvoice) error {
	invoice, err := a.invoices.Find(ctx, pay.ID)
	if err != nil {
		return err
	}

	if invoice.Status != models.InvoicePending {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be paid for")
	}

	invoice.Status = models.InvoicePaid

	if err = a.orders.Complete(ctx, invoice.ID, invoice.OrderID); err != nil {
		return err
	}

	return a.invoices.Update(ctx, invoice)
}

func (a Application) CancelInvoice(ctx context.Context, cancel CancelInvoice) error {
	invoice, err := a.invoices.Find(ctx, cancel.ID)
	if err != nil {
		return err
	}

	if invoice.Status != models.InvoicePending {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be paid for")
	}

	invoice.Status = models.InvoiceCanceled

	return a.invoices.Update(ctx, invoice)
}
