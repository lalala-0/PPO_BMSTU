package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"context"
	"github.com/google/uuid"
)

type IParticipantRepository interface {
	Create(ctx context.Context, participant *models.Participant) (*models.Participant, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, participant *models.Participant) (*models.Participant, error)
	GetParticipantDataByID(ctx context.Context, id uuid.UUID) (*models.Participant, error)
	GetParticipantsDataByCrewID(ctx context.Context, crewID uuid.UUID) ([]models.Participant, error)
	GetParticipantsDataByProtestID(ctx context.Context, crewID uuid.UUID) ([]models.Participant, error)
	GetAllParticipants(ctx context.Context) ([]models.Participant, error)
}
