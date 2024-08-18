package commands

import (
	"context"

	"eda-in-go/internal/ddd"
	"eda-in-go/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	stores          domain.StoreRepository
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher
}

func NewRemoveProductHandler(stores domain.StoreRepository, products domain.ProductRepository, domainPublisher ddd.EventPublisher) RemoveProductHandler {
	return RemoveProductHandler{
		stores:          stores,
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.products.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	// add remove product event
	if err = product.Remove(); err != nil {
		return err
	}

	if err = h.products.Delete(ctx, cmd.ID); err != nil {
		return err
	}

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, product.GetEvents()...); err != nil {
		return err
	}

	return nil
}
