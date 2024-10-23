package mocks

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// MockRaceRepository представляет мок-репозиторий для Race
type MockRaceRepository struct {
	CreateFunc                 func(*models.Race) (*models.Race, error)
	DeleteFunc                 func(uuid.UUID) error
	UpdateFunc                 func(*models.Race) (*models.Race, error)
	GetRaceDataByIDFunc        func(uuid.UUID) (*models.Race, error)
	GetRacesDataByRatingIDFunc func(uuid.UUID) ([]models.Race, error)
}

// Create вызывает функцию мока для создания гонки
func (m *MockRaceRepository) Create(race *models.Race) (*models.Race, error) {
	return m.CreateFunc(race)
}

// Delete вызывает функцию мока для удаления гонки по ID
func (m *MockRaceRepository) Delete(id uuid.UUID) error {
	return m.DeleteFunc(id)
}

// Update вызывает функцию мока для обновления данных гонки
func (m *MockRaceRepository) Update(race *models.Race) (*models.Race, error) {
	return m.UpdateFunc(race)
}

// GetRaceDataByID вызывает функцию мока для получения данных гонки по ID
func (m *MockRaceRepository) GetRaceDataByID(id uuid.UUID) (*models.Race, error) {
	return m.GetRaceDataByIDFunc(id)
}

// GetRacesDataByRatingID вызывает функцию мока для получения гонок по ID рейтинга
func (m *MockRaceRepository) GetRacesDataByRatingID(ratingID uuid.UUID) ([]models.Race, error) {
	return m.GetRacesDataByRatingIDFunc(ratingID)
}
