package database_test

import (
	"context"
	"testing"

	customerDatabase "github.com/davidchristie/cloud/pkg/customer-database"
	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
	CustomerRepository customerDatabase.CustomerRepository
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

func (suite *DatabaseSuite) SetupTest() {
	database := customerDatabase.Connect()
	suite.CustomerRepository = customerDatabase.NewCustomerRepository(database)
}

func (suite *DatabaseSuite) TestCreateCustomer() {
	createdCustomer := &entity.Customer{
		FirstName: fake.FirstName() + "+" + uuid.New().String(),
		ID:        uuid.New(),
		LastName:  fake.LastName() + "+" + uuid.New().String(),
	}

	err := suite.CustomerRepository.CreateCustomer(context.Background(), createdCustomer)

	suite.Assert().Nil(err)

	customers, err := suite.CustomerRepository.GetCustomers(context.Background())

	suite.Assert().Nil(err)
	suite.Assert().Contains(customers, createdCustomer)
}
