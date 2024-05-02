package service_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"time"
)

type IParticipantService interface {
	AddNewParticipant(participantID uuid.UUID, fio string, category int, gender int, birthDay time.Time, coach string) (*models.Participant, error)
	DeleteParticipantByID(id uuid.UUID) error
	UpdateParticipantByID(participantID uuid.UUID, fio string, category int, birthDay time.Time, coach string) (*models.Participant, error)
	GetParticipantDataByID(id uuid.UUID) (*models.Participant, error)
	GetParticipantsDataByCrewID(crewID uuid.UUID) ([]models.Participant, error)
	GetParticipantsDataByProtestID(protestID uuid.UUID) ([]models.Participant, error)
}
