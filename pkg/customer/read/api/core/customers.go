package core

import (
	"context"
	"log"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/google/uuid"
)

func (c *core) Customers(ctx context.Context, ids []string) (map[uuid.UUID]*customer.Customer, error) {
	customerIDs := []uuid.UUID{}
	for _, id := range ids {
		customerID, err := uuid.Parse(id)
		if err == nil {
			customerIDs = append(customerIDs, customerID)
		} else {
			log.Println(err)
		}
	}
	return c.CustomerRepository.FindCustomers(ctx, customerIDs)
}
