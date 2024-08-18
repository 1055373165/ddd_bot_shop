package commands

import (
	"context"
	"eda-in-go/depot/internal/domain"
	"eda-in-go/internal/ddd"
)

type CancelShoppingList struct {
	ID string
}

type CancelShoppingListHandler struct {
	shoppingLists  domain.ShoppingListRepository
	domainPushlier ddd.EventPublisher
}

func NewCancelShoppingListHandler(shoppingLists domain.ShoppingListRepository, domainPublisher ddd.EventPublisher) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingLists:  shoppingLists,
		domainPushlier: domainPublisher,
	}
}

func (h CancelShoppingListHandler) CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	// set list new status (list')
	err = list.Cancel()
	if err != nil {
		return err
	}

	// update shoppling list status with list'
	if err = h.shoppingLists.Update(ctx, list); err != nil {
		return err
	}

	// publish domain event
	if err = h.domainPushlier.Publish(ctx, list.GetEvents()...); err != nil {
		return err
	}

	return nil
}
