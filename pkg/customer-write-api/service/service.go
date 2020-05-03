package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer-write-api/core"
	"github.com/davidchristie/cloud/pkg/customer-write-api/handler"
)

func Start() {
	c := core.NewCore()

	// Add handle func for producer.
	http.HandleFunc("/customers", handler.CreateCustomerHandler(c))

	// Run the web server.
	fmt.Println("serving customer-write-api...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
