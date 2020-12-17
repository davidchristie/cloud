package api

import (
	"github.com/davidchristie/cloud/pkg/http"
	"github.com/davidchristie/cloud/pkg/order/write/api/core"
	"github.com/davidchristie/cloud/pkg/order/write/api/handler"
	"github.com/davidchristie/cloud/pkg/router"
)

func StartService() error {
	c := core.NewCore()
	r := router.NewRouter()
	r.HandleFunc("/orders", handler.CreateOrderHandler(c))
	return http.ListenAndServe(r)
}
