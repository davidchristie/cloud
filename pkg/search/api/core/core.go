package core

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Core interface {
	Customers(ctx context.Context, query string) ([]uuid.UUID, error)
	Products(ctx context.Context, query string) ([]uuid.UUID, error)
}

type core struct {
	Elasticsearch *elasticsearch.Client
	Specification specification
}

type specification struct {
	ElasticsearchCustomerIndex string `required:"true" split_words:"true"`
	ElasticsearchProductIndex  string `required:"true" split_words:"true"`
}

func NewCore() Core {
	spec := specification{}
	envconfig.MustProcess("", &spec)
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}
	return &core{
		Elasticsearch: es,
		Specification: spec,
	}
}
