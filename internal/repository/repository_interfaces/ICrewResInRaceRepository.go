package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type ICrewResInRaceRepository interface {
	Create(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error)
	Delete(raceID uuid.UUID, crewID uuid.UUID) error
	Update(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error)
	GetCrewResByRaceIDAndCrewID(raceID uuid.UUID, crewID uuid.UUID) (*models.CrewResInRace, error)
}
