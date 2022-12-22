package service

import (
	"github.com/Salam4nder/inventory/pkg/model"
	"github.com/Salam4nder/inventory/pkg/persistence"
)

// Service is an interface of basic CRUD operations.
type Service interface {
	Create(model.Item) (*model.Item, error)
	ReadSingleByName(string) (*model.Item, error)
	ReadSingleByFilter(model.Item) (*model.Item, error)
	ReadAll() ([]model.Item, error)
	ReadByFilter(model.Item) ([]model.Item, error)
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

// ReadSingleByName returns a model.Item by its Name.
func (i *Inventory) ReadSingleByName(name string) (*model.Item, error) {
	return &model.Item{}, nil
}

// ReadSingleByFilter returns a model.Item based off of a filter.
func (i *Inventory) ReadSingleByFilter(model.Item) (*model.Item, error) {
	return &model.Item{}, nil
}

// ReadAll returns all model.Items.
func (i *Inventory) ReadAll() ([]model.Item, error) {
	return []model.Item{}, nil
}

// ReadByFilter returns a model.Item based off of a filter.
func (i *Inventory) ReadByFilter(model.Item) ([]model.Item, error) {
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
