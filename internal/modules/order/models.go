package order

import (
	"time"

	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

type Order struct {
	ID        uint            `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	ClientID  uint            `json:"clientId" gorm:"OnDelete:CASCADE;"`
	ProductID uint            `json:"productId" gorm:"OnDelete:CASCADE;"`
	Product   product.Product `json:"product" gorm:"OnDelete:CASCADE;"`
	Quantity  uint            `json:"quantity"`
}
