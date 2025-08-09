package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	_ "os"

	"github.com/agnaldopidev/deputados-app/internal/handler"
	"github.com/agnaldopidev/deputados-app/internal/repository"
	_ "github.com/lib/pq"
)

// getEnv retorna o valor da variável ou um padrão se não existir
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// Carregar .env

	// Tenta carregar o arquivo .env (apenas se existir)
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente")
	}

	// Monta a string de conexão usando variáveis de ambiente
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "teste123"),
		getEnv("DB_NAME", "deputadosdb"),
	)

	fmt.Println("String de conexão:", dsn)
	// Conectar no PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}

	defer db.Close()

	repo := repository.NewDeputadoRepository(db)
	mux := handler.NovoRouter(repo)

	port := os.Getenv("REST_PORT")
	log.Println("REST server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
