package http

import (
	"time"

	"github.com/Salam4nder/inventory/internal/persistence"

	"github.com/google/uuid"
)

// CreateItemRequest is a request to create an item.
type CreateItemRequest struct {
	Name      string    `json:"name" binding:"required"`
	Unit      string    `json:"unit" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

// UpdateItemRequest is a request to update an item.
type UpdateItemRequest struct {
	ID        uuid.UUID `json:"uuid" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Unit      string    `json:"unit" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

// ToPersistenceItem converts CreateItemRequest
// to a persistence.Item.
func (r *CreateItemRequest) ToPersistenceItem() persistence.Item {
	return persistence.Item{
		Name:      r.Name,
		Unit:      r.Unit,
		Amount:    r.Amount,
		ExpiresAt: r.ExpiresAt,
	}
}

// ToPersistenceItem converts UpdateItemRequest
// to a persistence.Item.
func (r *UpdateItemRequest) ToPersistenceItem() persistence.Item {
	return persistence.Item{
		ID:        r.ID,
		Name:      r.Name,
		Unit:      r.Unit,
		Amount:    r.Amount,
		ExpiresAt: r.ExpiresAt,
	}
}
