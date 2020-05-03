package handler

import (
	"encoding/json"
	"net/http"

	customerDatabase "github.com/davidchristie/cloud/pkg/customer-database"
	"github.com/davidchristie/cloud/pkg/entity"
)

type customersResponseBody struct {
	Data []*entity.Customer `json:"data"`
}

func CustomersHandler(customerRepository customerDatabase.CustomerRepository) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		customers, err := customerRepository.GetCustomers(ctx)

		writer.Header().Add("Content-Type", "application/json")

		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte("{\"message\":\"" + err.Error() + "\"}"))
		} else {
			response := &customersResponseBody{
				Data: customers,
			}
			blob, _ := json.Marshal(response)
			writer.Write(blob)
		}
	})
}
