package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type ICrewRepository interface {
	Create(race *models.Crew) (*models.Crew, error)
	Delete(id uuid.UUID) error
	Update(crew *models.Crew) (*models.Crew, error)
	GetCrewDataByID(id uuid.UUID) (*models.Crew, error)
	GetCrewsDataByRatingID(id uuid.UUID) ([]models.Crew, error)
	GetCrewsDataByProtestID(id uuid.UUID) ([]models.Crew, error)
	GetCrewDataBySailNumAndRatingID(sailNum int, ratingID uuid.UUID) (*models.Crew, error)
	AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error
	DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error
	ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error
}
