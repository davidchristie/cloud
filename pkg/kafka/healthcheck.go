package kafka

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/segmentio/kafka-go"
)

type healthcheckSpecification struct {
	Topic string `required:"true"`
}

func WaitUntilHealthy() {
	log.Println("waiting for Kafka brokers to become healthy")

	spec := healthcheckSpecification{}
	envconfig.MustProcess("KAFKA_HEALTHCHECK", &spec)

	ctx := context.Background()

	// Produce a test message.
	ping := []byte("ping+" + uuid.New().String())
	writer := NewWriter(spec.Topic)
	defer writer.Close()
	for {
		err := writer.WriteMessages(ctx, kafka.Message{Value: ping})
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(5 * time.Second)
	}

	// Wait for the message to be consumed.
	reader := NewReader(spec.Topic)
	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		if bytes.Equal(msg.Value, ping) {
			break
		}
	}
}
