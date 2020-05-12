//go:generate go run github.com/99designs/gqlgen generate

package gateway

import (
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/gateway/router"
	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
	"github.com/kelseyhightower/envconfig"
)

type serviceSpecification struct {
	Port string `default:"8080"`
}

func StartService() error {
	spec := serviceSpecification{}
	envconfig.MustProcess("", &spec)
	r := router.NewRouter(productReadAPI.NewClient())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", spec.Port)
	return http.ListenAndServe(":"+spec.Port, r)
}
