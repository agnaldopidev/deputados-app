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

func main() {
	// Carregar .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	// Montar string de conexão
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

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
