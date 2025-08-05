package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "os"

	"github.com/agnaldopidev/deputados-app/internal/handler"
	"github.com/agnaldopidev/deputados-app/internal/repository"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=agnaldo password=teste123 dbname=deputados_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewDeputadoRepository(db)
	mux := handler.NovoRouter(repo)

	port := "8080"
	log.Println("REST server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
