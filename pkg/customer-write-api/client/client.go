package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/kelseyhightower/envconfig"
)

type CustomerWriteAPIClient interface {
	CreateCustomer(firstName string, lastName string) (*entity.Customer, error)
}

type client struct {
	URL string
}

type createCustomerResponseBody struct {
	Data    *entity.Customer `json:"data"`
	Message string           `json:"message"`
}

type createCustomerRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type specification struct {
	URL string `required:"true"`
}

func NewClient() CustomerWriteAPIClient {
	spec := specification{}
	envconfig.MustProcess("CUSTOMER_WRITE_API", &spec)
	return &client{
		URL: spec.URL,
	}
}

func (c *client) CreateCustomer(firstName string, lastName string) (*entity.Customer, error) {
	requestBodyBytes, err := json.Marshal(&createCustomerRequestBody{
		FirstName: firstName,
		LastName:  lastName,
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
	body, err := unmarshalCreateCustomerResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}

func unmarshalCreateCustomerResponseBody(response *http.Response) (*createCustomerResponseBody, error) {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := createCustomerResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
