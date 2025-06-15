package main

import (
	"log"

	"github.com/emanuelhardwell/go-user/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration and initialize DB
	db, err := config.ConnectPostgres()
	if err != nil {
		log.Fatalf("ERROR: failed to connect database: %v", err)
	}
	//defer db.Close()
	log.Println(db)
	// Initialize Fiber app
	app := fiber.New()

	// Register middleware
	app.Use("/api", config.RequestLogger)

	log.Fatal(app.Listen(":3000"))
}
