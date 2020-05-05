package system_test

import (
	productReadAPIClient "github.com/davidchristie/cloud/pkg/product-read-api/client"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

func (suite *SystemSuite) TestCreateProduct() {
	name := fake.ProductName() + "+" + uuid.New().String()
	description := fake.Sentences() + "+" + uuid.New().String()

	createdProduct, err := suite.ProductWriteAPI.CreateProduct(name, description, uuid.New())

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdProduct)
	suite.Assert().Equal(name, createdProduct.Name)
	suite.Assert().Equal(description, createdProduct.Description)

	suite.T().Log("wait for created product to be queryable")
	suite.WaitFor(func() bool {
		product, err := suite.ProductReadAPI.Product(createdProduct.ID)
		if err != productReadAPIClient.ErrProductNotFound {
			suite.Assert().Equal(createdProduct, product)
			return true
		}
		return false
	})

	suite.T().Log("wait for created product to appear in product list")
	suite.WaitFor(func() bool {
		products, err := suite.ProductReadAPI.Products()
		suite.Assert().Nil(err)
		for _, product := range products {
			if product.ID == createdProduct.ID {
				suite.Assert().Equal(createdProduct, product)
				return true
			}
		}
		return false
	})
}
