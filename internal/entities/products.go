package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"not null" faker:"name"`
	Price float64 `gorm:"not null" faker:"amount"`
}

func (p Product) GetId() uint {
	return p.ID
}
