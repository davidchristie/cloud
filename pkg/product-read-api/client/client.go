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

type ProductReadAPIClient interface {
	Products() ([]*entity.Product, error)
}

type client struct {
	url string
}

type productsResponseBody struct {
	Data    *[]*entity.Product `json:"data"`
	Message string             `json:"message"`
}

type specification struct {
	URL string `required:"true"`
}

func NewClient() ProductReadAPIClient {
	spec := specification{}
	envconfig.MustProcess("PRODUCT_READ_API", &spec)
	return &client{
		url: spec.URL,
	}
}

func (c *client) Products() ([]*entity.Product, error) {
	url := c.url + "/products"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	body, err := unmarshalProductsResponseBody(response)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return *body.Data, nil
}

func unmarshalProductsResponseBody(response *http.Response) (*productsResponseBody, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := productsResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Fatal(err)
	}
	return &body, nil
}
