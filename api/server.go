package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/graph"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/graph/generated"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/database"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dynamoDbClient := database.CreateDynamoDBClient()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DbClient: dynamoDbClient}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
