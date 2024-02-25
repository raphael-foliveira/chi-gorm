package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
)

type JwtPayload struct {
	UserID   uint   `json:"id"`
	Username string `json:"client_name"`
	Email    string `json:"email"`
}

type Claims struct {
	*JwtPayload
	*jwt.RegisteredClaims
}

type JwtService struct {
	secret []byte
}

func NewJwt() *JwtService {
	return &JwtService{[]byte(cfg.Cfg().JwtSecret)}
}

func (j *JwtService) Sign(payload *JwtPayload) (string, error) {
	claims := Claims{
		JwtPayload: payload,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(j.secret)
}

func (j *JwtService) Verify(token string) (*JwtPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	}
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, keyFunc)
	return &JwtPayload{
		UserID:   claims.UserID,
		Username: claims.Username,
		Email:    claims.Email,
	}, err
}
