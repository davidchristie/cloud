package system_test

import (
	"time"

	customerReadAPI "github.com/davidchristie/cloud/pkg/customer/read/api"
	"github.com/google/uuid"
)

func (suite *SystemSuite) TestCreateCustomer() {
	createdCustomer, err := suite.CreateCustomer()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdCustomer)

	suite.T().Log("wait for created customer to be queryable")
	suite.Assert().Eventually(func() bool {
		customer, err := suite.CustomerReadAPI.Customer(createdCustomer.ID)

		if err != customerReadAPI.ErrCustomerNotFound {
			suite.Assert().Nil(err)
			suite.Assert().Equal(createdCustomer, customer)

			return true
		}

		return false
	}, 10*time.Second, time.Second)

	suite.T().Log("wait for created customer to appear in customer list")
	suite.Assert().Eventually(func() bool {
		customers, err := suite.CustomerReadAPI.Customers([]uuid.UUID{createdCustomer.ID})

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
