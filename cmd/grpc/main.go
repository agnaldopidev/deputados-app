package main

import (
	"context"
	"database/sql"
	"github.com/agnaldopidev/deputados-app/internal/grpc/proto"
	"log"
	"net"

	"github.com/agnaldopidev/deputados-app/internal/repository"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	deputadopb.UnimplementedDeputadoServiceServer
	repo repository.DeputadoRepository
}

func (s *server) ListDeputados(ctx context.Context, _ *deputadopb.deputadopb) (*deputadopb.DeputadoList, error) {
	orders, err := s.repo.ListDeputados()
	if err != nil {
		return nil, err
	}

	var grpcDeputados []*deputadopb.Deputado
	for _, o := range orders {
		grpcDeputados = append(grpcDeputados, &deputadopb.Deputado{
			Id:          o.ID,
			Nome:        o.Nome,
			Partido:     o.Partido,
			NumeroVotos: o.NumeroVotos,
		})
	}

	return &deputadopb.DeputadoList{Deputados: grpcDeputados}, nil
}

func main() {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=agnaldo password=teste123 dbname=deputados_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewDeputadoRepository(db)

	lis, err := net.Listen("tcp", ":50051")
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
