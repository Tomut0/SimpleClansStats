package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

const (
	responseWriteTimeout = 10 * time.Second
	themeApplyTimeout    = 5 * time.Minute
)

func RunHttpServer(serverHandler *presentation.ServerHandler, sessions *scs.SessionManager, installationID string, serveStatic bool) {
	srv := &http.Server{
		Addr:           ":4000",
		Handler:        newRouter(serverHandler, sessions, installationID, serveStatic),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   responseWriteTimeout,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	log.Println("API сервер запущен на :4000...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func newRouter(serverHandler *presentation.ServerHandler, sessions *scs.SessionManager, installationID string, serveStatic bool) http.Handler {
	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		standard := middleware.Timeout(60 * time.Second)(next)
		applyTheme := middleware.Timeout(themeApplyTimeout)(next)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost || r.URL.Path != "/api/v1/settings/theme" {
				standard.ServeHTTP(w, r)
				return
			}
			// Leave time to send a response even if the build reaches its deadline.
			deadline := time.Now().Add(themeApplyTimeout + responseWriteTimeout)
			if err := http.NewResponseController(w).SetWriteDeadline(deadline); err != nil && !errors.Is(err, http.ErrNotSupported) {
				log.Printf("set theme response deadline: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			applyTheme.ServeHTTP(w, r)
		})
	})

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

	// Rate limiting requests by 300 per minute
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(httprate.LimitBy(300, time.Minute, func(r *http.Request) (string, error) {
		return httprate.CanonicalizeIP(middleware.GetClientIP(r.Context())), nil
	}))

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
	if telemetry := newPosthogTelemetry(installationID); telemetry != nil {
		r.Use(func(next http.Handler) http.Handler {
			return telemetry.middleware(next, serverHandler.TelemetryAllowed)
		})
	}

	api := chi.NewRouter()
	api.Post("/ingest", serverHandler.Ingest)
	api.Get("/dashboard/events", serverHandler.DashboardEvents)
	api.With(httprate.LimitBy(120, time.Minute, func(r *http.Request) (string, error) {
		return httprate.CanonicalizeIP(middleware.GetClientIP(r.Context())), nil
	})).Get("/search", serverHandler.Search)
	api.Get("/dashboard/layout", serverHandler.GetDashboardLayout)
	api.Put("/dashboard/layout", serverHandler.SaveDashboardLayout)
	r.Mount("/api/v1", presentation.HandlerFromMux(presentation.NewStrictHandler(serverHandler, nil), api))
	if serveStatic {
		r.NotFound(serveFrontend)
	}
	return r
}

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api" || strings.HasPrefix(r.URL.Path, "/api/") || (r.Method != http.MethodGet && r.Method != http.MethodHead) {
		http.NotFound(w, r)
		return
	}
	buildDir := os.Getenv("WEB_BUILD_DIR")
	if buildDir == "" {
		buildDir = filepath.Join("web", "build")
	}
	file := filepath.Join(buildDir, filepath.FromSlash(strings.TrimPrefix(path.Clean(r.URL.Path), "/")))
	if info, err := os.Stat(file); err == nil && !info.IsDir() {
		http.ServeFile(w, r, file)
		return
	}
	w.Header().Set("Cache-Control", "no-cache")
	http.ServeFile(w, r, filepath.Join(buildDir, "200.html"))
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
