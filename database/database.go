package database

import (
	"log"

	"github.com/HanmaDevin/image-processing/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&types.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
