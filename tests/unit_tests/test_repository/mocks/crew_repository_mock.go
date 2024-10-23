package mocks

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// MockICrewRepository представляет мок-репозиторий для экипажей
type MockICrewRepository struct {
	CreateFunc                          func(*models.Crew) (*models.Crew, error)
	DeleteFunc                          func(uuid.UUID) error
	GetCrewDataByIDFunc                 func(uuid.UUID) (*models.Crew, error)
	GetCrewDataBySailNumAndRatingIDFunc func(int, uuid.UUID) (*models.Crew, error)
	GetCrewsDataByRatingIDFunc          func(uuid.UUID) ([]models.Crew, error)
	GetCrewsDataByProtestIDFunc         func(uuid.UUID) ([]models.Crew, error)
	AttachParticipantToCrewFunc         func(uuid.UUID, uuid.UUID, int) error
	DetachParticipantFromCrewFunc       func(uuid.UUID, uuid.UUID) error
	ReplaceParticipantStatusInCrewFunc  func(uuid.UUID, uuid.UUID, int, int) error
	UpdateFunc                          func(*models.Crew) (*models.Crew, error)
}

// Create вызывает функцию мока для создания экипажа
func (m *MockICrewRepository) Create(crew *models.Crew) (*models.Crew, error) {
	return m.CreateFunc(crew)
}

// Delete вызывает функцию мока для удаления экипажа по ID
func (m *MockICrewRepository) Delete(id uuid.UUID) error {
	return m.DeleteFunc(id)
}

// GetCrewDataByID вызывает функцию мока для получения экипажа по ID
func (m *MockICrewRepository) GetCrewDataByID(id uuid.UUID) (*models.Crew, error) {
	return m.GetCrewDataByIDFunc(id)
}

// GetCrewDataBySailNumAndRatingID вызывает функцию мока для получения экипажа по номеру паруса и ID рейтинга
func (m *MockICrewRepository) GetCrewDataBySailNumAndRatingID(sailNum int, ratingID uuid.UUID) (*models.Crew, error) {
	return m.GetCrewDataBySailNumAndRatingIDFunc(sailNum, ratingID)
}

// GetCrewsDataByRatingID вызывает функцию мока для получения экипажей по ID рейтинга
func (m *MockICrewRepository) GetCrewsDataByRatingID(id uuid.UUID) ([]models.Crew, error) {
	return m.GetCrewsDataByRatingIDFunc(id)
}

// GetCrewsDataByProtestID вызывает функцию мока для получения экипажей по ID протеста
func (m *MockICrewRepository) GetCrewsDataByProtestID(id uuid.UUID) ([]models.Crew, error) {
	return m.GetCrewsDataByProtestIDFunc(id)
}

// AttachParticipantToCrew вызывает функцию мока для прикрепления участника к экипажу
func (m *MockICrewRepository) AttachParticipantToCrew(crewID, participantID uuid.UUID, helmsman int) error {
	return m.AttachParticipantToCrewFunc(crewID, participantID, helmsman)
}

// DetachParticipantFromCrew вызывает функцию мока для отсоединения участника от экипажа
func (m *MockICrewRepository) DetachParticipantFromCrew(crewID, participantID uuid.UUID) error {
	return m.DetachParticipantFromCrewFunc(crewID, participantID)
}

// ReplaceParticipantStatusInCrew вызывает функцию мока для замены статуса участника в экипаже
func (m *MockICrewRepository) ReplaceParticipantStatusInCrew(crewID, participantID uuid.UUID, helmsman int, active int) error {
	return m.ReplaceParticipantStatusInCrewFunc(crewID, participantID, helmsman, active)
}

// Update вызывает функцию мока для обновления данных экипажа
func (m *MockICrewRepository) Update(crew *models.Crew) (*models.Crew, error) {
	return m.UpdateFunc(crew)
}
