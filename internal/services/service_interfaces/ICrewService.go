package service_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type ICrewService interface {
	AddNewCrew(crewID uuid.UUID, ratingID uuid.UUID, class int, sailNum int) (*models.Crew, error)
	DeleteCrewByID(id uuid.UUID) error
	UpdateCrewByID(crewID uuid.UUID, ratingID uuid.UUID, class int, sailNum int) (*models.Crew, error)
	GetCrewDataByID(id uuid.UUID) (*models.Crew, error)
	GetCrewsDataByRatingID(id uuid.UUID) ([]models.Crew, error)
	AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error
	DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error
	ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error
}
