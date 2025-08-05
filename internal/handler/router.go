package handler

import (
	"encoding/json"
	"net/http"

	"github.com/agnaldopidev/deputados-app/internal/domain"
	"github.com/agnaldopidev/deputados-app/internal/repository"
)

func NovoRouter(repo repository.DeputadoRepository) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/deputados", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			deputado, err := repo.ListDeputados()
			if err != nil {
				http.Error(w, "Erro ao listar", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(deputado)

		case http.MethodPost:
			var deputado domain.Deputado
			if err := json.NewDecoder(r.Body).Decode(&deputado); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}
			if err := repo.CreateDeputados(deputado); err != nil {
				http.Error(w, "Erro ao salvar", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
