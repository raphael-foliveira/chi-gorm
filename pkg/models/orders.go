package models

import (
	"time"
)

type Order struct {
	ID        int64     `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	ClientID  int64     `json:"clientId" gorm:"OnDelete:CASCADE;"`
	ProductID int64     `json:"productId" gorm:"OnDelete:CASCADE;"`
	Quantity  int       `json:"quantity"`
}
