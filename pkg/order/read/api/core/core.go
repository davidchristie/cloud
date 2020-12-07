package core

import (
	"context"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/database"
	"github.com/google/uuid"
)

type Core interface {
	Order(context.Context, uuid.UUID) (*order.Order, error)
	Orders(context.Context, OrdersInput) ([]*order.Order, error)
}

type OrdersInput struct {
	CustomerID *uuid.UUID
	Limit      *int64
	Skip       *int64
}

type core struct {
	orderRepository database.OrderRepository
}

func NewCore(orderRepository database.OrderRepository) Core {
	return &core{
		orderRepository: orderRepository,
	}
}

func (c *core) Order(ctx context.Context, orderID uuid.UUID) (*order.Order, error) {
	return c.orderRepository.FindOrder(ctx, orderID)
}

func (c *core) Orders(ctx context.Context, input OrdersInput) ([]*order.Order, error) {
	return c.orderRepository.FindOrders(ctx, database.FindOrdersInput{
		CustomerID: input.CustomerID,
		Limit:      input.Limit,
		Skip:       input.Skip,
	})
}
