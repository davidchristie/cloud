package handler

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/cloud/pkg/product"
	"github.com/davidchristie/cloud/pkg/product/read/api/core"
)

type ProductsResponseBody struct {
	Data    []*product.Product `json:"data,omitempty"`
	Message string             `json:"message,omitempty"`
}

func ProductsHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := ProductsResponseBody{}

		writer.Header().Add("Content-Type", "application/json")

		products, err := c.Products(request.Context())
		switch err {
		case nil:
			response.Data = products
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
