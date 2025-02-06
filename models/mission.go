package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Mission struct {
	ID        uuid.UUID   `db:"id" json:"id"`
	Name      string      `db:"name" json:"name" validate:"required"`
	Completed bool        `db:"complited" json:"completed"` 
	Targets   []*Target   `json:"targets" db:"targets" validate:"required,dive,required,len=1|2|3"`
}

type FullMission struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Complited bool            `json:"complited"`
	Cats      json.RawMessage `json:"cats"`
	Targets   json.RawMessage `json:"targets"`
}