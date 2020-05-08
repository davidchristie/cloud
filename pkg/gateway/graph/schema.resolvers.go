package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	customerReadAPI "github.com/davidchristie/cloud/pkg/customer/read/api"
	"github.com/davidchristie/cloud/pkg/gateway/graph/generated"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
	"github.com/davidchristie/cloud/pkg/gateway/graph/utility"
	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/write/api"
	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	correlationID := uuid.New()
	customer, err := r.CustomerWriteAPI.CreateCustomer(input.FirstName, input.LastName, correlationID)
	if err != nil {
		return nil, err
	}
	return utility.ConvertCustomerToModel(customer), nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*model.Order, error) {
	customerID, err := uuid.Parse(input.CustomerID)
	if err != nil {
		return nil, err
	}
	lineItems := make([]*order.LineItem, len(input.LineItems))
	for i, inputLineItem := range input.LineItems {
		productID, err := uuid.Parse(inputLineItem.ProductID)
		if err != nil {
			return nil, err
		}
		lineItems[i] = &order.LineItem{
			ProductID: productID,
			Quantity:  inputLineItem.Quantity,
		}
	}
	createdOrder, err := r.OrderWriteAPI.CreateOrder(&api.CreateOrderInput{
		CorrelationID: uuid.New(),
		CustomerID:    customerID,
		LineItems:     lineItems,
	})
	// TODO: Support the rest of the order fields.
	return &model.Order{
		ID: createdOrder.ID.String(),
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	correlationID := uuid.New()
	product, err := r.ProductWriteAPI.CreateProduct(input.Name, input.Description, correlationID)
	if err != nil {
		return nil, err
	}
	return utility.ConvertProductToModel(product), nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	customers, err := r.CustomerReadAPI.Customers()
	if err != nil {
		return nil, err
	}
	modelCustomers := make([]*model.Customer, len(customers))
	for i, customer := range customers {
		modelCustomers[i] = utility.ConvertCustomerToModel(customer)
	}
	return modelCustomers, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.OrderReadAPI.Orders()
	if err != nil {
		return nil, err
	}
	modelOrders := make([]*model.Order, len(orders))
	for i, order := range orders {
		customer, err := r.CustomerReadAPI.Customer(order.CustomerID)
		if err != nil && err != customerReadAPI.ErrCustomerNotFound {
			return nil, err
		}
		lineItemModels := make([]*model.LineItem, len(order.LineItems))
		for i, lineItem := range order.LineItems {
			// TODO: Fetch all line item products in a single request.
			product, err := r.ProductReadAPI.Product(lineItem.ProductID)
			if err != nil && err != productReadAPI.ErrProductNotFound {
				return nil, err
			}
			lineItemModels[i] = &model.LineItem{
				Product:  utility.ConvertProductToModel(product),
				Quantity: lineItem.Quantity,
			}
		}
		modelOrders[i] = &model.Order{
			Customer:  utility.ConvertCustomerToModel(customer),
			CreatedAt: order.CreatedAt.String(),
			ID:        order.ID.String(),
			LineItems: lineItemModels,
		}
	}
	return modelOrders, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := r.ProductReadAPI.Products()
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
