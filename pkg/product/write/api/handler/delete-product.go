package handler

import (
	"net/http"

	"github.com/davidchristie/cloud/pkg/product/write/api/core"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func DeleteProductHandler(c core.Core) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		corrlationID, err := uuid.Parse(req.Header.Get("correlation_id"))
		if err != nil {
			wrt.WriteHeader(400)
			return
		}
		productID, err := uuid.Parse(mux.Vars(req)["id"])
		if err != nil {
			wrt.WriteHeader(400)
			return
		}
		c.DeleteProduct(&core.DeleteProductInput{
			Context:        req.Context(),
			CorreleationID: corrlationID,
			ProductID:      productID,
		})
	})
}
