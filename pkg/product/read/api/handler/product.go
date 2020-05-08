package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/product"
	"github.com/davidchristie/cloud/pkg/product/read/api/core"
	"github.com/gorilla/mux"
)

type ProductResponseBody struct {
	Data    *product.Product `json:"data,omitempty"`
	Message string           `json:"message,omitempty"`
}

func ProductHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := ProductResponseBody{}

		writer.Header().Add("Content-Type", "application/json")

		id := mux.Vars(request)["id"]
		product, err := c.Product(request.Context(), id)
		switch err {
		case nil:
			response.Data = product
			break

		case core.ErrProductNotFound:
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
