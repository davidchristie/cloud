package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/product/write/api/core"
	"github.com/google/uuid"
)

type CreateProductRequestBody struct {
	CorrelationID uuid.UUID `json:"correlation_id"`
	Description   string    `json:"description"`
	Name          string    `json:"name"`
}

type CreateProductResponseBody struct {
	Data    *entity.Product `json:"data,omitempty"`
	Message string          `json:"message,omitempty"`
}

func CreateProductHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Add("Content-Type", "application/json")
		requestBodyBytes, err := ioutil.ReadAll(req.Body)
		if err == nil {
			requestBody := CreateProductRequestBody{}
			err = json.Unmarshal(requestBodyBytes, &requestBody)
			if err == nil {
				output, err := c.CreateProduct(&core.CreateProductInput{
					Context:        req.Context(),
					CorreleationID: requestBody.CorrelationID,
					Description:    requestBody.Description,
					Name:           requestBody.Name,
				})
				if err == nil {
					responseBody := CreateProductResponseBody{
						Data: output.CreatedProduct,
					}
					responseBodyBytes, _ := json.Marshal(responseBody)
					wrt.Write(responseBodyBytes)
					return
				}
			}
		}
		fmt.Println(err)
		responseBody := CreateProductResponseBody{
			Message: err.Error(),
		}
		responseBodyBytes, _ := json.Marshal(responseBody)
		wrt.WriteHeader(500)
		wrt.Write(responseBodyBytes)
	})
}
