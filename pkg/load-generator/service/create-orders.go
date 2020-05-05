package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/write/api"
	"github.com/google/uuid"
)

func (s *service) CreateFakeOrder() (*order.Order, error) {
	customer, err := s.CreateFakeCustomer()
	if err != nil {
		return nil, err
	}
	lineItemCount := rand.Intn(10) + 1
	lineItems := make([]*order.LineItem, lineItemCount)
	for i, _ := range lineItems {
		product, err := s.CreateFakeProduct()
		if err != nil {
			return nil, err
		}
		lineItems[i] = &order.LineItem{
			ProductID: product.ID,
			Quantity:  rand.Intn(10) + 1,
		}
	}
	return s.OrderWriteAPI.CreateOrder(&api.CreateOrderInput{
		CorrelationID: uuid.New(),
		CustomerID:    customer.ID,
		LineItems:     lineItems,
	})
}

func (s *service) GenerateFakeOrders() {
	for i := 0; ; i++ {
		fmt.Println("create fake order")
		_, err := s.CreateFakeOrder()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
