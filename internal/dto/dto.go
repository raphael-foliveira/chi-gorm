package dto

import "github.com/golang-jwt/jwt/v5"

type JwtPayload struct {
	ClientID   uint   `json:"id"`
	ClientName string `json:"client_name"`
	Email      string `json:"email"`
}

type JwtClaims struct {
	*JwtPayload
	*jwt.RegisteredClaims
}
