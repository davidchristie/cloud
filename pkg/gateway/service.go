//go:generate go run github.com/99designs/gqlgen generate

package gateway

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	customerReadAPI "github.com/davidchristie/cloud/pkg/customer/read/api"
	"github.com/davidchristie/cloud/pkg/gateway/graph"
	"github.com/davidchristie/cloud/pkg/gateway/graph/dataloader"
	"github.com/davidchristie/cloud/pkg/gateway/graph/generated"
	"github.com/davidchristie/cloud/pkg/http"
	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
	"github.com/davidchristie/cloud/pkg/router"
)

func StartService() error {
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver()}))
	r := router.NewRouter()
	r.Handle("/", playground.Handler("GraphQL playground", "/query")).Methods("GET")
	r.Handle("/query", server).Methods("POST")
	r.Use(dataloader.Middleware(customerReadAPI.NewClient(), productReadAPI.NewClient()))
	return http.ListenAndServe(r)
}
