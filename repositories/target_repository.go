package repositories

import (
	"database/sql"
	"sca_api/database"
	"sca_api/models"

	"github.com/google/uuid"
)

func CreateTargetSafe(tx *sql.Tx, target *models.Target) (uuid.UUID, error) {
	query := `INSERT INTO app.target (id, name, country, notes, complited) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	res := tx.QueryRow(query, uuid.New(), target.Name, target.Country, target.Notes, target.Completed)

	var id uuid.UUID
	err := res.Scan(&id)
	return id, err
}
func CreateTarget(target *models.Target) (uuid.UUID, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()
	id, err := CreateTargetSafe(tx, target)
	if err != nil {
		return uuid.Nil, err
	}
	return id, tx.Commit()
}

func GetTargetByName(name string) (models.Target, error) {
	var target models.Target
	err := database.DB.Get(&target, "SELECT * FROM app.target WHERE name = $1", name)
	return target, err
}

func GetTargetById(id uuid.UUID) (models.Target, error) {
	var target models.Target
	err := database.DB.Get(&target, "SELECT * FROM app.target WHERE id = $1", id)
	return target, err
}

func DeleteTarget(id uuid.UUID) error {
	_, err := database.DB.Exec("DELETE FROM app.target WHERE id = $1", id)
	return err
}

func UpdateTargetNotes(id uuid.UUID, notes string) error {
	_, err := database.DB.Exec("UPDATE app.target SET notes = $1 WHERE id = $2", notes, id)
	return err
}

func UpdateTargetCompletion(id uuid.UUID, completed bool) error {
	_, err := database.DB.Exec("UPDATE app.target SET complited = $1 WHERE id = $2", completed, id)
	return err
}

func CanAssignTargetToMission(targetID, missionID uuid.UUID) (bool, error) {
	var count int
	var assigned bool
	err := database.DB.Get(&count, "SELECT COUNT(*) FROM app.mission_targets WHERE mission_id = $1", missionID)
	if err != nil {
		return false, err
	}
	err = database.DB.Get(&assigned, "SELECT COUNT(*) FROM app.mission_targets WHERE target_id = $1 AND mission_id = $2", targetID, missionID)
	if err != nil {
		return false, err
	}
	if assigned {
		return false, nil
	}
	if count == 3 {
		return false, nil
	}
	return true, nil
}

func GetTargetsByMissionID(missionID uuid.UUID) ([]models.Target, error) {
	var targets []models.Target
	err := database.DB.Select(&targets, "SELECT * FROM app.target WHERE id IN (SELECT target_id FROM app.mission_targets WHERE mission_id = $1)", missionID)
	return targets, err
}
