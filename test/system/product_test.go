package acceptance_test

import (
	"testing"
	"time"

	productReadAPIClient "github.com/davidchristie/cloud/pkg/product-read-api/client"
	productWriteAPIClient "github.com/davidchristie/cloud/pkg/product-write-api/client"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type ProductSuite struct {
	suite.Suite
	ProductReadAPI  productReadAPIClient.ProductReadAPIClient
	ProductWriteAPI productWriteAPIClient.ProductWriteAPIClient
}

func TestProductSuite(t *testing.T) {
	suite.Run(t, new(ProductSuite))
}

func (suite *ProductSuite) SetupTest() {
	suite.ProductReadAPI = productReadAPIClient.NewClient()
	suite.ProductWriteAPI = productWriteAPIClient.NewClient()
}

func (suite *ProductSuite) TestCreateProduct() {
	name := fake.ProductName() + "+" + uuid.New().String()
	description := fake.Sentences() + "+" + uuid.New().String()

	createdProduct, err := suite.ProductWriteAPI.CreateProduct(name, description)

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdProduct)
	suite.Assert().Equal(name, createdProduct.Name)
	suite.Assert().Equal(description, createdProduct.Description)

	suite.T().Log("wait for created product to appear in product list")
Retries:
	for {
		products, err := suite.ProductReadAPI.Products()
		suite.Assert().Nil(err)
		for _, product := range products {
			if product.ID == createdProduct.ID {
				break Retries
			}
		}
		suite.T().Log("...")
		time.Sleep(1 * time.Second)
	}
}
