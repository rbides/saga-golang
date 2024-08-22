package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	Processing = OrderStatus("PROCESSING")
	Completed  = OrderStatus("COMPLETED")
	Failed     = OrderStatus("FAILED")
)

// Inventory Item model
type Order struct {
	ID        uuid.UUID   `json: "id"`
	ItemID    uuid.UUID   `json: "item_id"`
	Quantity  uint        `json: "quantity"`
	Status    OrderStatus `json: "status"`
	CreatedAt time.Time   `json: "created_at"`
	UpdatedAt time.Time   `json: "updated_at"`
}
