package api

import (
	"github.com/davidchristie/cloud/pkg/customer/write/api/core"
	"github.com/davidchristie/cloud/pkg/customer/write/api/handler"
	"github.com/davidchristie/cloud/pkg/http"
	"github.com/davidchristie/cloud/pkg/router"
)

func StartService() error {
	c := core.NewCore()
	r := router.NewRouter()
	r.HandleFunc("/customers", handler.CreateCustomerHandler(c))
	return http.ListenAndServe(r)
}
