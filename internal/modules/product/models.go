package product

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        uint         `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time    `json:"-" faker:"-"`
	UpdatedAt time.Time    `json:"-" faker:"-"`
	DeletedAt sql.NullTime `gorm:"index" json:"-" faker:"-"`
	Name      string       `json:"name" gorm:"not null" faker:"name"`
	Price     float64      `json:"price" gorm:"not null" faker:"amount"`
}
