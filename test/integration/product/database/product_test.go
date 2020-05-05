package database_test

import (
	"context"
	"testing"

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
	createdProduct := &entity.Product{
		Description: fake.Sentences() + "+" + uuid.New().String(),
		ID:          uuid.New(),
		Name:        fake.ProductName() + "+" + uuid.New().String(),
	}

	err := suite.ProductRepository.CreateProduct(context.Background(), createdProduct)

	suite.Assert().Nil(err)

	products, err := suite.ProductRepository.GetProducts(context.Background())

	suite.Assert().Nil(err)
	suite.Assert().Contains(products, createdProduct)
}
