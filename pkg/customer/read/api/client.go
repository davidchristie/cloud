package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/read/api/handler"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Client interface {
	Customer(uuid.UUID) (*customer.Customer, error)
	Customers() ([]*customer.Customer, error)
}

type client struct {
	url string
}

type specification struct {
	URL string `required:"true"`
}

var ErrCustomerNotFound = errors.New("customer not found")

func NewClient() Client {
	spec := specification{}
	envconfig.MustProcess("CUSTOMER_READ_API", &spec)
	return &client{
		url: spec.URL,
	}
}

// TODO: Implement this properly.
func (c *client) Customer(id uuid.UUID) (*customer.Customer, error) {
	customers, err := c.Customers()
	if err != nil {
		return nil, err
	}
	for _, customer := range customers {
		if id == customer.ID {
			return customer, nil
		}
	}
	return nil, ErrCustomerNotFound
}

func (c *client) Customers() ([]*customer.Customer, error) {
	url := c.url + "/customers"
	response, err := http.Get(url)
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
	body := handler.CustomersResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}
