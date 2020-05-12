package database_test

import (
	"context"
	"testing"

	"github.com/davidchristie/cloud/pkg/product"
	productDatabase "github.com/davidchristie/cloud/pkg/product/database"
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
	createdProduct := &product.Product{
		Description: fake.Sentences() + "+" + uuid.New().String(),
		ID:          uuid.New(),
		Name:        fake.ProductName() + "+" + uuid.New().String(),
	}

	err := suite.ProductRepository.CreateProduct(context.Background(), createdProduct)

	suite.Assert().Nil(err)

	product, err := suite.ProductRepository.FindProduct(context.Background(), createdProduct.ID)

	suite.Assert().Nil(err)
	suite.Assert().Equal(createdProduct, product)

	products, err := suite.ProductRepository.FindProducts(context.Background(), []uuid.UUID{createdProduct.ID})

	suite.Assert().Nil(err)
	suite.Assert().Len(products, 1)
	suite.Assert().Equal(products[createdProduct.ID], createdProduct)
}
