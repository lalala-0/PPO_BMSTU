package mocks

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// MockRatingRepository представляет мок-репозиторий для Rating
type MockRatingRepository struct {
	CreateFunc                func(*models.Rating) (*models.Rating, error)
	DeleteFunc                func(uuid.UUID) error
	UpdateFunc                func(*models.Rating) (*models.Rating, error)
	GetRatingDataByIDFunc     func(uuid.UUID) (*models.Rating, error)
	AttachJudgeToRatingFunc   func(uuid.UUID, uuid.UUID) error
	DetachJudgeFromRatingFunc func(uuid.UUID, uuid.UUID) error
	GetAllRatingsFunc         func() ([]models.Rating, error)
	GetRatingTableFunc        func(uuid.UUID) ([]models.RatingTableLine, error)
}

// Create вызывает функцию мока для создания рейтинга
func (m *MockRatingRepository) Create(rating *models.Rating) (*models.Rating, error) {
	return m.CreateFunc(rating)
}

// Delete вызывает функцию мока для удаления рейтинга по ID
func (m *MockRatingRepository) Delete(id uuid.UUID) error {
	return m.DeleteFunc(id)
}

// Update вызывает функцию мока для обновления данных рейтинга
func (m *MockRatingRepository) Update(rating *models.Rating) (*models.Rating, error) {
	return m.UpdateFunc(rating)
}

// GetRatingDataByID вызывает функцию мока для получения данных рейтинга по ID
func (m *MockRatingRepository) GetRatingDataByID(id uuid.UUID) (*models.Rating, error) {
	return m.GetRatingDataByIDFunc(id)
}

// AttachJudgeToRating вызывает функцию мока для привязки судьи к рейтингу
func (m *MockRatingRepository) AttachJudgeToRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	return m.AttachJudgeToRatingFunc(ratingID, judgeID)
}

// DetachJudgeFromRating вызывает функцию мока для отвязки судьи от рейтинга
func (m *MockRatingRepository) DetachJudgeFromRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	return m.DetachJudgeFromRatingFunc(ratingID, judgeID)
}

// GetAllRatings вызывает функцию мока для получения всех рейтингов
func (m *MockRatingRepository) GetAllRatings() ([]models.Rating, error) {
	return m.GetAllRatingsFunc()
}

// GetRatingTable вызывает функцию мока для получения таблицы рейтинга по ID
func (m *MockRatingRepository) GetRatingTable(id uuid.UUID) ([]models.RatingTableLine, error) {
	return m.GetRatingTableFunc(id)
}
