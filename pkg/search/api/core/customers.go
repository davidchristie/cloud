package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

const customerSearch = `
	{
		"query": {
			"multi_match": {
				"fields": [
					"first_name",
					"last_name"
				],
				"query": "%s"
			}
		},
		"size": 25
	}
`

const customerSearchAll = `
	{
		"query": {
			"match_all": {}
		},
		"size": 25
	}	
`

func (c *core) Customers(ctx context.Context, query string) ([]uuid.UUID, error) {

	log.SetFlags(0)

	var (
		r map[string]interface{}
	)

	// Build the request body.
	var body string
	if query == "" {
		body = customerSearchAll
	} else {
		body = fmt.Sprintf(customerSearch, query)
	}

	// Perform the search request.
	res, err := c.Elasticsearch.Search(
		c.Elasticsearch.Search.WithContext(context.Background()),
		c.Elasticsearch.Search.WithIndex(c.Specification.ElasticsearchCustomerIndex),
		c.Elasticsearch.Search.WithBody(strings.NewReader(body)),
		c.Elasticsearch.Search.WithTrackTotalHits(true),
		c.Elasticsearch.Search.WithPretty(),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting response: %s", err))
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errors.New(fmt.Sprintf("Error parsing the response body: %s", err))
		} else {
			// Print the response status and error information.
			return nil, errors.New(fmt.Sprintf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			))
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing the response body: %s", err))
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})

	results := make([]uuid.UUID, len(hits))

	for i, hit := range hits {
		id, err := uuid.Parse(hit.(map[string]interface{})["_id"].(string))
		if err != nil {
			log.Printf("Error parsing customer ID: %v", err)
		}
		results[i] = id
	}

	return results, nil
}
