package product

import (
	"github.com/google/uuid"
)

type Product struct {
	Description string    `json:"description"`
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
}
