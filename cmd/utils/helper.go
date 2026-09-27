package utils

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"github.com/Tomut0/SimpleClansStats/internal/application"
)

func UpdateEnvVar(key, value string) error {
	envData, err := godotenv.Read(".env")
	if err != nil {
		return err
	}
	envData[key] = value
	return godotenv.Write(envData, ".env")
}

func UpdateEnvVars(envData map[string]string) error {
	for key, value := range envData {
		if err := UpdateEnvVar(key, value); err != nil {
			return err
		}
	}
	return nil
}

func CreateAdminIfNotExists(authService *application.AuthService) error {
	if os.Getenv("ADMIN_USERNAME") == "" || os.Getenv("ADMIN_PASSWORD_HASH") == "" {
		log.Println("No admin credentials found, creating default admin user...")
		hash, err := authService.CreateAdminUser("admin", "admin")
		if err != nil {
			return err
		}

		credentials := map[string]string{"ADMIN_USERNAME": "admin", "ADMIN_PASSWORD_HASH": hash}
		
		if err := UpdateEnvVars(credentials); err != nil {
			return err
		}

		log.Println("Admin user created successfully with username 'admin' and password 'admin'")
		return nil
	}

	return nil
}