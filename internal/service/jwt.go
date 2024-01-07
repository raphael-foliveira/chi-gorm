package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
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

type Jwt interface {
	Sign(*Payload) (string, error)
	Verify(string) (*Payload, error)
}

type jwtService struct {
	secret string
}

func NewJwt() Jwt {
	return &jwtService{cfg.Cfg.JwtSecret}
}

func (j *jwtService) Sign(payload *Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, Claims{
		Payload: payload,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(j.secret))
}

func (j *jwtService) Verify(token string) (*Payload, error) {
	data, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if !t.Valid {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodES256.Name}))
	return &Payload{
		ClientID:   data.Claims.(Claims).ClientID,
		ClientName: data.Claims.(Claims).ClientName,
		Email:      data.Claims.(Claims).Email,
	}, err
}
