package elasticsearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
)

func FormatError(res *esapi.Response) error {
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return errors.New(fmt.Sprintf("error parsing the response body: %s", err))
		} else {
			return errors.New(fmt.Sprintf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			))
		}
	}
	return nil
}

func HitIDs(body map[string]interface{}) []uuid.UUID {
	hits := body["hits"].(map[string]interface{})["hits"].([]interface{})
	ids := make([]uuid.UUID, len(hits))
	for i, hit := range hits {
		id, err := uuid.Parse(hit.(map[string]interface{})["_id"].(string))
		if err != nil {
			log.Printf("error parsing hit ID: %v", err)
		}
		ids[i] = id
	}
	return ids
}

func ParseResponseBody(res *esapi.Response) (map[string]interface{}, error) {
	var (
		body map[string]interface{}
	)
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing response body: %s", err))
	}
	return body, nil
}

func PrintHitInfo(body map[string]interface{}) {
	log.Printf(
		"%d hits; took: %dms",
		int(body["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(body["took"].(float64)),
	)
	for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
	log.Println(strings.Repeat("=", 37))
}

func PrintResponseStatus(res *esapi.Response) {
	log.Printf("[%s]", res.Status())
}
