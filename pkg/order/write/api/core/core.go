package core

import (
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/kelseyhightower/envconfig"
)

type Core interface {
	CreateOrder(*CreateOrderInput) (*CreateOrderOutput, error)
}

type core struct {
	orderCreatedWriter *kafka.Writer
}

type specification struct {
	KafkaOrderCreatedTopic string `required:"true" split_words:"true"`
}

func NewCore() Core {
	spec := specification{}
	envconfig.MustProcess("", &spec)
	return &core{
		orderCreatedWriter: kafka.NewWriter(spec.KafkaOrderCreatedTopic),
	}
}
