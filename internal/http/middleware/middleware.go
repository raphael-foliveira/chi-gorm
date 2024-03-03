package middleware

import "github.com/raphael-foliveira/chi-gorm/internal/service"

func Auth() *AuthMiddleware {
	return NewAuth(service.Jwt(), service.Users())
}
