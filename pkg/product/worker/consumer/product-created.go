package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/davidchristie/cloud/pkg/message"
	productDatabase "github.com/davidchristie/cloud/pkg/product/database"
)

func ProductCreatedConsumer(productRepository productDatabase.ProductRepository) {
	topic := spec.KafkaProductCreatedTopic

	reader := kafka.NewReader(topic)

	defer reader.Close()

	fmt.Println("reading messages from topic: " + topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("message value: %v\n", string(msg.Value))

		event := message.ProductCreatedEvent{}
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			fmt.Println("error consuming message: ", err)
			continue
		}

		fmt.Printf("product created: %+v\n", event.Data)

		productRepository.CreateProduct(context.Background(), event.Data)
		if err != nil {
			fmt.Println("error writing product to database: ", err)
			continue
		}
	}
}
