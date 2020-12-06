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

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	customerID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	customer, err := r.CustomerReadAPI.Customer(customerID)
	if err != nil {
		return nil, err
	}
	return convert.Customer(customer), nil
}

func (r *queryResolver) Customers(ctx context.Context, query *string) ([]*model.Customer, error) {
	q := ""
	if query != nil {
		q = *query
	}
	results, err := r.SearchAPI.Customers(q)
	if err != nil {
		return nil, err
	}
	customers, errs := dataloader.For(ctx).Customer.LoadAll(results)

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	// Filter out any nil elements. A customer will be nil if it has been indexed
	// in Elasticsearch but isn't found in the customer database.
	//
	// FIXME: Guarantee all customers indexed in Elasticsearch exist in the customer database.
	nonNilCustomers := make([]*model.Customer, 0)
	for _, customer := range customers {
		if customer != nil {
			nonNilCustomers = append(nonNilCustomers, customer)
		}
	}

	return nonNilCustomers, nil
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	orderID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	order, err := r.OrderReadAPI.Order(orderID)
	if err != nil {
		return nil, err
	}
	return convert.Order(order), nil
}

func (r *queryResolver) Orders(ctx context.Context, customerID *string) ([]*model.Order, error) {
	orders, err := r.OrderReadAPI.Orders(customerID)
	if err != nil {
		return nil, err
	}
	modelOrders := make([]*model.Order, len(orders))
	for i, order := range orders {
		modelOrders[i] = convert.Order(order)
	}
	return modelOrders, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	productID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	product, err := r.ProductReadAPI.Product(productID)
	if err != nil {
		return nil, err
	}
	return convert.Product(product), nil
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
	products, errs := dataloader.For(ctx).Product.LoadAll(results)

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	// Filter out any nil elements. A product will be nil if it has been indexed
	// in Elasticsearch but isn't found in the product database.
	//
	// FIXME: Guarantee all products indexed in Elasticsearch exist in the product database.
	nonNilProducts := make([]*model.Product, 0)
	for _, product := range products {
		if product != nil {
			nonNilProducts = append(nonNilProducts, product)
		}
	}

	return nonNilProducts, nil
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
