package database

import (
	"back/internal/models/user"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	DB.AutoMigrate(&user.User{})
}
