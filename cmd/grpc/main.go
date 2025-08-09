package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/agnaldopidev/deputados-app/internal/grpc/proto"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"

	"github.com/agnaldopidev/deputados-app/internal/repository"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	deputadopb.UnimplementedDeputadoServiceServer
	repo repository.DeputadoRepository
}

func (s *server) ListDeputados(ctx context.Context, _ *deputadopb.Empty) (*deputadopb.DeputadoList, error) {
	deputados, err := s.repo.ListDeputados()
	if err != nil {
		return nil, err
	}

	var grpcDeputados []*deputadopb.Deputado
	for _, o := range deputados {
		grpcDeputados = append(grpcDeputados, &deputadopb.Deputado{
			Id:          o.ID,
			Nome:        o.Nome,
			Partido:     o.Partido,
			NumeroVotos: o.Votos,
		})
	}

	return &deputadopb.DeputadoList{Deputados: grpcDeputados}, nil
}

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

	repo := repository.NewDeputadoRepository(db)

	port := os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Erro ao escutar:", err)
	}

	s := grpc.NewServer()
	deputadopb.RegisterDeputadoServiceServer(s, &server{repo: repo})

	// Ativa suporte a reflection (para usar grpcurl, Postman, etc)
	reflection.Register(s)

	log.Println("gRPC rodando na porta 50051 (DeputadoService)...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
