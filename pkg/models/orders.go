package models

import (
	"time"
)

type Order struct {
	ID        int64 `gorm:"primarykey" faker:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ClientID  int64 ` gorm:"OnDelete:CASCADE;"`
	Client    Client
	ProductID int64 ` gorm:"OnDelete:CASCADE;"`
	Product   Product
	Quantity  int
}
