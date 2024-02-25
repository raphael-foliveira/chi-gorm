package service

import "testing"

func TestJwt(t *testing.T) {
	t.Run("should sign a token", func(t *testing.T) {
		jwt := NewJwtService()
		payload := &JwtPayload{
			UserID:   1,
			Username: "John Doe",
			Email:    "john@doe.com",
		}
		tokenString, err := jwt.Sign(payload)
		if err != nil {
			t.Error(err)
		}
		if tokenString == "" {
			t.Error("tokenString should not be empty")
		}
	})

	t.Run("should verify a token", func(t *testing.T) {
		jwt := NewJwtService()
		payload := &JwtPayload{
			UserID:   1,
			Username: "John Doe",
			Email:    "john@doe.com",
		}
		tokenString, err := jwt.Sign(payload)
		if err != nil {
			t.Error(err)
		}
		tokenPayload, err := jwt.Verify(tokenString)
		if err != nil {
			t.Error(err)
		}
		if tokenPayload.UserID != payload.UserID {
			t.Errorf("Expected client id %d, got %d", payload.UserID, tokenPayload.UserID)
		}
		if tokenPayload.Username != payload.Username {
			t.Errorf("Expected client name %s, got %s", payload.Username, tokenPayload.Username)
		}
		if tokenPayload.Email != payload.Email {
			t.Errorf("Expected email %s, got %s", payload.Email, tokenPayload.Email)
		}
	})
}
