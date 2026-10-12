package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func RunHttpServer(serverHandler *presentation.ServerHandler) {
	r := chi.NewRouter()

	crossOriginProtection := http.NewCrossOriginProtection()
	allowedOrigins := corsAllowedOrigins()
	for _, origin := range allowedOrigins {
		if err := crossOriginProtection.AddTrustedOrigin(origin); err != nil {
			log.Fatalf("Incorrect CORS origin %q: %v", origin, err)
		}
	}
	r.Use(crossOriginProtection.Handler)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Prefer", "Authorization"},
		AllowCredentials: true,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	swaggerHandler := presentation.Handler(serverHandler)
	r.Mount("/api/v1", swaggerHandler)

	log.Println("API сервер запущен на :4000...")
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func corsAllowedOrigins() []string {
	value := os.Getenv("CORS_ALLOWED_ORIGINS")
	if value == "" {
		return []string{"http://localhost:5173"}
	}

	var origins []string
	for _, origin := range strings.Split(value, ",") {
		if origin = strings.TrimSpace(origin); origin != "" {
			origins = append(origins, origin)
		}
	}
	return origins
}
