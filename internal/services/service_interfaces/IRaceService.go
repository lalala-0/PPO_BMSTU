package service_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"time"
)

type IRaceService interface {
	AddNewRace(raceID uuid.UUID, ratingID uuid.UUID, number int, date time.Time, class int) (*models.Race, error)
	DeleteRaceByID(id uuid.UUID) error
	UpdateRaceByID(raceID uuid.UUID, ratingID uuid.UUID, number int, date time.Time, class int) (*models.Race, error)
	GetRaceDataByID(id uuid.UUID) (*models.Race, error)
	GetRacesDataByRatingID(ratingID uuid.UUID) ([]models.Race, error)
	GetAllCrewResInRace(race *models.Race) ([]models.CrewResInRace, error)
	MakeStartProcedure(raceID uuid.UUID, falseStartYachtList map[int]int) error
	MakeFinishProcedure(raceID uuid.UUID, finishersList map[int]int, nonFinishersList map[int]int) error
}
