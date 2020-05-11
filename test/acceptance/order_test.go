package acceptance_test

import "time"

func (suite *AcceptanceSuite) TestCreateOrder() {
	createdOrder, err := suite.CreateOrder()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdOrder)

	suite.Assert().Eventually(func() bool {
		orders, err := suite.Gateway.Orders()

		suite.Assert().Nil(err)

		for _, order := range orders {
			if order.ID == createdOrder.ID {

				suite.Assert().Equal(createdOrder, order)

				return true
			}
		}
		return false
	}, 30*time.Second, time.Second)
}
