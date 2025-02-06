package main

import (
	"log"
	"os"

	"sca_api/database"
	"sca_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok"})
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen("0.0.0.0:" + os.Getenv("APP_PORT")))
}
