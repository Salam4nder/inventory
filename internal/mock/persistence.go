package mock

import (
	"context"
	"errors"

	"github.com/Salam4nder/inventory/internal/entity"
	"github.com/google/uuid"
)

// Persistence is a mock implementation of persistence.Repository.
type Persistence struct {
	storage []entity.Item
	fails   bool
}

// NewPersistence is a constructor for mock.Persistence.
func NewPersistence() *Persistence {
	return &Persistence{
		storage: make([]entity.Item, 0),
	}
}

// Create is a mock implementation of domain.Inventory.Create.
func (i *Persistence) Create(
	ctx context.Context, item entity.Item,
) (uuid.UUID, error) {
	if i.fails {
		return uuid.Nil, errors.New("failed")
	}
	i.storage = append(i.storage, item)
	return item.ID, nil
}

// Read is a mock implementation of domain.Inventory.Read.
func (i *Persistence) Read(ctx context.Context, uuid string) (
	*entity.Item, error) {
	if i.fails {
		return nil, errors.New("failed")
	}
	for _, item := range i.storage {
		if item.ID.String() == uuid {
			return &item, nil
		}
	}

	return &entity.Item{}, errors.New("not found")
}

// ReadAll is a mock implementation of domain.Inventory.ReadAll.
func (i *Persistence) ReadAll(ctx context.Context) (
	[]entity.Item, error) {
	return i.storage, nil
}

// ReadBy is a mock implementation of domain.Inventory.Readby.
func (i *Persistence) ReadBy(ctx context.Context,
	filter entity.ItemFilter) (
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
func (i *Persistence) Update(ctx context.Context, item *entity.Item) (*entity.Item, error) {
	for idx, val := range i.storage {
		if val.ID == item.ID {
			i.storage[idx] = *item
			return item, nil
		}
	}

	return nil, errors.New("not found")
}

// Delete is a mock implementation of domain.Inventory.Delete.
func (i *Persistence) Delete(ctx context.Context, uuid string) error {
	for idx, item := range i.storage {
		if item.ID.String() == uuid {
			i.storage = append(i.storage[:idx], i.storage[idx+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}
