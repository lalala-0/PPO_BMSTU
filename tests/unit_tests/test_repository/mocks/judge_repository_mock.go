package mocks

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// MockJudgeRepository представляет мок-репозиторий для Judge
type MockJudgeRepository struct {
	CreateProfileFunc           func(*models.Judge) (*models.Judge, error)
	DeleteProfileFunc           func(uuid.UUID) error
	UpdateProfileFunc           func(*models.Judge) (*models.Judge, error)
	GetJudgeDataByIDFunc        func(uuid.UUID) (*models.Judge, error)
	GetJudgeDataByProtestIDFunc func(uuid.UUID) (*models.Judge, error)
	GetJudgesDataByRatingIDFunc func(uuid.UUID) ([]models.Judge, error)
	GetAllJudgesFunc            func() ([]models.Judge, error)
	GetJudgeDataByLoginFunc     func(string) (*models.Judge, error)
}

// CreateProfile вызывает функцию мока для создания профиля судьи
func (m *MockJudgeRepository) CreateProfile(judge *models.Judge) (*models.Judge, error) {
	return m.CreateProfileFunc(judge)
}

// DeleteProfile вызывает функцию мока для удаления профиля судьи по ID
func (m *MockJudgeRepository) DeleteProfile(id uuid.UUID) error {
	return m.DeleteProfileFunc(id)
}

// UpdateProfile вызывает функцию мока для обновления данных профиля судьи
func (m *MockJudgeRepository) UpdateProfile(judge *models.Judge) (*models.Judge, error) {
	return m.UpdateProfileFunc(judge)
}

// GetJudgeDataByID вызывает функцию мока для получения данных судьи по ID
func (m *MockJudgeRepository) GetJudgeDataByID(id uuid.UUID) (*models.Judge, error) {
	return m.GetJudgeDataByIDFunc(id)
}

// GetJudgeDataByProtestID вызывает функцию мока для получения данных судьи по ID протеста
func (m *MockJudgeRepository) GetJudgeDataByProtestID(protestID uuid.UUID) (*models.Judge, error) {
	return m.GetJudgeDataByProtestIDFunc(protestID)
}

// GetJudgesDataByRatingID вызывает функцию мока для получения данных судей по ID рейтинга
func (m *MockJudgeRepository) GetJudgesDataByRatingID(ratingID uuid.UUID) ([]models.Judge, error) {
	return m.GetJudgesDataByRatingIDFunc(ratingID)
}

// GetAllJudges вызывает функцию мока для получения всех судей
func (m *MockJudgeRepository) GetAllJudges() ([]models.Judge, error) {
	return m.GetAllJudgesFunc()
}

// GetJudgeDataByLogin вызывает функцию мока для получения данных судьи по логину
func (m *MockJudgeRepository) GetJudgeDataByLogin(login string) (*models.Judge, error) {
	return m.GetJudgeDataByLoginFunc(login)
}
