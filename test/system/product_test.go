package system_test

import (
	"time"

	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
)

func (suite *SystemSuite) TestCreateProduct() {
	createdProduct, err := suite.CreateProduct()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdProduct)

	suite.T().Log("wait for created product to be queryable")
	suite.Eventually(func() bool {
		product, err := suite.ProductReadAPI.Product(createdProduct.ID)

		if err != productReadAPI.ErrProductNotFound {
			suite.Assert().Equal(createdProduct, product)

			return true
		}
		return false
	}, 10*time.Second, time.Second)

	suite.T().Log("wait for created product to appear in product list")
	suite.Eventually(func() bool {
		products, err := suite.ProductReadAPI.Products()

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
