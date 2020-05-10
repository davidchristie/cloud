package core

import (
	"context"
	"errors"
	"log"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/database"
	"github.com/google/uuid"
)

var ErrCustomerNotFound = errors.New("customer not found")

func (c *core) Customer(ctx context.Context, id string) (*customer.Customer, error) {
	customerID, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return nil, ErrCustomerNotFound
	}
	customer, err := c.CustomerRepository.FindCustomer(ctx, customerID)
	switch err {
	case nil:
		return customer, nil

	case database.ErrCustomerNotFound:
		return nil, ErrCustomerNotFound

	default:
		return nil, err
	}
}
