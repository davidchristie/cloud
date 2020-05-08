package database

import (
	"context"
	"fmt"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository interface {
	CreateCustomer(context.Context, *customer.Customer) error
	HasCustomer(context.Context, uuid.UUID) (bool, error)
}

type customerRepository struct {
	collection *mongo.Collection
}

type customerSpecification struct {
	CustomerCollectionName string `required:"true" split_words:"true"`
}

func NewCustomerRepository(database *mongo.Database) CustomerRepository {
	spec := customerSpecification{}
	envconfig.MustProcess("", &spec)
	collection := database.Collection(spec.CustomerCollectionName)
	return &customerRepository{
		collection: collection,
	}
}

func (p *customerRepository) CreateCustomer(ctx context.Context, customer *customer.Customer) error {
	document := encodeCustomer(customer)
	insertResult, err := p.collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single customer: ", insertResult.InsertedID)
	return nil
}

func (p *customerRepository) HasCustomer(ctx context.Context, id uuid.UUID) (bool, error) {
	result := p.collection.FindOne(ctx, bson.M{
		"_id": id.String(),
	})
	err := result.Err()
	switch err {
	case nil:
		return true, nil

	case mongo.ErrNoDocuments:
		return false, bson.ErrDecodeToNil

	default:
		return false, err
	}
}

func encodeCustomer(customer *customer.Customer) *map[string]interface{} {
	return &map[string]interface{}{
		"_id": customer.ID.String(),
	}
}
