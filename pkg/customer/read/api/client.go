package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/davidchristie/cloud/pkg/customer"
	"github.com/davidchristie/cloud/pkg/customer/read/api/handler"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Client interface {
	Customer(uuid.UUID) (*customer.Customer, error)
	Customers([]uuid.UUID) (map[uuid.UUID]*customer.Customer, error)
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

func (c *client) Customer(id uuid.UUID) (*customer.Customer, error) {
	url := c.url + "/customers/" + id.String()
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 404 {
		return nil, ErrCustomerNotFound
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := handler.CustomerResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}

func (c *client) Customers(ids []uuid.UUID) (map[uuid.UUID]*customer.Customer, error) {
	url := c.url + "/customers?ids=" + joinIDs(ids)
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if response.Header.Get("Content-Type") != "application/json" {
		err = errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
		log.Println(err)
		return nil, err
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	body := handler.CustomersResponseBody{}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if response.StatusCode != 200 {
		err = errors.New(body.Message)
		log.Println(err)
		return nil, err
	}
	return body.Data, nil
}

func joinIDs(ids []uuid.UUID) string {
	s := make([]string, len(ids))
	for i, id := range ids {
		s[i] = id.String()
	}
	return strings.Join(s, ",")
}
