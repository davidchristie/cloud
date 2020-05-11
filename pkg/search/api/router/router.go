package router

import (
	"net/http"
	"os"

	"github.com/davidchristie/cloud/pkg/search/api/core"
	"github.com/davidchristie/cloud/pkg/search/api/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(c core.Core) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/products", handler.ProductsHandler(c)).Methods("GET")
	return handlers.LoggingHandler(os.Stdout, r)
}
