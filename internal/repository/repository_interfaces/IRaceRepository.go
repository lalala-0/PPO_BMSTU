package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"context"
	"github.com/google/uuid"
)

type IRaceRepository interface {
	Create(ctx context.Context, race *models.Race) (*models.Race, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, race *models.Race) (*models.Race, error)
	GetRaceDataByID(ctx context.Context, id uuid.UUID) (*models.Race, error)
	GetRacesDataByRatingID(ctx context.Context, rating_id uuid.UUID) ([]models.Race, error)
}
