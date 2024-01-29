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
}
