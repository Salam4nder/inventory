package service

import (
	"github.com/Salam4nder/inventory/pkg/model"
	"github.com/Salam4nder/inventory/pkg/persistence"
)

// Service is an interface of basic CRUD operations.
type Service interface {
	Create(model.Item) (*model.Item, error)
	GetSingleByName(name string) (*model.Item, error)
	GetSingleByFilter(model.Item) (*model.Item, error)
	GetAll() ([]model.Item, error)
	GetFilter(model.Item) ([]model.Item, error)
	Update() (*model.Item, error)
	DeleteByName(name string) error
	DeleteByFilter(model.Item) error
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
func (i *Inventory) Create(model.Item) (*model.Item, error) {
	return &model.Item{}, nil
}

// GetSingleByName returns a model.Item by its Name.
func (i *Inventory) GetSingleByName(name string) (*model.Item, error) {
	return &model.Item{}, nil
}

// GetSingleByFilter returns a model.Item based off of a filter.
func (i *Inventory) GetSingleByFilter(model.Item) (*model.Item, error) {
	return &model.Item{}, nil
}

// GetAll returns all model.Items.
func (i *Inventory) GetAll() ([]model.Item, error) {
	return []model.Item{}, nil
}

// GetFilter returns a model.Item based off of a filter.
func (i *Inventory) GetFilter(model.Item) ([]model.Item, error) {
	return []model.Item{}, nil
}

// Update updates a model.Item.
func (i *Inventory) Update() (*model.Item, error) {
	return &model.Item{}, nil
}

// DeleteByName deletes a model.Item based off of a name.
func (i *Inventory) DeleteByName(name string) error {
	return nil
}

// DeleteByFilter deletes model.Item(s) based off of a filter.
func (i *Inventory) DeleteByFilter(model.Item) error {
	return nil
}
