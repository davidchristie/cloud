package handler

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	productDatabase "github.com/davidchristie/cloud/pkg/product-database"
)

type productsResponseBody struct {
	Data []*entity.Product `json:"data"`
}

func ProductsHandler(productRespository productDatabase.ProductRepository) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		products, err := productRespository.GetProducts(ctx)

		writer.Header().Add("Content-Type", "application/json")

		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte("{\"message\":\"" + err.Error() + "\"}"))
		} else {
			response := &productsResponseBody{
				Data: products,
			}
			blob, _ := json.Marshal(response)
			writer.Write(blob)
		}
	})
}
