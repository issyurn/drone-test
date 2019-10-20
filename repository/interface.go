package repository

import (
	"context"
	"drone-test/domain"
)

type IItemRepository interface {
	ResolveByID(context context.Context, itemID domain.ItemID) *domain.Item
}
