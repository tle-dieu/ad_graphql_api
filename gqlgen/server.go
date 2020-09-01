package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tle-dieu/gql_test/gqlgen/graph"
	"github.com/tle-dieu/gql_test/gqlgen/graph/generated"
	"github.com/tle-dieu/gql_test/gqlgen/pkg/db/mysql"
)

const defaultPort = "8080"

func main() {
	mysqlClient := mysql.NewMySQLClient()
	mysqlClient.Migrate()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := "8080"
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}