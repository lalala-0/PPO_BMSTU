package rating_repository_tests

import (
	"PPO_BMSTU/internal/models"
	builders2 "PPO_BMSTU/tests/unit_tests/builders"
	"PPO_BMSTU/tests/unit_tests/test_repository/mocks"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

// RatingRepositoryTestSuite - структура для тестового сьюта репозитория рейтингов
type RatingRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockRatingRepository
}

// SetupTest - метод, вызываемый перед каждым тестом
func (suite *RatingRepositoryTestSuite) SetupTest() {
	suite.mockRepo = &mocks.MockRatingRepository{} // Создание нового мока перед каждым тестом
}

// Пример теста для успешного создания рейтинга
func (suite *RatingRepositoryTestSuite) TestCreateRating_Success() {
	// Используем Object Mother для создания рейтинга
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)

	suite.mockRepo.CreateFunc = func(r *models.Rating) (*models.Rating, error) {
		return r, nil
	}

	createdRating, err := suite.mockRepo.Create(rating)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), rating.Name, createdRating.Name)
	assert.Equal(suite.T(), rating.Class, createdRating.Class)
	assert.Equal(suite.T(), rating.BlowoutCnt, createdRating.BlowoutCnt)
}

// Пример теста для неудачного создания рейтинга
func (suite *RatingRepositoryTestSuite) TestCreateRating_Failure() {
	// Используем Object Mother для создания рейтинга
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)

	suite.mockRepo.CreateFunc = func(r *models.Rating) (*models.Rating, error) {
		return nil, errors.New("creation failed")
	}

	createdRating, err := suite.mockRepo.Create(rating)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdRating)
}

// Пример позитивного теста для получения рейтинга по ID
func (suite *RatingRepositoryTestSuite) TestGetRatingByID_Success() {
	// Используем Object Mother для создания рейтинга
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)

	suite.mockRepo.GetRatingDataByIDFunc = func(id uuid.UUID) (*models.Rating, error) {
		return rating, nil
	}

	receivedRating, err := suite.mockRepo.GetRatingDataByID(rating.ID)

	require.NoError(suite.T(), err)
	require.Equal(suite.T(), rating.ID, receivedRating.ID)
}

// Пример негативного теста для получения рейтинга по ID
func (suite *RatingRepositoryTestSuite) TestGetRatingByID_Failure() {
	// Используем Object Mother для создания рейтинга
	ratingID := uuid.New()

	suite.mockRepo.GetRatingDataByIDFunc = func(id uuid.UUID) (*models.Rating, error) {
		return nil, errors.New("rating not found")
	}

	receivedRating, err := suite.mockRepo.GetRatingDataByID(ratingID)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), receivedRating)
}

// Пример позитивного теста для удаления рейтинга
func (suite *RatingRepositoryTestSuite) TestDeleteRating_Success() {
	ratingID := uuid.New()

	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return nil
	}

	// Удаление рейтинга
	err := suite.mockRepo.Delete(ratingID)
	require.NoError(suite.T(), err)
}

// Пример негативного теста для удаления рейтинга
func (suite *RatingRepositoryTestSuite) TestDeleteRating_Failure() {
	ratingID := uuid.New()

	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return errors.New("deletion failed")
	}

	// Попытка удаления, которая должна завершиться ошибкой
	err := suite.mockRepo.Delete(ratingID)
	assert.Error(suite.T(), err)
}

// Пример позитивного теста для обновления рейтинга
func (suite *RatingRepositoryTestSuite) TestUpdateRating_Success() {

	// Используем Builder для создания начального рейтинга
	initialRating := builders2.NewRatingBuilder().
		WithName("InitialName").
		WithClass(models.Laser).
		WithBlowoutCnt(1).
		Build()

	// Моковая функция для успешного обновления
	suite.mockRepo.UpdateFunc = func(r *models.Rating) (*models.Rating, error) {
		// Имитация изменения значений полей
		r.Name = "UpdatedName"
		r.Class = models.LaserRadial
		r.BlowoutCnt = 2
		return r, nil
	}

	// Обновляем рейтинг
	updatedRating, err := suite.mockRepo.Update(initialRating)

	// Проверяем, что обновление прошло успешно
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "UpdatedName", updatedRating.Name)
	assert.Equal(suite.T(), models.LaserRadial, updatedRating.Class)
	assert.Equal(suite.T(), 2, updatedRating.BlowoutCnt)
}

