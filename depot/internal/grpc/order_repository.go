package grpc

import (
	"context"
	"eda-in-go/depot/internal/domain"
	"eda-in-go/ordering/orderingpb"

	"google.golang.org/grpc"
)

type OrderRepository struct {
	client orderingpb.OrderingServiceClient
}

var _ (domain.OrderRepository) = (*OrderRepository)(nil)

func NewOrderRepository(conn *grpc.ClientConn) OrderRepository {
	return OrderRepository{
		client: orderingpb.NewOrderingServiceClient(conn),
	}
}

func (r OrderRepository) Ready(ctx context.Context, orderID string) error {
	_, err := r.client.ReadyOrder(ctx, &orderingpb.ReadyOrderRequest{Id: orderID})
	return err
}
