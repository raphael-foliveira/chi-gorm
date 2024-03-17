package entities

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name   string  `faker:"name"`
	Email  string  `faker:"email"`
	Orders []Order `faker:"-"`
}

func (p Client) GetId() uint {
	return p.ID
}
