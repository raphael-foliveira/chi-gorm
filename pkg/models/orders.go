package models

import (
	"time"
)

type Order struct {
	ID        uint      `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	ClientID  uint      `json:"clientId" gorm:"OnDelete:CASCADE;"`
	ProductID uint      `json:"productId" gorm:"OnDelete:CASCADE;"`
	Product   Product   `json:"product" gorm:"OnDelete:CASCADE;"`
	Quantity  uint      `json:"quantity"`
}
