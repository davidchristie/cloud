package database

import (
	"context"
	"fmt"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(context.Context, *entity.Product) error
	HasProduct(context.Context, uuid.UUID) (bool, error)
}

type productRepository struct {
	collection *mongo.Collection
}

type productSpecification struct {
	ProductCollectionName string `required:"true" split_words:"true"`
}

func NewProductRepository(database *mongo.Database) ProductRepository {
	spec := productSpecification{}
	envconfig.MustProcess("", &spec)
	collection := database.Collection(spec.ProductCollectionName)
	return &productRepository{
		collection: collection,
	}
}

func (p *productRepository) CreateProduct(ctx context.Context, product *entity.Product) error {
	document := encodeProduct(product)
	insertResult, err := p.collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single product: ", insertResult.InsertedID)
	return nil
}

func (p *productRepository) HasProduct(ctx context.Context, id uuid.UUID) (bool, error) {
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

func encodeProduct(product *entity.Product) *map[string]interface{} {
	return &map[string]interface{}{
		"_id": product.ID.String(),
	}
}
