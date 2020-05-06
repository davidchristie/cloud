package system_test

import (
	"time"

	customerReadAPIClient "github.com/davidchristie/cloud/pkg/customer-read-api/client"
)

func (suite *SystemSuite) TestCreateCustomer() {
	createdCustomer, err := suite.CreateCustomer()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdCustomer)

	suite.T().Log("wait for created customer to be queryable")
	suite.Assert().Eventually(func() bool {
		customer, err := suite.CustomerReadAPI.Customer(createdCustomer.ID)

		if err != customerReadAPIClient.ErrCustomerNotFound {
			suite.Assert().Equal(createdCustomer, customer)
			return true
		}

		return false
	}, 10*time.Second, time.Second)

	suite.T().Log("wait for created customer to appear in customer list")
	suite.Assert().Eventually(func() bool {
		customers, err := suite.CustomerReadAPI.Customers()

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
