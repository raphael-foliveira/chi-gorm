package controller

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type authController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *authController {
	return &authController{authService}
}

func (a *authController) Login(w http.ResponseWriter, r *http.Request) error {
	payload, err := parseBody(r, &schemas.Login{})
	if err != nil {
		return err
	}
	token, err := a.authService.Login(payload.Email, payload.Password)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(&schemas.LoginResponse{Token: token})
}
