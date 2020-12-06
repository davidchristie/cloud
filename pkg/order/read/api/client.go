package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/read/api/handler"
	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type OrderReadAPIClient interface {
	Order(uuid.UUID) (*order.Order, error)
	Orders(customerID *string) ([]*order.Order, error)
}

type OrdersOptions struct {
	CustomerID *string `url:"customer_id"`
}

type client struct {
	url string
}

type specification struct {
	URL string `required:"true"`
}

func NewClient() OrderReadAPIClient {
	spec := specification{}
	envconfig.MustProcess("ORDER_READ_API", &spec)
	return &client{
		url: spec.URL,
	}
}

func (c *client) Order(id uuid.UUID) (*order.Order, error) {
	url := c.url + "/orders/" + id.String()
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	body, err := unmarshalOrderResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}

func (c *client) Orders(customerID *string) ([]*order.Order, error) {
	opt := OrdersOptions{
		CustomerID: customerID,
	}
	v, _ := query.Values(opt)
	url := c.url + "/orders?" + v.Encode()
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	body, err := unmarshalOrdersResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}

func unmarshalOrderResponseBody(response *http.Response) (*handler.OrderResponseBody, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := handler.OrderResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func unmarshalOrdersResponseBody(response *http.Response) (*handler.OrdersResponseBody, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := handler.OrdersResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
