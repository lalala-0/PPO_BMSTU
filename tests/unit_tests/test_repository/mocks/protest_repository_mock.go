package mocks

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// MockIProtestRepository представляет мок-репозиторий для Protest
type MockIProtestRepository struct {
	CreateFunc                       func(*models.Protest) (*models.Protest, error)
	DeleteFunc                       func(uuid.UUID) error
	UpdateFunc                       func(*models.Protest) (*models.Protest, error)
	GetProtestDataByIDFunc           func(uuid.UUID) (*models.Protest, error)
	GetProtestsDataByRaceIDFunc      func(uuid.UUID) ([]models.Protest, error)
	GetProtestParticipantsIDByIDFunc func(uuid.UUID) (map[uuid.UUID]int, error)
	AttachCrewToProtestFunc          func(protestID, crewID uuid.UUID, crewStatus int) error
	DetachCrewFromProtestFunc        func(protestID, crewID uuid.UUID) error
}

// Create вызывает функцию мока для создания протеста
func (m *MockIProtestRepository) Create(protest *models.Protest) (*models.Protest, error) {
	return m.CreateFunc(protest)
}

// Delete вызывает функцию мока для удаления протеста по ID
func (m *MockIProtestRepository) Delete(id uuid.UUID) error {
	return m.DeleteFunc(id)
}

// Update вызывает функцию мока для обновления данных протеста
func (m *MockIProtestRepository) Update(protest *models.Protest) (*models.Protest, error) {
	return m.UpdateFunc(protest)
}

// GetProtestDataByID вызывает функцию мока для получения данных протеста по ID
func (m *MockIProtestRepository) GetProtestDataByID(id uuid.UUID) (*models.Protest, error) {
	return m.GetProtestDataByIDFunc(id)
}

// GetProtestsDataByRaceID вызывает функцию мока для получения протестов по ID гонки
func (m *MockIProtestRepository) GetProtestsDataByRaceID(raceID uuid.UUID) ([]models.Protest, error) {
	return m.GetProtestsDataByRaceIDFunc(raceID)
}

// GetProtestParticipantsIDByID вызывает функцию мока для получения ID участников протеста по ID
func (m *MockIProtestRepository) GetProtestParticipantsIDByID(id uuid.UUID) (map[uuid.UUID]int, error) {
	return m.GetProtestParticipantsIDByIDFunc(id)
}

// AttachCrewToProtest вызывает функцию мока для привязки команды к протесту
func (m *MockIProtestRepository) AttachCrewToProtest(protestID, crewID uuid.UUID, crewStatus int) error {
	return m.AttachCrewToProtestFunc(protestID, crewID, crewStatus)
}

// DetachCrewFromProtest вызывает функцию мока для отвязывания команды от протеста
func (m *MockIProtestRepository) DetachCrewFromProtest(protestID, crewID uuid.UUID) error {
	return m.DetachCrewFromProtestFunc(protestID, crewID)
}
