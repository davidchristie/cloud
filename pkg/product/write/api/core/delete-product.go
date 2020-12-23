package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/davidchristie/cloud/pkg/message"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type DeleteProductInput struct {
	Context        context.Context
	CorreleationID uuid.UUID
	ProductID      uuid.UUID
}

func (c *core) DeleteProduct(input *DeleteProductInput) error {
	fmt.Printf("product deleted: %+v\n", input.ProductID)
	event := message.ProductDeletedEvent{
		Metadata: &message.EventMetadata{
			CorrelationID: input.CorreleationID,
		},
		ProductID: input.ProductID,
	}
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Printf("message value: %+v\n", string(value))
	msg := kafka.Message{
		Value: value,
	}
	fmt.Println("writing to topic: ", c.productDeletedWriter.Stats().Topic)
	return c.productDeletedWriter.WriteMessages(input.Context, msg)
}
