package config

import (
	"fmt"
	"log"
	"os"

	"github.com/emanuelhardwell/go-user/model"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// ConnectPostgres initializes the global DB
func ConnectPostgres() error {

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error cargando el archivo .env: %w", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	// Auto-migrate
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Printf("migration failed: %v", err)
	}

	return nil
}

// RequestLogger is a Fiber middleware for logging
var RequestLogger = logger.New(logger.Config{
	Format: "${time} | ${method} ${path} | ${status}\n",
})
