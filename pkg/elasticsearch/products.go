package elasticsearch

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
)

const productSearch = `
	{
		"query": {
			"multi_match": {
				"fields": [
					"description",
					"name"
				],
				"query": "%s"
			}
		},
		"size": 25
	}
`

const productSearchAll = `
	{
		"query": {
			"match_all": {}
		},
		"size": 25
	}	
`

func FormatProductResults(res *esapi.Response) ([]uuid.UUID, error) {
	err := FormatError(res)
	if err != nil {
		return nil, err
	}
	body, err := ParseResponseBody(res)
	if err != nil {
		return nil, err
	}
	PrintResponseStatus(res)
	PrintHitInfo(body)
	return HitIDs(body), nil
}

func NewProductRequestBody(query string) string {
	if query == "" {
		return productSearchAll
	} else {
		return fmt.Sprintf(productSearch, query)
	}
}
