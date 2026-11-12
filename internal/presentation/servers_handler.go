package presentation

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/google/uuid"
)

func (s *ServerHandler) GetServers(ctx context.Context, _ GetServersRequestObject) (GetServersResponseObject, error) {
	servers, err := s.serverService.List(ctx)
	if err != nil {
		log.Printf("list servers: %v", err)
		return GetServers500Response{}, nil
	}
	response := make([]Server, 0, len(servers))
	for _, server := range servers {
		id, err := uuid.Parse(server.ID)
		if err != nil {
			log.Printf("invalid stored server ID: %v", err)
			return GetServers500Response{}, nil
		}
		response = append(response, Server{Id: id, DisplayName: server.DisplayName, Disabled: server.Disabled, Online: server.Online, LastIngestAt: server.LastIngestAt})
	}
	return GetServers200JSONResponse(response), nil
}

func (s *ServerHandler) CreateServer(ctx context.Context, request CreateServerRequestObject) (CreateServerResponseObject, error) {
	switch s.adminStatus(ctx) {
	case http.StatusUnauthorized:
		return CreateServer401Response{}, nil
	case http.StatusForbidden:
		return CreateServer403Response{}, nil
	}
	registration, err := s.serverService.Create(ctx, request.Body.DisplayName)
	if errors.Is(err, application.ErrInvalidServerName) {
		return CreateServer400Response{}, nil
	}
	if err != nil {
		log.Printf("create server: %v", err)
		return CreateServer500Response{}, nil
	}
	id, _ := uuid.Parse(registration.Server.ID)
	cacheControl := "no-store"
	return CreateServer201JSONResponse{Body: ServerRegistration{Server: Server{Id: id, DisplayName: registration.Server.DisplayName}, ApiKey: registration.APIKey}, Headers: CreateServer201ResponseHeaders{CacheControl: &cacheControl}}, nil
}

func (s *ServerHandler) DeleteServer(ctx context.Context, request DeleteServerRequestObject) (DeleteServerResponseObject, error) {
	switch s.adminStatus(ctx) {
	case http.StatusUnauthorized:
		return DeleteServer401Response{}, nil
	case http.StatusForbidden:
		return DeleteServer403Response{}, nil
	}
	if err := s.serverService.Delete(ctx, request.ServerId.String()); err != nil {
		log.Printf("delete server: %v", err)
		return DeleteServer500Response{}, nil
	}
	return DeleteServer204Response{}, nil
}

func (s *ServerHandler) adminStatus(ctx context.Context) int {
	userID, ok := s.sessions.Get(ctx, "user_id").(int64)
	if !ok {
		return http.StatusUnauthorized
	}
	admin, err := s.authService.IsAdmin(userID)
	if err != nil {
		return http.StatusUnauthorized
	}
	if !admin {
		return http.StatusForbidden
	}
	return http.StatusOK
}
