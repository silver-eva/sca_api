package models

import "github.com/google/uuid"

type Mission struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" validate:"required"`
	Completed bool      `db:"complited" json:"completed"`
}
