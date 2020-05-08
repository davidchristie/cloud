package database

import (
	"context"
	"errors"

	"github.com/davidchristie/cloud/pkg/product"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(context.Context, *product.Product) error
	FindProduct(context.Context, uuid.UUID) (*product.Product, error)
	FindProducts(context.Context) ([]*product.Product, error)
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
	_, err := p.collection.InsertOne(context.Background(), product)
	if err != nil {
		return err
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

func (p *productRepository) FindProducts(ctx context.Context) ([]*product.Product, error) {
	cursor, err := p.collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	var results []*product.Product

	for cursor.Next(ctx) {
		product := &product.Product{}

		err := cursor.Decode(product)
		if err != nil {
			return nil, err
		}
		results = append(results, product)
	}

	return results, nil
}
