package repository

import (
	"database/sql"
	"log"

	"github.com/agnaldopidev/deputados-app/internal/domain"
)

type DeputadoRepository interface {
	ListDeputados() ([]domain.Deputado, error)
	CreateDeputados(order domain.Deputado) error
}

type deputadoRepository struct {
	db *sql.DB
}

func NewDeputadoRepository(db *sql.DB) DeputadoRepository {
	return &deputadoRepository{db: db}
}

func (r *deputadoRepository) ListDeputados() ([]domain.Deputado, error) {
	sql := `SELECT id, nome, partido, votos FROM deputados ORDER BY votos DESC`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Deputado
	for rows.Next() {
		var o domain.Deputado
		if err := rows.Scan(&o.ID, &o.Nome, &o.Partido, &o.Votos); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *deputadoRepository) CreateDeputados(deputado domain.Deputado) error {
	sql := `INSERT INTO deputados (nome, partido, votos) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(sql, deputado.Nome, deputado.Partido, deputado.Votos)
	if err != nil {
		log.Println("Erro ao inserir:", err)
	}
	return err
}
