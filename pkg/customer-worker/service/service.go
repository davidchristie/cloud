package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	customerDatabase "github.com/davidchristie/cloud/pkg/customer-database"
	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/kelseyhightower/envconfig"
)

type specificiation struct {
	KafkaCustomerCreatedTopic string `required:"true" split_words:"true"`
}

func Start() {
	spec := specificiation{}
	envconfig.MustProcess("", &spec)

	customerRepository := customerDatabase.NewCustomerRepository(customerDatabase.Connect())

	topic := spec.KafkaCustomerCreatedTopic

	reader := kafka.NewKafkaReader(topic)

	defer reader.Close()

	fmt.Println("reading messages from topic: " + topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("message value: %v\n", string(msg.Value))

		customer := decodeCustomer(msg.Value)
		fmt.Printf("customer created: %+v\n", customer)

		customerRepository.CreateCustomer(context.Background(), customer)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func decodeCustomer(data []byte) *entity.Customer {
	customer := entity.Customer{}
	err := json.Unmarshal(data, &customer)
	if err != nil {
		log.Fatal(err)
	}
	return &customer
}
