package repositories

import (
	"sca_api/database"
	"sca_api/models"

	"github.com/google/uuid"
)

func CreateMission(mission *models.Mission) error {
	query := `INSERT INTO app.mission (id, name, complited) VALUES ($1, $2, $3)`
	_, err := database.DB.Exec(query, uuid.New(), mission.Name, mission.Completed)
	return err
}

func GetMissions() ([]models.Mission, error) {
	var missions []models.Mission
	err := database.DB.Select(&missions, "SELECT * FROM app.mission")
	return missions, err
}

func GetMissionByID(id uuid.UUID) (models.Mission, error) {
	var mission models.Mission
	err := database.DB.Get(&mission, "SELECT * FROM app.mission WHERE id = $1", id)
	return mission, err
}

func DeleteMission(id uuid.UUID) error {
	_, err := database.DB.Exec("DELETE FROM app.mission WHERE id = $1", id)
	return err
}

func UpdateMissionCompletion(id uuid.UUID, completed bool) error {
	_, err := database.DB.Exec("UPDATE app.mission SET complited = $1 WHERE id = $2", completed, id)
	return err
}
