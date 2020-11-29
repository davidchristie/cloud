package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/search/api/core"
	"github.com/google/uuid"
)

type CustomersResponseBody struct {
	Data    []uuid.UUID `json:"data"`
	Message string      `json:"message,omitempty"`
}

func CustomersHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := CustomersResponseBody{}

		query := request.URL.Query().Get("q")

		writer.Header().Add("Content-Type", "application/json")

		results, err := c.Customers(request.Context(), query)
		switch err {
		case nil:
			response.Data = results
			break

		default:
			log.Printf("Error handling request: %s\n", err)
			writer.WriteHeader(500)
			response.Message = "Internal Server Error"
			break
		}

		data, _ := json.Marshal(response)
		writer.Write(data)
	})
}
