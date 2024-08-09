package main

import (
	"log"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/pkg/logger"
	"task-management-system/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.Load()
	logger.Init()

	err := db.Init()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
