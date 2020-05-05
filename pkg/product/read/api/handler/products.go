package handler

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	productDatabase "github.com/davidchristie/cloud/pkg/product/database"
)

type ProductsResponseBody struct {
	Data    []*entity.Product `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
}

func ProductsHandler(productRespository productDatabase.ProductRepository) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		products, err := productRespository.GetProducts(ctx)

		writer.Header().Add("Content-Type", "application/json")

		if err != nil {
			writer.WriteHeader(500)
			response := &ProductsResponseBody{
				Message: err.Error(),
			}
			data, _ := json.Marshal(response)
			writer.Write(data)
		} else {
			response := &ProductsResponseBody{
				Data: products,
			}
			data, _ := json.Marshal(response)
			writer.Write(data)
		}
	})
}
