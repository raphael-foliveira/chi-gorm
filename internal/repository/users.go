package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type usersRepository struct {
	*repository[entities.User]
}

func NewUsersRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{newRepository[entities.User](db)}
}

func (u *usersRepository) FindOneByEmail(email string) (*entities.User, error) {
	user := &entities.User{Email: email}
	return user, u.db.Model(&entities.User{}).First(&user).Error
}
