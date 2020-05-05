package utility

import (
	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
)

func ConvertCustomerToModel(customer *entity.Customer) *model.Customer {
	return &model.Customer{
		FirstName: customer.FirstName,
		ID:        customer.ID.String(),
		LastName:  customer.LastName,
	}
}
