package database_test

import (
	"context"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

func (suite *DatabaseSuite) TestCreateProduct() {
	createdProduct := entity.Product{
		Description: fake.Sentences() + "+" + uuid.New().String(),
		ID:          uuid.New(),
		Name:        fake.ProductName() + "+" + uuid.New().String(),
	}

	err := suite.ProductRepository.CreateProduct(context.Background(), &createdProduct)

	suite.Assert().Nil(err)

	hasProduct, err := suite.ProductRepository.HasProduct(context.Background(), createdProduct.ID)

	suite.Assert().Nil(err)
	suite.Assert().True(hasProduct)
}
