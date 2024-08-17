package commands

import (
	"context"
	"eda-in-go/depot/internal/domain"

	"github.com/pkg/errors"
)

// 创建购物单
type CreateShoppingList struct {
	ID      string
	OrderID string
	Items   []OrderItem
}

type CreateShoppingListHandler struct {
	shoppingLists domain.ShoppingListRepository
	stores        domain.StoreRepository
	products      domain.ProductRepository
}

func NewCreateShoppingListHandler(shopplingLists domain.ShoppingListRepository, stores domain.StoreRepository, products domain.ProductRepository) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingLists: shopplingLists,
		stores:        stores,
		products:      products,
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

	return errors.Wrap(h.shoppingLists.Save(ctx, list), "scheduling shopping")
}
