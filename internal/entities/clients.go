package entities

import (
	"time"
)

type Client struct {
	ID        int64     `gorm:"primarykey" faker:"-"`
	CreatedAt time.Time `faker:"-"`
	UpdatedAt time.Time `faker:"-"`
	Name      string    `gorm:"not null" faker:"name"`
	Email     string    `gorm:"not null" faker:"email"`
	Orders    []Order   `faker:"-"`
}

func (p Client) GetId() int64 {
	return p.ID
}
