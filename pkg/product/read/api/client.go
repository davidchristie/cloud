package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
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

func (c *client) Product(id uuid.UUID) (*entity.Product, error) {
	url := c.url + "/products/" + id.String()
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 404 {
		return nil, ErrProductNotFound
	}
	if response.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid response content type: " + response.Header.Get("Content-Type"))
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body := handler.ProductResponseBody{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(body.Message)
	}
	return body.Data, nil
}

func (c *client) Products() ([]*entity.Product, error) {
	url := c.url + "/products"
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
	body := handler.ProductsResponseBody{}
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
