package message

import "github.com/google/uuid"

type EventMetadata struct {
	CorrelationID uuid.UUID
}
