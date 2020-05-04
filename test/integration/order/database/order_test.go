package database_test

import (
	"context"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/google/uuid"
)

func (suite *DatabaseSuite) TestCreateOrder() {
	createdOrder := order.Order{
		ID: uuid.New(),
		LineItems: []*order.LineItem{
			{
				ProductID: uuid.New(),
				Quantity:  1,
			},
			{
				ProductID: uuid.New(),
				Quantity:  2,
			},
			{
				ProductID: uuid.New(),
				Quantity:  3,
			},
		},
	}

	err := suite.OrderRepository.CreateOrder(context.Background(), &createdOrder)

	suite.Assert().Nil(err)

	orders, err := suite.OrderRepository.GetOrders(context.Background())

	suite.Assert().Nil(err)

	includesCreatedOrder := false
	for _, order := range orders {
		if order.ID == createdOrder.ID {
			// Time values are not equal here due to the loss of nanosecond precision.
			// - Times in BSON are represented as UTC milliseconds since the Unix epoch.
			// - Time values in Go have nanosecond precision.
			createdAtDiff := createdOrder.CreatedAt.Sub(order.CreatedAt).Milliseconds()
			suite.Assert().Less(createdAtDiff, int64(1))

			suite.Assert().Equal(createdOrder.LineItems, order.LineItems)
			includesCreatedOrder = true
			break
		}
	}

	suite.Assert().True(includesCreatedOrder)
}
