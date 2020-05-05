package system_test

import (
	"math/rand"
	"testing"
	"time"

	customerReadAPIClient "github.com/davidchristie/cloud/pkg/customer-read-api/client"
	customerWriteAPIClient "github.com/davidchristie/cloud/pkg/customer-write-api/client"
	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/order"
	orderReadAPI "github.com/davidchristie/cloud/pkg/order/read/api"
	"github.com/davidchristie/cloud/pkg/order/write/api"
	orderWriteAPI "github.com/davidchristie/cloud/pkg/order/write/api"
	productReadAPIClient "github.com/davidchristie/cloud/pkg/product-read-api/client"
	productWriteAPIClient "github.com/davidchristie/cloud/pkg/product-write-api/client"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type SystemSuite struct {
	suite.Suite
	CustomerReadAPI  customerReadAPIClient.CustomerReadAPIClient
	CustomerWriteAPI customerWriteAPIClient.CustomerWriteAPIClient
	OrderReadAPI     orderReadAPI.OrderReadAPIClient
	OrderWriteAPI    orderWriteAPI.OrderWriteAPIClient
	ProductReadAPI   productReadAPIClient.ProductReadAPIClient
	ProductWriteAPI  productWriteAPIClient.ProductWriteAPIClient
}

func (suite *SystemSuite) SetupTest() {
	suite.CustomerReadAPI = customerReadAPIClient.NewClient()
	suite.CustomerWriteAPI = customerWriteAPIClient.NewClient()
	suite.OrderReadAPI = orderReadAPI.NewClient()
	suite.OrderWriteAPI = orderWriteAPI.NewClient()
	suite.ProductReadAPI = productReadAPIClient.NewClient()
	suite.ProductWriteAPI = productWriteAPIClient.NewClient()
}

func (suite *SystemSuite) CreateCustomer() (*entity.Customer, error) {
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

func (suite *SystemSuite) CreateProduct() (*entity.Product, error) {
	return suite.ProductWriteAPI.CreateProduct(fake.ProductName(), fake.Sentences(), uuid.New())
}

func (suite *SystemSuite) WaitFor(done func() bool) {
	for {
		if done() {
			break
		}
		suite.T().Log("...")
		time.Sleep(1 * time.Second)
	}
}

func TestSystemSuite(t *testing.T) {
	suite.Run(t, new(SystemSuite))
}
