package models

import (
	"time"
)

type Client struct {
	ID        uint      `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time `json:"-" faker:"-"`
	UpdatedAt time.Time `json:"-" faker:"-"`
	Name      string    `json:"name" gorm:"not null" faker:"name"`
	Email     string    `json:"email" gorm:"not null" faker:"email"`
}
