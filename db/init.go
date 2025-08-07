package db

import (
	"github.com/pento-zh/gin-test/db/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db
}

func Init() {
	db := GetDb()
	// 迁移 schema
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Order{})
}
