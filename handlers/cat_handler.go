package handlers

import (
	"sca_api/models"
	"sca_api/repositories"
	"sca_api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateCat creates a new spy cat
//
//	@Summary		Create a spy cat
//	@Description	Add a new spy cat to the database
//	@Tags			Cats
//	@Accept			json
//	@Produce		json
//	@Param			cat	body		models.Cat	true	"Spy Cat Data"
//	@Success		201	{object}	models.Cat
//	@Failure		400	{object}	map[string]string
//	@Router			/cats [post]
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

// GetCats retrieves all spy cats
//
//	@Summary	List all spy cats
//	@Tags		Cats
//	@Produce	json
//	@Param		limit	query		int	false	"Limit"
//	@Param		page	query		int	false	"Page"
//	@Success	200		{array}		models.Cat
//	@Failure	500		{object}	map[string]string
//	@Router		/cats [get]
func GetCats(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 0)
	cats, err := repositories.GetCats(limit, page)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch cats"})
	}
	return c.JSON(cats)
}

// GetCat retrieves a single spy cat by ID
//
//	@Summary	Retrieve a spy cat
//	@Tags		Cats
//	@Produce	json
//	@Param		id	path		string	true	"Cat ID"
//	@Success	200	{object}	models.Cat
//	@Failure	400	{object}	map[string]string
//	@Failure	404	{object}	map[string]string
//	@Router		/cats/{id} [get]
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

// DeleteCat deletes a spy cat by ID
//
//	@Summary	Delete a spy cat
//	@Tags		Cats
//	@Param		id	path	string	true	"Cat ID"
//	@Success	204
//	@Failure	400	{object}	map[string]string
//	@Failure	500	{object}	map[string]string
//	@Router		/cats/{id} [delete]
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

// UpdateCatSalary updates the salary of a spy cat
//
//	@Summary	Update the salary of a spy cat
//	@Tags		Cats
//	@Param		id		path	string	true	"Cat ID"
//	@Param		salary	body	string	true	"Salary data"
//	@Success	200
//	@Failure	400	{object}	map[string]string
//	@Failure	500	{object}	map[string]string
//	@Router		/cats/{id}/salary [put]
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
