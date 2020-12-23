package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidchristie/cloud/pkg/product"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(context.Context, *product.Product) error
	DeleteProduct(context.Context, uuid.UUID) error
	FindProduct(context.Context, uuid.UUID) (*product.Product, error)
	FindProducts(context.Context, []uuid.UUID) (map[uuid.UUID]*product.Product, error)
}

type productRepository struct {
	collection *mongo.Collection
}

type productSpecification struct {
	ProductCollectionName string `required:"true" split_words:"true"`
}

var ErrProductNotFound = errors.New("product not found")

func NewProductRepository(database *mongo.Database) ProductRepository {
	spec := productSpecification{}

	envconfig.MustProcess("", &spec)

	collection := database.Collection(spec.ProductCollectionName)

	return &productRepository{
		collection: collection,
	}
}

func (p *productRepository) CreateProduct(ctx context.Context, product *product.Product) error {
	_, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	result, err := p.collection.DeleteOne(ctx, bson.D{
		{Key: "id", Value: id},
	})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("product not found: %v", id)
	}
	return nil
}

func (p *productRepository) FindProduct(ctx context.Context, id uuid.UUID) (*product.Product, error) {
	result := p.collection.FindOne(ctx, bson.D{
		{Key: "id", Value: id},
	})
	err := result.Err()
	if err == mongo.ErrNoDocuments {
		return nil, ErrProductNotFound
	}
	if err != nil {
		return nil, err
	}
	product := product.Product{}
	err = result.Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) FindProducts(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]*product.Product, error) {
	cursor, err := p.collection.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	results := map[uuid.UUID]*product.Product{}

	for cursor.Next(ctx) {
		product := &product.Product{}

		err := cursor.Decode(product)
		if err != nil {
			return nil, err
		}
		results[product.ID] = product
	}

	return results, nil
}
