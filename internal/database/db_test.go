package database

import (
	"testing"

	"gorm.io/gorm"
)

func TestInitDb(t *testing.T) {
	t.Run("should return a db instance when db is already initialized", func(t *testing.T) {
		db = &gorm.DB{}
		if InitDb(nil) != db {
			t.Error("Should return db instance")
		}
		db = nil
	})
}
