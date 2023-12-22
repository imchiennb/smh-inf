package database

import (
	"github.com/imchiennb/acmex/internal/app/model"
	_ "github.com/mattn/go-sqlite3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Entity{})

	DB = db

	return db
}
