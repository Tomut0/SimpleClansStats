package application

import (
    "errors"
    "github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type AuthService struct {
    hasher           ports.PasswordHasher
    credentialsStore     ports.CredentialsStore
    adminUsername    string
    adminPasswordHash string
}

func NewAuthService(hasher ports.PasswordHasher, credentialsStore ports.CredentialsStore) *AuthService {
    return &AuthService{
        hasher:            hasher,
        credentialsStore:  credentialsStore,
    }
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *AuthService) CreateAdminUser(username string, password string) (string, error) {
    hashedPassword, err := s.hasher.HashPassword(password)
    if err != nil {
        return "", err
    }

    s.adminUsername = username
    s.adminPasswordHash = hashedPassword
    
    s.credentialsStore.Put(ports.Credentials{Username: username, PasswordHash: hashedPassword})
    
    return hashedPassword, nil
}

func (s *AuthService) Login(username, password string) error {
    if username != s.adminUsername {
        return ErrInvalidCredentials
    }

    ok, err := s.hasher.VerifyPassword([]byte(password), []byte(s.adminPasswordHash))
    if err != nil {
        return err
    }
    
    if !ok {
        return ErrInvalidCredentials
    }

    return nil
}