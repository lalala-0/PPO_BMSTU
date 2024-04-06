package models

import (
	"github.com/google/uuid"
	"time"
)

type Protest struct {
	ID          uuid.UUID
	RaceID      uuid.UUID
	ProtesteeID uuid.UUID
	RuleNum     int
	ReviewDate  time.Time
	Status      int
	Comment     string
	RatingID    uuid.UUID
}

const PendingReview = 1
const Reviewed = 2

const Protestor = 1
const Protestee = 2
const Witness = 3
