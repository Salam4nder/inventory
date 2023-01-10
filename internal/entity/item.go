package entity

import (
	"time"

	"github.com/google/uuid"
)

// Item represents a single item in the inventory.
type Item struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Unit      string    `json:"unit"`
	Amount    float64   `json:"amount"`
	ExpiresAt time.Time `json:"expires_at"`
}

// ItemFilter represents a filter for items.
// Used to build a query to the database.
type ItemFilter struct {
	Name      string    `json:"name"`
	Unit      string    `json:"unit"`
	Amount    float64   `json:"amount"`
	ExpiresAt time.Time `json:"expires_at"`
}
