package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type IParticipantRepository interface {
	Create(participant *models.Participant) (*models.Participant, error)
	Delete(id uuid.UUID) error
	Update(participant *models.Participant) (*models.Participant, error)
	GetParticipantDataByID(id uuid.UUID) (*models.Participant, error)
	GetParticipantsDataByCrewID(crewID uuid.UUID) ([]models.Participant, error)
}
