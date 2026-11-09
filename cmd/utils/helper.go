package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

// CreateAdminIfNotExists bootstraps the configured admin account once.
func CreateAdminIfNotExists(authService *application.AuthService, users ports.UserRepository) error {
	email := os.Getenv("BOOTSTRAP_ADMIN_EMAIL")
	username := os.Getenv("BOOTSTRAP_ADMIN_USERNAME")
	password := os.Getenv("BOOTSTRAP_ADMIN_PASSWORD")

	if email == "" && username == "" && password == "" {
		log.Print("Admin bootstrap skipped: BOOTSTRAP_ADMIN_* variables are not configured")
		return nil
	}
	if email == "" || username == "" || password == "" {
		return errors.New("BOOTSTRAP_ADMIN_EMAIL, BOOTSTRAP_ADMIN_USERNAME, and BOOTSTRAP_ADMIN_PASSWORD must all be set")
	}

	user, err := users.GetUserByEmail(email)
	if err == nil {
		if user.Role() != domain.RoleAdmin {
			return fmt.Errorf("configured admin email %q belongs to a non-admin user", email)
		}
		log.Printf("Admin account %q already exists", email)
		return nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("look up admin account %q: %w", email, err)
	}

	if err := authService.CreateAdminUser(email, username, password); err != nil {
		return fmt.Errorf("create admin account %q: %w", email, err)
	}

	log.Printf("Created admin account %q", email)
	return nil
}
