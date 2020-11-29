package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository interface {
	CreateCustomer(context.Context, *customer.Customer) error
	FindCustomer(context.Context, uuid.UUID) (*customer.Customer, error)
	FindCustomers(context.Context, []uuid.UUID) (map[uuid.UUID]*customer.Customer, error)
}

type customerRepository struct {
	collection *mongo.Collection
}

type customerSpecification struct {
	CustomerCollectionName string `required:"true" split_words:"true"`
}

var ErrCustomerNotFound = errors.New("customer not found")

func NewCustomerRepository(database *mongo.Database) CustomerRepository {
	spec := customerSpecification{}

	envconfig.MustProcess("", &spec)

	collection := database.Collection(spec.CustomerCollectionName)

	return &customerRepository{
		collection: collection,
	}
}

func (p *customerRepository) CreateCustomer(ctx context.Context, customer *customer.Customer) error {
	insertResult, err := p.collection.InsertOne(context.Background(), customer)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single customer: ", insertResult.InsertedID)
	return nil
}

func (c *customerRepository) FindCustomer(ctx context.Context, id uuid.UUID) (*customer.Customer, error) {
	result := c.collection.FindOne(ctx, bson.D{
		{Key: "id", Value: id},
	})
	err := result.Err()
	if err == mongo.ErrNoDocuments {
		return nil, ErrCustomerNotFound
	}
	if err != nil {
		return nil, err
	}
	customer := customer.Customer{}
	err = result.Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customerRepository) FindCustomers(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]*customer.Customer, error) {
	cursor, err := c.collection.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	results := map[uuid.UUID]*customer.Customer{}

	for cursor.Next(ctx) {
		customer := &customer.Customer{}

		err := cursor.Decode(customer)
		if err != nil {
			return nil, err
		}
		results[customer.ID] = customer
	}

	return results, nil
}
