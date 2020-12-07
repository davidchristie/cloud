package system_test

import (
	"time"

	"github.com/davidchristie/cloud/pkg/order/read/api"
)

func (suite *SystemSuite) TestCreateOrder() {
	createdOrder, err := suite.CreateOrder()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdOrder)

	suite.T().Log("wait for created order to appear in order list")
	suite.Assert().Eventually(func() bool {
		customerID := createdOrder.CustomerID.String()
		orders, err := suite.OrderReadAPI.Orders(api.OrdersInput{
			CustomerID: &customerID,
		})

		suite.Assert().Nil(err)

		for _, order := range orders {
			if order.ID == createdOrder.ID {
				suite.Assert().Equal(createdOrder, order)

				return true
			}
		}
		return false
	}, 10*time.Second, time.Second)
}
