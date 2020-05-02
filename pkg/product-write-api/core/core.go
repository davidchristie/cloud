package core

import (
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/kelseyhightower/envconfig"
)

type Core interface {
	CreateProduct(*CreateProductInput) (*CreateProductOutput, error)
}

type core struct {
	productCreatedWriter *kafka.Writer
}

type specification struct {
	KafkaProductCreatedTopic string `required:"true" split_words:"true"`
}

func NewCore() Core {
	spec := specification{}
	envconfig.MustProcess("", &spec)
	return &core{
		productCreatedWriter: kafka.NewWriter(spec.KafkaProductCreatedTopic),
	}
}
