package service

import "testing"

func TestJwt(t *testing.T) {
	t.Run("should sign a token", func(t *testing.T) {
		jwt := NewJwt()
		payload := &Payload{
			ClientID:   1,
			ClientName: "John Doe",
			Email:      "john@doe.com",
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
		jwt := NewJwt()
		payload := &Payload{
			ClientID:   1,
			ClientName: "John Doe",
			Email:      "john@doe.com",
		}
		tokenString, err := jwt.Sign(payload)
		if err != nil {
			t.Error(err)
		}
		tokenPayload, err := jwt.Verify(tokenString)
		if err != nil {
			t.Error(err)
		}
		if tokenPayload.ClientID != payload.ClientID {
			t.Errorf("Expected client id %d, got %d", payload.ClientID, tokenPayload.ClientID)
		}
		if tokenPayload.ClientName != payload.ClientName {
			t.Errorf("Expected client name %s, got %s", payload.ClientName, tokenPayload.ClientName)
		}
		if tokenPayload.Email != payload.Email {
			t.Errorf("Expected email %s, got %s", payload.Email, tokenPayload.Email)
		}
	})
}
