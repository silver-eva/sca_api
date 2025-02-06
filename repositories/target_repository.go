package repositories

import (
	"database/sql"
	"sca_api/database"
	"sca_api/models"

	"github.com/google/uuid"
)

func CreateTarget(tx *sql.Tx, target *models.Target) (uuid.UUID, error) {
	query := `INSERT INTO app.target (id, name, country, notes, complited) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	res := tx.QueryRow(query, uuid.New(), target.Name, target.Country, target.Notes, target.Completed)

	var id uuid.UUID
	err := res.Scan(&id)
	return id, err
}

func GetTargets() ([]models.Target, error) {
	var targets []models.Target
	err := database.DB.Select(&targets, "SELECT * FROM app.target")
	return targets, err
}

func GetTargetByName(name string) (models.Target, error) {
	var target models.Target
	err := database.DB.Get(&target, "SELECT * FROM app.target WHERE name = $1", name)
	return target, err
}

func DeleteTarget(id uuid.UUID) error {
	_, err := database.DB.Exec("DELETE FROM app.target WHERE id = $1", id)
	return err
}

func UpdateTargetNotes(id uuid.UUID, notes string) error {
	_, err := database.DB.Exec("UPDATE app.target SET notes = $1 WHERE id = $2 AND complited = false", notes, id)
	return err
}

func UpdateTargetCompletion(id uuid.UUID, completed bool) error {
	_, err := database.DB.Exec("UPDATE app.target SET complited = $1 WHERE id = $2", completed, id)
	return err
}
