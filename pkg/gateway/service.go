//go:generate go run github.com/99designs/gqlgen generate

package gateway

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davidchristie/cloud/pkg/gateway/graph"
	"github.com/davidchristie/cloud/pkg/gateway/graph/generated"
	"github.com/kelseyhightower/envconfig"
)

type serviceSpecification struct {
	Port string `default:"8080"`
}

func StartService() error {
	spec := serviceSpecification{}
	envconfig.MustProcess("", &spec)

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", spec.Port)
	return http.ListenAndServe(":"+spec.Port, nil)
}
