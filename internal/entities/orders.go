package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model `faker:"-" json:"-"`
	ClientID   uint
	Client     Client `gorm:"OnDelete:CASCADE;"`
	ProductID  uint
	Product    Product `gorm:"OnDelete:CASCADE;"`
	Quantity   uint
}
