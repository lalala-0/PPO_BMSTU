package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type IProtestRepository interface {
	Create(protest *models.Protest) (*models.Protest, error)
	Delete(id uuid.UUID) error
	Update(protest *models.Protest) (*models.Protest, error)
	GetProtestDataByID(id uuid.UUID) (*models.Protest, error)
	GetProtestsDataByRaceID(raceID uuid.UUID) ([]models.Protest, error)
	AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error
	DetachCrewFromProtest(crewID uuid.UUID, protestID uuid.UUID) error
	GetProtestParticipantsIDByID(id uuid.UUID) (map[uuid.UUID]int, error)
}
