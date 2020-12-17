package api

import (
	"github.com/davidchristie/cloud/pkg/customer/read/api/core"
	"github.com/davidchristie/cloud/pkg/customer/read/api/handler"
	"github.com/davidchristie/cloud/pkg/http"
	"github.com/davidchristie/cloud/pkg/router"
)

func StartService() error {
	c := core.NewCore()
	r := router.NewRouter()
	r.HandleFunc("/customers", handler.CustomersHandler(c)).Methods("GET")
	r.HandleFunc("/customers/{id}", handler.CustomerHandler(c)).Methods("GET")
	return http.ListenAndServe(r)
}
