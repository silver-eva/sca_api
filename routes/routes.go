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
	app.Post("/missions/:id/complete", handlers.UpdateMissionCompletion)

	app.Post("/missions/:mid/cats/:cid", handlers.AssignCatToMission)
	app.Put("/missions/:mid/targets", handlers.AddTargetToMission)
	app.Delete("/missions/:mid/targets/:tid", handlers.RemoveTargetFromMission)

	app.Post("/missions/:mid/targets/:tid/complete", handlers.UpdateTargetCompletion)
	app.Put("/missions/:mid/targets/:tid/note", handlers.UpdateTargetNotes)
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Spy Cats
	setupCatsRoutes(api)
	// Missions
	setupMissionsRoutes(api)

}
