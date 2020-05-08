package api

import (
	"net/http"

	customerDatabase "github.com/davidchristie/cloud/pkg/customer/database"
	"github.com/davidchristie/cloud/pkg/customer/read/api/handler"
)

func StartService() error {
	db := customerDatabase.Connect()

	customerRepository := customerDatabase.NewCustomerRepository(db)

	http.HandleFunc("/customers", handler.CustomersHandler(customerRepository))

	return http.ListenAndServe(":8080", nil)
}
