package acceptance_test

func (suite *AcceptanceSuite) TestCreateOrder() {
	createdOrder, err := suite.CreateOrder()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdOrder)

	suite.AssertOrderAppearsInSearchResults(createdOrder)
}
