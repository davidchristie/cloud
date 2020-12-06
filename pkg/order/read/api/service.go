package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/davidchristie/cloud/pkg/order/core"
	"github.com/davidchristie/cloud/pkg/order/database"
	"github.com/davidchristie/cloud/pkg/order/read/api/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartService() error {
	orderRepository := database.NewOrderRepository(database.Connect())
	c := core.NewCore(orderRepository)
	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.OrdersHandler(c)).Methods("GET")
	r.HandleFunc("/orders/{id}", handler.OrderHandler(c)).Methods("GET")
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	fmt.Println("serving order-read-api...")
	return http.ListenAndServe(":8080", loggedRouter)
}
