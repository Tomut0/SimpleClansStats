package server

import (
	"log"
	"net/http"

	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RunHttpServer(serverHandler *presentation.ServerHandler) {
	r := chi.NewRouter()
	
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	swaggerHandler := presentation.Handler(serverHandler)
	r.Mount("/api/v1", swaggerHandler)

	log.Println("API сервер запущен на :4000...")
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}