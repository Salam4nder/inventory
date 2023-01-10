package inventory

import (
	"context"

	"github.com/Salam4nder/inventory/pkg/entity"
	"github.com/Salam4nder/inventory/pkg/persistence"

	"github.com/google/uuid"
)

// Service is an interface of basic CRUD operations.
type Service interface {
	Create(ctx context.Context, item entity.Item) (
		uuid.UUID, error)
	Read(ctx context.Context, uuid string) (
		*entity.Item, error)
	ReadBy(ctx context.Context, filter entity.ItemFilter) (
		[]*entity.Item, error)
	Update(ctx context.Context, item *entity.Item) (
		*entity.Item, error)
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
