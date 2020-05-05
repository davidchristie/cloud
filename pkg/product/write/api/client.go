package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/product/write/api/handler"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Client interface {
	CreateProduct(name, description string, correlationID uuid.UUID) (*entity.Product, error)
}

type client struct {
	URL string
}

type clientSpecification struct {
	URL string `required:"true"`
}

func NewClient() Client {
	spec := clientSpecification{}
	envconfig.MustProcess("PRODUCT_WRITE_API", &spec)
	return &client{
		URL: spec.URL,
	}
}

func (c *client) CreateProduct(name, description string, correlationID uuid.UUID) (*entity.Product, error) {
	requestBodyBytes, err := json.Marshal(&handler.CreateProductRequestBody{
		CorrelationID: correlationID,
		Description:   description,
		Name:          name,
	})
	if err != nil {
		return nil, err
	}
	url := c.URL + "/products"
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
	body, err := unmarshalCreateProductResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}

func unmarshalCreateProductResponseBody(response *http.Response) (*handler.CreateProductResponseBody, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := handler.CreateProductResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
