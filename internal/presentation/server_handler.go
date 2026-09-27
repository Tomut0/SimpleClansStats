package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/Tomut0/SimpleClansStats/internal/application"
)

type ServerHandler struct {
	validator *validator.Validate
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
    
    w.Write([]byte("login successful"))
}

// Logout Log out
// (POST /auth/logout)
func (s *ServerHandler) Logout(w http.ResponseWriter, r *http.Request) {
	
}