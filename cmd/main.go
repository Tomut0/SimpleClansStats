package main

import (
    "log"
	"github.com/Tomut0/SimpleClansStats/cmd/server"
	"github.com/Tomut0/SimpleClansStats/cmd/utils"
	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/security"
	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	argonHasher := security.NewArgonHasher(3, 64*1024, 4)
	authService := application.NewAuthService(argonHasher)

	if err := utils.CreateAdminIfNotExists(authService); err != nil {
		log.Fatal(err)
	}

	serverHandler := presentation.NewServerHandler(authService)
	server.RunHttpServer(serverHandler)
}
