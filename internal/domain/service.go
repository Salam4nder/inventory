package domain

import (
	"context"

	"github.com/Salam4nder/inventory/internal/persistence"

	"github.com/google/uuid"
)

// Service is an interface of basic CRUD operations.
type Service interface {
	Create(ctx context.Context, item persistence.Item) (
		uuid.UUID, error)
	Read(ctx context.Context, uuid string) (
		*persistence.Item, error)
	ReadAll(ctx context.Context) ([]*persistence.Item, error)
	ReadBy(ctx context.Context, filter persistence.ItemFilter) (
		[]*persistence.Item, error)
	Update(ctx context.Context, item *persistence.Item) (
		*persistence.Item, error)
	Delete(ctx context.Context, uuid string) error
}

// Inventory is a service that implements the Service interface.
type Inventory struct {
	storage persistence.Repository
}

// New returns a new Inventory service.
func New(r persistence.Repository) *Inventory {
	return &Inventory{storage: r}
}
