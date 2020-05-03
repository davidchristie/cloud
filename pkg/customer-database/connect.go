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
	Name string `required:"true"`
	URL  string `required:"true"`
}

func Connect() *mongo.Database {
	fmt.Println("connecting to customer database")

	spec := databaseSpecification{}

	envconfig.MustProcess("CUSTOMER_DATABASE", &spec)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(spec.URL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to customer database")

	db := client.Database(spec.Name)

	return db
}
