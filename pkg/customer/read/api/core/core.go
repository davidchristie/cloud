package core

import (
	"context"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/database"
)

type Core interface {
	Customer(context.Context, string) (*customer.Customer, error)
	Customers(context.Context) ([]*customer.Customer, error)
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
