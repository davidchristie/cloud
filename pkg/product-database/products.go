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

type ProductRepository interface {
	CreateProduct(context.Context, *entity.Product) error
	GetProducts(context.Context) ([]*entity.Product, error)
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
	document, err := encodeProduct(product)
	if err != nil {
		return err
	}
	insertResult, err := p.collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single product: ", insertResult.InsertedID)
	return nil
}

func (p *productRepository) GetProducts(ctx context.Context) ([]*entity.Product, error) {
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
	products, err := convertDocumentsToProducts(documents)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func convertDocumentsToProducts(documents *[]bson.Raw) ([]*entity.Product, error) {
	products := make([]*entity.Product, len(*documents))
	for i, document := range *documents {
		product, err := decodeProduct(&document)
		if err != nil {
			return nil, err
		}
		products[i] = product
	}
	return products, nil
}

func decodeProduct(document *bson.Raw) (*entity.Product, error) {
	id, err := uuid.Parse(document.Lookup("id").StringValue())
	if err != nil {
		return nil, err
	}
	product := entity.Product{
		Description: document.Lookup("description").StringValue(),
		ID:          id,
		Name:        document.Lookup("name").StringValue(),
	}
	return &product, nil
}

func encodeProduct(product *entity.Product) (*map[string]interface{}, error) {
	document := make(map[string]interface{})
	data, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &document)
	document["_id"] = product.ID.String()
	return &document, err
}
