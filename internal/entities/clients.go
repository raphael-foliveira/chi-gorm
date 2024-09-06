package entities

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model `faker:"-"`
	Name       string  `faker:"name"`
	Email      string  `faker:"email"`
	Orders     []Order `faker:"-"`
}
