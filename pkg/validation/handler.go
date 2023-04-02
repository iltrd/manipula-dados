package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iltrd/manipula-dados/internal/csv"
	"github.com/iltrd/manipula-dados/internal/database/postgres"
	"github.com/iltrd/manipula-dados/pkg/validation"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (h *UploadHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	// verifica se o método HTTP é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método HTTP inválido", http.StatusMethodNotAllowed)
		return
	}

	// lê o arquivo enviado pelo usuário
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// faz o parse do arquivo CSV/TXT
	data, err := csv.ParseCSV(file)
	if err != nil {
		http.Error(w, "Erro ao fazer o parse do arquivo", http.StatusBadRequest)
		return
	}

	// valida os CPFs contidos no arquivo
	for _, datum := range data {
		if !validation.IsValidCPF(datum.CPF) {
			http.Error(w, "CPF inválido", http.StatusBadRequest)
			return
		}
	}

	// insere os dados no banco de dados
	for _, datum := range data {
		err = postgres.InsertData(datum)
		if err != nil {
			http.Error(w, "Erro ao inserir os dados no banco de dados", http.StatusInternalServerError)
			return
		}
	}

	// retorna a resposta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Dados inseridos com sucesso",
	})
}
