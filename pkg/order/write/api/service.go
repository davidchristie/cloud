package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order/write/api/core"
	"github.com/davidchristie/cloud/pkg/order/write/api/handler"
)

func StartService() {
	c := core.NewCore()

	// Add handle func for producer.
	http.HandleFunc("/orders", handler.CreateOrderHandler(c))

	// Run the web server.
	fmt.Println("serving order-write-api...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
