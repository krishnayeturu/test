package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/go-graphql/graph"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/go-graphql/graph/generated"
)

const defaultPort = "8080"

func envLoad() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
}

func main() {
	envLoad()

	graph.ConnectDB()
	defer graph.CloseDB()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	// Set up the server struct
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(transport.Options{
		AllowedMethods: []string{},
	})
	if os.Getenv("ENVIRONMENT") == "development" {
		srv.Use(extension.Introspection{})
	}

	// Enable query endpoint
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Expose query endpoint to localhost
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
