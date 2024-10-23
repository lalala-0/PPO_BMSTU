package test_repository

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/tests/unit_tests/builders"
	"PPO_BMSTU/tests/unit_tests/test_repository/mocks"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"runtime"
	"testing"
)

// CrewRepositoryTestSuite описывает тестовый набор для ICrewRepository
type CrewRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockICrewRepository
}

// SetupSuite выполняется один раз перед запуском тестов в этом наборе
func (suite *CrewRepositoryTestSuite) SetupSuite() {
	suite.mockRepo = &mocks.MockICrewRepository{}
}

// TestICrewRepository запускает тесты в наборе
func TestICrewRepository(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	suite.Run(t, new(CrewRepositoryTestSuite))
}

// TestCrewRepositoryCreate_Success тестирует успешное создание записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryCreate_Success() {
	// Создаем запись о команде с помощью Builder
	inputCrew := builders.NewCrewBuilder().
		WithSailNum(2).
		WithClass(1).
		Build()

	// Настраиваем мок для создания записи о команде
	suite.mockRepo.CreateFunc = func(crew *models.Crew) (*models.Crew, error) {
		return crew, nil // Возвращаем созданную запись
	}

	// Пытаемся создать запись о команде
	createdCrew, err := suite.mockRepo.Create(inputCrew)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputCrew.ID, createdCrew.ID)
	require.Equal(suite.T(), inputCrew.RatingID, createdCrew.RatingID)
	require.Equal(suite.T(), inputCrew.SailNum, createdCrew.SailNum)
	require.Equal(suite.T(), inputCrew.Class, createdCrew.Class)
}

// TestCrewRepositoryCreate_Failure тестирует ошибку при создании записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryCreate_Failure() {
	// Создаем запись о команде с помощью Object Mother
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 0, 0) // Некорректное значение для номера паруса

	// Настраиваем мок для создания записи о команде с ошибкой
	suite.mockRepo.CreateFunc = func(crew *models.Crew) (*models.Crew, error) {
		return nil, errors.New("invalid sail number") // Возвращаем ошибку
	}

	// Пытаемся создать запись о команде
	createdCrew, err := suite.mockRepo.Create(inputCrew)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), createdCrew)                       // Убедимся, что созданная запись равна nil
	require.EqualError(suite.T(), err, "invalid sail number") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryDelete_Success тестирует успешное удаление записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDelete_Success() {
	// Настраиваем мок для удаления записи о команде
	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return nil // Успешно удаляем запись
	}

	// Пытаемся удалить запись о команде
	err := suite.mockRepo.Delete(uuid.New())
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryDelete_Failure тестирует ошибку при удалении записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDelete_Failure() {
	// Настраиваем мок для удаления записи о команде с ошибкой
	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return errors.New("crew not found") // Возвращаем ошибку
	}

	// Пытаемся удалить запись о команде
	err := suite.mockRepo.Delete(uuid.New())
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "crew not found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryUpdate_Success тестирует успешное обновление записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryUpdate_Success() {
	// Создаем запись о команде с помощью Builder
	inputCrew := builders.NewCrewBuilder().
		WithSailNum(3).
		WithClass(2).
		Build()

	// Настраиваем мок для обновления записи о команде
	suite.mockRepo.UpdateFunc = func(crew *models.Crew) (*models.Crew, error) {
		return crew, nil // Возвращаем обновленную запись
	}

	// Пытаемся обновить запись о команде
	updatedCrew, err := suite.mockRepo.Update(inputCrew)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputCrew.ID, updatedCrew.ID)
	require.Equal(suite.T(), inputCrew.RatingID, updatedCrew.RatingID)
	require.Equal(suite.T(), inputCrew.SailNum, updatedCrew.SailNum)
	require.Equal(suite.T(), inputCrew.Class, updatedCrew.Class)
}

// TestCrewRepositoryUpdate_Failure тестирует ошибку при обновлении записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryUpdate_Failure() {
	// Создаем запись о команде с помощью Object Mother
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), -1, 0) // Некорректное значение для номера паруса

	// Настраиваем мок для обновления записи о команде с ошибкой
	suite.mockRepo.UpdateFunc = func(crew *models.Crew) (*models.Crew, error) {
		return nil, errors.New("invalid sail number") // Возвращаем ошибку
	}

	// Пытаемся обновить запись о команде
	updatedCrew, err := suite.mockRepo.Update(inputCrew)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), updatedCrew)                       // Убедимся, что обновленная запись равна nil
	require.EqualError(suite.T(), err, "invalid sail number") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryGetCrewDataByID_Success тестирует успешное получение записи о команде по ID
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataByID_Success() {
	// Создаем запись о команде с помощью Builder
	crew := builders.NewCrewBuilder().
		WithSailNum(4).
		WithClass(3).
		Build()

	// Настраиваем мок для получения записи о команде по ID
	suite.mockRepo.GetCrewDataByIDFunc = func(id uuid.UUID) (*models.Crew, error) {
		return crew, nil // Возвращаем созданную запись
	}

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.mockRepo.GetCrewDataByID(crew.ID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), crew.ID, receivedCrew.ID)
	require.Equal(suite.T(), crew.RatingID, receivedCrew.RatingID)
	require.Equal(suite.T(), crew.SailNum, receivedCrew.SailNum)
	require.Equal(suite.T(), crew.Class, receivedCrew.Class)
}

