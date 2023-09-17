package store

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func InitSqlDb(conn gorm.Dialector) {
	if db != nil {
		return
	}
	var err error
	db, err = gorm.Open(conn, &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
