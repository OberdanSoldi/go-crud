package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"simple-crud/src/model/entity"
	"sync"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func InitDb() {
	once.Do(func() {
		var err error
		dbConn, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database!")
		}

		if err = dbConn.AutoMigrate(&entity.Book{}); err != nil {
			panic("Failed to migrate database!")
		}
	})
}

func GetDb() *gorm.DB {
	if dbConn == nil {
		InitDb()
	}
	return dbConn
}
