package repository

import (
	"context"
	"drone-test/domain"
)

type ItemRepository struct {
	id domain.ItemID
}

func NewItemRepository() IItemRepository {
	return new(ItemRepository)
}

func (r *ItemRepository) ResolveByID(context context.Context, itemID domain.ItemID) *domain.Item {
	// FIXME Get from any data source
	return &domain.Item{
		ID: itemID,
	}
}
