package consumer

import "github.com/kelseyhightower/envconfig"

type specification struct {
	ElasticsearchCustomerIndex string `required:"true" split_words:"true"`
	ElasticsearchProductIndex  string `required:"true" split_words:"true"`
	KafkaCustomerCreatedTopic  string `required:"true" split_words:"true"`
	KafkaProductCreatedTopic   string `required:"true" split_words:"true"`
}

var spec specification

func init() {
	envconfig.MustProcess("", &spec)
}
