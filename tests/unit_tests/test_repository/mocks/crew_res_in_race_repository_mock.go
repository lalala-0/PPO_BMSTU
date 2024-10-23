package mocks

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// MockCrewResInRaceRepository представляет мок-репозиторий для CrewResInRace
type MockCrewResInRaceRepository struct {
	CreateFunc                      func(*models.CrewResInRace) (*models.CrewResInRace, error)
	DeleteFunc                      func(uuid.UUID, uuid.UUID) error
	UpdateFunc                      func(*models.CrewResInRace) (*models.CrewResInRace, error)
	GetCrewResByRaceIDAndCrewIDFunc func(uuid.UUID, uuid.UUID) (*models.CrewResInRace, error)
	GetAllCrewResInRaceFunc         func(uuid.UUID) ([]models.CrewResInRace, error)
}

// Create вызывает функцию мока для создания записи CrewResInRace
func (m *MockCrewResInRaceRepository) Create(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	return m.CreateFunc(crewResInRace)
}

// Delete вызывает функцию мока для удаления записи CrewResInRace по ID гонки и экипажа
func (m *MockCrewResInRaceRepository) Delete(raceID uuid.UUID, crewID uuid.UUID) error {
	return m.DeleteFunc(raceID, crewID)
}

// Update вызывает функцию мока для обновления данных CrewResInRace
func (m *MockCrewResInRaceRepository) Update(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	return m.UpdateFunc(crewResInRace)
}

// GetCrewResByRaceIDAndCrewID вызывает функцию мока для получения записи CrewResInRace по ID гонки и экипажа
func (m *MockCrewResInRaceRepository) GetCrewResByRaceIDAndCrewID(raceID uuid.UUID, crewID uuid.UUID) (*models.CrewResInRace, error) {
	return m.GetCrewResByRaceIDAndCrewIDFunc(raceID, crewID)
}

// GetAllCrewResInRace вызывает функцию мока для получения всех записей CrewResInRace для указанной гонки
func (m *MockCrewResInRaceRepository) GetAllCrewResInRace(raceID uuid.UUID) ([]models.CrewResInRace, error) {
	return m.GetAllCrewResInRaceFunc(raceID)
}
