package queries

import (
	"context"
	"eda-in-go/ordering/internal/domain"

	"github.com/stackus/errors"
)

type GetOrders struct{}

type GetOrdersHandler struct {
	repo domain.OrderRepository
}

func NewGetOrdersHandler(repo domain.OrderRepository) GetOrdersHandler {
	return GetOrdersHandler{repo: repo}
}

func (h GetOrdersHandler) GetOrders(ctx context.Context, query GetOrders) ([]*domain.Order, error) {
	orders, err := h.repo.FindAll(ctx)
	return orders, errors.Wrap(err, "get order list query")
}
