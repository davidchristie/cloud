package handler

import (
	"encoding/json"
	"net/http"

	"github.com/davidchristie/cloud/pkg/order"
	"github.com/davidchristie/cloud/pkg/order/core"
	"github.com/google/uuid"
)

type OrdersResponseBody struct {
	Data    []*order.Order `json:"data"`
	Message string         `json:"message"`
}

func OrdersHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		orders, err := c.Orders(ctx, core.OrdersInput{
			CustomerID: parseUUID(request.URL.Query().Get("customer_id")),
		})

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

func parseUUID(s string) *uuid.UUID {
	u, err := uuid.Parse(s)
	if err != nil {
		return nil
	}
	return &u
}
