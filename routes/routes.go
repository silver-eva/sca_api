package routes

import (
	"sca_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Spy Cats
	api.Post("/cats", handlers.CreateCat)
	api.Get("/cats", handlers.GetCats)
	api.Get("/cats/:id", handlers.GetCat)
	api.Delete("/cats/:id", handlers.DeleteCat)
	api.Put("/cats/:id/salary", handlers.UpdateCatSalary)

	// Missions
	// api.Post("/missions", handlers.CreateMission)
	// api.Get("/missions", handlers.GetMissions)
	// api.Get("/missions/:id", handlers.GetMission)
	// api.Delete("/missions/:id", handlers.DeleteMission)
	// api.Put("/missions/:id/complited", handlers.UpdateMissionComplited)
}