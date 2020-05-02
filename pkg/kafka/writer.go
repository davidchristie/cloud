package kafka

import (
	"github.com/kelseyhightower/envconfig"
	kafka "github.com/segmentio/kafka-go"
)

type Writer = kafka.Writer

func NewWriter(topic string) *kafka.Writer {
	spec := kafkaSpecification{}
	envconfig.MustProcess(prefix, &spec)
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{spec.Broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}
