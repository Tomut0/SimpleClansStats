package ports

import "github.com/Tomut0/SimpleClansStats/internal/domain"

type UserRepository interface {
	Create(email, username, password string, role domain.Role) error
	GetUserById(id int64) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	GetUserCredentialsByEmail(email string) (domain.UserCredentials, error)
}
