package model

import "github.com/google/uuid"

type LineItem struct {
	ProductID uuid.UUID
	Quantity  int
}

type Order struct {
	CustomerID uuid.UUID
	CreatedAt  string      `json:"createdAt"`
	ID         string      `json:"id"`
	LineItems  []*LineItem `json:"lineItems"`
}
