package api

import (
	"fmt"
	"net/http"

	orderDatabase "github.com/davidchristie/cloud/pkg/order/database"
	"github.com/davidchristie/cloud/pkg/order/read/api/handler"
	"github.com/gorilla/mux"
)

func StartService() error {
	db := orderDatabase.Connect()

	orderRepository := orderDatabase.NewOrderRepository(db)

	r := mux.NewRouter()

	r.HandleFunc("/orders", handler.OrdersHandler(orderRepository)).Methods("GET")
	r.HandleFunc("/orders/{id}", handler.OrderHandler(orderRepository)).Methods("GET")

	fmt.Println("serving order-read-api...")
	return http.ListenAndServe(":8080", r)
}
