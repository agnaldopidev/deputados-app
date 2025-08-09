package main

import (
	"database/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/agnaldopidev/deputados-app/internal/graph"
	"github.com/agnaldopidev/deputados-app/internal/graph/generated"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

// getEnv retorna o valor da variável ou um padrão se não existir
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {

	// Tenta carregar o arquivo .env (apenas se existir)
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente")
	}

	// Montar string de conexão
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "teste123"),
		getEnv("DB_NAME", "deputadosdb"),
	)

	// Conectar no PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}

	defer db.Close()

	r := &graph.Resolver{
		DB: db,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	port := os.Getenv("GRAPHQL_PORT")
	log.Println("🚀 servidor iniciado em http://localhost:8081/")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
