package middleware

import (
	"context"
	"log/slog"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/api"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type AuthMiddleware struct{ jwtService service.Jwt }

func NewAuthMiddleware(jwtService service.Jwt) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (am *AuthMiddleware) Auth(next api.ControllerFunc) api.ControllerFunc {
	return func(ctx *api.Context) error {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			return exceptions.Unauthorized("Missing authorization header not present")
		}
		payload, err := am.jwtService.Verify(ctx.Request.Header.Get("Authorization"))
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
