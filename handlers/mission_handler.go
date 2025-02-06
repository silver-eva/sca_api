package handlers

import (
	"sca_api/models"
	"sca_api/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateMission(c *fiber.Ctx) error {
	var mission models.Mission
	if err := c.BodyParser(&mission); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := repositories.CreateMission(&mission)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create mission"})
	}
	return c.Status(201).JSON(mission)
}

func GetMissions(c *fiber.Ctx) error {
	missions, err := repositories.GetMissions()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch missions"})
	}
	return c.JSON(missions)
}

func GetMission(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	mission, err := repositories.GetMissionByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mission not found"})
	}
	return c.JSON(mission)
}

func DeleteMission(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = repositories.DeleteMission(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete mission"})
	}
	return c.SendStatus(204)
}

func UpdateMissionCompletion(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var data struct {
		Completed bool `json:"completed"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = repositories.UpdateMissionCompletion(id, data.Completed)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update mission completion"})
	}
	return c.SendStatus(200)
}

func AssignCatToMission(c *fiber.Ctx) error {
	var data struct {
		CatID uuid.UUID `json:"cat_id"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	missionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
	}

	err = repositories.AssignCatToMission(missionID, data.CatID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(200)
}

// func AddTargetToMission(c *fiber.Ctx) error {
// 	var data struct {
// 		TargetID uuid.UUID `json:"target_id"`
// 	}

// 	if err := c.BodyParser(&data); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
// 	}

// 	missionID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
// 	}

// 	err = repositories.AddTargetToMission(missionID, data.TargetID)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return c.SendStatus(200)
// }

func RemoveTargetFromMission(c *fiber.Ctx) error {
	missionID, err := uuid.Parse(c.Params("mission_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
	}

	targetID, err := uuid.Parse(c.Params("target_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid target ID"})
	}

	err = repositories.RemoveTargetFromMission(missionID, targetID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(200)
}

func MarkTargetCompleted(c *fiber.Ctx) error {
	targetID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid target ID"})
	}

	err = repositories.UpdateTargetCompletion(targetID, true)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(200)
}
