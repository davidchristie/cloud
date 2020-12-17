package api

import (
	"github.com/davidchristie/cloud/pkg/http"
	"github.com/davidchristie/cloud/pkg/product/read/api/core"
	"github.com/davidchristie/cloud/pkg/product/read/api/handler"
	"github.com/davidchristie/cloud/pkg/router"
)

func StartService() error {
	c := core.NewCore()
	r := router.NewRouter()
	r.HandleFunc("/products", handler.ProductsHandler(c)).Methods("GET")
	r.HandleFunc("/products/{id}", handler.ProductHandler(c)).Methods("GET")
	return http.ListenAndServe(r)
}
