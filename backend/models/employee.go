package models

import (
	"time"
)

type Employee struct {
	ID        string    `bson:"_id,omitempty"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required"`
	Position  string    `json:"position" validate:"required"`
	Department string   `json:"department" validate:"required"`
	DateOfHire time.Time `json:"date_of_hire" validate:"required"`
}
