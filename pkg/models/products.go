package models

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"-" faker:"-"`
	UpdatedAt time.Time `json:"-" faker:"-"`
	Name      string    `json:"name" gorm:"not null" faker:"name"`
	Price     float64   `json:"price" gorm:"not null" faker:"amount"`
}
