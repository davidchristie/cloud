package kafka

type kafkaSpecification struct {
	Broker string `required:"true" split_words:"true"`
}

const prefix = "kafka"
