package convert

import (
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
	"github.com/davidchristie/cloud/pkg/order"
)

func Order(order *order.Order) *model.Order {
	if order == nil {
		return nil
	}
	lineItemModels := make([]*model.LineItem, len(order.LineItems))
	for i, lineItem := range order.LineItems {
		lineItemModels[i] = &model.LineItem{
			ProductID: lineItem.ProductID,
			Quantity:  lineItem.Quantity,
		}
	}
	modelOrder := model.Order{
		CustomerID: order.CustomerID,
		CreatedAt:  order.CreatedAt.String(),
		ID:         order.ID.String(),
		LineItems:  lineItemModels,
	}
	return &modelOrder
}
