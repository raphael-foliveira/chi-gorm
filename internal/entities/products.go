package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model `faker:"-"`
	Name       string  `gorm:"not null" faker:"word"`
	Price      float64 `gorm:"not null" faker:"amount"`
}
