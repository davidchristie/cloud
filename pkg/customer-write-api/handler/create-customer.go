package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer-write-api/core"
	"github.com/davidchristie/cloud/pkg/entity"
)

type createCustomerRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type createCustomerResponseBody struct {
	Data    *entity.Customer `json:"data"`
	Message string           `json:"message"`
}

func CreateCustomerHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Add("Content-Type", "application/json")
		requestBodyBytes, err := ioutil.ReadAll(req.Body)
		if err == nil {
			requestBody := createCustomerRequestBody{}
			err = json.Unmarshal(requestBodyBytes, &requestBody)
			if err == nil {
				output, err := c.CreateCustomer(&core.CreateCustomerInput{
					Context:   req.Context(),
					FirstName: requestBody.FirstName,
					LastName:  requestBody.LastName,
				})
				if err == nil {
					responseBody := createCustomerResponseBody{
						Data: output.CreatedCustomer,
					}
					responseBodyBytes, _ := json.Marshal(responseBody)
					wrt.Write(responseBodyBytes)
					return
				}
			}
		}
		fmt.Println(err)
		responseBody := createCustomerResponseBody{
			Message: err.Error(),
		}
		responseBodyBytes, _ := json.Marshal(responseBody)
		wrt.WriteHeader(500)
		wrt.Write(responseBodyBytes)
	})
}
