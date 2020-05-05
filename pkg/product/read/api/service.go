package api

import (
	"fmt"
	"net/http"

	productDatabase "github.com/davidchristie/cloud/pkg/product/database"
	"github.com/davidchristie/cloud/pkg/product/read/api/handler"
	"github.com/kelseyhightower/envconfig"
)

type serviceSpecification struct {
	Port string `default:"8080"`
}

func StartService() error {
	spec := serviceSpecification{}
	envconfig.MustProcess("", &spec)

	db := productDatabase.Connect()

	productRepository := productDatabase.NewProductRepository(db)

	http.HandleFunc("/products", handler.ProductsHandler(productRepository))

	fmt.Println("serving")

	address := ":" + spec.Port

	return http.ListenAndServe(address, nil)
}
