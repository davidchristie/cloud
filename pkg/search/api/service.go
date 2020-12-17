package api

import (
	"github.com/davidchristie/cloud/pkg/http"
	"github.com/davidchristie/cloud/pkg/router"
	"github.com/davidchristie/cloud/pkg/search/api/core"
	"github.com/davidchristie/cloud/pkg/search/api/handler"
)

func StartService() error {
	c := core.NewCore()
	r := router.NewRouter()
	r.HandleFunc("/customers", handler.CustomersHandler(c)).Methods("GET")
	r.HandleFunc("/products", handler.ProductsHandler(c)).Methods("GET")
	return http.ListenAndServe(r)
}
