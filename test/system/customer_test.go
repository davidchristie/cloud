package system_test

import customerReadAPIClient "github.com/davidchristie/cloud/pkg/customer-read-api/client"

func (suite *SystemSuite) TestCreateCustomer() {
	createdCustomer, err := suite.CreateCustomer()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdCustomer)

	suite.T().Log("wait for created customer to be queryable")
	suite.WaitFor(func() bool {
		customer, err := suite.CustomerReadAPI.Customer(createdCustomer.ID)
		if err != customerReadAPIClient.ErrCustomerNotFound {
			suite.Assert().Equal(createdCustomer, customer)
			return true
		}
		return false
	})

	suite.T().Log("wait for created customer to appear in customer list")
	suite.WaitFor(func() bool {
		customers, err := suite.CustomerReadAPI.Customers()
		suite.Assert().Nil(err)
		for _, customer := range customers {
			if customer.ID == createdCustomer.ID {
				suite.Assert().Equal(createdCustomer, customer)
				return true
			}
		}
		return false
	})
}
