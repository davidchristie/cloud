package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type databaseSpecification struct {
	ProductDatabaseName string `required:"true" split_words:"true"`
	ProductDatabaseURL  string `required:"true" split_words:"true"`
}

func Connect() *mongo.Database {
	fmt.Println("connecting to product database")

	spec := databaseSpecification{}

	envconfig.MustProcess("", &spec)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(spec.ProductDatabaseURL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to product database")

	db := client.Database(spec.ProductDatabaseName)

	return db
}
