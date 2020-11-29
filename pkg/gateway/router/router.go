package router

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	customerReadAPI "github.com/davidchristie/cloud/pkg/customer/read/api"
	"github.com/davidchristie/cloud/pkg/gateway/graph"
	"github.com/davidchristie/cloud/pkg/gateway/graph/dataloader"
	"github.com/davidchristie/cloud/pkg/gateway/graph/generated"
	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(customerReadAPI customerReadAPI.Client, productReadAPI productReadAPI.Client) http.Handler {
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver()}))
	r := mux.NewRouter()
	r.Handle("/", playground.Handler("GraphQL playground", "/query")).Methods("GET")
	r.Handle("/query", server).Methods("POST")
	r.Use(dataloader.Middleware(customerReadAPI, productReadAPI))
	return handlers.LoggingHandler(os.Stdout, r)
}
