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
