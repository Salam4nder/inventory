package domain

import (
	"context"

	"github.com/Salam4nder/inventory/internal/entity"

	"github.com/google/uuid"
)

// Create creates a new Item from an entity.Item structure.
func (i *Inventory) Create(
	ctx context.Context, item entity.Item) (uuid.UUID, error) {
	uuID, err := i.storage.Create(ctx, item)
	if err != nil {
		return uuid.Nil, err
	}

	return uuID, nil
}

// Read returns an Item based off of an uuid from storage.
func (i *Inventory) Read(
	ctx context.Context, uuid string) (*entity.Item, error) {
	item, err := i.storage.Read(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// ReadBy returns Items from storage that match the filter.
func (i *Inventory) ReadBy(
	ctx context.Context, filter entity.ItemFilter) (
	[]*entity.Item, error) {
	items, err := i.storage.ReadBy(ctx, filter)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Update updates the given Item and returns it.
func (i *Inventory) Update(
	ctx context.Context, item *entity.Item) (
	*entity.Item, error) {
	updatedItem, err := i.storage.Update(ctx, item)
	if err != nil {
		return nil, err
	}

	return updatedItem, nil
}

// Delete deletes an Item based off of an uuid.
func (i *Inventory) Delete(
	ctx context.Context, uuid string) error {
	return i.storage.Delete(ctx, uuid)
}
