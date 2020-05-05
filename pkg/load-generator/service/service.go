package service

import (
	customerWriteAPI "github.com/davidchristie/cloud/pkg/customer-write-api/client"
	orderWriteAPI "github.com/davidchristie/cloud/pkg/order/write/api"
	productWriteAPI "github.com/davidchristie/cloud/pkg/product-write-api/client"
)

type service struct {
	CustomerWriteAPI customerWriteAPI.CustomerWriteAPIClient
	OrderWriteAPI    orderWriteAPI.OrderWriteAPIClient
	ProductWriteAPI  productWriteAPI.ProductWriteAPIClient
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
