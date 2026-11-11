package server

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func RunHttpServer(serverHandler *presentation.ServerHandler, sessions *scs.SessionManager) {
	srv := &http.Server{
		Addr:           ":4000",
		Handler:        newRouter(serverHandler, sessions),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	log.Println("API сервер запущен на :4000...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func newRouter(serverHandler *presentation.ServerHandler, sessions *scs.SessionManager) http.Handler {
	r := chi.NewRouter()

	r.Use(func(h http.Handler) http.Handler {
		settingsUpload := http.MaxBytesHandler(h, presentation.MaxSettingsRequestBytes)
		defaultLimit := http.MaxBytesHandler(h, 4096)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost && r.URL.Path == "/api/v1/ingest" {
				h.ServeHTTP(w, r)
				return
			}
			if r.Method == http.MethodPost && r.URL.Path == "/api/v1/settings" {
				settingsUpload.ServeHTTP(w, r)
				return
			}
			defaultLimit.ServeHTTP(w, r)
		})
	})

	// Rate limiting requests by 100 per minute
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(httprate.LimitBy(100, time.Minute, func(r *http.Request) (string, error) {
		return httprate.CanonicalizeIP(middleware.GetClientIP(r.Context())), nil
	}))

	// Timeout any requests after 60s
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS+CSRF protection
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
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Prefer", "Authorization"},
		AllowCredentials: true,
	}))

	// Log any request and recover from panic
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Store sessions between requests
	r.Use(sessions.LoadAndSave)

	api := chi.NewRouter()
	api.Post("/ingest", serverHandler.Ingest)
	api.Get("/dashboard/events", serverHandler.DashboardEvents)
	api.Get("/dashboard/layout", serverHandler.GetDashboardLayout)
	api.Put("/dashboard/layout", serverHandler.SaveDashboardLayout)
	r.Mount("/api/v1", presentation.HandlerFromMux(presentation.NewStrictHandler(serverHandler, nil), api))
	return r
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
