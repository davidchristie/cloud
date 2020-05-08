package api

import (
	"net/http"

	"github.com/davidchristie/cloud/pkg/customer/write/api/core"
	"github.com/davidchristie/cloud/pkg/customer/write/api/handler"
)

func StartService() error {
	c := core.NewCore()

	http.HandleFunc("/customers", handler.CreateCustomerHandler(c))

	return http.ListenAndServe(":8080", nil)
}
