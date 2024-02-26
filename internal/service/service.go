package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

func Orders() *OrdersService {
	return NewOrdersService(repository.Orders())
}

func Products() *ProductsService {
	return NewProductsService(repository.Products())
}

func Clients() *ClientsService {
	return NewClientsService(repository.Clients(), repository.Orders())
}

func Jwt() *JwtService {
	return NewJwtService()
}

func Encryption() *EncryptionService {
	return NewEncryptionService(10)
}

func Users() *UsersService {
	return NewUsersService(repository.Users(), Encryption())
}

func Auth() *AuthService {
	return NewAuthService(repository.Users(), Encryption(), Jwt())
}
