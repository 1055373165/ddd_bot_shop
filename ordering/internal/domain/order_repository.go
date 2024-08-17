package domain

import (
	"context"
)

type OrderRepository interface {
	Find(ctx context.Context, orderID string) (*Order, error)
	FindAll(ctx context.Context) ([]*Order, error)
	Save(ctx context.Context, order *Order) error
	Update(ctx context.Context, order *Order) error
}
