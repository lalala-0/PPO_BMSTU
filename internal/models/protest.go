package models

import (
	"github.com/google/uuid"
	"time"
)

type Protest struct {
	ID         uuid.UUID
	RaceID     uuid.UUID
	JudgeID    uuid.UUID
	RuleNum    int
	ReviewDate time.Time
	Status     int
	Comment    string
	RatingID   uuid.UUID
}

//// Protest status vars
//const PendingReview = 1
//const Reviewed = 2
//
//// Protest participants role vars
//const Protestor = 1
//const Protestee = 2
//const Witness = 3
