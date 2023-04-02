package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iltrd/manipula-dados/internal/handler"
)

func UploadRoutes(r *mux.Router, uploadHandler *handler.UploadHandler) {
	r.HandleFunc("/upload", uploadHandler.HandleUpload).Methods(http.MethodPost)
}
