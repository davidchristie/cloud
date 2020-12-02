package database_test

import (
	"context"
	"time"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/google/uuid"
)

func (suite *DatabaseSuite) TestCreateOrder() {
	createdOrder := &order.Order{
		CustomerID: uuid.New(),
		CreatedAt:  time.Now().In(time.UTC).Truncate(time.Second),
		ID:         uuid.New(),
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

	err := suite.OrderRepository.CreateOrder(context.Background(), createdOrder)

	suite.Assert().Nil(err)

	orders, err := suite.OrderRepository.FindOrders(context.Background())

	suite.Assert().Nil(err)
	suite.Assert().Contains(orders, createdOrder)
}
