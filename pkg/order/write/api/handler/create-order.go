package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/write/api/core"
	"github.com/google/uuid"
)

type createOrderRequestBody struct {
	CorrelationID uuid.UUID         `json:"correlation_id"`
	CustomerID    uuid.UUID         `json:"customer_id"`
	LineItems     []*order.LineItem `json:"line_items"`
}

type createOrderResponseBody struct {
	Data    *order.Order `json:"data"`
	Message string       `json:"message"`
}

func CreateOrderHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Add("Content-Type", "application/json")
		requestBodyBytes, err := ioutil.ReadAll(req.Body)
		if err == nil {
			requestBody := createOrderRequestBody{}
			err = json.Unmarshal(requestBodyBytes, &requestBody)
			if err == nil {
				output, err := c.CreateOrder(&core.CreateOrderInput{
					Context:       req.Context(),
					CorrelationID: requestBody.CorrelationID,
					CustomerID:    requestBody.CustomerID,
					LineItems:     requestBody.LineItems,
				})
				if err == nil {
					responseBody := createOrderResponseBody{
						Data: output.CreatedOrder,
					}
					responseBodyBytes, _ := json.Marshal(responseBody)
					wrt.Write(responseBodyBytes)
					return
				}
			}
		}
		fmt.Println(err)
		responseBody := createOrderResponseBody{
			Message: err.Error(),
		}
		responseBodyBytes, _ := json.Marshal(responseBody)
		wrt.WriteHeader(500)
		wrt.Write(responseBodyBytes)
	})
}
