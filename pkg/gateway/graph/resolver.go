package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	customerReadAPI "github.com/davidchristie/cloud/pkg/customer-read-api/client"
	customerWriteAPI "github.com/davidchristie/cloud/pkg/customer-write-api/client"
	orderReadAPI "github.com/davidchristie/cloud/pkg/order/read/api"
	orderWriteAPI "github.com/davidchristie/cloud/pkg/order/write/api"
	productReadAPI "github.com/davidchristie/cloud/pkg/product-read-api/client"
	productWriteAPI "github.com/davidchristie/cloud/pkg/product-write-api/client"
)

type Resolver struct {
	CustomerReadAPI  customerReadAPI.CustomerReadAPIClient
	CustomerWriteAPI customerWriteAPI.CustomerWriteAPIClient
	OrderReadAPI     orderReadAPI.OrderReadAPIClient
	OrderWriteAPI    orderWriteAPI.OrderWriteAPIClient
	ProductReadAPI   productReadAPI.ProductReadAPIClient
	ProductWriteAPI  productWriteAPI.ProductWriteAPIClient
}

func NewResolver() *Resolver {
	return &Resolver{
		CustomerReadAPI:  customerReadAPI.NewClient(),
		CustomerWriteAPI: customerWriteAPI.NewClient(),
		OrderReadAPI:     orderReadAPI.NewClient(),
		OrderWriteAPI:    orderWriteAPI.NewClient(),
		ProductReadAPI:   productReadAPI.NewClient(),
		ProductWriteAPI:  productWriteAPI.NewClient(),
	}
}
