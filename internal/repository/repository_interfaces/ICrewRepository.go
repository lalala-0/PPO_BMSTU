package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"context"
	"github.com/google/uuid"
)

type ICrewRepository interface {
	Create(ctx context.Context, race *models.Crew) (*models.Crew, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, crew *models.Crew) (*models.Crew, error)
	GetCrewDataByID(ctx context.Context, id uuid.UUID) (*models.Crew, error)
	GetCrewsDataByRatingID(ctx context.Context, id uuid.UUID) ([]models.Crew, error)
	GetCrewsDataByProtestID(ctx context.Context, id uuid.UUID) ([]models.Crew, error)
	GetCrewDataBySailNumAndRatingID(ctx context.Context, sailNum int, ratingID uuid.UUID) (*models.Crew, error)
	AttachParticipantToCrew(ctx context.Context, participantID uuid.UUID, crewID uuid.UUID, helmsman int) error
	DetachParticipantFromCrew(ctx context.Context, participantID uuid.UUID, crewID uuid.UUID) error
	ReplaceParticipantStatusInCrew(ctx context.Context, participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error
}
