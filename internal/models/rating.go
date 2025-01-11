package models

import (
	"github.com/google/uuid"
	"time"
)

type Rating struct {
	ID         uuid.UUID
	Name       string
	Class      int
	BlowoutCnt int
}

type RatingTableLine struct {
	CrewID                string
	SailNum               int
	ParticipantNames      []string
	ParticipantBirthDates []time.Time
	ParticipantCategories []int
	ResInRace             map[int]int // map[raceNum]Points
	PointsSum             int
	Rank                  int
	CoachNames            []string
}
