package models

import (
	uuid "github.com/google/uuid"
)

type Cat struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Name       string    `db:"name" json:"name" validate:"required"`
	Experience float64   `db:"experience" json:"experience" validate:"gte=1,lte=15"`
	Breed      string    `db:"breed" json:"breed" validate:"required"`
	Salary     float64   `db:"salary" json:"salary"`
}