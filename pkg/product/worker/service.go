package worker

import (
	"sync"

	productDatabase "github.com/davidchristie/cloud/pkg/product/database"
	"github.com/davidchristie/cloud/pkg/product/worker/consumer"
	"github.com/kelseyhightower/envconfig"
)

type specificiation struct {
	KafkaProductCreatedTopic string `required:"true" split_words:"true"`
}

func StartService() error {
	var wg sync.WaitGroup

	spec := specificiation{}
	envconfig.MustProcess("", &spec)

	productRepository := productDatabase.NewProductRepository(productDatabase.Connect())

	wg.Add(1)

	go consumer.ProductCreatedConsumer(productRepository)
	go consumer.ProductDeletedConsumer(productRepository)

	wg.Wait()

	return nil
}
