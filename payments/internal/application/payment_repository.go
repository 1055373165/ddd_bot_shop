package application

import (
	"context"

	"eda-in-go/payments/internal/models"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *models.Payment) error
	Find(ctx context.Context, paymentID string) (*models.Payment, error)
	FindAll(ctx context.Context) ([]*models.Payment, error)
}
