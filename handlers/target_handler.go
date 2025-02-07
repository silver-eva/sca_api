package handlers

import (
	"sca_api/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UpdateTargetNotes updates the notes of a target
//
//	@Summary		Update target notes
//	@Description	Update the notes of a target
//	@Tags			Targets
//	@Accept			json
//	@Produce		json
//	@Param			mid		path	string				true	"Mission ID"
//	@Param			tid		path	string				true	"Target ID"
//	@Param			data	body	map[string]string	true	"Target notes"
//	@Success		200
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/missions/{mid}/targets/{tid}/note [put]
func UpdateTargetNotes(c *fiber.Ctx) error {
	miss_id, err := uuid.Parse(c.Params("mid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
	}

	targetId, err := uuid.Parse(c.Params("tid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid target ID"})
	}

	mission, err := repositories.GetMissionByID(miss_id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mission not found"})
	}

	if mission.Complited {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot update notes for completed mission"})
	}

	target, err := repositories.GetTargetById(targetId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Target not found"})
	}

	if target.Completed {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot update notes for completed target"})
	}

	type UpdateTargetNotesRequest struct {
		Notes string `json:"notes"`
	}

	var data UpdateTargetNotesRequest

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = repositories.UpdateTargetNotes(targetId, data.Notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update target notes"})
	}
	return c.SendStatus(200)
}

// UpdateTargetCompletion marks a target as completed and if all targets in the mission are completed marks the mission as completed
//
//	@Summary		Mark target completed
//	@Description	Mark a specific target as completed
//	@Tags			Targets
//	@Produce		json
//	@Param			mission_id	path	string	true	"Mission ID"
//	@Param			target_id	path	string	true	"Target ID"
//	@Success		200
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/missions/{mission_id}/targets/{target_id}/complete [post]
func UpdateTargetCompletion(c *fiber.Ctx) error {
	// Parse the mission ID and target ID from the request URL
	missionID, err := uuid.Parse(c.Params("mid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
	}

	targetID, err := uuid.Parse(c.Params("tid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid target ID"})
	}

	// Retrieve all targets for the mission
	targets, err := repositories.GetTargetsByMissionID(missionID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get targets"})
	}

	// Check if all targets are completed
	completed := true
	for _, target := range targets {
		if !target.Completed {
			completed = false
			break
		}
	}

	// If all targets are completed mark the mission as completed
	if completed {
		err = repositories.UpdateMissionCompletion(missionID, true)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to complete mission"})
		}
		return c.SendStatus(200)
	}

	// Otherwise mark the target as completed
	err = repositories.UpdateTargetCompletion(targetID, true)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to complete target"})
	}
	return c.SendStatus(200)
}
