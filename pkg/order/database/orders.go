package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	CreateOrder(context.Context, *order.Order) error
	FindOrder(context.Context, uuid.UUID) (*order.Order, error)
	FindOrders(context.Context) ([]*order.Order, error)
}

type orderRepository struct {
	collection *mongo.Collection
}

type orderSpecification struct {
	OrderCollectionName string `required:"true" split_words:"true"`
}

var ErrOrderNotFound = errors.New("order not found")

func NewOrderRepository(database *mongo.Database) OrderRepository {
	spec := orderSpecification{}

	envconfig.MustProcess("", &spec)

	collection := database.Collection(spec.OrderCollectionName)

	return &orderRepository{
		collection: collection,
	}
}

func (o *orderRepository) CreateOrder(ctx context.Context, order *order.Order) error {
	insertResult, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single order: ", insertResult.InsertedID)
	return nil
}

func (o *orderRepository) FindOrder(ctx context.Context, id uuid.UUID) (*order.Order, error) {
	result := o.collection.FindOne(ctx, bson.D{
		{Key: "id", Value: id},
	})
	err := result.Err()
	if err == mongo.ErrNoDocuments {
		return nil, ErrOrderNotFound
	}
	if err != nil {
		return nil, err
	}
	order := order.Order{}
	err = result.Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *orderRepository) FindOrders(ctx context.Context) ([]*order.Order, error) {
	cursor, err := o.collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	var results []*order.Order

	for cursor.Next(ctx) {
		order := &order.Order{}

		err := cursor.Decode(order)
		if err != nil {
			return nil, err
		}
		results = append(results, order)
	}

	return results, nil
}