// TestCrewRepositoryGetCrewDataByID_Failure тестирует ошибку при получении записи о команде по ID
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataByID_Failure() {
	// Настраиваем мок для получения записи о команде по ID с ошибкой
	suite.mockRepo.GetCrewDataByIDFunc = func(id uuid.UUID) (*models.Crew, error) {
		return nil, errors.New("crew not found") // Возвращаем ошибку
	}

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.mockRepo.GetCrewDataByID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrew)                 // Убедимся, что полученная запись равна nil
	require.EqualError(suite.T(), err, "crew not found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryGetCrewsDataByRatingID_Success тестирует успешное получение списка команд по ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByRatingID_Success() {
	crews := []models.Crew{
		*builders.NewCrewBuilder().WithSailNum(5).WithClass(1).Build(),
		*builders.NewCrewBuilder().WithSailNum(6).WithClass(2).Build(),
	}

	// Настраиваем мок для получения списка команд по ID рейтинга
	suite.mockRepo.GetCrewsDataByRatingIDFunc = func(id uuid.UUID) ([]models.Crew, error) {
		return crews, nil // Возвращаем список команд
	}

	// Пытаемся получить список команд по ID рейтинга
	receivedCrews, err := suite.mockRepo.GetCrewsDataByRatingID(uuid.New())
	require.NoError(suite.T(), err)
	require.Len(suite.T(), receivedCrews, len(crews))
	require.Equal(suite.T(), crews[0].SailNum, receivedCrews[0].SailNum)
	require.Equal(suite.T(), crews[1].SailNum, receivedCrews[1].SailNum)
}

// TestCrewRepositoryGetCrewsDataByRatingID_Failure тестирует ошибку при получении списка команд по ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByRatingID_Failure() {
	// Настраиваем мок для получения списка команд по ID рейтинга с ошибкой
	suite.mockRepo.GetCrewsDataByRatingIDFunc = func(id uuid.UUID) ([]models.Crew, error) {
		return nil, errors.New("no crews found") // Возвращаем ошибку
	}

	// Пытаемся получить список команд по ID рейтинга
	receivedCrews, err := suite.mockRepo.GetCrewsDataByRatingID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrews)                // Убедимся, что полученный список команд равен nil
	require.EqualError(suite.T(), err, "no crews found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryGetCrewsDataByProtestID_Success тестирует успешное получение списка команд по ID протеста
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByProtestID_Success() {
	crews := []models.Crew{
		*builders.NewCrewBuilder().WithSailNum(1).WithClass(1).Build(),
		*builders.NewCrewBuilder().WithSailNum(2).WithClass(2).Build(),
	}

	// Настраиваем мок для получения списка команд по ID протеста
	suite.mockRepo.GetCrewsDataByProtestIDFunc = func(id uuid.UUID) ([]models.Crew, error) {
		return crews, nil // Возвращаем список команд
	}

	// Пытаемся получить список команд по ID протеста
	receivedCrews, err := suite.mockRepo.GetCrewsDataByProtestID(uuid.New())
	require.NoError(suite.T(), err)
	require.Len(suite.T(), receivedCrews, len(crews))
	require.Equal(suite.T(), crews[0].SailNum, receivedCrews[0].SailNum)
	require.Equal(suite.T(), crews[1].SailNum, receivedCrews[1].SailNum)
}

// TestCrewRepositoryGetCrewsDataByProtestID_Failure тестирует ошибку при получении списка команд по ID протеста
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByProtestID_Failure() {
	// Настраиваем мок для получения списка команд по ID протеста с ошибкой
	suite.mockRepo.GetCrewsDataByProtestIDFunc = func(id uuid.UUID) ([]models.Crew, error) {
		return nil, errors.New("no crews found for protest") // Возвращаем ошибку
	}

	// Пытаемся получить список команд по ID протеста
	receivedCrews, err := suite.mockRepo.GetCrewsDataByProtestID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrews)                            // Убедимся, что полученный список команд равен nil
	require.EqualError(suite.T(), err, "no crews found for protest") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Success тестирует успешное получение записи о команде по номеру паруса и ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Success() {
	crew := builders.NewCrewBuilder().
		WithSailNum(1).
		WithRatingID(uuid.New()).
		Build()

	// Настраиваем мок для получения записи о команде по номеру паруса и ID рейтинга
	suite.mockRepo.GetCrewDataBySailNumAndRatingIDFunc = func(sailNum int, ratingID uuid.UUID) (*models.Crew, error) {
		return crew, nil // Возвращаем созданную запись
	}

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.mockRepo.GetCrewDataBySailNumAndRatingID(1, crew.RatingID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), crew.ID, receivedCrew.ID)
	require.Equal(suite.T(), crew.SailNum, receivedCrew.SailNum)
	require.Equal(suite.T(), crew.RatingID, receivedCrew.RatingID)
}

// TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Failure тестирует ошибку при получении записи о команде по номеру паруса и ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Failure() {
	// Настраиваем мок для получения записи о команде по номеру паруса и ID рейтинга с ошибкой
	suite.mockRepo.GetCrewDataBySailNumAndRatingIDFunc = func(sailNum int, ratingID uuid.UUID) (*models.Crew, error) {
		return nil, errors.New("crew not found") // Возвращаем ошибку
	}

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.mockRepo.GetCrewDataBySailNumAndRatingID(1, uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrew)                 // Убедимся, что полученная запись равна nil
	require.EqualError(suite.T(), err, "crew not found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryAttachParticipantToCrew_Success тестирует успешное добавление участника в команду
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryAttachParticipantToCrew_Success() {
	// Настраиваем мок для добавления участника в команду
	suite.mockRepo.AttachParticipantToCrewFunc = func(participantID, crewID uuid.UUID, helmsman int) error {
		return nil // Успешно добавляем участника
	}

	// Пытаемся добавить участника в команду
	err := suite.mockRepo.AttachParticipantToCrew(uuid.New(), uuid.New(), 1)
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryAttachParticipantToCrew_Failure тестирует ошибку при добавлении участника в команду
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryAttachParticipantToCrew_Failure() {
	// Настраиваем мок для добавления участника в команду с ошибкой
	suite.mockRepo.AttachParticipantToCrewFunc = func(participantID, crewID uuid.UUID, helmsman int) error {
		return errors.New("failed to attach participant") // Возвращаем ошибку
	}

	// Пытаемся добавить участника в команду
	err := suite.mockRepo.AttachParticipantToCrew(uuid.New(), uuid.New(), 1)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "failed to attach participant") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryDetachParticipantFromCrew_Success тестирует успешное удаление участника из команды
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDetachParticipantFromCrew_Success() {
	// Настраиваем мок для удаления участника из команды
	suite.mockRepo.DetachParticipantFromCrewFunc = func(participantID, crewID uuid.UUID) error {
		return nil // Успешно удаляем участника
	}

	// Пытаемся удалить участника из команды
	err := suite.mockRepo.DetachParticipantFromCrew(uuid.New(), uuid.New())
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryDetachParticipantFromCrew_Failure тестирует ошибку при удалении участника из команды
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDetachParticipantFromCrew_Failure() {
	// Настраиваем мок для удаления участника из команды с ошибкой
	suite.mockRepo.DetachParticipantFromCrewFunc = func(participantID, crewID uuid.UUID) error {
		return errors.New("failed to detach participant") // Возвращаем ошибку
	}

	// Пытаемся удалить участника из команды
	err := suite.mockRepo.DetachParticipantFromCrew(uuid.New(), uuid.New())
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "failed to detach participant") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryReplaceParticipantStatusInCrew_Success тестирует успешную замену статуса участника в команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryReplaceParticipantStatusInCrew_Success() {
	// Настраиваем мок для замены статуса участника в команде
	suite.mockRepo.ReplaceParticipantStatusInCrewFunc = func(participantID, crewID uuid.UUID, helmsman int, active int) error {
		return nil // Успешно заменяем статус
	}

	// Пытаемся заменить статус участника в команде
	err := suite.mockRepo.ReplaceParticipantStatusInCrew(uuid.New(), uuid.New(), 1, 1)
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryReplaceParticipantStatusInCrew_Failure тестирует ошибку при замене статуса участника в команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryReplaceParticipantStatusInCrew_Failure() {
	// Настраиваем мок для замены статуса участника в команде с ошибкой
	suite.mockRepo.ReplaceParticipantStatusInCrewFunc = func(participantID, crewID uuid.UUID, helmsman int, active int) error {
		return errors.New("failed to replace participant status") // Возвращаем ошибку
	}

	// Пытаемся заменить статус участника в команде
	err := suite.mockRepo.ReplaceParticipantStatusInCrew(uuid.New(), uuid.New(), 1, 1)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "failed to replace participant status") // Проверяем, что ошибка соответствует ожиданиям
}
