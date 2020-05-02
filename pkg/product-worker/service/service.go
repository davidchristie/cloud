package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/kafka"
	productDatabase "github.com/davidchristie/cloud/pkg/product-database"
	"github.com/kelseyhightower/envconfig"
)

type specificiation struct {
	KafkaProductCreatedTopic string `required:"true" split_words:"true"`
}

func Start() {
	spec := specificiation{}
	envconfig.MustProcess("", &spec)

	productRepository := productDatabase.NewProductRespository(productDatabase.Connect())

	topic := spec.KafkaProductCreatedTopic

	reader := kafka.NewKafkaReader(topic)

	defer reader.Close()

	fmt.Println("reading messages from topic: " + topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("message value: %v\n", string(msg.Value))

		product := decodeProduct(msg.Value)
		fmt.Printf("product created: %+v\n", product)

		productRepository.CreateProduct(context.Background(), product)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func decodeProduct(data []byte) *entity.Product {
	product := entity.Product{}
	err := json.Unmarshal(data, &product)
	if err != nil {
		log.Fatal(err)
	}
	return &product
}
