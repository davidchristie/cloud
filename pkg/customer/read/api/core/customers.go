package core

import (
	"context"
	"log"

	"github.com/davidchristie/cloud/pkg/customer"
)

func (c *core) Customers(ctx context.Context) ([]*customer.Customer, error) {
	customers, err := c.CustomerRepository.FindCustomers(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return customers, nil
}
