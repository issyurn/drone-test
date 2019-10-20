package usecase

import (
	"context"
	"drone-test/domain"
	"drone-test/repository"
	"reflect"
	"testing"
)

func TestFindItem_SearchItemByID(t *testing.T) {
	type args struct {
		itemID domain.ItemID
	}
	tests := []struct {
		name string
		args args
		want *FindItemDto
	}{
		// cases.
		{"item ID 1", args{domain.ItemID(1)}, &FindItemDto{&domain.Item{1}}},
		{"item ID 2", args{domain.ItemID(2)}, &FindItemDto{&domain.Item{2}}},
	}

	itemRepository := newItemRepositoryMock()
	s := NewFindItemUseCase(itemRepository)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.SearchItemByID(context.Background(), tt.args.itemID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItemByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

type itemRepositoryMock struct {
}

func (r *itemRepositoryMock) ResolveByID(context context.Context, itemID domain.ItemID) *domain.Item {
	return &domain.Item{ID: itemID}
}

func newItemRepositoryMock() repository.IItemRepository {
	return &itemRepositoryMock{}
}
