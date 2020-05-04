package database

import (
	"context"
	"fmt"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	CreateOrder(context.Context, *order.Order) error
	GetOrders(context.Context) ([]*order.Order, error)
}

type orderRepository struct {
	collection *mongo.Collection
}

type orderSpecification struct {
	OrderCollectionName string `required:"true" split_words:"true"`
}

func NewOrderRepository(database *mongo.Database) OrderRepository {
	spec := orderSpecification{}

	envconfig.MustProcess("", &spec)

	collection := database.Collection(spec.OrderCollectionName)

	return &orderRepository{
		collection: collection,
	}
}

func (p *orderRepository) CreateOrder(ctx context.Context, order *order.Order) error {
	insertResult, err := p.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single order: ", insertResult.InsertedID)
	return nil
}

func (p *orderRepository) GetOrders(ctx context.Context) ([]*order.Order, error) {
	cursor, err := p.collection.Find(ctx, bson.M{})
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
