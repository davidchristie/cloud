package handler

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/read/api/core"
)

type CustomersResponseBody struct {
	Data    []*customer.Customer `json:"data,omitempty"`
	Message string               `json:"message,omitempty"`
}

func CustomersHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := CustomersResponseBody{}

		writer.Header().Add("Content-Type", "application/json")

		customers, err := c.Customers(request.Context())
		switch err {
		case nil:
			response.Data = customers
			break

		default:
			writer.WriteHeader(500)
			response.Message = "Internal Server Error"
			break
		}

		data, _ := json.Marshal(response)
		writer.Write(data)
	})
}
