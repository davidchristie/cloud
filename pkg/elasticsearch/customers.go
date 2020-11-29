package elasticsearch

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
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

func FormatCustomerResults(res *esapi.Response) ([]uuid.UUID, error) {
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

func NewCustomerRequestBody(query string) string {
	if query == "" {
		return customerSearchAll
	} else {
		return fmt.Sprintf(customerSearch, query)
	}
}
