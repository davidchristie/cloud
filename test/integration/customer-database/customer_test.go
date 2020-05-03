package database_test

import (
	"context"
	"testing"
	"time"

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
	createdCustomer := entity.Customer{
		FirstName: fake.FirstName() + "+" + uuid.New().String(),
		ID:        uuid.New(),
		LastName:  fake.LastName() + "+" + uuid.New().String(),
	}

	err := suite.CustomerRepository.CreateCustomer(context.Background(), &createdCustomer)

	suite.Assert().Nil(err)

	suite.T().Log("wait for created customer to appear in customer list")
Retries:
	for {
		customers, err := suite.CustomerRepository.GetCustomers(context.Background())
		suite.Assert().Nil(err)
		for _, customer := range customers {
			if customer.ID == createdCustomer.ID {
				break Retries
			}
		}
		suite.T().Log("...")
		time.Sleep(1 * time.Second)
	}
}
