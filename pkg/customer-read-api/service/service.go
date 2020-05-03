package service

import (
	"fmt"
	"log"
	"net/http"

	customerDatabase "github.com/davidchristie/cloud/pkg/customer-database"
	"github.com/davidchristie/cloud/pkg/customer-read-api/handler"
)

func Start() {
	db := customerDatabase.Connect()

	customerRepository := customerDatabase.NewCustomerRepository(db)

	http.HandleFunc("/customers", handler.CustomersHandler(customerRepository))

	fmt.Println("serving customer-read-api...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
