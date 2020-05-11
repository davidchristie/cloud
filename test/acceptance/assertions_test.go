package acceptance_test

import (
	"time"

	"github.com/davidchristie/cloud/pkg/product"
)

func (suite *AcceptanceSuite) AssertProductAppearsInSearchResults(product *product.Product) {
	suite.Assert().Eventually(func() bool {
		results, err := suite.Gateway.Products(&product.Name)
		suite.Assert().Nil(err)
		for _, result := range results {
			if result.ID == product.ID {
				suite.Assert().Equal(result, product)
				return true
			}
		}
		return false
	}, 10*time.Second, time.Second)
}
