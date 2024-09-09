package db

import (
	"log"

	"github.com/jdotcurs/go-todo-cli/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
}
