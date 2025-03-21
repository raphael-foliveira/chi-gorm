package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model `faker:"-" json:"-"`
	ClientID   uint    `json:"-"`
	Client     Client  `gorm:"OnDelete:CASCADE;" json:"client"`
	ProductID  uint    `json:"-"`
	Product    Product `gorm:"OnDelete:CASCADE;" json:"product"`
	Quantity   uint    `json:"quantity"`
}
