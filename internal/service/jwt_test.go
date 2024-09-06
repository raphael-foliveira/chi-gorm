package service_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	t.Run("should sign a token", func(t *testing.T) {
		jwt := service.NewJwt()
		payload := &service.Payload{
			ClientID:   1,
			ClientName: "John Doe",
			Email:      "john@doe.com",
		}
		tokenString, err := jwt.Sign(payload)
		assert.NoError(t, err)
		assert.NotEqual(t, "", tokenString)
	})

	t.Run("should verify a token", func(t *testing.T) {
		jwt := service.NewJwt()
		payload := &service.Payload{
			ClientID:   1,
			ClientName: "John Doe",
			Email:      "john@doe.com",
		}
		tokenString, err := jwt.Sign(payload)
		assert.NoError(t, err)
		tokenPayload, err := jwt.Verify(tokenString)
		assert.NoError(t, err)
		assert.Equal(t, payload.ClientID, tokenPayload.ClientID)
		assert.Equal(t, payload.ClientName, tokenPayload.ClientName)
		assert.Equal(t, payload.Email, tokenPayload.Email)
	})
}
