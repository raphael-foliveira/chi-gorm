package middleware

import (
	"context"
	"log/slog"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Auth(next controller.ControllerFunc) controller.ControllerFunc {
	return func(ctx *controller.Context) error {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			return exceptions.Unauthorized("Missing authorization header not present")
		}
		payload, err := service.Jwt.Verify(ctx.Request.Header.Get("Authorization"))
		if err != nil {
			slog.Error(err.Error())
			return exceptions.Unauthorized("Invalid token")
		}
		payloadContext := context.WithValue(ctx.Request.Context(), contextKey("user"), payload)
		ctx.Request = ctx.Request.WithContext(payloadContext)
		err = next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}

type contextKey string
