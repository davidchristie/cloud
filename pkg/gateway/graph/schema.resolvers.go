package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	customerReadAPI "github.com/davidchristie/cloud/pkg/customer/read/api"
	"github.com/davidchristie/cloud/pkg/gateway/graph/convert"
	"github.com/davidchristie/cloud/pkg/gateway/graph/dataloader"
	"github.com/davidchristie/cloud/pkg/gateway/graph/generated"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/write/api"
	"github.com/google/uuid"
)

func (r *lineItemResolver) Product(ctx context.Context, obj *model.LineItem) (*model.Product, error) {
	return dataloader.For(ctx).Product.Load(obj.ProductID)
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	correlationID := uuid.New()
	customer, err := r.CustomerWriteAPI.CreateCustomer(input.FirstName, input.LastName, correlationID)
	if err != nil {
		return nil, err
	}
	return convert.Customer(customer), nil
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
	return convert.Order(createdOrder), nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	correlationID := uuid.New()
	product, err := r.ProductWriteAPI.CreateProduct(input.Name, input.Description, correlationID)
	if err != nil {
		return nil, err
	}
	return convert.Product(product), nil
}

func (r *orderResolver) Customer(ctx context.Context, obj *model.Order) (*model.Customer, error) {
	customer, err := r.CustomerReadAPI.Customer(obj.CustomerID)
	switch err {
	case nil:
		return convert.Customer(customer), nil

	case customerReadAPI.ErrCustomerNotFound:
		return nil, nil

	default:
		return nil, err
	}
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	customers, err := r.CustomerReadAPI.Customers()
	if err != nil {
		return []*model.Customer{}, err
	}
	modelCustomers := make([]*model.Customer, len(customers))
	for i, customer := range customers {
		modelCustomers[i] = convert.Customer(customer)
	}
	return modelCustomers, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.OrderReadAPI.Orders()
	if err != nil {
		return []*model.Order{}, err
	}
	modelOrders := make([]*model.Order, len(orders))
	for i, order := range orders {
		modelOrders[i] = convert.Order(order)
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
		return []*model.Product{}, err
	}
	products, errs := dataloader.For(ctx).Product.LoadAll(results)
	for _, err := range errs {
		if err != nil {
			return []*model.Product{}, err
		}
	}

	return products, nil
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
