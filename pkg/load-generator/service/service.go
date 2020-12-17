package service

import (
	"errors"

	customerWriteAPI "github.com/davidchristie/cloud/pkg/customer/write/api"
	orderWriteAPI "github.com/davidchristie/cloud/pkg/order/write/api"
	productWriteAPI "github.com/davidchristie/cloud/pkg/product/write/api"
)

type service struct {
	CustomerWriteAPI customerWriteAPI.Client
	OrderWriteAPI    orderWriteAPI.OrderWriteAPIClient
	ProductWriteAPI  productWriteAPI.Client
}

func StartService() error {
	s := service{
		CustomerWriteAPI: customerWriteAPI.NewClient(),
		OrderWriteAPI:    orderWriteAPI.NewClient(),
		ProductWriteAPI:  productWriteAPI.NewClient(),
	}
	go s.GenerateFakeCustomers()
	go s.GenerateFakeProducts()
	s.GenerateFakeOrders()
	return errors.New("error in load-generator service")
}
