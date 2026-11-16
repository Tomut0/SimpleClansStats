package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/Tomut0/SimpleClansStats/cmd/server"
	"github.com/Tomut0/SimpleClansStats/cmd/utils"
	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/security"
	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/Tomut0/SimpleClansStats/sql/migrations"
	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pressly/goose/v3"
)

func main() {
	mode := flag.String("mode", "api", "server mode: web or api")
	flag.Parse()
	if *mode != "web" && *mode != "api" {
		log.Fatalf("invalid mode %q: use web or api", *mode)
	}

	db := sqlite.New(os.Getenv("GOOSE_DBSTRING"))
	defer db.Close()
	goose.SetBaseFS(migrations.Files)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("configure database migrations: %v", err)
	}

	if err := goose.Up(db.Db, "."); err != nil {
		log.Fatalf("run database migrations: %v", err)
	}

	installationID, err := sqlite.EnsureInstallationID(db.Db)
	if err != nil {
		log.Fatalf("initialize installation ID: %v", err)
	}

	startIngestCleanup(db.Queries)

	argonHasher := security.NewArgonHasher(3, 64*1024, 4)
	userRepository := sqlite.NewSQLiteUserRepository(db.Queries, argonHasher)

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Store = sqlite3store.New(db.Db)

	authService := application.NewAuthService(userRepository, argonHasher)
	if err := utils.CreateAdminIfNotExists(authService, userRepository); err != nil {
		log.Fatalf("bootstrap admin user: %v", err)
	}

	dashboardService := application.NewDashboardService(sqlite.NewSQLiteDashboardRepository(db.Db))
	ingestService := application.NewIngestService(sqlite.NewSQLiteIngestRepository(db.Db))
	searchService := application.NewSearchService(sqlite.NewSQLiteSearchRepository(db.Db))
	settingsService := application.NewSettingsService(sqlite.NewSQLiteSettingsRepository(db.Db))
	serverService := application.NewServerService(sqlite.NewSQLiteServerRepository(db.Db))

	serverHandler := presentation.NewServerHandler(authService, sessionManager, dashboardService, ingestService, settingsService, serverService, searchService)
	server.RunHttpServer(serverHandler, sessionManager, installationID, *mode == "web")
}
