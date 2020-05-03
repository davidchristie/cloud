package core

import (
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/kelseyhightower/envconfig"
)

type Core interface {
	CreateCustomer(*CreateCustomerInput) (*CreateCustomerOutput, error)
}

type core struct {
	customerCreatedWriter *kafka.Writer
}

type specification struct {
	KafkaCustomerCreatedTopic string `required:"true" split_words:"true"`
}

func NewCore() Core {
	spec := specification{}
	envconfig.MustProcess("", &spec)
	return &core{
		customerCreatedWriter: kafka.NewWriter(spec.KafkaCustomerCreatedTopic),
	}
}
