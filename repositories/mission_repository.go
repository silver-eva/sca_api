package repositories

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"sca_api/database"
	"sca_api/models"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func CreateMission(mission *models.Mission) error {
	var miss_id uuid.UUID
	tx, err := database.DB.Begin()
	if err != nil {
		log.Errorf("Error starting transaction: %v", err)
	}

	query := `INSERT INTO app.mission (name) VALUES ($1) RETURNING id`
	res := tx.QueryRow(query, mission.Name)

	err = res.Scan(&miss_id)
	if err != nil {
		tx.Rollback()
		log.Errorf("Error scanning mission ID: %v", err)
	}

	fmt.Println(miss_id)
	mission.ID = miss_id

	targetIDs := make(map[string]bool)
	for _, target := range mission.Targets {
		// Check for duplicate targets
		if targetIDs[target.Name] {
			return fmt.Errorf("duplicate target found: %s", target.Name)
		}
		targetIDs[target.Name] = true

		// Check if target already exists
		existingTarget, err := GetTargetByName(target.Name)
		if err != nil {
			if err.Error() != "sql: no rows in result set" {
				return err	
			}
		}
		log.Infof("Existing target: %v", existingTarget)
		// Create target if it doesn't exist
		if existingTarget.ID == uuid.Nil {
			target_id, err := CreateTarget(tx, target)
			if err != nil {
				tx.Rollback()
				return err
			} else {
				target.ID = target_id
			}
		} else {
			target.ID = existingTarget.ID
		}
		
		// Attach target to mission
		err = AddTargetToMission(tx, mission.ID, target.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	
	return tx.Commit()
}

func GetMissions() ([]models.Mission, error) {
	var missions []models.Mission
	err := database.DB.Select(&missions, "SELECT * FROM app.mission")
	return missions, err
}

func GetMissionByID(id uuid.UUID) (models.FullMission, error) {
	var mission models.FullMission
	err := database.DB.Get(&mission, `WITH mission_data AS (
    SELECT 
        m.id AS mission_id,
        m.name AS mission_name,
        m.complited AS mission_completed
    FROM app.mission m
    WHERE m.id = $1
),
cats_data AS (
    SELECT 
        mc.mission_id,
        json_agg(json_build_object(
            'id', c.id,
            'name', c.name
        )) AS cats
    FROM app.mission_cats mc
    JOIN app.cat c ON mc.cat_id = c.id
    WHERE mc.mission_id = $1
    GROUP BY mc.mission_id
),
targets_data AS (
    SELECT 
        mt.mission_id,
        json_agg(json_build_object(
            'id', t.id,
            'name', t.name,
            'notes', t.notes,
            'completed', t.complited
        )) AS targets
    FROM app.mission_targets mt
    JOIN app.target t ON mt.target_id = t.id
    WHERE mt.mission_id = $1
    GROUP BY mt.mission_id
)
SELECT 
    md.mission_id AS id,
    md.mission_name AS name,
    md.mission_completed AS complited,
    COALESCE(cd.cats, '[]') AS cats,
    COALESCE(td.targets, '[]') AS targets
FROM mission_data md
LEFT JOIN cats_data cd ON md.mission_id = cd.mission_id
LEFT JOIN targets_data td ON md.mission_id = td.mission_id;
`, id)
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

func AssignCatToMission(missionID, catID uuid.UUID) error {
	var completed bool
	err := database.DB.Get(&completed, "SELECT complited FROM app.mission WHERE id = $1", missionID)
	if err != nil {
		return err
	}
	if completed {
		return fmt.Errorf("mission is already completed, cannot assign cat")
	}

	query := `INSERT INTO app.mission_cats (id, mission_id, cat_id) VALUES ($1, $2, $3)`
	_, err = database.DB.Exec(query, uuid.New(), missionID, catID)
	return err
}

func AddTargetToMission(tx *sql.Tx, missionID, targetID uuid.UUID) error {
	res := tx.QueryRow("SELECT complited FROM app.mission WHERE id = $1", missionID)

	var completed bool
	err := res.Scan(&completed)

	if err != nil {
		return err
	}
	if completed {
		return fmt.Errorf("mission is already completed, cannot add target")
	}

	query := `INSERT INTO app.mission_targets (id, mission_id, target_id) VALUES ($1, $2, $3)`
	_, err = tx.Exec(query, uuid.New(), missionID, targetID)
	return err
}

func RemoveTargetFromMission(missionID, targetID uuid.UUID) error {
	var completed bool
	err := database.DB.Get(&completed, "SELECT complited FROM app.target WHERE id = $1", targetID)
	if err != nil {
		return err
	}
	if completed {
		return fmt.Errorf("target is already completed, cannot be removed")
	}

	query := `DELETE FROM app.mission_targets WHERE mission_id = $1 AND target_id = $2`
	_, err = database.DB.Exec(query, missionID, targetID)
	return err
}
