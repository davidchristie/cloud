package service

import (
	"fmt"
	"log"
	"net/http"

	productDatabase "github.com/davidchristie/cloud/pkg/product-database"
	"github.com/davidchristie/cloud/pkg/product-read-api/handler"
)

func Start() {
	db := productDatabase.Connect()

	productRepository := productDatabase.NewProductRepository(db)

	http.HandleFunc("/products", handler.ProductsHandler(productRepository))

	fmt.Println("serving product-read-api...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
