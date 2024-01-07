package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
)

type Claims struct {
	ID         uint   `json:"id"`
	ClientName string `json:"client_name"`
	Email      string `json:"email"`
}

type Jwt interface {
	Sign(Claims) (string, error)
	Verify(string) (*Claims, error)
}

type jwtService struct {
	secret string
}

func NewJwt() Jwt {
	return &jwtService{cfg.Cfg.JwtSecret}
}

func (j *jwtService) Sign(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":          claims.ID,
		"client_name": claims.ClientName,
		"email":       claims.Email,
	})
	return token.SignedString([]byte(j.secret))

}

func (j *jwtService) Verify(token string) (*Claims, error) {
	data, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if !t.Valid {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodES256.Name}))
	return &Claims{
		ID:         data.Claims.(jwt.MapClaims)["id"].(uint),
		ClientName: data.Claims.(jwt.MapClaims)["client_name"].(string),
		Email:      data.Claims.(jwt.MapClaims)["email"].(string),
	}, err
}
