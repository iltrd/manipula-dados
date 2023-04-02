package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iltrd/manipula-dados/internal/handler"
	"github.com/iltrd/manipula-dados/internal/middleware"
	"github.com/iltrd/manipula-dados/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	// carrega as variáveis de ambiente a partir do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar as variáveis de ambiente")
	}

	// cria o roteador da aplicação
	r := mux.NewRouter()

	// cria os handlers da aplicação
	uploadHandler := handler.NewUploadHandler()

	// adiciona os middlewares da aplicação
	r.Use(middleware.LoggerMiddleware)

	// adiciona as rotas da aplicação
	router.UploadRoutes(r, uploadHandler)

	// inicia o servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Servidor iniciado na porta %s\n", port)
	http.ListenAndServe(addr, r)
}
