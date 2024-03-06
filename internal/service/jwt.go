package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
)

type Payload struct {
	ClientID   uint   `json:"id"`
	ClientName string `json:"client_name"`
	Email      string `json:"email"`
}

type Claims struct {
	*Payload
	*jwt.RegisteredClaims
}

type Jwt struct {
	secret []byte
}

func NewJwt() *Jwt {
	return &Jwt{[]byte(config.Cfg().JwtSecret)}
}

func (j *Jwt) Sign(payload *Payload) (string, error) {
	claims := Claims{
		Payload: payload,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(j.secret)
}

func (j *Jwt) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	}
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, keyFunc)
	return &Payload{
		ClientID:   claims.ClientID,
		ClientName: claims.ClientName,
		Email:      claims.Email,
	}, err
}
