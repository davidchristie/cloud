package worker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/davidchristie/cloud/pkg/message"
	orderDatabase "github.com/davidchristie/cloud/pkg/order/database"
	"github.com/kelseyhightower/envconfig"
)

type specificiation struct {
	KafkaOrderCreatedTopic string `required:"true" split_words:"true"`
}

func StartService() error {
	spec := specificiation{}
	envconfig.MustProcess("", &spec)

	orderRepository := orderDatabase.NewOrderRepository(orderDatabase.Connect())

	topic := spec.KafkaOrderCreatedTopic

	reader := kafka.NewReader(topic)

	defer reader.Close()

	fmt.Println("reading messages from topic: " + topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("message value: %v\n", string(msg.Value))

		event := message.OrderCreatedEvent{}
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			fmt.Println("error consuming message, ignoring: ", err)
		}

		fmt.Printf("order created: %+v\n", event.Data)

		orderRepository.CreateOrder(context.Background(), event.Data)
	}

	return errors.New("error in order-worker service")
}
