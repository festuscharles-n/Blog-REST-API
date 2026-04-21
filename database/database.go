package database

import (
	"fmt"
	"log"
	"os"

	"gofiber-blog/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Connect() {
	host := getenv("DB_HOST", "localhost")
	user := getenv("DB_USER", "mac")
	password := os.Getenv("DB_PASSWORD")
	dbName := getenv("DB_NAME", "blog_db")
	port := getenv("DB_PORT", "5432")

	var dsn string
	if password != "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	} else {
		dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", user, host, port, dbName)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Post{})

	DB = db
	log.Println("Database connected and migrated")
}
