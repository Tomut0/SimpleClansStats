package application

import (
	"errors"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type AuthService struct {
	repository ports.UserRepository
	hasher     ports.PasswordHasher
}

func NewAuthService(repository ports.UserRepository, hasher ports.PasswordHasher) *AuthService {
	return &AuthService{
		repository: repository,
		hasher:     hasher,
	}
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *AuthService) CreateAdminUser(email, username string, password string) error {
	return s.repository.Create(email, username, password, domain.RoleAdmin)
}

func (s *AuthService) Login(email, password string) (int64, error) {
	credentials, err := s.repository.GetUserCredentialsByEmail(email)
	if err != nil {
		return 0, err
	}

	ok, err := s.hasher.VerifyPassword([]byte(password), []byte(credentials.PasswordHash))
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, ErrInvalidCredentials
	}

	return credentials.UserID, nil
}

func (s *AuthService) GetByID(id int64) (domain.User, error) {
	return s.repository.GetUserById(id)
}