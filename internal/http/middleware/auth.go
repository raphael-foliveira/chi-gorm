package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type AuthMiddleware struct {
	jwtService  *service.JwtService
	userService *service.UsersService
}

type ctxKey string

func NewAuth(jwtService *service.JwtService, usersService *service.UsersService) *AuthMiddleware {
	return &AuthMiddleware{jwtService, usersService}
}

func (a *AuthMiddleware) CheckToken(fn controller.ControllerFunc) controller.ControllerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		token, err := a.extractTokenFromHeader(r)
		if err != nil {
			return exceptions.Unauthorized()
		}
		claims, err := a.jwtService.Verify(token)
		if err != nil {
			return exceptions.Unauthorized()
		}
		user, err := a.userService.Get(claims.UserID)
		if err != nil {
			return exceptions.Unauthorized()
		}
		r = r.WithContext(context.WithValue(r.Context(), ctxKey("user"), user))
		return fn(w, r)
	}
}

func (a *AuthMiddleware) extractTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errInvalidAuthorizationHeader
	}
	if !strings.Contains(header, "Bearer ") {
		return "", errInvalidAuthorizationHeader
	}
	splitHeader := strings.Split(header, " ")
	if len(splitHeader) < 2 {
		return "", errInvalidAuthorizationHeader
	}
	token := splitHeader[1]
	return token, nil
}

var errInvalidAuthorizationHeader = errors.New("invalid authentication header")
