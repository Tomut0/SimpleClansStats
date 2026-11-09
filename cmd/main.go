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

	serverHandler := presentation.NewServerHandler(authService, sessionManager, db.Queries)
	server.RunHttpServer(serverHandler, sessionManager)
}
