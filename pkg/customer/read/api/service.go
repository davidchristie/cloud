package api

import (
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer/read/api/core"
	"github.com/davidchristie/cloud/pkg/customer/read/api/router"
)

func StartService() error {
	c := core.NewCore()

	r := router.NewRouter(c)

	return http.ListenAndServe(":8080", r)
}
