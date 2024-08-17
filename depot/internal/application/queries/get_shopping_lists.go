package queries

import (
	"context"
	"eda-in-go/depot/internal/domain"
)

type GetShoppingLists struct {
}

type GetShoppingListsHandler struct {
	shoppingLists domain.ShoppingListRepository
}

func NewGetShoppingListsHandler(shoppingLists domain.ShoppingListRepository) GetShoppingListsHandler {
	return GetShoppingListsHandler{shoppingLists: shoppingLists}
}

func (h GetShoppingListHandler) GetShoppingLists(ctx context.Context, _ GetShoppingLists) ([]*domain.ShoppingList, error) {
	return h.shoppingLists.FindAll(ctx)
}
