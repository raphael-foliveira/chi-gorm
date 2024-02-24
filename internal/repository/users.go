package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Repository[entities.User]
	FindOneByEmail(string) (*entities.User, error)
}

type users struct {
	*repository[entities.User]
}

func NewUsers(db *gorm.DB) UsersRepository {
	return &users{newRepository[entities.User](db)}
}

func (u *users) FindOneByEmail(email string) (*entities.User, error) {
	user := &entities.User{Email: email}
	return user, u.db.Model(&entities.User{}).First(&user).Error
}
