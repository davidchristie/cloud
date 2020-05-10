package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/read/api/core"
	"github.com/gorilla/mux"
)

type CustomerResponseBody struct {
	Data    *customer.Customer `json:"data,omitempty"`
	Message string             `json:"message,omitempty"`
}

func CustomerHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := CustomerResponseBody{}

		writer.Header().Add("Content-Type", "application/json")

		id := mux.Vars(request)["id"]
		customer, err := c.Customer(request.Context(), id)
		switch err {
		case nil:
			response.Data = customer
			break

		case core.ErrCustomerNotFound:
			writer.WriteHeader(404)
			response.Message = "Not Found"
			break

		default:
			log.Println(err)
			writer.WriteHeader(500)
			response.Message = "Internal Server Error"
			break
		}

		data, _ := json.Marshal(response)
		writer.Write(data)
	})
}
