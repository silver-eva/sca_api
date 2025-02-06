package models

import "github.com/google/uuid"

type Target struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" validate:"required"`
	Country   string    `db:"country" json:"country" validate:"required"`
	Notes     string    `db:"notes" json:"notes"`
	Completed bool      `db:"complited" json:"completed"`
}
