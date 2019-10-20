package usecase

import (
	"context"
	"drone-test/domain"
	"drone-test/repository"
)

type FindItemUseCase struct {
	itemRepository repository.IItemRepository
}

func NewFindItemUseCase(itemRepository repository.IItemRepository) *FindItemUseCase {
	return &FindItemUseCase{itemRepository}
}

type FindItemDto struct {
	Item *domain.Item
}

func (s *FindItemUseCase) SearchItemByID(context context.Context, itemID domain.ItemID) *FindItemDto {
	item := s.itemRepository.ResolveByID(context, itemID)
	return &FindItemDto{
		item,
	}
}
