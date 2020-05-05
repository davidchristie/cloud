package api

import (
	"fmt"
	"net/http"

	"github.com/davidchristie/cloud/pkg/product/write/api/core"
	"github.com/davidchristie/cloud/pkg/product/write/api/handler"
	"github.com/kelseyhightower/envconfig"
)

type serviceSpecification struct {
	Port string `default:"8080"`
}

func StartService() error {
	spec := serviceSpecification{}
	envconfig.MustProcess("", &spec)

	c := core.NewCore()

	http.HandleFunc("/products", handler.CreateProductHandler(c))

	fmt.Println("serving")

	address := ":" + spec.Port

	return http.ListenAndServe(address, nil)
}
