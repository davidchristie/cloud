package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/search/api/handler"
	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Client interface {
	Customers(query string) ([]uuid.UUID, error)
	Products(query string) ([]uuid.UUID, error)
}

type client struct {
	url string
}

type clientSpecification struct {
	URL string `required:"true"`
}

type customersOptions struct {
	Query string `url:"q"`
}

type productsOptions struct {
	Query string `url:"q"`
}

func NewClient() Client {
	spec := clientSpecification{}
	envconfig.MustProcess("SEARCH_API", &spec)
	return &client{
		url: spec.URL,
	}
}

func (c *client) Customers(searchQuery string) ([]uuid.UUID, error) {
	opt := customersOptions{
		Query: searchQuery,
	}
	v, _ := query.Values(&opt)
	url := c.url + "/customers?" + v.Encode()
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

func (c *client) Products(searchQuery string) ([]uuid.UUID, error) {
	opt := productsOptions{
		Query: searchQuery,
	}
	v, _ := query.Values(&opt)
	url := c.url + "/products?" + v.Encode()
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
