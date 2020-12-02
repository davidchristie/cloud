package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order"
	orderDatabase "github.com/davidchristie/cloud/pkg/order/database"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type OrderResponseBody struct {
	Data    *order.Order `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
}

func OrderHandler(orderRepository orderDatabase.OrderRepository) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := OrderResponseBody{}

		writer.Header().Add("Content-Type", "application/json")

		id := mux.Vars(request)["id"]

		orderID, err := uuid.Parse(id)
		if err != nil {
			writer.WriteHeader(404)
			response.Message = "Not Found"
			return
		}

		order, err := orderRepository.FindOrder(request.Context(), orderID)

		switch err {
		case nil:
			response.Data = order
			break

		case orderDatabase.ErrOrderNotFound:
			writer.WriteHeader(404)
			response.Message = "Not Found"
			break

		default:
			log.Println(err)
			writer.WriteHeader(500)
			response.Message = "Internal Server Error"
			break
		}

		data, _ := json.Marshal(response)
		writer.Write(data)
	})
}
