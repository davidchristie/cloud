package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type CreateCustomerInput struct {
	Context   context.Context
	FirstName string
	LastName  string
}

type CreateCustomerOutput struct {
	CreatedCustomer *entity.Customer
}

func (c *core) CreateCustomer(input *CreateCustomerInput) (*CreateCustomerOutput, error) {
	customer := entity.Customer{
		FirstName: input.FirstName,
		ID:        uuid.New(),
		LastName:  input.LastName,
	}
	fmt.Printf("customer created: %+v\n", customer)
	value, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	fmt.Printf("message value: %+v\n", string(value))
	msg := kafka.Message{
		Value: value,
	}
	fmt.Println("writing to topic: ", c.customerCreatedWriter.Stats().Topic)
	err = c.customerCreatedWriter.WriteMessages(input.Context, msg)
	return &CreateCustomerOutput{
		CreatedCustomer: &customer,
	}, nil
}
