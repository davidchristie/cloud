package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	customerReadAPI "github.com/davidchristie/cloud/pkg/customer/read/api"
	customerWriteAPI "github.com/davidchristie/cloud/pkg/customer/write/api"
	orderReadAPI "github.com/davidchristie/cloud/pkg/order/read/api"
	orderWriteAPI "github.com/davidchristie/cloud/pkg/order/write/api"
	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
	productWriteAPI "github.com/davidchristie/cloud/pkg/product/write/api"
	searchAPI "github.com/davidchristie/cloud/pkg/search/api"
)

type Resolver struct {
	CustomerReadAPI  customerReadAPI.Client
	CustomerWriteAPI customerWriteAPI.Client
	OrderReadAPI     orderReadAPI.OrderReadAPIClient
	OrderWriteAPI    orderWriteAPI.OrderWriteAPIClient
	ProductReadAPI   productReadAPI.Client
	ProductWriteAPI  productWriteAPI.Client
	SearchAPI        searchAPI.Client
}

func NewResolver() *Resolver {
	return &Resolver{
		CustomerReadAPI:  customerReadAPI.NewClient(),
		CustomerWriteAPI: customerWriteAPI.NewClient(),
		OrderReadAPI:     orderReadAPI.NewClient(),
		OrderWriteAPI:    orderWriteAPI.NewClient(),
		ProductReadAPI:   productReadAPI.NewClient(),
		ProductWriteAPI:  productWriteAPI.NewClient(),
		SearchAPI:        searchAPI.NewClient(),
	}
}
