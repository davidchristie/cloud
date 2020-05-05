package api

import (
	"fmt"
	"log"
	"net/http"

	orderDatabase "github.com/davidchristie/cloud/pkg/order/database"
	"github.com/davidchristie/cloud/pkg/order/read/api/handler"
)

func StartService() {
	db := orderDatabase.Connect()

	orderRepository := orderDatabase.NewOrderRepository(db)

	http.HandleFunc("/orders", handler.OrdersHandler(orderRepository))

	fmt.Println("serving order-read-api...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
