package system_test

import (
	"math/rand"
	"testing"

	"github.com/davidchristie/cloud/pkg/customer"
	customerReadAPI "github.com/davidchristie/cloud/pkg/customer/read/api"
	customerWriteAPI "github.com/davidchristie/cloud/pkg/customer/write/api"
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/davidchristie/cloud/pkg/order"
	orderReadAPI "github.com/davidchristie/cloud/pkg/order/read/api"
	"github.com/davidchristie/cloud/pkg/order/write/api"
	orderWriteAPI "github.com/davidchristie/cloud/pkg/order/write/api"
	"github.com/davidchristie/cloud/pkg/product"
	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
	productWriteAPI "github.com/davidchristie/cloud/pkg/product/write/api"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type SystemSuite struct {
	suite.Suite
	CustomerReadAPI  customerReadAPI.Client
	CustomerWriteAPI customerWriteAPI.Client
	OrderReadAPI     orderReadAPI.OrderReadAPIClient
	OrderWriteAPI    orderWriteAPI.OrderWriteAPIClient
	ProductReadAPI   productReadAPI.Client
	ProductWriteAPI  productWriteAPI.Client
}

func (suite *SystemSuite) SetupSuite() {
	suite.CustomerReadAPI = customerReadAPI.NewClient()
	suite.CustomerWriteAPI = customerWriteAPI.NewClient()
	suite.OrderReadAPI = orderReadAPI.NewClient()
	suite.OrderWriteAPI = orderWriteAPI.NewClient()
	suite.ProductReadAPI = productReadAPI.NewClient()
	suite.ProductWriteAPI = productWriteAPI.NewClient()
	kafka.WaitUntilHealthy()
}

func (suite *SystemSuite) CreateCustomer() (*customer.Customer, error) {
	return suite.CustomerWriteAPI.CreateCustomer(fake.FirstName(), fake.LastName(), uuid.New())
}

func (suite *SystemSuite) CreateOrder() (*order.Order, error) {
	customer, err := suite.CreateCustomer()
	if err != nil {
		return nil, err
	}
	lineItemCount := rand.Intn(10) + 1
	lineItems := make([]*order.LineItem, lineItemCount)
	for i, _ := range lineItems {
		product, err := suite.CreateProduct()
		if err != nil {
			return nil, err
		}
		lineItems[i] = &order.LineItem{
			ProductID: product.ID,
			Quantity:  rand.Intn(10) + 1,
		}
	}
	return suite.OrderWriteAPI.CreateOrder(&api.CreateOrderInput{
		CorrelationID: uuid.New(),
		CustomerID:    customer.ID,
		LineItems:     lineItems,
	})
}

func (suite *SystemSuite) CreateProduct() (*product.Product, error) {
	return suite.ProductWriteAPI.CreateProduct(fake.ProductName(), fake.Sentences(), uuid.New())
}

func TestSystemSuite(t *testing.T) {
	suite.Run(t, new(SystemSuite))
}
