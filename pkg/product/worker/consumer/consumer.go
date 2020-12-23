package consumer

import "github.com/kelseyhightower/envconfig"

type specification struct {
	KafkaProductCreatedTopic string `required:"true" split_words:"true"`
	KafkaProductDeletedTopic string `required:"true" split_words:"true"`
}

var spec specification

func init() {
	envconfig.MustProcess("", &spec)
}
