package core

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/davidchristie/cloud/pkg/message"
	"github.com/davidchristie/cloud/pkg/order"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type CreateOrderInput struct {
	Context       context.Context
	CorrelationID uuid.UUID
	CustomerID    uuid.UUID
	LineItems     []*order.LineItem
}

type CreateOrderOutput struct {
	CreatedOrder *order.Order
}

func (c *core) CreateOrder(input *CreateOrderInput) (*CreateOrderOutput, error) {
	order := order.Order{
		CreatedAt:  time.Now().In(time.UTC).Truncate(time.Millisecond),
		CustomerID: input.CustomerID,
		ID:         uuid.New(),
		LineItems:  input.LineItems,
	}
	fmt.Printf("order created: %+v\n", order)
	event := message.OrderCreatedEvent{
		Data: &order,
		Metadata: &message.EventMetadata{
			CorrelationID: input.CorrelationID,
		},
	}
	value, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	fmt.Printf("message value: %+v\n", string(value))
	msg := kafka.Message{
		Value: value,
	}
	fmt.Println("writing to topic: ", c.orderCreatedWriter.Stats().Topic)
	err = c.orderCreatedWriter.WriteMessages(input.Context, msg)
	return &CreateOrderOutput{
		CreatedOrder: &order,
	}, nil
}
