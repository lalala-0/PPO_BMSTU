package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"context"
	"github.com/google/uuid"
)

type IProtestRepository interface {
	Create(ctx context.Context, protest *models.Protest) (*models.Protest, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, protest *models.Protest) (*models.Protest, error)
	GetProtestDataByID(ctx context.Context, id uuid.UUID) (*models.Protest, error)
	GetProtestsDataByRaceID(ctx context.Context, raceID uuid.UUID) ([]models.Protest, error)
	AttachCrewToProtest(ctx context.Context, crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error
	DetachCrewFromProtest(ctx context.Context, crewID uuid.UUID, protestID uuid.UUID) error
	GetProtestParticipantsIDByID(ctx context.Context, id uuid.UUID) (map[uuid.UUID]int, error)
}
