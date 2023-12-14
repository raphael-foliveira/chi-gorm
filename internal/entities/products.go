package entities

import "time"

type Product struct {
	ID        int64     `gorm:"primarykey" faker:"-"`
	CreatedAt time.Time `faker:"-"`
	UpdatedAt time.Time `faker:"-"`
	Name      string    `gorm:"not null" faker:"name"`
	Price     float64   `gorm:"not null" faker:"amount"`
}

func (p Product) GetId() int64 {
	return p.ID
}
