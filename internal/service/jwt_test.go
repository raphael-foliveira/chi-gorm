package service_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/dto"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	jwtService := service.NewJwt()
	payload := &dto.JwtPayload{
		ClientID:   1,
		ClientName: "John Doe",
		Email:      "john@doe.com",
	}

	t.Run("should sign a token", func(t *testing.T) {
		tokenString, err := jwtService.Sign(payload)
		assert.NoError(t, err)
		assert.NotEqual(t, "", tokenString)
	})

	t.Run("should verify a token", func(t *testing.T) {
		tokenString, err := jwtService.Sign(payload)
		assert.NoError(t, err)
		tokenPayload, err := jwtService.Verify(tokenString)
		assert.NoError(t, err)
		assert.Equal(t, payload.ClientID, tokenPayload.ClientID)
		assert.Equal(t, payload.ClientName, tokenPayload.ClientName)
		assert.Equal(t, payload.Email, tokenPayload.Email)
	})
}
