package models

import (
	"github.com/google/uuid"
	"time"
)

type Race struct {
	ID       uuid.UUID
	RatingID uuid.UUID
	Date     time.Time
	Number   int
	Class    string
}
