package main

import (
	"log"
	"os"

	"gofiber-blog/database"
	"gofiber-blog/handlers"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Overload(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	database.Connect()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		},
	})

	app.Use(logger.New())
	app.Use(recover.New())

	config := huma.DefaultConfig("Blog API", "1.0.0")
	config.Info.Description = "A Blog REST API built with Go Fiber and PostgreSQL."
	config.Components.SecuritySchemes = nil

	api := humafiber.New(app, config)

	handlers.RegisterRoutes(api)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
