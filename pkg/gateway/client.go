package gateway

import (
	"context"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/order"
	"github.com/kelseyhightower/envconfig"
	"github.com/machinebox/graphql"
)

type Client interface {
	CreateCustomer(*CreateCustomerInput) (*entity.Customer, error)
	CreateOrder(*CreateOrderInput) (*order.Order, error)
	CreateProduct(*CreateProductInput) (*entity.Product, error)
	Customers() ([]*entity.Customer, error)
	Orders() ([]*order.Order, error)
	Products() ([]*entity.Product, error)
}

type client struct {
	GraphQL *graphql.Client
}

type clientSpecification struct {
	URL string `required:"true"`
}

// Inputs

type CreateCustomerInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateOrderInput struct {
	CustomerID string           `json:"customerID"`
	LineItems  []*LineItemInput `json:"lineItems"`
}

type CreateProductInput struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type LineItemInput struct {
	ProductID string `json:"productID"`
	Quantity  int    `json:"quantity"`
}

// Responses

type createCustomerResponse struct {
	CreateCustomer *entity.Customer
}

type createOrderResponse struct {
	CreateOrder *order.Order
}

type createProductResponse struct {
	CreateProduct *entity.Product
}

type customersResponse struct {
	Customers []*entity.Customer
}

type ordersResponse struct {
	Orders []*order.Order
}

type productsResponse struct {
	Products []*entity.Product
}

func NewClient() Client {
	spec := clientSpecification{}
	envconfig.MustProcess("GATEWAY", &spec)
	return &client{
		GraphQL: graphql.NewClient(spec.URL + "/query"),
	}
}

func (c *client) CreateCustomer(input *CreateCustomerInput) (*entity.Customer, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		mutation ($input: CreateCustomerInput!) {
			createCustomer(input: $input) {
				firstName
				id
				lastName
			}
		}
	`)
	request.Var("input", input)
	response := createCustomerResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.CreateCustomer, nil
}

func (c *client) CreateOrder(input *CreateOrderInput) (*order.Order, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		mutation ($input: CreateOrderInput!) {
			createOrder(input: $input) {
				id
			}
		}
	`)
	request.Var("input", input)
	response := createOrderResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.CreateOrder, nil
}

func (c *client) CreateProduct(input *CreateProductInput) (*entity.Product, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		mutation ($input: CreateProductInput!) {
			createProduct(input: $input) {
				description
				id
				name
			}
		}
	`)
	request.Var("input", input)
	response := createProductResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.CreateProduct, nil
}

func (c *client) Customers() ([]*entity.Customer, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		query {
			customers {
				firstName
				id
				lastName
			}
		}
	`)
	response := customersResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.Customers, nil
}

func (c *client) Orders() ([]*order.Order, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		query {
			orders {
				customer {
					firstName
					id
					lastName
				}
				id
				lineItems {
					product {
						description
						id
						name
					}
					quantity
				}
			}
		}
	`)
	response := ordersResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.Orders, nil
}

func (c *client) Products() ([]*entity.Product, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		query {
			products {
				description
				id
				name
			}
		}
	`)
	response := productsResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.Products, nil
}
