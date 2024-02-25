package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestUsers(t *testing.T) {
	encryptionService := Encryption()
	jwtService := Jwt()
	usersService := NewUsersService(mocks.UsersStore, encryptionService, jwtService)
	setUp := func() {
		mocks.UsersStore.Populate()
		for i := range mocks.UsersStore.Store {
			user := &mocks.UsersStore.Store[i]
			user.Password, _ = encryptionService.Hash(user.Password)
		}
	}
	tearDown := func() {
		mocks.UsersStore.Clear()
	}
	t.Run("Login", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			setUp()
			expectedUser := &mocks.UsersStore.Store[0]
			password, hash := getPasswordAndHash(encryptionService)
			expectedUser.Password = hash
			token, err := usersService.Login(expectedUser.Email, password)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			tokenUser, err := jwtService.Verify(token)
			if err != nil {
				panic(err)
			}
			if expectedUser.Username != tokenUser.Username {
				t.Errorf("Expected %v, got %v", expectedUser.Username, tokenUser.Username)
			}
			tearDown()
		})
	})
}

func getPasswordAndHash(encryptionService *EncryptionService) (string, string) {
	password := faker.Password()
	hash, _ := encryptionService.Hash(password)
	return password, hash
}
