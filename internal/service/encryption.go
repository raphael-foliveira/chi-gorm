package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type EncryptionService struct {
	cost int
}

func NewEncryptionService(cost int) *EncryptionService {
	return &EncryptionService{cost}
}

func (e *EncryptionService) Hash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), e.cost)
	return string(b), err
}

func (e *EncryptionService) Compare(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return ErrPasswordsDoNotMatch
	}
	return nil
}

var ErrPasswordsDoNotMatch = errors.New("passwords do not match")
