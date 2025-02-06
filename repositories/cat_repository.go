package repositories

import (
	"sca_api/database"
	"sca_api/models"

	"github.com/google/uuid"
)

func CreateCat(cat *models.Cat) (models.Cat,error) {
	query := `INSERT INTO app.cat (id, name, experience, breed, salary) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, experience, breed, salary`
	new_cat := models.Cat{}
	err := database.DB.Get(&new_cat, query, uuid.New(), cat.Name, cat.Experience, cat.Breed, cat.Salary)
	return new_cat, err
}

func GetCats() ([]models.Cat, error) {
	var cats []models.Cat
	err := database.DB.Select(&cats, "SELECT * FROM app.cat")
	return cats, err
}

func GetCatByID(id uuid.UUID) (models.Cat, error) {
	var cat models.Cat
	err := database.DB.Get(&cat, "SELECT * FROM app.cat WHERE id = $1", id)
	return cat, err
}

func DeleteCat(id uuid.UUID) error {
	_, err := database.DB.Exec("DELETE FROM app.cat WHERE id = $1", id)
	return err
}

func UpdateCatSalary(id uuid.UUID, salary float64) error {
	_, err := database.DB.Exec("UPDATE app.cat SET salary = $1 WHERE id = $2", salary, id)
	return err
}