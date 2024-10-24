package protest_repository_tests

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/tests/unit_tests/builders"
	"PPO_BMSTU/tests/unit_tests/test_repository/mocks"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// ProtestRepositoryTestSuite описывает тестовый набор для IProtestRepository
type ProtestRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockIProtestRepository
}

// SetupSuite выполняется один раз перед запуском тестов в этом наборе
func (suite *ProtestRepositoryTestSuite) SetupSuite() {
	suite.mockRepo = &mocks.MockIProtestRepository{}
}

// TestIProtestRepository запускает тесты в наборе
func TestIProtestRepository(t *testing.T) {
	suite.Run(t, new(ProtestRepositoryTestSuite))
}

// TestProtestRepositoryCreate_Success тестирует успешное создание записи о протесте
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryCreate_Success() {
	// Создаем запись о протесте с помощью Builder
	inputProtest := builders.NewProtestBuilder().
		WithRuleNum(23).
		WithReviewDate(time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC)).
		WithStatus(1).
		WithComment("").
		Build()

	// Настраиваем мок для создания записи о протесте
	suite.mockRepo.CreateFunc = func(protest *models.Protest) (*models.Protest, error) {
		return protest, nil // Возвращаем созданную запись
	}

	// Пытаемся создать запись о протесте
	createdProtest, err := suite.mockRepo.Create(inputProtest)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputProtest.RaceID, createdProtest.RaceID)
	require.Equal(suite.T(), inputProtest.JudgeID, createdProtest.JudgeID)
	require.Equal(suite.T(), inputProtest.RatingID, createdProtest.RatingID)
	require.Equal(suite.T(), inputProtest.RuleNum, createdProtest.RuleNum)
	require.Equal(suite.T(), inputProtest.ReviewDate, createdProtest.ReviewDate)
	require.Equal(suite.T(), inputProtest.Status, createdProtest.Status)
	require.Equal(suite.T(), inputProtest.Comment, createdProtest.Comment)
}

// TestProtestRepositoryCreate_Failure тестирует ошибку при создании записи о протесте
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryCreate_Failure() {
	// Создаем запись о протесте с помощью Object Mother
	inputProtest := builders.ProtestMother.CustomProtest(
		uuid.New(), // ID
		uuid.Nil,   // RaceID (недопустимое значение)
		uuid.New(), // JudgeID
		uuid.New(), // RatingID
		23,         // RuleNum
		1,          // Status
		time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC), // ReviewDate
		"", // Comment
	)
	// Настраиваем мок для создания записи о протесте с ошибкой
	suite.mockRepo.CreateFunc = func(protest *models.Protest) (*models.Protest, error) {
		return nil, errors.New("invalid rule number") // Возвращаем ошибку
	}

	// Пытаемся создать запись о протесте
	createdProtest, err := suite.mockRepo.Create(inputProtest)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), createdProtest)                    // Убедимся, что созданная запись равна nil
	require.EqualError(suite.T(), err, "invalid rule number") // Проверяем, что ошибка соответствует ожиданиям
}

// TestProtestRepositoryGetByID_Success тестирует успешное получение протеста по ID
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryGetByID_Success() {
	// Создаем протест с помощью Builder
	inputProtest := builders.NewProtestBuilder().
		WithJudgeID(uuid.New()).
		WithRaceID(uuid.New()).
		WithRatingID(uuid.New()).
		WithRuleNum(1).
		WithStatus(0).
		WithReviewDate(time.Now()).
		WithComment("Test comment").
		Build()

	// Настраиваем мок для получения протеста по ID
	suite.mockRepo.GetProtestDataByIDFunc = func(id uuid.UUID) (*models.Protest, error) {
		return inputProtest, nil // Возвращаем созданный протест
	}

	// Пытаемся получить протест по ID
	receivedProtest, err := suite.mockRepo.GetProtestDataByID(inputProtest.ID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputProtest.ID, receivedProtest.ID)
	require.Equal(suite.T(), inputProtest.RuleNum, receivedProtest.RuleNum)
	require.Equal(suite.T(), inputProtest.ReviewDate, receivedProtest.ReviewDate)
	require.Equal(suite.T(), inputProtest.Status, receivedProtest.Status)
	require.Equal(suite.T(), inputProtest.Comment, receivedProtest.Comment)
}

