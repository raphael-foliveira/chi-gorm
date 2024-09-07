package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/dto"
)

type jwtS struct {
	secret []byte
}

func NewJwt() *jwtS {
	return &jwtS{[]byte(config.JwtSecret)}
}

func (j *jwtS) Sign(payload *dto.JwtPayload) (string, error) {
	claims := dto.JwtClaims{
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

func (j *jwtS) Verify(token string) (*dto.JwtPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	}
	claims := &dto.JwtClaims{}
	_, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	return &dto.JwtPayload{
		ClientID:   claims.ClientID,
		ClientName: claims.ClientName,
		Email:      claims.Email,
	}, nil
}
