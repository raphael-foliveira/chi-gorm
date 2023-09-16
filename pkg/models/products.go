package models

import "time"

type Product struct {
	ID        int64     `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time `json:"-" faker:"-"`
	UpdatedAt time.Time `json:"-" faker:"-"`
	Name      string    `json:"name" gorm:"not null" faker:"name"`
	Price     float64   `json:"price" gorm:"not null" faker:"amount"`
}
