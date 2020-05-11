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

func (r *lineItemResolver) Product(ctx context.Context, obj *model.LineItem) (*model.Product, error) {
	product, err := r.ProductReadAPI.Product(obj.ProductID)
	switch err {
	case nil:
		return utility.ConvertProductToModel(product), nil

	case productReadAPI.ErrProductNotFound:
		return nil, nil

	default:
		return nil, err
	}
}

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
	return utility.ConvertOrderToModel(createdOrder), nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	correlationID := uuid.New()
	product, err := r.ProductWriteAPI.CreateProduct(input.Name, input.Description, correlationID)
	if err != nil {
		return nil, err
	}
	return utility.ConvertProductToModel(product), nil
}

func (r *orderResolver) Customer(ctx context.Context, obj *model.Order) (*model.Customer, error) {
	customer, err := r.CustomerReadAPI.Customer(obj.CustomerID)
	switch err {
	case nil:
		return utility.ConvertCustomerToModel(customer), nil

	case customerReadAPI.ErrCustomerNotFound:
		return nil, nil

	default:
		return nil, err
	}
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
		modelOrders[i] = utility.ConvertOrderToModel(order)
	}
	return modelOrders, nil
}

func (r *queryResolver) Products(ctx context.Context, query *string) ([]*model.Product, error) {
	q := ""
	if query != nil {
		q = *query
	}
	results, err := r.SearchAPI.Products(q)
	if err != nil {
		return nil, err
	}
	modelProducts := make([]*model.Product, len(results))
	for i, productID := range results {
		product, err := r.ProductReadAPI.Product(productID)
		if err != nil {
			return nil, err
		}
		modelProducts[i] = utility.ConvertProductToModel(product)
	}
	return modelProducts, nil
}

// LineItem returns generated.LineItemResolver implementation.
func (r *Resolver) LineItem() generated.LineItemResolver { return &lineItemResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type lineItemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
