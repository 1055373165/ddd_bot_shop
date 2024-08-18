package commands

import (
	"context"
	"eda-in-go/depot/internal/domain"
	"eda-in-go/internal/ddd"

	"github.com/pkg/errors"
)

// 创建购物单
type CreateShoppingList struct {
	ID      string
	OrderID string
	Items   []OrderItem
}

type CreateShoppingListHandler struct {
	shoppingLists   domain.ShoppingListRepository
	stores          domain.StoreRepository
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher
}

func NewCreateShoppingListHandler(shopplingLists domain.ShoppingListRepository, stores domain.StoreRepository,
	products domain.ProductRepository, domainPublisher ddd.EventPublisher) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingLists:   shopplingLists,
		stores:          stores,
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h CreateShoppingListHandler) CreateShoppingList(ctx context.Context, cmd CreateShoppingList) error {
	list := domain.CreateShopping(cmd.ID, cmd.OrderID)

	for _, item := range cmd.Items {
		// horribly inefficient 效率极低
		store, err := h.stores.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
		product, err := h.products.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
		err = list.AddItem(store, product, item.Quantity)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
	}

	if err := h.shoppingLists.Save(ctx, list); err != nil {
		return errors.Wrap(err, "scheduling shopping")
	}

	// public domain event
	if err := h.domainPublisher.Publish(ctx, list.GetEvents()...); err != nil {
		return err
	}

	return nil
}
