package config

import (
	"fmt"
	"log"
	"os"

	"github.com/emanuelhardwell/go-user/dao"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectPostgres sets up the GORM connection
func ConnectPostgres() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando el archivo .env: %w", err)
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
		return nil, err
	}

	// Auto-migrate
	if err := db.AutoMigrate(&dao.User{}); err != nil {
		log.Printf("migration failed: %v", err)
	}

	return db, nil
}

// RequestLogger is a Fiber middleware for logging
var RequestLogger = logger.New(logger.Config{
	Format: "${time} | ${method} ${path} | ${status}\n",
})
