package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ip75/media.server.go/graph"
	"github.com/ip75/media.server.go/graph/generated"

	//	"github.com/99designs/gqlgen/example/federation/products/graph/generated"
	//  "github.com/99designs/gqlgen/example/federation/reviews/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
