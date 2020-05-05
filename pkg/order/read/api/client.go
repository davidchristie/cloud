package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/read/api/handler"
	"github.com/kelseyhightower/envconfig"
)

type OrderReadAPIClient interface {
	Orders() ([]*order.Order, error)
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

func (c *client) Orders() ([]*order.Order, error) {
	url := c.url + "/orders"
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
