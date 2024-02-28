package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type UsersService struct {
	repo              UsersRepository
	encryptionService *EncryptionService
}

func NewUsersService(repo UsersRepository, encryptionService *EncryptionService) *UsersService {
	return &UsersService{repo, encryptionService}
}

func (s *UsersService) Create(payload *schemas.CreateUser) (*entities.User, error) {
	hashedPassword, err := s.encryptionService.Hash(payload.Password)
	if err != nil {
		return nil, err
	}
	entity := payload.ToEntity()
	entity.Password = hashedPassword
	err = s.repo.Create(entity)
	return entity, err
}

func (s *UsersService) Update(id uint, payload *schemas.UpdateUser) (*entities.User, error) {
	entity, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Username = payload.Username
	entity.Email = payload.Email
	if payload.Password != "" {
		hashedPassword, err := s.encryptionService.Hash(payload.Password)
		if err != nil {
			return nil, err
		}
		entity.Password = hashedPassword
	}
	err = s.repo.Update(entity)
	return entity, err
}

func (s *UsersService) Get(id uint) (*entities.User, error) {
	return s.repo.Get(id)
}
