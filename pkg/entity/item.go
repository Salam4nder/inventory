package entity

import "time"

// Item represents a single item in the database.
type Item struct {
	Name      string    `bson:"name"`
	Unit      string    `bson:"unit"`
	Amount    float64   `bson:"amount"`
	ExpiresAt time.Time `bson:"expires_at"`
}
