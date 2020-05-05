package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/message"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type CreateProductInput struct {
	Context        context.Context
	CorreleationID uuid.UUID
	Description    string
	Name           string
}

type CreateProductOutput struct {
	CreatedProduct *entity.Product
}

func (c *core) CreateProduct(input *CreateProductInput) (*CreateProductOutput, error) {
	product := entity.Product{
		Description: input.Description,
		ID:          uuid.New(),
		Name:        input.Name,
	}
	fmt.Printf("product created: %+v\n", product)
	event := message.ProductCreatedEvent{
		Data: &product,
		Metadata: &message.EventMetadata{
			CorrelationID: input.CorreleationID,
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
	fmt.Println("writing to topic: ", c.productCreatedWriter.Stats().Topic)
	err = c.productCreatedWriter.WriteMessages(input.Context, msg)
	return &CreateProductOutput{
		CreatedProduct: &product,
	}, nil
}
