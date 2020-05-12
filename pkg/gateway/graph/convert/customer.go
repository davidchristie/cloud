package convert

import (
	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
)

func Customer(customer *customer.Customer) *model.Customer {
	if customer == nil {
		return nil
	}
	return &model.Customer{
		FirstName: customer.FirstName,
		ID:        customer.ID.String(),
		LastName:  customer.LastName,
	}
}
