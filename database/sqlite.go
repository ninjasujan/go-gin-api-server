package database

import (
	"ginserver/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	var err error
	Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	err = Database.AutoMigrate(&models.Book{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
