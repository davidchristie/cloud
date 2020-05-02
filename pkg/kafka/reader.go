package kafka

import (
	"github.com/kelseyhightower/envconfig"
	kafka "github.com/segmentio/kafka-go"
)

type readerSpecification struct {
	GroupID string `required:"true" split_words:"true"`
}

type Reader = kafka.Reader

func NewKafkaReader(topic string) *Reader {
	kafkaSpec := kafkaSpecification{}
	readerSpec := readerSpecification{}
	envconfig.MustProcess(prefix, &kafkaSpec)
	envconfig.MustProcess(prefix, &readerSpec)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaSpec.Broker},
		GroupID:  readerSpec.GroupID,
		MaxBytes: 10e6, // 10MB
		MinBytes: 10e3, // 10KB
		Topic:    topic,
	})
}
