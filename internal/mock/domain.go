package mock

import (
	"context"
	"errors"

	"github.com/Salam4nder/inventory/internal/entity"
	"github.com/google/uuid"
)

// Inventory is a mock implementation of domain.Inventory.
type Inventory struct {
	storage []entity.Item
}

// NewInventory is a constructor for mock.Inventory.
func NewInventory() *Inventory {
	return &Inventory{
		storage: make([]entity.Item, 0),
	}
}

// Create is a mock implementation of domain.Inventory.Create.
func (i *Inventory) Create(
	ctx context.Context, item entity.Item,
) (uuid.UUID, error) {
	i.storage = append(i.storage, item)
	return item.ID, nil
}

// Read is a mock implementation of domain.Inventory.Read.
func (i *Inventory) Read(ctx context.Context, uuid string) (
	*entity.Item, error,
) {
	for _, item := range i.storage {
		if item.ID.String() == uuid {
			return &item, nil
		}
	}

	return &entity.Item{}, errors.New("not found")
}

// ReadAll is a mock implementation of domain.Inventory.ReadAll.
func (i *Inventory) ReadAll(ctx context.Context) (
	[]entity.Item, error) {
	return i.storage, nil
}

// ReadBy is a mock implementation of domain.Inventory.Readby.
func (i *Inventory) ReadBy(ctx context.Context, filter entity.ItemFilter) (
	[]*entity.Item, error) {
	var result []*entity.Item

	for _, item := range i.storage {
		if item.Name == filter.Name {
			result = append(result, &item)
			continue
		}
		if item.Amount == filter.Amount {
			result = append(result, &item)
			continue
		}
		if item.ExpiresAt == filter.ExpiresAt {
			result = append(result, &item)
			continue
		}
	}

	if len(result) == 0 {
		return nil, errors.New("nothing found")
	}

	return result, nil
}

// Update is a mock implementation of domain.Inventory.Update.
func (i *Inventory) Update(ctx context.Context, item *entity.Item) error {
	for idx, val := range i.storage {
		if val.ID == item.ID {
			i.storage[idx] = *item
			return nil
		}
	}

	return errors.New("not found")
}

// Delete is a mock implementation of domain.Inventory.Delete.
func (i *Inventory) Delete(ctx context.Context, uuid string) error {
	for idx, item := range i.storage {
		if item.ID.String() == uuid {
			i.storage = append(i.storage[:idx], i.storage[idx+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}
