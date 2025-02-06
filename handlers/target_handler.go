package handlers

import (
	"sca_api/models"
	"sca_api/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTarget(c *fiber.Ctx) error {
	var target models.Target
	if err := c.BodyParser(&target); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := repositories.CreateTarget(&target)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create target"})
	}
	return c.Status(201).JSON(target)
}

func GetTargets(c *fiber.Ctx) error {
	targets, err := repositories.GetTargets()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch targets"})
	}
	return c.JSON(targets)
}

func UpdateTargetNotes(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var data struct {
		Notes string `json:"notes"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = repositories.UpdateTargetNotes(id, data.Notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update target notes"})
	}
	return c.SendStatus(200)
}

func UpdateTargetCompletion(c *fiber.Ctx) error {
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

	err = repositories.UpdateTargetCompletion(id, data.Completed)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update target completion"})
	}
	return c.SendStatus(200)
}