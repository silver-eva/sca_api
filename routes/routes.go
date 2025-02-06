package routes

import (
	"sca_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupCatsRoutes(app fiber.Router) {
	app.Post("/cats", handlers.CreateCat)
	app.Get("/cats", handlers.GetCats)
	app.Get("/cats/:id", handlers.GetCat)
	app.Delete("/cats/:id", handlers.DeleteCat)
	app.Put("/cats/:id/salary", handlers.UpdateCatSalary)
}

func setupMissionsRoutes(app fiber.Router) {
	app.Post("/missions", handlers.CreateMission)
	app.Get("/missions", handlers.GetMissions)
	app.Get("/missions/:id", handlers.GetMission)
	app.Delete("/missions/:id", handlers.DeleteMission)
	app.Put("/missions/:id/complited", handlers.UpdateMissionCompletion)

	app.Put("/missions/:id/cats", handlers.AssignCatToMission)
	// app.Put("/missions/:id/targets", handlers.AddTargetToMission)
	app.Delete("/missions/:mission_id/targets/:target_id", handlers.RemoveTargetFromMission)
}

func setupTargetsRoutes(app fiber.Router) {
	app.Put("/targets/:id/notes", handlers.UpdateTargetNotes)
	app.Put("/targets/:id/complited", handlers.UpdateTargetCompletion)
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Spy Cats
	setupCatsRoutes(api)
	// Missions
	setupMissionsRoutes(api)
	// Targets
	setupTargetsRoutes(api)

}