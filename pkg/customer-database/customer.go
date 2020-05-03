package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository interface {
	CreateCustomer(context.Context, *entity.Customer) error
	GetCustomers(context.Context) ([]*entity.Customer, error)
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

func (p *customerRepository) CreateCustomer(ctx context.Context, customer *entity.Customer) error {
	document, err := encodeCustomer(customer)
	if err != nil {
		return err
	}
	insertResult, err := p.collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single customer: ", insertResult.InsertedID)
	return nil
}

func (p *customerRepository) GetCustomers(ctx context.Context) ([]*entity.Customer, error) {
	cursor, err := p.collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	documents := &[]bson.Raw{}
	err = cursor.All(ctx, documents)
	if err != nil {
		return nil, err
	}
	customers, err := convertDocumentsToCustomers(documents)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func convertDocumentsToCustomers(documents *[]bson.Raw) ([]*entity.Customer, error) {
	customers := make([]*entity.Customer, len(*documents))
	for i, document := range *documents {
		customer, err := decodeCustomer(&document)
		if err != nil {
			return nil, err
		}
		customers[i] = customer
	}
	return customers, nil
}

func decodeCustomer(document *bson.Raw) (*entity.Customer, error) {
	id, err := uuid.Parse(document.Lookup("id").StringValue())
	if err != nil {
		return nil, err
	}
	customer := entity.Customer{
		FirstName: document.Lookup("first_name").StringValue(),
		ID:        id,
		LastName:  document.Lookup("last_name").StringValue(),
	}
	return &customer, nil
}

func encodeCustomer(customer *entity.Customer) (*map[string]interface{}, error) {
	document := make(map[string]interface{})
	data, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &document)
	document["_id"] = customer.ID.String()
	return &document, err
}