// Пример негативного теста для обновления рейтинга
func (suite *RatingRepositoryTestSuite) TestUpdateRating_Failure() {

	// Используем Object Mother для создания рейтинга
	rating := builders2.RatingMother.CustomRating(uuid.New(), "Name", models.Laser, 1)

	// Моковая функция для имитации ошибки обновления
	suite.mockRepo.UpdateFunc = func(r *models.Rating) (*models.Rating, error) {
		return nil, errors.New("update failed")
	}

	// Попытка обновления, которая должна завершиться ошибкой
	updatedRating, err := suite.mockRepo.Update(rating)

	// Проверяем, что ошибка произошла
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedRating)
}

// Пример позитивного теста для AttachJudgeToRating
func (suite *RatingRepositoryTestSuite) TestAttachJudgeToRating_Success() {

	// Используем Object Mother для создания рейтинга и судьи
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	judge := builders2.JudgeMother.Default() // Используем Object Mother для судьи

	// Моковая функция для успешного прикрепления судьи к рейтингу
	suite.mockRepo.AttachJudgeToRatingFunc = func(ratingID uuid.UUID, judgeID uuid.UUID) error {
		return nil
	}

	// Пробуем прикрепить судью к рейтингу
	err := suite.mockRepo.AttachJudgeToRating(rating.ID, judge.ID)

	// Проверяем, что прикрепление прошло успешно
	assert.NoError(suite.T(), err)
}

// Пример негативного теста для AttachJudgeToRating
func (suite *RatingRepositoryTestSuite) TestAttachJudgeToRating_Failure() {

	// Используем Object Mother для создания рейтинга и судьи
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	judge := builders2.JudgeMother.Default()

	// Моковая функция для имитации ошибки прикрепления судьи
	suite.mockRepo.AttachJudgeToRatingFunc = func(ratingID uuid.UUID, judgeID uuid.UUID) error {
		return errors.New("attach judge failed")
	}

	// Попытка прикрепления судьи к рейтингу, которая должна завершиться ошибкой
	err := suite.mockRepo.AttachJudgeToRating(rating.ID, judge.ID)

	// Проверяем, что произошла ошибка
	assert.Error(suite.T(), err)
}

// Пример позитивного теста для DetachJudgeFromRating
func (suite *RatingRepositoryTestSuite) TestDetachJudgeFromRating_Success() {

	// Используем Object Mother для создания рейтинга и судьи
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	judge := builders2.JudgeMother.Default()

	// Моковая функция для успешного отсоединения судьи от рейтинга
	suite.mockRepo.DetachJudgeFromRatingFunc = func(ratingID uuid.UUID, judgeID uuid.UUID) error {
		return nil
	}

	// Пробуем отсоединить судью от рейтинга
	err := suite.mockRepo.DetachJudgeFromRating(rating.ID, judge.ID)

	// Проверяем, что отсоединение прошло успешно
	assert.NoError(suite.T(), err)
}

// Пример негативного теста для DetachJudgeFromRating
func (suite *RatingRepositoryTestSuite) TestDetachJudgeFromRating_Failure() {

	// Используем Object Mother для создания рейтинга и судьи
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	judge := builders2.JudgeMother.Default()

	// Моковая функция для имитации ошибки отсоединения судьи
	suite.mockRepo.DetachJudgeFromRatingFunc = func(ratingID uuid.UUID, judgeID uuid.UUID) error {
		return errors.New("detach judge failed")
	}

	// Попытка отсоединения судьи от рейтинга, которая должна завершиться ошибкой
	err := suite.mockRepo.DetachJudgeFromRating(rating.ID, judge.ID)

	// Проверяем, что произошла ошибка
	assert.Error(suite.T(), err)
}

// Запуск тестов
func TestRatingRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RatingRepositoryTestSuite))
}
