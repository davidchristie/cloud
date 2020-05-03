package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type ProductWriteAPIClient interface {
	CreateProduct(name, description string, correlationID uuid.UUID) (*entity.Product, error)
}

type client struct {
	URL string
}

type createProductResponseBody struct {
	Data    *entity.Product `json:"data"`
	Message string          `json:"message"`
}

type createProductRequestBody struct {
	CorrelationID uuid.UUID `json:"correlation_id"`
	Description   string    `json:"description"`
	Name          string    `json:"name"`
}

type specification struct {
	URL string `required:"true"`
}

func NewClient() ProductWriteAPIClient {
	spec := specification{}
	envconfig.MustProcess("PRODUCT_WRITE_API", &spec)
	return &client{
		URL: spec.URL,
	}
}

func (c *client) CreateProduct(name, description string, correlationID uuid.UUID) (*entity.Product, error) {
	requestBodyBytes, err := json.Marshal(&createProductRequestBody{
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

func unmarshalCreateProductResponseBody(response *http.Response) (*createProductResponseBody, error) {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := createProductResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
