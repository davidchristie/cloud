package acceptance_test

import (
	"time"

	"github.com/davidchristie/cloud/pkg/gateway"
)

func (suite *AcceptanceSuite) AssertCustomerAppearsInSearchResults(customer *gateway.Customer) {
	suite.Assert().Eventually(func() bool {
		query := customer.FirstName + " " + customer.LastName
		results, err := suite.Gateway.Customers(&query)
		suite.Assert().Nil(err)
		for _, result := range results {
			if result.ID == customer.ID {
				suite.Assert().Equal(result.FirstName, customer.FirstName)
				suite.Assert().Equal(result.LastName, customer.LastName)
				return true
			}
		}
		return false
	}, 10*time.Second, time.Second)
}

func (suite *AcceptanceSuite) AssertOrderAppearsInSearchResults(order *gateway.Order) {
	suite.Assert().Eventually(func() bool {
		results, err := suite.Gateway.Orders()
		suite.Assert().Nil(err)
		for _, result := range results {
			if result.ID == order.ID {
				return true
			}
		}
		return false
	}, 10*time.Second, time.Second)
}

func (suite *AcceptanceSuite) AssertProductAppearsInSearchResults(product *gateway.Product) {
	suite.Assert().Eventually(func() bool {
		query := product.Name
		results, err := suite.Gateway.Products(&query)
		suite.Assert().Nil(err)
		for _, result := range results {
			if result.ID == product.ID {
				suite.Assert().Equal(result.Description, product.Description)
				suite.Assert().Equal(result.Name, product.Name)
				return true
			}
		}
		return false
	}, 10*time.Second, time.Second)
}
