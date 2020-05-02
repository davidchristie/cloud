package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/product-write-api/core"
	"github.com/davidchristie/cloud/pkg/product-write-api/handler"
)

func Start() {
	c := core.NewCore()

	// Add handle func for producer.
	http.HandleFunc("/products", handler.CreateProductHandler(c))

	// Run the web server.
	fmt.Println("start producer-api ... !!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
