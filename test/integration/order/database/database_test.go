package database_test

import (
	"testing"

	"github.com/davidchristie/cloud/pkg/order/database"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
	CustomerRepository database.CustomerRepository
	OrderRepository    database.OrderRepository
	ProductRepository  database.ProductRepository
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

func (suite *DatabaseSuite) SetupTest() {
	db := database.Connect()
	suite.CustomerRepository = database.NewCustomerRepository(db)
	suite.OrderRepository = database.NewOrderRepository(db)
	suite.ProductRepository = database.NewProductRepository(db)
}
