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

	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
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
