package controller

import "github.com/raphael-foliveira/chi-gorm/internal/service"

type UsersController struct {
	service *service.UsersService
}

func NewUsers(service *service.UsersService) *UsersController {
	return &UsersController{service}
}
