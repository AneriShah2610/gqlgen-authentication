package main

import (
	log "log"
	http "net/http"
	os "os"

	"github.com/aneri/gqlgen-authentication/api"

	"github.com/gorilla/mux"

	handler "github.com/99designs/gqlgen/handler"
	gqlgen_authentication "github.com/aneri/gqlgen-authentication"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := mux.NewRouter()
	router.Use(api.MiddleWareHandler)
	router.Use(api.AuthHandlerMiddleware)
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(gqlgen_authentication.NewExecutableSchema(gqlgen_authentication.Config{Resolvers: &gqlgen_authentication.Resolver{}})))

	log.Printf("connect to   http://localhost:1234/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
