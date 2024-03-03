package schemas

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
)

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cu *CreateUser) Validate() error {
	return validate.Rules(
		validate.Required("username", cu.Username),
		validate.Required("password", cu.Password),
		validate.Required("email", cu.Email),
		validate.Email("email", cu.Email),
		validate.MinLength("password", cu.Password, 6),
	)
}

func (cu *CreateUser) ToEntity() *entities.User {
	return &entities.User{
		Username: cu.Username,
		Password: cu.Password,
		Email:    cu.Email,
	}
}

type UpdateUser struct {
	CreateUser
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) Validate() error {
	return nil
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	User *UserResponse `json:"user"`
}