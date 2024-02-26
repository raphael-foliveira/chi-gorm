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


func (u *usersController) Register(w http.ResponseWriter, r *http.Request) error {
	payload, err := parseBody(r, &schemas.CreateUser{})
	if err != nil {
		return err
	}
	user, err := u.service.Create(payload)
	if err != nil {
		return err
	}
	userResponse := &schemas.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return json.NewEncoder(w).Encode(&schemas.RegisterResponse{
		User: userResponse,
	})
}
