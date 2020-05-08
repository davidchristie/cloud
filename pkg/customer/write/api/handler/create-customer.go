package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/write/api/core"
	"github.com/google/uuid"
)

type CreateCustomerRequestBody struct {
	CorrelationID uuid.UUID `json:"correlation_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
}

type CreateCustomerResponseBody struct {
	Data    *customer.Customer `json:"data,omitempty"`
	Message string             `json:"message,omitempty"`
}

func CreateCustomerHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		requestBodyBytes, err := ioutil.ReadAll(request.Body)
		if err == nil {
			requestBody := CreateCustomerRequestBody{}
			err = json.Unmarshal(requestBodyBytes, &requestBody)
			if err == nil {
				output, err := c.CreateCustomer(&core.CreateCustomerInput{
					Context:       request.Context(),
					CorrelationID: requestBody.CorrelationID,
					FirstName:     requestBody.FirstName,
					LastName:      requestBody.LastName,
				})
				if err == nil {
					responseBody := CreateCustomerResponseBody{
						Data: output.CreatedCustomer,
					}
					responseBodyBytes, _ := json.Marshal(responseBody)
					writer.Write(responseBodyBytes)
					return
				}
			}
		}
		fmt.Println(err)
		responseBody := CreateCustomerResponseBody{
			Message: err.Error(),
		}
		responseBodyBytes, _ := json.Marshal(responseBody)
		writer.WriteHeader(500)
		writer.Write(responseBodyBytes)
	})
}
