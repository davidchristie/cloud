package acceptance_test

import (
	"testing"
	"time"

	customerReadAPIClient "github.com/davidchristie/cloud/pkg/customer-read-api/client"
	customerWriteAPIClient "github.com/davidchristie/cloud/pkg/customer-write-api/client"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type CustomerSuite struct {
	suite.Suite
	CustomerReadAPI  customerReadAPIClient.CustomerReadAPIClient
	CustomerWriteAPI customerWriteAPIClient.CustomerWriteAPIClient
}

func TestCustomerSuite(t *testing.T) {
	suite.Run(t, new(CustomerSuite))
}

func (suite *CustomerSuite) SetupTest() {
	suite.CustomerReadAPI = customerReadAPIClient.NewClient()
	suite.CustomerWriteAPI = customerWriteAPIClient.NewClient()
}

func (suite *CustomerSuite) TestCreateCustomer() {
	firstName := fake.FirstName() + "+" + uuid.New().String()
	lastName := fake.LastName() + "+" + uuid.New().String()

	createdCustomer, err := suite.CustomerWriteAPI.CreateCustomer(firstName, lastName)

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdCustomer)
	suite.Assert().Equal(firstName, createdCustomer.FirstName)
	suite.Assert().Equal(lastName, createdCustomer.LastName)

	suite.T().Log("wait for created customer to appear in customer list")
Retries:
	for {
		customers, err := suite.CustomerReadAPI.Customers()
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
