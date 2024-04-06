package service_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"time"
)

type IRaceService interface {
	AddNewRace(raceID uuid.UUID, ratingID uuid.UUID, number int, date time.Time, class string) (*models.Race, error)
	DeleteRaceByID(id uuid.UUID) error
	UpdateRaceByID(raceID uuid.UUID, number int, date time.Time, class string) (*models.Race, error)
	GetRaceDataByID(id uuid.UUID) (*models.Race, error)
	GetRacesDataByRatingID(ratingID uuid.UUID) ([]models.Race, error)
	MakeStartProcedure(raceID uuid.UUID, falseStartYachtList map[int]string) error
	MakeFinishProcedure(raceID uuid.UUID, finishersList map[int]int, nonFinishersList map[int]string) error
}
