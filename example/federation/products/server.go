//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thienohs/gqlgen/example/federation/products/graph"
	"github.com/thienohs/gqlgen/example/federation/products/graph/generated"
	"github.com/thienohs/gqlgen/graphql/handler"
	"github.com/thienohs/gqlgen/graphql/handler/debug"
	"github.com/thienohs/gqlgen/graphql/playground"
)

const defaultPort = "4002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
