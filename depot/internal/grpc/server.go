package grpc

import (
	"context"
	"eda-in-go/depot/depotpb"
	"eda-in-go/depot/internal/application"
	"eda-in-go/depot/internal/application/commands"
	"eda-in-go/depot/internal/application/queries"
	"eda-in-go/depot/internal/domain"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	depotpb.UnimplementedDepotServiceServer
}

var _ depotpb.DepotServiceServer = (*server)(nil)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	depotpb.RegisterDepotServiceServer(registrar, server{app: app})
	return nil
}

func (s server) GetShoppingLists(ctx context.Context, request *depotpb.GetShoppingListsRequest) (*depotpb.GetShoppingListsResposne, error) {
	domainShoppingLists, err := s.app.GetShoppingLists(ctx, queries.GetShoppingLists{})
	if err != nil {
		return nil, err
	}

	pbShoppingLists := make([]*depotpb.ShoppingList, len(domainShoppingLists))
	for _, domainShopplist := range domainShoppingLists {
		pbShoppingLists = append(pbShoppingLists, s.shoppingListFromDomain(domainShopplist))
	}
	return &depotpb.GetShoppingListsResposne{
		Shoplists: pbShoppingLists,
	}, nil
}

func (s server) CreateShoppingList(ctx context.Context, request *depotpb.CreateShoppingListRequest) (*depotpb.CreateShoppingListResponse, error) {
	id := uuid.New().String()

	items := make([]commands.OrderItem, 0, len(request.GetItems()))
	for _, item := range request.Items {
		items = append(items, s.itemToDomain(item))
	}

	err := s.app.CreateShoppingList(ctx, commands.CreateShoppingList{
		ID:      id,
		OrderID: request.GetOrderId(),
		Items:   items,
	})

	return &depotpb.CreateShoppingListResponse{}, err
}

func (s server) CancelShoppingList(ctx context.Context, request *depotpb.CancelShoppingListRequest) (*depotpb.CancelShoppingListResponse, error) {
	err := s.app.CancelShoppingList(ctx, commands.CancelShoppingList{
		ID: request.GetId(),
	})
	return &depotpb.CancelShoppingListResponse{}, err
}

func (s server) AssignShoppingList(ctx context.Context, request *depotpb.AssignShoppingListRequest) (*depotpb.AssignShoppingListResponse, error) {
	err := s.app.AssignShoppingList(ctx, commands.AssignShoppingList{
		ID:    request.GetId(),
		BotID: request.GetBotId(),
	})
	return &depotpb.AssignShoppingListResponse{}, err
}

func (s server) CompleteShoppingList(ctx context.Context, request *depotpb.CompleteShoppingListRequest) (*depotpb.CompleteShoppingListResponse, error) {
	err := s.app.CompleteShoppingList(ctx, commands.CompleteShoppingList{
		ID: request.GetId(),
	})
	return &depotpb.CompleteShoppingListResponse{}, err
}

func (s server) itemToDomain(item *depotpb.OrderItem) commands.OrderItem {
	return commands.OrderItem{
		StoreID:   item.GetStoreId(),
		ProductID: item.GetProductId(),
		Quantity:  int(item.GetQuantity()),
	}
}

func (s server) shoppingListFromDomain(shoppingList *domain.ShoppingList) *depotpb.ShoppingList {
	return &depotpb.ShoppingList{
		Id:            shoppingList.ID,
		OrderId:       shoppingList.OrderID,
		Stops:         s.stopsFromDomain(shoppingList.Stops),
		AssignedBotId: shoppingList.AssignedBotID,
		Status:        shoppingList.Status.String(),
	}
}

func (s server) stopsFromDomain(stops domain.Stops) map[string]*depotpb.Stop {
	pbStops := make(map[string]*depotpb.Stop, len(stops))
	for stopName, stop := range stops {
		pbStops[stopName] = s.stopFromDomain(stop)
	}
	return pbStops
}

func (s server) stopFromDomain(stop *domain.Stop) *depotpb.Stop {
	pbItems := make(map[string]*depotpb.Item, len(stop.Items))
	for ItemName, domainItem := range stop.Items {
		pbItems[ItemName] = s.itemFromDomain(domainItem)
	}
	return &depotpb.Stop{
		StoreName:     stop.StoreName,
		StoreLocation: stop.StoreLocation,
		Items:         pbItems,
	}
}

func (s server) itemFromDomain(item *domain.Item) *depotpb.Item {
	return &depotpb.Item{
		Name:     item.ProductName,
		Quantity: int32(item.Quantity),
	}
}
