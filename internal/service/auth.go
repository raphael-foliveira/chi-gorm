package service

import "github.com/raphael-foliveira/chi-gorm/internal/repository"

type AuthService struct {
	usersRepository   repository.UsersRepository
	encryptionService *EncryptionService
	jwtService        *JwtService
}

func NewAuthService(
	usersRepository repository.UsersRepository,
	encryptionService *EncryptionService,
	jwtService *JwtService,
) *AuthService {
	return &AuthService{
		usersRepository,
		encryptionService,
		jwtService,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.usersRepository.FindOneByEmail(email)
	if err != nil {
		return "", err
	}
	err = s.encryptionService.Compare(user.Password, password)
	if err != nil {
		return "", err
	}
	return s.jwtService.Sign(&JwtPayload{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
}
