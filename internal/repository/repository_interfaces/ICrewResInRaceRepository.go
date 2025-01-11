package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"context"
	"github.com/google/uuid"
)

type ICrewResInRaceRepository interface {
	Create(ctx context.Context, crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error)
	Delete(ctx context.Context, raceID uuid.UUID, crewID uuid.UUID) error
	Update(ctx context.Context, crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error)
	GetCrewResByRaceIDAndCrewID(ctx context.Context, raceID uuid.UUID, crewID uuid.UUID) (*models.CrewResInRace, error)
	GetAllCrewResInRace(ctx context.Context, raceID uuid.UUID) ([]models.CrewResInRace, error)
}
