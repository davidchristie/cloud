package core

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/davidchristie/cloud/pkg/elasticsearch"
	"github.com/google/uuid"
)

func (c *core) Customers(ctx context.Context, query string) ([]uuid.UUID, error) {
	res, err := c.Elasticsearch.Search(
		c.Elasticsearch.Search.WithContext(context.Background()),
		c.Elasticsearch.Search.WithIndex(c.Specification.ElasticsearchCustomerIndex),
		c.Elasticsearch.Search.WithBody(strings.NewReader(elasticsearch.NewCustomerRequestBody(query))),
		c.Elasticsearch.Search.WithTrackTotalHits(true),
		c.Elasticsearch.Search.WithPretty(),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting response: %s", err))
	}
	defer res.Body.Close()
	return elasticsearch.FormatCustomerResults(res)
}
