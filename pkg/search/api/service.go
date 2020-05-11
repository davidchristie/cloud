package api

import (
	"net/http"

	"github.com/davidchristie/cloud/pkg/search/api/core"
	"github.com/davidchristie/cloud/pkg/search/api/router"
	"github.com/kelseyhightower/envconfig"
)

type serviceSpecification struct {
	Port string `default:"8080"`
}

func StartService() error {
	spec := serviceSpecification{}
	envconfig.MustProcess("", &spec)
	c := core.NewCore()
	r := router.NewRouter(c)
	return http.ListenAndServe(":"+spec.Port, r)
}
