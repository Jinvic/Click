package db

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var mutex sync.Mutex

func init() {
	db, err := gorm.Open(sqlite.Open("click.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	DB = db
}
