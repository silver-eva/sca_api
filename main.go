package main

import (
	"log"
	"os"

	"sca_api/database"
	"sca_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/swaggo/fiber-swagger"
	_ "sca_api/docs"
)

// @title			Spy Cats API
// @version		1.0
// @description	A simple REST API for managing spy cats and missions
// @host			localhost:8000
// @BasePath		/api
func main() {
	database.ConnectDB()
	defer database.DB.Close()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok"})
	})
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	routes.SetupRoutes(app)

	log.Fatal(app.Listen("0.0.0.0:" + os.Getenv("APP_PORT")))
}

// "8000"))//
