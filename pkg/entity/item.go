package entity

import "time"

// Item represents a single item in the database.
type Item struct {
	Name      string    `json:"name"`
	Unit      string    `json:"unit"`
	Amount    float64   `json:"amount"`
	ExpiresAt time.Time `json:"expires_at"`
}
