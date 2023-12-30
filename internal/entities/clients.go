package entities

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name   string  `gorm:"not null" faker:"name"`
	Email  string  `gorm:"not null" faker:"email"`
	Orders []Order `faker:"-"`
}

func (p Client) GetId() uint {
	return p.ID
}
