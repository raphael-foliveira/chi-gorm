package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb(dialector gorm.Dialector) *gorm.DB {
	if db != nil {
		return db
	}
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	return db
}
