package commands

import (
	"context"
	"eda-in-go/depot/internal/domain"
)

type CompleteShoppingList struct {
	ID string
}

type CompleteShoppingListHandler struct {
	shoppingList domain.ShoppingListRepository
	orders       domain.OrderRepository
}

func NewCompleteShoppingListHandler(shoppingLists domain.ShoppingListRepository, oders domain.OrderRepository) CompleteShoppingListHandler {
	return CompleteShoppingListHandler{
		shoppingList: shoppingLists,
		orders:       oders,
	}
}

func (r CompleteShoppingListHandler) CompleteShoppingList(ctx context.Context, cmd CompleteShoppingList) error {
	list, err := r.shoppingList.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Complete()
	if err != nil {
		return err
	}

	// 购买的物品已由机器人挑拣完毕，订单可以就绪
	err = r.orders.Ready(ctx, list.OrderID)
	if err != nil {
		return err
	}

	return r.shoppingList.Update(ctx, list)
}
