package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/product/read/api/handler"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Client interface {
	Product(uuid.UUID) (*entity.Product, error)
	Products() ([]*entity.Product, error)
}

type client struct {
	url string
}

type clientSpecification struct {
	URL string `required:"true"`
}

var ErrProductNotFound = errors.New("product not found")

func NewClient() Client {
	spec := clientSpecification{}
	envconfig.MustProcess("PRODUCT_READ_API", &spec)
	return &client{
		url: spec.URL,
	}
}

// TODO: Implement this properly.
func (c *client) Product(id uuid.UUID) (*entity.Product, error) {
	customers, err := c.Products()
	if err != nil {
		return nil, err
	}
	for _, customer := range customers {
		if id == customer.ID {
			return customer, nil
		}
	}
	return nil, ErrProductNotFound
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
	return body.Data, nil
}

func unmarshalProductsResponseBody(response *http.Response) (*handler.ProductsResponseBody, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := &handler.ProductsResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
