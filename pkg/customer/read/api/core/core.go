package core

import (
	"context"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/database"
	"github.com/google/uuid"
)

type Core interface {
	Customer(context.Context, string) (*customer.Customer, error)
	Customers(ctx context.Context, ids []string) (map[uuid.UUID]*customer.Customer, error)
}

type core struct {
	CustomerRepository database.CustomerRepository
}

func NewCore() Core {
	db := database.Connect()
	return &core{
		CustomerRepository: database.NewCustomerRepository(db),
	}
}
