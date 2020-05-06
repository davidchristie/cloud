package system_test

import "time"

func (suite *AcceptanceSuite) TestCreateProduct() {
	createdProduct, err := suite.CreateProduct()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdProduct)

	suite.Assert().Eventually(func() bool {
		products, err := suite.Gateway.Products()

		suite.Assert().Nil(err)

		for _, product := range products {
			if product.ID == createdProduct.ID {

				suite.Assert().Equal(createdProduct, product)

				return true
			}
		}
		return false
	}, 10*time.Second, time.Second)
}
