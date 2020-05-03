package entity

import (
	"github.com/google/uuid"
)

type Customer struct {
	FirstName string    `json:"first_name"`
	ID        uuid.UUID `json:"id"`
	LastName  string    `json:"last_name"`
}
