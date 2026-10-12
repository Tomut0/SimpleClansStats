package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/go-playground/validator/v10"
)

type ServerHandler struct {
	validator   *validator.Validate
	authService *application.AuthService
}

func NewServerHandler(authService *application.AuthService) *ServerHandler {
	validate := validator.New()
	return &ServerHandler{validator: validate, authService: authService}
}

// Login Log in with username/password
// (POST /auth/login)
func (s *ServerHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := s.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.authService.Login(req.Username, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login successful"))
}

// GetCurrentUser Current user
// (GET /auth/me)
func (s *ServerHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"id":    "user-id-here",
		"email": "user@example.com",
	})
}

// GetServers Get servers
// (GET /servers)
func (s *ServerHandler) GetServers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]string{"server1", "server2"})
}

// Logout Log out
// (POST /auth/logout)
func (s *ServerHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
