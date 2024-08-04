package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		log.Panic(err)
	}
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
  log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func setupRoutes(app *fiber.App) {
	// app.Get("/:url", routes.ResolveURL)
	// app.Post("/api/v1", routes.ShortenURL)
}
