package api

import (
	"github.com/davidchristie/cloud/pkg/http"
	"github.com/davidchristie/cloud/pkg/product/write/api/core"
	"github.com/davidchristie/cloud/pkg/product/write/api/handler"
	"github.com/davidchristie/cloud/pkg/router"
)

func StartService() error {
	c := core.NewCore()
	r := router.NewRouter()
	r.HandleFunc("/products", handler.CreateProductHandler(c)).Methods("POST")
	r.HandleFunc("/products/{id}", handler.DeleteProductHandler(c)).Methods("DELETE")
	return http.ListenAndServe(r)
}
