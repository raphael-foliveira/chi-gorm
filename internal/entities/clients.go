package entities

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model `faker:"-"`
	Name       string  `faker:"name" json:"name"`
	Email      string  `faker:"email" json:"email"`
	Orders     []Order `faker:"-" json:"orders"`
}
