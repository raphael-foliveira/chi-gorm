package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model `faker:"-"`
	ClientID   uint    `faker:"-"`
	Client     Client  `faker:"-" gorm:"OnDelete:CASCADE;"`
	ProductID  uint    `faker:"-"`
	Product    Product `faker:"-" gorm:"OnDelete:CASCADE;"`
	Quantity   uint
}

func (p Order) GetId() uint {
	return p.ID
}
