package presentation

import (
	"context"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/alexedwards/scs/v2"
)

type ServerHandler struct {
	authService     *application.AuthService
	dashboard       *application.DashboardService
	ingest          *application.IngestService
	settingsService *application.SettingsService
	serverService   *application.ServerService
	sessions        *scs.SessionManager
}

func NewServerHandler(authService *application.AuthService, sessions *scs.SessionManager, dashboard *application.DashboardService, ingest *application.IngestService, settingsService *application.SettingsService, serverService *application.ServerService) *ServerHandler {
	return &ServerHandler{authService: authService, dashboard: dashboard, ingest: ingest, settingsService: settingsService, serverService: serverService, sessions: sessions}
}

// Login Log in with username/password
// (POST /auth/login)
func (s *ServerHandler) Login(ctx context.Context, request LoginRequestObject) (LoginResponseObject, error) {
	userID, err := s.authService.Login(string(request.Body.Email), request.Body.Password)
	if err != nil {
		return Login401JSONResponse{}, err
	}

	// prevent session fixation attacks
	// @see https://github.com/alexedwards/scs#preventing-session-fixation
	if err := s.sessions.RenewToken(ctx); err != nil {
		return Login500Response{}, err
	}

	s.sessions.Put(ctx, "user_id", userID)
	return Login200JSONResponse{}, nil
}

// GetCurrentUser Current user
// (GET /auth/me)
func (s *ServerHandler) GetCurrentUser(ctx context.Context, request GetCurrentUserRequestObject) (GetCurrentUserResponseObject, error) {
	userID, ok := s.sessions.Get(ctx, "user_id").(int64)
	if !ok {
		return GetCurrentUser401JSONResponse{}, nil
	}

	user, err := s.authService.GetByID(userID)
	if err != nil {
		return GetCurrentUser401JSONResponse{}, err
	}

	// todo: probably create dto-mapper?
	apiUser := User{
		Email:    user.Email(),
		Username: user.Username(),
		Role:     UserRole(user.Role()),
	}

	return GetCurrentUser200JSONResponse(apiUser), nil
}

// Logout Log out
// (POST /auth/logout)
func (s *ServerHandler) Logout(ctx context.Context, request LogoutRequestObject) (LogoutResponseObject, error) {
	err := s.sessions.Destroy(ctx)
	if err != nil {
		return Logout500Response{}, err
	}

	return Logout204Response{}, nil
}
