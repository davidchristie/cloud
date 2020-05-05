package system_test

func (suite *SystemSuite) TestCreateOrder() {
	createdOrder, err := suite.CreateOrder()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdOrder)

	suite.T().Log("wait for created order to appear in order list")
	suite.WaitFor(func() bool {
		orders, err := suite.OrderReadAPI.Orders()
		suite.Assert().Nil(err)
		for _, order := range orders {
			if order.ID == createdOrder.ID {
				suite.Assert().Equal(createdOrder, order)
				return true
			}
		}
		return false
	})
}
