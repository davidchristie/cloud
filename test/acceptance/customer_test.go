package system_test

import "time"

func (suite *AcceptanceSuite) TestCreateCustomer() {
	createdCustomer, err := suite.CreateCustomer()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdCustomer)

	suite.Assert().Eventually(func() bool {
		customers, err := suite.Gateway.Customers()

		suite.Assert().Nil(err)

		for _, customer := range customers {
			if customer.ID == createdCustomer.ID {

				suite.Assert().Equal(createdCustomer, customer)

				return true
			}
		}
		return false
	}, 10*time.Second, time.Second)
}
