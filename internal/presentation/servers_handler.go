package presentation

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
	"github.com/google/uuid"
)

func (s *ServerHandler) GetServers(ctx context.Context, _ GetServersRequestObject) (GetServersResponseObject, error) {
	rows, err := s.settings.ListServers(ctx)
	if err != nil {
		log.Printf("list servers: %v", err)
		return GetServers500Response{}, nil
	}

	now := time.Now()
	servers := make([]Server, 0, len(rows))
	for _, row := range rows {
		id, err := uuid.Parse(row.ID)
		if err != nil {
			log.Printf("invalid stored server ID: %v", err)
			return GetServers500Response{}, nil
		}
		server := Server{Id: id, DisplayName: row.DisplayName, Disabled: row.Disabled}
		if row.LastReceivedAt != "" {
			receivedAt, err := time.Parse(time.RFC3339Nano, row.LastReceivedAt)
			if err != nil {
				log.Printf("invalid stored ingest time: %v", err)
				return GetServers500Response{}, nil
			}
			server.LastIngestAt = &receivedAt
			server.Online = !row.Disabled && now.Sub(receivedAt) <= time.Duration(row.ServerOnlineTimeoutMinutes)*time.Minute
		}
		servers = append(servers, server)
	}
	return GetServers200JSONResponse(servers), nil
}

func (s *ServerHandler) CreateServer(ctx context.Context, request CreateServerRequestObject) (CreateServerResponseObject, error) {
	switch s.serverAdminStatus(ctx) {
	case http.StatusUnauthorized:
		return CreateServer401Response{}, nil
	case http.StatusForbidden:
		return CreateServer403Response{}, nil
	}

	name := strings.TrimSpace(request.Body.DisplayName)
	if name == "" || utf8.RuneCountInString(name) > 100 || strings.ContainsFunc(name, unicode.IsControl) {
		return CreateServer400Response{}, nil
	}
	id := uuid.New()
	key := "sc-" + rand.Text()
	hash := sha256.Sum256([]byte(key))
	err := s.settings.CreateServer(ctx, generated.CreateServerParams{
		ID: id.String(), DisplayName: name, ApiKeyHash: hash[:],
	})
	if err != nil {
		log.Printf("create server: %v", err)
		return CreateServer500Response{}, nil
	}
	cacheControl := "no-store"
	return CreateServer201JSONResponse{
		Body:    ServerRegistration{Server: Server{Id: id, DisplayName: name}, ApiKey: key},
		Headers: CreateServer201ResponseHeaders{CacheControl: &cacheControl},
	}, nil
}

func (s *ServerHandler) DeleteServer(ctx context.Context, request DeleteServerRequestObject) (DeleteServerResponseObject, error) {
	switch s.serverAdminStatus(ctx) {
	case http.StatusUnauthorized:
		return DeleteServer401Response{}, nil
	case http.StatusForbidden:
		return DeleteServer403Response{}, nil
	}
	if err := s.settings.DeleteServer(ctx, request.ServerId.String()); err != nil {
		log.Printf("delete server: %v", err)
		return DeleteServer500Response{}, nil
	}
	return DeleteServer204Response{}, nil
}

func (s *ServerHandler) serverAdminStatus(ctx context.Context) int {
	userID, ok := s.sessions.Get(ctx, "user_id").(int64)
	if !ok {
		return http.StatusUnauthorized
	}
	user, err := s.authService.GetByID(userID)
	if err != nil {
		return http.StatusUnauthorized
	}
	if user.Role() != domain.RoleAdmin {
		return http.StatusForbidden
	}
	return http.StatusOK
}
