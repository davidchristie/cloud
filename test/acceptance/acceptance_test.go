package acceptance_test

import (
	"math/rand"
	"testing"

	"github.com/davidchristie/cloud/pkg/gateway"
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/suite"
)

type AcceptanceSuite struct {
	suite.Suite
	Gateway gateway.Client
}

func (suite *AcceptanceSuite) SetupSuite() {
	suite.Gateway = gateway.NewClient()
	kafka.WaitUntilHealthy()
}

func (suite *AcceptanceSuite) CreateCustomer() (*gateway.Customer, error) {
	return suite.Gateway.CreateCustomer(&gateway.CreateCustomerInput{
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
	})
}

func (suite *AcceptanceSuite) CreateOrder() (*gateway.Order, error) {
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

func (suite *AcceptanceSuite) CreateProduct() (*gateway.Product, error) {
	return suite.Gateway.CreateProduct(&gateway.CreateProductInput{
		Description: fake.Sentences(),
		Name:        fake.ProductName(),
	})
}

func TestAcceptanceSuite(t *testing.T) {
	suite.Run(t, new(AcceptanceSuite))
}