// TestProtestRepositoryGetByID_Failure тестирует ошибку при получении протеста по несуществующему ID
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryGetByID_Failure() {
	// Генерируем несуществующий ID
	nonExistentID := uuid.New()

	// Настраиваем мок для получения протеста по несуществующему ID
	suite.mockRepo.GetProtestDataByIDFunc = func(id uuid.UUID) (*models.Protest, error) {
		return nil, errors.New("protest not found") // Возвращаем ошибку
	}

	// Пытаемся получить протест по несуществующему ID
	receivedProtest, err := suite.mockRepo.GetProtestDataByID(nonExistentID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedProtest)                 // Убедимся, что полученный протест равен nil
	require.EqualError(suite.T(), err, "protest not found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestProtestRepositoryGetProtestsDataByRaceID_Success тестирует успешное получение протестов по ID гонки
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryGetProtestsDataByRaceID_Success() {
	// Создаем массив протестов с помощью Builder
	inputProtests := []models.Protest{
		*builders.NewProtestBuilder().WithRaceID(uuid.New()).Build(),
		*builders.NewProtestBuilder().WithRaceID(uuid.New()).Build(),
	}

	// Настраиваем мок для получения протестов по ID гонки
	suite.mockRepo.GetProtestsDataByRaceIDFunc = func(raceID uuid.UUID) ([]models.Protest, error) {
		return inputProtests, nil // Возвращаем созданные протесты
	}

	// Пытаемся получить протесты по ID гонки
	receivedProtests, err := suite.mockRepo.GetProtestsDataByRaceID(inputProtests[0].RaceID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(inputProtests), len(receivedProtests))
}

// TestProtestRepositoryGetProtestsDataByRaceID_Failure тестирует ошибку при получении протестов по несуществующему ID гонки
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryGetProtestsDataByRaceID_Failure() {
	// Генерируем несуществующий ID гонки
	nonExistentRaceID := uuid.New()

	// Настраиваем мок для получения протестов по несуществующему ID гонки
	suite.mockRepo.GetProtestsDataByRaceIDFunc = func(raceID uuid.UUID) ([]models.Protest, error) {
		return nil, errors.New("no protests found for this race") // Возвращаем ошибку
	}

	// Пытаемся получить протесты по несуществующему ID гонки
	receivedProtests, err := suite.mockRepo.GetProtestsDataByRaceID(nonExistentRaceID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedProtests)                              // Убедимся, что полученные протесты равны nil
	require.EqualError(suite.T(), err, "no protests found for this race") // Проверяем, что ошибка соответствует ожиданиям
}

// TestProtestRepositoryAttachCrewToProtest_Success тестирует успешное присоединение экипажа к протесту
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryAttachCrewToProtest_Success() {
	crewID := uuid.New()
	protestID := uuid.New()
	crewStatus := 1

	// Настраиваем мок для присоединения экипажа к протесту
	suite.mockRepo.AttachCrewToProtestFunc = func(crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error {
		return nil // Успешное выполнение
	}

	// Пытаемся присоединить экипаж к протесту
	err := suite.mockRepo.AttachCrewToProtest(crewID, protestID, crewStatus)
	require.NoError(suite.T(), err)
}

// TestProtestRepositoryAttachCrewToProtest_Failure тестирует ошибку при присоединении экипажа к протесту
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryAttachCrewToProtest_Failure() {
	crewID := uuid.New()
	protestID := uuid.New()
	crewStatus := 1

	// Настраиваем мок для присоединения экипажа к протесту с ошибкой
	suite.mockRepo.AttachCrewToProtestFunc = func(crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error {
		return errors.New("failed to attach crew") // Возвращаем ошибку
	}

	// Пытаемся присоединить экипаж к протесту
	err := suite.mockRepo.AttachCrewToProtest(crewID, protestID, crewStatus)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "failed to attach crew") // Проверяем, что ошибка соответствует ожиданиям
}

// TestProtestRepositoryDetachCrewFromProtest_Success тестирует успешное отсоединение экипажа от протеста
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryDetachCrewFromProtest_Success() {
	crewID := uuid.New()
	protestID := uuid.New()

	// Настраиваем мок для отсоединения экипажа от протеста
	suite.mockRepo.DetachCrewFromProtestFunc = func(crewID uuid.UUID, protestID uuid.UUID) error {
		return nil // Успешное выполнение
	}

	// Пытаемся отсоединить экипаж от протеста
	err := suite.mockRepo.DetachCrewFromProtest(crewID, protestID)
	require.NoError(suite.T(), err)
}

// TestProtestRepositoryDetachCrewFromProtest_Failure тестирует ошибку при отсоединении экипажа от протеста
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryDetachCrewFromProtest_Failure() {
	crewID := uuid.New()
	protestID := uuid.New()

	// Настраиваем мок для отсоединения экипажа от протеста с ошибкой
	suite.mockRepo.DetachCrewFromProtestFunc = func(crewID uuid.UUID, protestID uuid.UUID) error {
		return errors.New("failed to detach crew") // Возвращаем ошибку
	}

	// Пытаемся отсоединить экипаж от протеста
	err := suite.mockRepo.DetachCrewFromProtest(crewID, protestID)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "failed to detach crew") // Проверяем, что ошибка соответствует ожиданиям
}

// TestProtestRepositoryGetProtestParticipantsIDByID_Success тестирует успешное получение участников протеста по ID
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryGetProtestParticipantsIDByID_Success() {
	protestID := uuid.New()
	expectedParticipants := map[uuid.UUID]int{
		uuid.New(): 1,
		uuid.New(): 2,
	}

	// Настраиваем мок для получения участников протеста
	suite.mockRepo.GetProtestParticipantsIDByIDFunc = func(id uuid.UUID) (map[uuid.UUID]int, error) {
		return expectedParticipants, nil // Возвращаем участников
	}

	// Пытаемся получить участников протеста
	receivedParticipants, err := suite.mockRepo.GetProtestParticipantsIDByID(protestID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), expectedParticipants, receivedParticipants)
}

// TestProtestRepositoryGetProtestParticipantsIDByID_Failure тестирует ошибку при получении участников протеста по несуществующему ID
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryGetProtestParticipantsIDByID_Failure() {
	nonExistentID := uuid.New()

	// Настраиваем мок для получения участников протеста по несуществующему ID
	suite.mockRepo.GetProtestParticipantsIDByIDFunc = func(id uuid.UUID) (map[uuid.UUID]int, error) {
		return nil, errors.New("no participants found") // Возвращаем ошибку
	}

	// Пытаемся получить участников протеста по несуществующему ID
	receivedParticipants, err := suite.mockRepo.GetProtestParticipantsIDByID(nonExistentID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedParticipants)                // Убедимся, что полученные участники равны nil
	require.EqualError(suite.T(), err, "no participants found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestProtestRepositoryDelete_Success тестирует успешное удаление протеста
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryDelete_Success() {
	protestID := uuid.New()

	// Настраиваем мок для удаления протеста
	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return nil // Успешное выполнение
	}

	// Пытаемся удалить протест
	err := suite.mockRepo.Delete(protestID)
	require.NoError(suite.T(), err)
}

// TestProtestRepositoryDelete_Failure тестирует ошибку при удалении протеста
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryDelete_Failure() {
	protestID := uuid.New()

	// Настраиваем мок для удаления протеста с ошибкой
	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return errors.New("failed to delete protest") // Возвращаем ошибку
	}

	// Пытаемся удалить протест
	err := suite.mockRepo.Delete(protestID)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "failed to delete protest") // Проверяем, что ошибка соответствует ожиданиям
}

