package application

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
	"github.com/google/uuid"
)

var ErrInvalidServerName = errors.New("invalid server name")

type ServerService struct{ repository ports.ServerRepository }

func NewServerService(repository ports.ServerRepository) *ServerService {
	return &ServerService{repository: repository}
}

func (s *ServerService) List(ctx context.Context) ([]domain.Server, error) {
	records, err := s.repository.ListServers(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	servers := make([]domain.Server, 0, len(records))
	for _, record := range records {
		server := domain.Server{ID: record.ID, DisplayName: record.DisplayName, Disabled: record.Disabled}
		if record.LastReceivedAt != "" {
			receivedAt, err := time.Parse(time.RFC3339Nano, record.LastReceivedAt)
			if err != nil {
				return nil, err
			}
			server.LastIngestAt = &receivedAt
			server.Online = !record.Disabled && now.Sub(receivedAt) <= time.Duration(record.OnlineTimeoutMinutes)*time.Minute
		}
		servers = append(servers, server)
	}
	return servers, nil
}

func (s *ServerService) Create(ctx context.Context, name string) (domain.ServerRegistration, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > 100 || strings.ContainsFunc(name, unicode.IsControl) {
		return domain.ServerRegistration{}, ErrInvalidServerName
	}
	id, key := uuid.NewString(), "sc-"+rand.Text()
	hash := sha256.Sum256([]byte(key))
	if err := s.repository.CreateServer(ctx, id, name, hash[:]); err != nil {
		return domain.ServerRegistration{}, err
	}
	return domain.ServerRegistration{Server: domain.Server{ID: id, DisplayName: name}, APIKey: key}, nil
}

func (s *ServerService) Delete(ctx context.Context, id string) error {
	return s.repository.DeleteServer(ctx, id)
}
