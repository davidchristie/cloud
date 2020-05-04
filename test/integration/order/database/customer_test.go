package database_test

import (
	"context"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

func (suite *DatabaseSuite) TestCreateCustomer() {
	createdCustomer := entity.Customer{
		FirstName: fake.FirstName() + "+" + uuid.New().String(),
		ID:        uuid.New(),
		LastName:  fake.LastName() + "+" + uuid.New().String(),
	}

	err := suite.CustomerRepository.CreateCustomer(context.Background(), &createdCustomer)

	suite.Assert().Nil(err)

	hasCustomer, err := suite.CustomerRepository.HasCustomer(context.Background(), createdCustomer.ID)

	suite.Assert().Nil(err)
	suite.Assert().True(hasCustomer)
}
