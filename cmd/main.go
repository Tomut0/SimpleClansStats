package main

import (
	"os"

	"github.com/Tomut0/SimpleClansStats/cmd/server"
	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/security"
	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := sqlite.New(os.Getenv("GOOSE_DBSTRING"))
	defer db.Close()

	argonHasher := security.NewArgonHasher(3, 64*1024, 4)
	userRepository := sqlite.NewSQLiteUserRepository(db.Queries, argonHasher)
	authService := application.NewAuthService(userRepository, argonHasher)
	
	// 1. Задача AuthService: аутентифицировать пользователей
	// 2. Задача UserRepository:  (создание, обновление, удаление) хранить пользователей. AuthService(UserRepository)
	// 3. Нужно иметь сессии для хранения состояния аутентификации (alexedwards/scs)

	serverHandler := presentation.NewServerHandler(authService)
	server.RunHttpServer(serverHandler)
}
