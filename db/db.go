package db

import (
	"book-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=nicetas password=yourpassword dbname=bookcrud port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автоматически создаем таблицу для модели Book
	DB.AutoMigrate(&models.Book{})
}
