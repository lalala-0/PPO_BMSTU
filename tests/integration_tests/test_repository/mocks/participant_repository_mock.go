package mocks

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// MockIParticipantRepository представляет мок-репозиторий для Participant
type MockIParticipantRepository struct {
	CreateFunc                         func(*models.Participant) (*models.Participant, error)
	DeleteFunc                         func(uuid.UUID) error
	UpdateFunc                         func(*models.Participant) (*models.Participant, error)
	GetParticipantDataByIDFunc         func(uuid.UUID) (*models.Participant, error)
	GetAllParticipantsFunc             func() ([]models.Participant, error)
	GetParticipantsDataByCrewIDFunc    func(uuid.UUID) ([]models.Participant, error)
	GetParticipantsDataByProtestIDFunc func(uuid.UUID) ([]models.Participant, error)
}

// Create вызывает функцию мока для создания участника
func (m *MockIParticipantRepository) Create(participant *models.Participant) (*models.Participant, error) {
	return m.CreateFunc(participant)
}

// Delete вызывает функцию мока для удаления участника по ID
func (m *MockIParticipantRepository) Delete(id uuid.UUID) error {
	return m.DeleteFunc(id)
}

// Update вызывает функцию мока для обновления данных участника
func (m *MockIParticipantRepository) Update(participant *models.Participant) (*models.Participant, error) {
	return m.UpdateFunc(participant)
}

// GetParticipantDataByID вызывает функцию мока для получения данных участника по ID
func (m *MockIParticipantRepository) GetParticipantDataByID(id uuid.UUID) (*models.Participant, error) {
	return m.GetParticipantDataByIDFunc(id)
}

// GetAllParticipants вызывает функцию мока для получения всех участников
func (m *MockIParticipantRepository) GetAllParticipants() ([]models.Participant, error) {
	return m.GetAllParticipantsFunc()
}

// GetParticipantsDataByCrewID вызывает функцию мока для получения участников по ID команды
func (m *MockIParticipantRepository) GetParticipantsDataByCrewID(id uuid.UUID) ([]models.Participant, error) {
	return m.GetParticipantsDataByCrewIDFunc(id)
}

// GetParticipantsDataByProtestID вызывает функцию мока для получения участников по ID протеста
func (m *MockIParticipantRepository) GetParticipantsDataByProtestID(id uuid.UUID) ([]models.Participant, error) {
	return m.GetParticipantsDataByProtestIDFunc(id)
}
