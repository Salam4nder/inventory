package inventory

import (
	"github.com/Salam4nder/inventory/pkg/entity"
	"github.com/Salam4nder/inventory/pkg/persistence"
)

// Service is an interface of basic CRUD operations.
type Service interface {
	Create(entity.Item) (*entity.Item, error)
	ReadSingleByName(string) (*entity.Item, error)
	ReadSingleByFilter(entity.Item) (*entity.Item, error)
	ReadAll() ([]entity.Item, error)
	ReadByFilter(entity.Item) ([]entity.Item, error)
	Update() (*entity.Item, error)
	DeleteByName(name string) error
	DeleteByFilter(entity.Item) error
}

// Inventory is a service that implements the Service interface.
type Inventory struct {
	storage persistence.Repository
}

// NewInventory returns a new Inventory service.
func NewInventory(storage persistence.Repository) *Inventory {
	return &Inventory{storage: storage}
}

// Create creates a new Item.
func (i *Inventory) Create(entity.Item) (*entity.Item, error) {
	return &entity.Item{}, nil
}

// ReadSingleByName returns a Item by its Name.
func (i *Inventory) ReadSingleByName(name string) (*entity.Item, error) {
	return &entity.Item{}, nil
}

// ReadSingleByFilter returns a Item based off of a filter.
func (i *Inventory) ReadSingleByFilter(entity.Item) (*entity.Item, error) {
	return &entity.Item{}, nil
}

// ReadAll returns all Items.
func (i *Inventory) ReadAll() ([]entity.Item, error) {
	return []entity.Item{}, nil
}

// ReadByFilter returns a Item based off of a filter.
func (i *Inventory) ReadByFilter(entity.Item) ([]entity.Item, error) {
	return []entity.Item{}, nil
}

// Update updates a Item.
func (i *Inventory) Update() (*entity.Item, error) {
	return &entity.Item{}, nil
}

// DeleteByName deletes a Item based off of a name.
func (i *Inventory) DeleteByName(name string) error {
	return nil
}

// DeleteByFilter deletes Item(s) based off of a filter.
func (i *Inventory) DeleteByFilter(entity.Item) error {
	return nil
}
