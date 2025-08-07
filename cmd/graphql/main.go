package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/agnaldopidev/deputados-app/internal/graph"
	"github.com/agnaldopidev/deputados-app/internal/graph/generated"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=agnaldo password=teste123 dbname=deputadosdb sslmode=disable")
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	defer db.Close()

	r := &graph.Resolver{
		DB: db,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("🚀 servidor iniciado em http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
