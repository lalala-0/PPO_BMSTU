package models

import (
	"github.com/google/uuid"
)

type CrewResInRace struct {
	CrewID           uuid.UUID
	RaceID           uuid.UUID
	Points           int
	SpecCircumstance string
}
