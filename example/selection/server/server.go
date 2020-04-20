package main

import (
	"log"
	"net/http"

	"github.com/thienohs/gqlgen/example/selection"
	"github.com/thienohs/gqlgen/graphql/handler"
	"github.com/thienohs/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
