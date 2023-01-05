package service

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
	storage persistence.Storage
}

// NewInventory returns a new Inventory service.
func NewInventory(storage persistence.Storage) *Inventory {
	return &Inventory{storage: storage}
}

// Create creates a new model.Item.
func (i *Inventory) Create(entity.Item) (*entity.Item, error) {
	return &entity.Item{}, nil
}

// ReadSingleByName returns a model.Item by its Name.
func (i *Inventory) ReadSingleByName(name string) (*entity.Item, error) {
	return &entity.Item{}, nil
}

// ReadSingleByFilter returns a model.Item based off of a filter.
func (i *Inventory) ReadSingleByFilter(entity.Item) (*entity.Item, error) {
	return &entity.Item{}, nil
}

// ReadAll returns all model.Items.
func (i *Inventory) ReadAll() ([]entity.Item, error) {
	return []entity.Item{}, nil
}

// ReadByFilter returns a model.Item based off of a filter.
func (i *Inventory) ReadByFilter(entity.Item) ([]entity.Item, error) {
	return []entity.Item{}, nil
}

// Update updates a model.Item.
func (i *Inventory) Update() (*entity.Item, error) {
	return &entity.Item{}, nil
}

// DeleteByName deletes a model.Item based off of a name.
func (i *Inventory) DeleteByName(name string) error {
	return nil
}

// DeleteByFilter deletes model.Item(s) based off of a filter.
func (i *Inventory) DeleteByFilter(entity.Item) error {
	return nil
}
