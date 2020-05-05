package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type OrderWriteAPIClient interface {
	CreateOrder(*CreateOrderInput) (*order.Order, error)
}

type CreateOrderInput struct {
	CorrelationID uuid.UUID
	CustomerID    uuid.UUID
	LineItems     []*order.LineItem
}

type client struct {
	URL string
}

type createOrderResponseBody struct {
	Data    *order.Order `json:"data"`
	Message string       `json:"message"`
}

type createOrderRequestBody struct {
	CorrelationID uuid.UUID         `json:"correlation_id"`
	CustomerID    uuid.UUID         `json:"customer_id"`
	LineItems     []*order.LineItem `json:"line_items"`
}

type specification struct {
	URL string `required:"true"`
}

func NewClient() OrderWriteAPIClient {
	spec := specification{}
	envconfig.MustProcess("ORDER_WRITE_API", &spec)
	return &client{
		URL: spec.URL,
	}
}

func (c *client) CreateOrder(input *CreateOrderInput) (*order.Order, error) {
	requestBodyBytes, err := json.Marshal(&createOrderRequestBody{
		CorrelationID: input.CorrelationID,
		CustomerID:    input.CustomerID,
		LineItems:     input.LineItems,
	})
	if err != nil {
		return nil, err
	}
	url := c.URL + "/orders"
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
	body, err := unmarshalCreateOrderResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}

func unmarshalCreateOrderResponseBody(response *http.Response) (*createOrderResponseBody, error) {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := createOrderResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
