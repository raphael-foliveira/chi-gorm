package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model `faker:"-"`
	Username   string
	Email      string
	Password   string
}

func (u *User) GetId() uint {
	return u.ID
}
