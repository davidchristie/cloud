package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/machinebox/graphql"
)

type Client interface {
	CreateCustomer(*CreateCustomerInput) (*Customer, error)
	CreateOrder(*CreateOrderInput) (*Order, error)
	CreateProduct(*CreateProductInput) (*Product, error)
	Customers(query *string) ([]*Customer, error)
	Orders() ([]*Order, error)
	Products(query *string) ([]*Product, error)
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

type Customer struct {
	FirstName string    `json:"firstName"`
	ID        uuid.UUID `json:"id"`
	LastName  string    `json:"lastName"`
}

type LineItemInput struct {
	ProductID string `json:"productID"`
	Quantity  int    `json:"quantity"`
}

type Order struct {
	ID uuid.UUID `json:"id"`
}

type Product struct {
	Description string    `json:"description"`
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
}

// Responses

type createCustomerResponse struct {
	CreateCustomer *Customer
}

type createOrderResponse struct {
	CreateOrder *Order
}

type createProductResponse struct {
	CreateProduct *Product
}

type customersResponse struct {
	Customers []*Customer
}

type ordersResponse struct {
	Orders []*Order
}

type productsResponse struct {
	Products []*Product
}

func NewClient() Client {
	spec := clientSpecification{}
	envconfig.MustProcess("GATEWAY", &spec)
	return &client{
		GraphQL: graphql.NewClient(spec.URL + "/query"),
	}
}

func (c *client) CreateCustomer(input *CreateCustomerInput) (*Customer, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		mutation CreateCustomer($input: CreateCustomerInput!) {
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

func (c *client) CreateOrder(input *CreateOrderInput) (*Order, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		mutation CreateOrder($input: CreateOrderInput!) {
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

func (c *client) CreateProduct(input *CreateProductInput) (*Product, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		mutation CreateProduct($input: CreateProductInput!) {
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

func (c *client) Customers(query *string) ([]*Customer, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		query Customers($query: String) {
			customers(query: $query) {
				firstName
				id
				lastName
			}
		}
	`)
	request.Var("query", query)
	response := customersResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.Customers, nil
}

func (c *client) Orders() ([]*Order, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		query Orders{
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

func (c *client) Products(query *string) ([]*Product, error) {
	ctx := context.Background()
	request := graphql.NewRequest(`
		query Products($query: String) {
			products(query: $query) {
				description
				id
				name
			}
		}
	`)
	request.Var("query", query)
	response := productsResponse{}
	if err := c.GraphQL.Run(ctx, request, &response); err != nil {
		return nil, err
	}
	return response.Products, nil
}
