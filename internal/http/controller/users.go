package controller

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type usersController struct {
	service *service.UsersService
}

func NewUsersController(service *service.UsersService) *usersController {
	return &usersController{service}
}

func (u *usersController) Login(w http.ResponseWriter, r *http.Request) error {
	payload, err := parseBody(r, &schemas.Login{})
	if err != nil {
		return err
	}
	token, err := u.service.Login(payload.Email, payload.Password)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(&schemas.LoginResponse{Token: token})

}
