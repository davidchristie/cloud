package api

import (
	"github.com/davidchristie/cloud/pkg/http"
	"github.com/davidchristie/cloud/pkg/order/database"
	"github.com/davidchristie/cloud/pkg/order/read/api/core"
	"github.com/davidchristie/cloud/pkg/order/read/api/handler"
	"github.com/davidchristie/cloud/pkg/router"
)

func StartService() error {
	c := core.NewCore(database.NewOrderRepository(database.Connect()))
	r := router.NewRouter()
	r.HandleFunc("/orders", handler.OrdersHandler(c)).Methods("GET")
	r.HandleFunc("/orders/{id}", handler.OrderHandler(c)).Methods("GET")
	return http.ListenAndServe(r)
}
