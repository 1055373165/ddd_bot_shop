package grpc

import (
	"context"
	"eda-in-go/notifications/internal/application"
	"eda-in-go/notifications/notificationspb"

	"google.golang.org/grpc"
)

type server struct {
	app application.App
	notificationspb.UnimplementedNotificationsServiceServer
}

var _ notificationspb.NotificationsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	notificationspb.RegisterNotificationsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) NotifyOrderCreated(ctx context.Context, request *notificationspb.NotifyOrderCreatedRequset) (*notificationspb.NotifyOrderCreatedResponse, error) {
	err := s.app.NotifyOrderCreated(ctx, application.OrderCreated{
		OrderID:    request.GetOrderId(),
		CustomerID: request.GetCustomerId(),
	})
	return &notificationspb.NotifyOrderCreatedResponse{}, err
}

func (s server) NotifyOrderCanceled(ctx context.Context, request *notificationspb.NotifyOrderCanceledRequest) (*notificationspb.NotifyOrderCanceledResponse, error) {
	err := s.app.NotifyOrderCanceled(ctx, application.OrderCanceled{
		OrderID:    request.GetOrderId(),
		CustomerID: request.GetCustomerId(),
	})
	return &notificationspb.NotifyOrderCanceledResponse{}, err
}

func (s server) NotifyOrderReady(ctx context.Context, request *notificationspb.NotifyOrderReadyRequest) (*notificationspb.NotifyOrderReadyResponse, error) {
	err := s.app.NotifyOrderReady(ctx, application.OrderReady{
		OrderID:    request.GetOrderId(),
		CustomerID: request.GetCustomerId(),
	})
	return &notificationspb.NotifyOrderReadyResponse{}, err
}
