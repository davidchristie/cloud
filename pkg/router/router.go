package router

import (
	"github.com/davidchristie/cloud/pkg/router/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logging)
	return r
}
