package system_test

import (
	"math/rand"
	"testing"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/gateway"
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/davidchristie/cloud/pkg/order"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type AcceptanceSuite struct {
	suite.Suite
	Gateway gateway.Client
}

func (suite *AcceptanceSuite) SetupTest() {
	suite.Gateway = gateway.NewClient()
	kafka.WaitUntilHealthy()
}

func (suite *AcceptanceSuite) CreateCustomer() (*entity.Customer, error) {
	return suite.Gateway.CreateCustomer(&gateway.CreateCustomerInput{
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
	})
}

func (suite *AcceptanceSuite) CreateOrder() (*order.Order, error) {
	customer, err := suite.CreateCustomer()
	if err != nil {
		return nil, err
	}
	lineItemCount := rand.Intn(10) + 1
	lineItems := make([]*gateway.LineItemInput, lineItemCount)
	for i, _ := range lineItems {
		product, err := suite.CreateProduct()
		if err != nil {
			return nil, err
		}
		lineItems[i] = &gateway.LineItemInput{
			ProductID: product.ID.String(),
			Quantity:  rand.Intn(10) + 1,
		}
	}
	return suite.Gateway.CreateOrder(&gateway.CreateOrderInput{
		CustomerID: customer.ID.String(),
		LineItems:  lineItems,
	})
}

func (suite *AcceptanceSuite) CreateProduct() (*entity.Product, error) {
	return suite.Gateway.CreateProduct(&gateway.CreateProductInput{
		Description: fake.Sentences(),
		Name:        fake.ProductName(),
	})
}

func TestAcceptanceSuite(t *testing.T) {
	suite.Run(t, new(AcceptanceSuite))
}
