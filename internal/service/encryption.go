package service

import "golang.org/x/crypto/bcrypt"

type EncryptionService struct {
	cost int
}

func NewEncryption(cost int) *EncryptionService {
	return &EncryptionService{cost}
}

func (e *EncryptionService) Hash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), e.cost)
	return string(b), err
}

func (e *EncryptionService) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
