package main

import (
	"log"
	"os"

	"github.com/akhiltn/go-url-shortener/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		log.Panic(err)
	}
	app := fiber.New()
	app.Use(logger.New(), healthcheck.New())
	log.Printf("Starting server on port %s", os.Getenv("APP_PORT"))
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenUrl)
}
