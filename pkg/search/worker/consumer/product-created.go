package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/davidchristie/cloud/pkg/kafka"
	"github.com/davidchristie/cloud/pkg/message"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func ProductCreatedConsumer(es *elasticsearch.Client) {
	topic := spec.KafkaProductCreatedTopic

	reader := kafka.NewReader(topic)

	defer reader.Close()

	log.Println("Consuming events from topic: " + topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		event := message.ProductCreatedEvent{}
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			fmt.Println("Error consuming event (ignoring): ", err)
		}

		product := event.Data

		// Build the request body.
		body, err := json.Marshal(product)
		if err != nil {
			log.Printf("Error building request body: %v\n", err)
		}

		// Set up the request object.
		req := esapi.IndexRequest{
			Index:      spec.ElasticsearchProductIndex,
			DocumentID: product.ID.String(),
			Body:       strings.NewReader(string(body)),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("[%s] Error indexing document ID=%d", res.Status(), product.ID)
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				// Print the response status and indexed document version.
				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}
	}
}
