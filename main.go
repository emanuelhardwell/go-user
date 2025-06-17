package main

import (
	"log"

	"github.com/emanuelhardwell/go-user/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration and initialize DB
	err := config.ConnectPostgres()
	if err != nil {
		log.Fatalf("ERROR: failed to connect database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Register middleware
	app.Use("/", config.RequestLogger)

	log.Fatal(app.Listen(":3000"))
}
