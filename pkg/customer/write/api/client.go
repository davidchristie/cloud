package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/write/api/handler"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Client interface {
	CreateCustomer(firstName, lastName string, correlationID uuid.UUID) (*customer.Customer, error)
}

type client struct {
	URL string
}

type specification struct {
	URL string `required:"true"`
}

func NewClient() Client {
	spec := specification{}
	envconfig.MustProcess("CUSTOMER_WRITE_API", &spec)
	return &client{
		URL: spec.URL,
	}
}

func (c *client) CreateCustomer(firstName, lastName string, correlationID uuid.UUID) (*customer.Customer, error) {
	requestBodyBytes, err := json.Marshal(&handler.CreateCustomerRequestBody{
		CorrelationID: correlationID,
		FirstName:     firstName,
		LastName:      lastName,
	})
	if err != nil {
		return nil, err
	}
	url := c.URL + "/customers"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := handler.CreateCustomerResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}
