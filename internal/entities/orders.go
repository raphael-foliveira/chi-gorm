package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ClientID  uint
	Client    Client `faker:"-" gorm:"OnDelete:CASCADE;"`
	ProductID uint
	Product   Product `faker:"-" gorm:"OnDelete:CASCADE;"`
	Quantity  int
}

func (p Order) GetId() uint {
	return p.ID
}
