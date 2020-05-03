package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/kelseyhightower/envconfig"
)

type CustomerReadAPIClient interface {
	Customers() ([]*entity.Customer, error)
}

type client struct {
	url string
}

type customersResponseBody struct {
	Data    *[]*entity.Customer `json:"data"`
	Message string              `json:"message"`
}

type specification struct {
	URL string `required:"true"`
}

func NewClient() CustomerReadAPIClient {
	spec := specification{}
	envconfig.MustProcess("CUSTOMER_READ_API", &spec)
	return &client{
		url: spec.URL,
	}
}

func (c *client) Customers() ([]*entity.Customer, error) {
	url := c.url + "/customers"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	body, err := unmarshalCustomersResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return *body.Data, nil
}

func unmarshalCustomersResponseBody(response *http.Response) (*customersResponseBody, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := customersResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Fatal(err)
	}
	return &body, nil
}
