package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
)

type sqliteUserRepository struct {
	queries *generated.Queries
	hasher  ports.PasswordHasher
}

func NewSQLiteUserRepository(queries *generated.Queries, hasher ports.PasswordHasher) *sqliteUserRepository {
	return &sqliteUserRepository{
		queries: queries,
		hasher:  hasher,
	}
}

func (r *sqliteUserRepository) All() ([]domain.User, error) {
	rows, err := r.queries.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, 0, len(rows))
	for _, row := range rows {
		domainUser, err := toDomainUser(row)
		if err != nil {
			continue
		}
		users = append(users, domainUser)
	}
	return users, nil
}

func (r *sqliteUserRepository) Create(email, username, password string, role domain.Role) error {
	hash, err := r.hasher.HashPassword(password)
	if err != nil {
		return fmt.Errorf("create hash (%q) for %q: %w", password, username, err)
	}

	params := generated.CreateUserParams{
		Email:        email,
		Username:     username,
		PasswordHash: hash,
		Role:         string(role),
	}

	return r.queries.CreateUser(context.Background(), params)
}

func (r *sqliteUserRepository) GetUserById(id int64) (domain.User, error) {
	user, err := r.queries.GetUserById(context.Background(), id)
	if err != nil {
		return domain.User{}, err
	}

	return toDomainUser(user)
}

func (r *sqliteUserRepository) GetUserByEmail(email string) (domain.User, error) {
	user, err := r.getUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}

	return toDomainUser(user)
}

func (r *sqliteUserRepository) GetUserCredentialsByEmail(email string) (domain.UserCredentials, error) {
	user, err := r.getUserByEmail(email)
	if err != nil {
		return domain.UserCredentials{}, err
	}

	return toDomainUserCredentials(user), nil
}

func (r *sqliteUserRepository) getUserByEmail(email string) (generated.User, error) {
	user, err := r.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return generated.User{}, fmt.Errorf("get user by username %q: %w", email, err)
	}
	return user, nil
}

func toDomainUserCredentials(u generated.User) domain.UserCredentials {
	return domain.UserCredentials{
		UserID:       u.ID,
		PasswordHash: u.PasswordHash,
	}
}

func parseTime(field, value string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse %s %q: %w", field, value, err)
	}
	return t, nil
}

func toDomainUser(u generated.User) (domain.User, error) {
	createdAt, err := parseTime("created_at", u.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	updatedAt, err := parseTime("updated_at", u.UpdatedAt)
	if err != nil {
		return domain.User{}, err
	}

	var emailVerifiedAt *time.Time
	if u.EmailVerifiedAt.Valid {
		t, err := parseTime("email_verified_at", u.EmailVerifiedAt.String)
		if err != nil {
			return domain.User{}, err
		}
		emailVerifiedAt = &t
	}

	return *domain.NewUser(
    	u.Email,
    	u.Username,
    	domain.Role(u.Role),
    	emailVerifiedAt,
    	createdAt,
    	updatedAt,
	), nil
}
