package handler

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order"
	orderDatabase "github.com/davidchristie/cloud/pkg/order/database"
)

type OrdersResponseBody struct {
	Data    []*order.Order `json:"data"`
	Message string         `json:"message"`
}

func OrdersHandler(orderRepository orderDatabase.OrderRepository) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		orders, err := orderRepository.GetOrders(ctx)

		writer.Header().Add("Content-Type", "application/json")

		if err != nil {
			response := &OrdersResponseBody{
				Message: err.Error(),
			}
			data, _ := json.Marshal(response)
			writer.WriteHeader(500)
			writer.Write(data)
		} else {
			response := &OrdersResponseBody{
				Data: orders,
			}
			data, _ := json.Marshal(response)
			writer.Write(data)
		}
	})
}
