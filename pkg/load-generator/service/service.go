package service

import (
	customerWriteAPI "github.com/davidchristie/cloud/pkg/customer/write/api"
	orderWriteAPI "github.com/davidchristie/cloud/pkg/order/write/api"
	productWriteAPI "github.com/davidchristie/cloud/pkg/product/write/api"
)

type service struct {
	CustomerWriteAPI customerWriteAPI.Client
	OrderWriteAPI    orderWriteAPI.OrderWriteAPIClient
	ProductWriteAPI  productWriteAPI.Client
}

func Start() {
	s := service{
		CustomerWriteAPI: customerWriteAPI.NewClient(),
		OrderWriteAPI:    orderWriteAPI.NewClient(),
		ProductWriteAPI:  productWriteAPI.NewClient(),
	}

	go s.GenerateFakeCustomers()
	go s.GenerateFakeProducts()
	s.GenerateFakeOrders()
}