// TestProtestRepositoryUpdate_Success тестирует успешное обновление протеста
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryUpdate_Success() {
	protest := builders.NewProtestBuilder().Build()

	// Настраиваем мок для обновления протеста
	suite.mockRepo.UpdateFunc = func(protest *models.Protest) (*models.Protest, error) {
		return protest, nil // Возвращаем обновленный протест
	}

	// Пытаемся обновить протест
	updatedProtest, err := suite.mockRepo.Update(protest)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), protest.ID, updatedProtest.ID)
}

// TestProtestRepositoryUpdate_Failure тестирует ошибку при обновлении протеста
func (suite *ProtestRepositoryTestSuite) TestProtestRepositoryUpdate_Failure() {
	protest := builders.NewProtestBuilder().Build()

	// Настраиваем мок для обновления протеста с ошибкой
	suite.mockRepo.UpdateFunc = func(protest *models.Protest) (*models.Protest, error) {
		return nil, errors.New("failed to update protest") // Возвращаем ошибку
	}

	// Пытаемся обновить протест
	updatedProtest, err := suite.mockRepo.Update(protest)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), updatedProtest)                         // Убедимся, что обновленный протест равен nil
	require.EqualError(suite.T(), err, "failed to update protest") // Проверяем, что ошибка соответствует ожиданиям
}
