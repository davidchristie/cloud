package router

import (
	"net/http"
	"os"

	"github.com/davidchristie/cloud/pkg/customer/read/api/core"
	"github.com/davidchristie/cloud/pkg/customer/read/api/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(c core.Core) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/customers", handler.CustomersHandler(c)).Methods("GET")
	r.HandleFunc("/customers/{id}", handler.CustomerHandler(c)).Methods("GET")
	return handlers.LoggingHandler(os.Stdout, r)
}
