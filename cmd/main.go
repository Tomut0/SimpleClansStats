package main

import (
	"log"
	"os"
	"time"

	"github.com/Tomut0/SimpleClansStats/cmd/server"
	"github.com/Tomut0/SimpleClansStats/cmd/utils"
	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/security"
	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := sqlite.New(os.Getenv("GOOSE_DBSTRING"))
	defer db.Close()
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

	// 1. Задача AuthService: аутентифицировать пользователей
	// 2. Задача UserRepository:  (создание, обновление, удаление) хранить пользователей. AuthService(UserRepository)
	// 3. Нужно иметь сессии для хранения состояния аутентификации (alexedwards/scs)

	dashboardService := application.NewDashboardService(sqlite.NewSQLiteDashboardRepository(db.Db))
	ingestService := application.NewIngestService(sqlite.NewSQLiteIngestRepository(db.Db))
	settingsService := application.NewSettingsService(sqlite.NewSQLiteSettingsRepository(db.Db))
	serverService := application.NewServerService(sqlite.NewSQLiteServerRepository(db.Db))
	serverHandler := presentation.NewServerHandler(authService, sessionManager, dashboardService, ingestService, settingsService, serverService)
	server.RunHttpServer(serverHandler, sessionManager)
}
