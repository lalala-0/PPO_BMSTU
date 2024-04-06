package models

import (
	"github.com/google/uuid"
)

type Crew struct {
	ID       uuid.UUID
	RatingID uuid.UUID
	SailNum  int
	Class    string
}
