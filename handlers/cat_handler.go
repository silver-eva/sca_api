package handlers

import (
	"sca_api/models"
	"sca_api/repositories"
	"sca_api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateCat(c *fiber.Ctx) error {
	var cat models.Cat
	if err := c.BodyParser(&cat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validate breed using external API (you will implement this in utils/validation.go)
	if !utils.IsValidBreed(cat.Breed) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid breed"})
	}

	new_cat, err := repositories.CreateCat(&cat)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create cat"})
	}
	return c.Status(201).JSON(new_cat)
}

func GetCats(c *fiber.Ctx) error {
	cats, err := repositories.GetCats()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch cats"})
	}
	return c.JSON(cats)
}

func GetCat(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	cat, err := repositories.GetCatByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Cat not found"})
	}
	return c.JSON(cat)
}

func DeleteCat(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = repositories.DeleteCat(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete cat"})
	}
	return c.SendStatus(204)
}

func UpdateCatSalary(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var data struct {
		Salary float64 `json:"salary"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = repositories.UpdateCatSalary(id, data.Salary)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update salary"})
	}
	return c.SendStatus(200)
}
