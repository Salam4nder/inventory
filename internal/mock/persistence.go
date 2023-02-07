package mock

import (
	"context"
	"errors"

	"github.com/Salam4nder/inventory/internal/persistence"

	"github.com/google/uuid"
)

// Persistence is a mock implementation of persistence.Repository.
type Persistence struct {
	storage []persistence.Item
	fails   bool
}

// NewPersistence is a constructor for mock.Persistence.
func NewPersistence() *Persistence {
	return &Persistence{
		storage: make([]persistence.Item, 0),
	}
}

// Create is a mock implementation of domain.Inventory.Create.
func (p *Persistence) Create(
	ctx context.Context, item persistence.Item,
) (uuid.UUID, error) {
	if p.fails {
		return uuid.Nil, errors.New("failed")
	}

	p.storage = append(p.storage, item)
	return item.ID, nil
}

// Read is a mock implementation of domain.Inventory.Read.
func (p *Persistence) Read(ctx context.Context, uuid string) (
	*persistence.Item, error) {
	if p.fails {
		return nil, errors.New("failed")
	}

	for _, item := range p.storage {
		if item.ID.String() == uuid {
			return &item, nil
		}
	}

	return &persistence.Item{}, errors.New("not found")
}

// ReadAll is a mock implementation of domain.Inventory.ReadAll.
func (p *Persistence) ReadAll(ctx context.Context) (
	[]persistence.Item, error) {
	if p.fails {
		return nil, errors.New("failed")
	}

	return p.storage, nil
}

// ReadBy is a mock implementation of domain.Inventory.Readby.
func (p *Persistence) ReadBy(ctx context.Context,
	filter persistence.ItemFilter) (
	[]*persistence.Item, error) {
	if p.fails {
		return nil, errors.New("failed")
	}

	var result []*persistence.Item

	for _, item := range p.storage {
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
func (p *Persistence) Update(
	ctx context.Context, item *persistence.Item) (
	*persistence.Item, error) {
	if p.fails {
		return nil, errors.New("failed")
	}

	for idx, val := range p.storage {
		if val.ID == item.ID {
			p.storage[idx] = *item
			return item, nil
		}
	}

	return nil, errors.New("not found")
}

// Delete is a mock implementation of domain.Inventory.Delete.
func (p *Persistence) Delete(ctx context.Context, uuid string) error {
	if p.fails {
		return errors.New("failed")
	}

	for idx, item := range p.storage {
		if item.ID.String() == uuid {
			p.storage = append(p.storage[:idx], p.storage[idx+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}
