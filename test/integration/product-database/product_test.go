package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/davidchristie/cloud/pkg/entity"
	productDatabase "github.com/davidchristie/cloud/pkg/product-database"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
	ProductRepository productDatabase.ProductRepository
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

func (suite *DatabaseSuite) SetupTest() {
	database := productDatabase.Connect()
	suite.ProductRepository = productDatabase.NewProductRepository(database)
}

func (suite *DatabaseSuite) TestCreateProduct() {
	createdProduct := entity.Product{
		Description: fake.Sentences() + "+" + uuid.New().String(),
		ID:          uuid.New(),
		Name:        fake.ProductName() + "+" + uuid.New().String(),
	}

	err := suite.ProductRepository.CreateProduct(context.Background(), &createdProduct)

	suite.Assert().Nil(err)

	suite.T().Log("wait for created product to appear in product list")
Retries:
	for {
		products, err := suite.ProductRepository.GetProducts(context.Background())
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
