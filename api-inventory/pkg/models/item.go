package models

import (
	"time"

	"github.com/google/uuid"
)

// Inventory Item model
type Item struct {
	ID        uuid.UUID `json: "id"`
	Name      string    `json: "name"`
	Stock     uint      `json: "stock"`
	Price     float64   `json: "price"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}
