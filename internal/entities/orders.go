package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model `faker:"-"`
	ClientID   uint
	Client     Client `gorm:"OnDelete:CASCADE;"`
	ProductID  uint
	Product    Product `gorm:"OnDelete:CASCADE;"`
	Quantity   uint
}

func (p Order) GetId() uint {
	return p.ID
}
