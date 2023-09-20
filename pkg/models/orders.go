package models

import (
	"time"
)

type Order struct {
	ID        int64   `gorm:"primarykey" faker:"-"`
	ClientID  int64   ` gorm:"OnDelete:CASCADE;"`
	Client    Client  `faker:"-"`
	ProductID int64   ` gorm:"OnDelete:CASCADE;"`
	Product   Product `faker:"-"`
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
