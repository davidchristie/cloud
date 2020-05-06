package utility

import (
	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
)

func ConvertCustomerToModel(customer *entity.Customer) *model.Customer {
	if customer == nil {
		return nil
	}
	return &model.Customer{
		FirstName: customer.FirstName,
		ID:        customer.ID.String(),
		LastName:  customer.LastName,
	}
}
