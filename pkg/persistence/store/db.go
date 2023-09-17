package store

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func InitSqlDb(d gorm.Dialector) {
	if db != nil {
		return
	}
	var err error
	db, err = gorm.Open(d, &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
