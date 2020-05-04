package order

import (
	"time"

	"github.com/google/uuid"
)

type LineItem struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type Order struct {
	CreatedAt time.Time   `json:"created_at"`
	ID        uuid.UUID   `json:"id"`
	LineItems []*LineItem `json:"line_items"`
}
