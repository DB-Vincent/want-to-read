package database

import (
	"log"

	"github.com/DB-Vincent/want-to-read/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("want-to-read.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	// Enable foreign keys for SQLite
	DB.Exec("PRAGMA foreign_keys = ON")

	// Migrate the database schema
	if err := DB.AutoMigrate(&models.Book{}, &models.User{}, &models.Health{}); err != nil {
		log.Printf("Failed to migrate database: %v", err)
	}

	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
