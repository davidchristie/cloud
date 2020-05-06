package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	customerDatabase "github.com/davidchristie/cloud/pkg/customer-database"
	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/davidchristie/cloud/pkg/message"
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

	reader := kafka.NewReader(topic)

	defer reader.Close()

	fmt.Println("reading messages from topic: " + topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("message value: %v\n", string(msg.Value))

		event := message.CustomerCreatedEvent{}
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			fmt.Println("error consuming message, ignoring: ", err)
		}

		fmt.Printf("customer created: %+v\n", event.Data)

		customerRepository.CreateCustomer(context.Background(), event.Data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
