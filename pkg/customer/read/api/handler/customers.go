package handler

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer"
	customerDatabase "github.com/davidchristie/cloud/pkg/customer/database"
)

type CustomersResponseBody struct {
	Data    []*customer.Customer `json:"data,omitempty"`
	Message string               `json:"string,omitempty"`
}

func CustomersHandler(customerRepository customerDatabase.CustomerRepository) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		customers, err := customerRepository.GetCustomers(ctx)

		writer.Header().Add("Content-Type", "application/json")

		response := CustomersResponseBody{}

		if err != nil {
			writer.WriteHeader(500)
			response.Message = err.Error()
		} else {
			response.Data = customers
		}

		data, _ := json.Marshal(response)
		writer.Write(data)
	})
}
