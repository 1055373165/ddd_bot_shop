package queries

import (
	"context"

	"eda-in-go/stores/internal/domain"
)

type GetStores struct {
}

type GetStoresHandler struct {
	stores domain.StoreRepository
}

func NewGetStoresHandler(stores domain.StoreRepository) GetStoresHandler {
	return GetStoresHandler{stores: stores}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.Store, error) {
	return h.stores.FindAll(ctx)
}
