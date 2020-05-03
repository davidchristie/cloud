package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	customerReadAPIClient "github.com/davidchristie/cloud/pkg/customer-read-api/client"
	customerWriteAPIClient "github.com/davidchristie/cloud/pkg/customer-write-api/client"
	"github.com/davidchristie/cloud/pkg/gateway/graph/generated"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
	productReadAPIClient "github.com/davidchristie/cloud/pkg/product-read-api/client"
	productWriteAPIClient "github.com/davidchristie/cloud/pkg/product-write-api/client"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	correlationID := uuid.New()
	customer, err := customerWriteAPIClient.NewClient().CreateCustomer(input.FirstName, input.LastName, correlationID)
	if err != nil {
		return nil, err
	}
	modelCustomer := model.Customer{
		FirstName: customer.FirstName,
		ID:        customer.ID.String(),
		LastName:  customer.LastName,
	}
	return &modelCustomer, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	correlationID := uuid.New()
	product, err := productWriteAPIClient.NewClient().CreateProduct(input.Name, input.Description, correlationID)
	if err != nil {
		return nil, err
	}
	modelProduct := model.Product{
		Description: product.Description,
		ID:          product.ID.String(),
		Name:        product.Name,
	}
	return &modelProduct, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	customers, err := customerReadAPIClient.NewClient().Customers()
	if err != nil {
		return nil, err
	}
	modelCustomers := make([]*model.Customer, len(customers))
	for i, customer := range customers {
		modelCustomers[i] = &model.Customer{
			FirstName: customer.FirstName,
			ID:        customer.ID.String(),
			LastName:  customer.LastName,
		}
	}
	return modelCustomers, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := productReadAPIClient.NewClient().Products()
	if err != nil {
		return nil, err
	}
	modelProducts := make([]*model.Product, len(products))
	for i, product := range products {
		modelProducts[i] = &model.Product{
			Description: product.Description,
			ID:          product.ID.String(),
			Name:        product.Name,
		}
	}
	return modelProducts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
