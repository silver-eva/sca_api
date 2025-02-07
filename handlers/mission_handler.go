package handlers

import (
	"sca_api/models"
	"sca_api/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateMission creates a new mission
//
//	@Summary		Create a mission
//	@Description	Add a new mission to the database
//	@Tags			Missions
//	@Accept			json
//	@Produce		json
//	@Param			mission	body		models.Mission	true	"Mission Data"
//	@Success		201		{object}	models.Mission
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/missions [post]
func CreateMission(c *fiber.Ctx) error {
	var mission models.Mission
	// Parse the request body into the mission struct
	if err := c.BodyParser(&mission); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Call the repository function to create a new mission
	err := repositories.CreateMission(&mission)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create mission"})
	}

	// Return the created mission with a 201 status code
	return c.Status(201).JSON(mission)
}

// GetMissions retrieves all missions
//
//	@Summary		List all missions
//	@Description	Get a list of all missions
//	@Tags			Missions
//	@Produce		json
//	@Param			limit	query		int	false	"Limit"
//	@Param			page	query		int	false	"Page"
//	@Success		200		{array}		models.Mission
//	@Failure		500		{object}	map[string]string
//	@Router			/missions [get]
func GetMissions(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("page", 0)

	missions, err := repositories.GetMissions(limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch missions"})
	}
	// Return the list of missions with a 200 status code
	return c.JSON(missions)
}

// GetMission retrieves a single mission by ID
//
//	@Summary		Retrieve a mission
//	@Description	Get a single mission by ID
//	@Tags			Missions
//	@Produce		json
//	@Param			id	path		string	true	"Mission ID"
//	@Success		200	{object}	models.Mission
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/missions/{id} [get]
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

// DeleteMission deletes a mission by ID
//
//	@Summary		Delete a mission
//	@Description	Delete a single mission by ID
//	@Tags			Missions
//	@Produce		json
//	@Param			id	path	string	true	"Mission ID"
//	@Success		204
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/missions/{id} [delete]
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

// UpdateMissionCompletion updates the completion status of a mission
//
//	@Summary		Update mission completion
//	@Description	Update the completion status of a mission
//	@Tags			Missions
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Mission ID"
//	@Success		200
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/missions/{id}/complete [post]
func UpdateMissionCompletion(c *fiber.Ctx) error {
	// Parse the mission ID from the request URL
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = repositories.UpdateMissionCompletion(id, true)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to complete mission"})
	}

	// Return a 200 status code if successful
	return c.SendStatus(200)
}

// AssignCatToMission assigns a spy cat to a mission
//
//	@Summary		Assign a spy cat to a mission
//	@Description	Assign a spy cat to a mission
//	@Tags			Missions
//	@Accept			json
//	@Produce		json
//	@Param			mid	path	string	true	"Mission ID"
//	@Param			cid	path	string	true	"Cat ID"
//	@Success		200
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/missions/{mid}/cats/{cid} [post]
func AssignCatToMission(c *fiber.Ctx) error {

	// Parse the mission ID from the request URL
	missionID, err := uuid.Parse(c.Params("mid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
	}
	catID, err := uuid.Parse(c.Params("cid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid cat ID"})
	}

	// Call the repository function to assign the cat to the mission
	err = repositories.AssignCatToMission(missionID, catID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Return a 200 status code if successful
	return c.SendStatus(200)
}

// AddTargetToMission handles adding a target to a mission
//
//	@Summary		Add target to mission
//	@Description	Add a target to a specific mission
//	@Tags			Missions
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string			true	"Mission ID"
//	@Param			target	body	models.Target	true	"Target Data"
//	@Success		200
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/missions/{id}/targets [post]
func AddTargetToMission(c *fiber.Ctx) error {
	var target models.Target
	// Parse the request body into the target struct
	if err := c.BodyParser(&target); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Parse the mission ID from the request URL
	missionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
	}

	// Retrieve the mission by ID
	mission, err := repositories.GetMissionByID(missionID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mission not found"})
	}

	// Check if the mission is already completed
	if mission.Complited { // TODO: refactor naming
		return c.Status(400).JSON(fiber.Map{"error": "Cannot add target to completed mission"})
	}

	// Retrieve the target by name
	target, err = repositories.GetTargetByName(target.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// If target does not exist, create a new target
	if target.ID == uuid.Nil {
		target.ID, err = repositories.CreateTarget(&target)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	} else {
		// Check if the target is already assigned to the mission
		can, err := repositories.CanAssignTargetToMission(missionID, target.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// If target is already assigned, return an error
		if !can {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot assign target to mission"})
		}
	}

	// Add the target to the mission
	err = repositories.AddTargetToMission(missionID, target.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Return a 200 status code if successful
	return c.SendStatus(200)
}

// RemoveTargetFromMission removes a target from a mission
//
//	@Summary		Remove target from mission
//	@Description	Remove a target from a mission
//	@Tags			Missions
//	@Param			mission_id	path	string	true	"Mission ID"
//	@Param			target_id	path	string	true	"Target ID"
//	@Success		200
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/missions/{mission_id}/targets/{target_id} [delete]
func RemoveTargetFromMission(c *fiber.Ctx) error {
	missionID, err := uuid.Parse(c.Params("mid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid mission ID"})
	}

	targetID, err := uuid.Parse(c.Params("tid"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid target ID"})
	}

	mission, err := repositories.GetMissionByID(missionID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Mission not found"})
	}

	if mission.Complited {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot remove target from completed mission"})
	}

	err = repositories.RemoveTargetFromMission(missionID, targetID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(200)
}
