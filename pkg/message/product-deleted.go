package message

import (
	"github.com/google/uuid"
)

type ProductDeletedEvent struct {
	Metadata  *EventMetadata `json:"metadata"`
	ProductID uuid.UUID      `json:"product_id"`
}
