package acceptance_test

func (suite *AcceptanceSuite) TestCreateCustomer() {
	createdCustomer, err := suite.CreateCustomer()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdCustomer)

	suite.AssertCustomerAppearsInSearchResults(createdCustomer)
}
