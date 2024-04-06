package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type IRaceRepository interface {
	Create(race *models.Race) (*models.Race, error)
	Delete(id uuid.UUID) error
	Update(race *models.Race) (*models.Race, error)
	GetRaceDataByID(id uuid.UUID) (*models.Race, error)
	GetRacesDataByRatingID(rating_id uuid.UUID) ([]models.Race, error)
}
