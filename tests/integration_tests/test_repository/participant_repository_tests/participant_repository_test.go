package participant_repository_tests

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

// ParticipantRepositoryTestSuite описывает тестовый набор для IParticipantRepository
type ParticipantRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockIParticipantRepository
}

// SetupSuite выполняется один раз перед запуском тестов в этом наборе
func (suite *ParticipantRepositoryTestSuite) SetupSuite() {
	suite.mockRepo = &mocks.MockIParticipantRepository{}
}

// TestIParticipantRepository запускает тесты в наборе
func TestIParticipantRepository(t *testing.T) {
	suite.Run(t, new(ParticipantRepositoryTestSuite))
}

// TestParticipantRepositoryCreate_Success тестирует успешное создание участника
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryCreate_Success() {
	inputParticipant := builders.NewParticipantBuilder().
		WithID(uuid.New()).
		WithFIO("John Doe").
		WithCategory(1).
		WithGender(0).
		WithBirthday(time.Now()).
		WithCoach("Coach A").
		Build()

	// Настраиваем мок для создания участника
	suite.mockRepo.CreateFunc = func(participant *models.Participant) (*models.Participant, error) {
		return participant, nil // Возвращаем созданного участника
	}

	// Пытаемся создать участника
	createdParticipant, err := suite.mockRepo.Create(inputParticipant)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputParticipant.ID, createdParticipant.ID)
	require.Equal(suite.T(), inputParticipant.FIO, createdParticipant.FIO)
	require.Equal(suite.T(), inputParticipant.Category, createdParticipant.Category)
}

// TestParticipantRepositoryCreate_Failure тестирует ошибку при создании участника
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryCreate_Failure() {
	inputParticipant := builders.ParticipantMother.Default() // Используем значение по умолчанию

	// Настраиваем мок для создания участника с ошибкой
	suite.mockRepo.CreateFunc = func(participant *models.Participant) (*models.Participant, error) {
		return nil, errors.New("participant creation failed") // Возвращаем ошибку
	}

	// Пытаемся создать участника
	createdParticipant, err := suite.mockRepo.Create(inputParticipant)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), createdParticipant)                        // Убедимся, что созданный участник равен nil
	require.EqualError(suite.T(), err, "participant creation failed") // Проверяем, что ошибка соответствует ожиданиям
}

// TestParticipantRepositoryDelete_Success тестирует успешное удаление участника
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryDelete_Success() {
	participantID := uuid.New() // Генерируем новый UUID

	// Настраиваем мок для удаления участника
	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return nil // Успешно удаляем участника
	}

	// Пытаемся удалить участника
	err := suite.mockRepo.Delete(participantID)
	require.NoError(suite.T(), err)
}

// TestParticipantRepositoryDelete_Failure тестирует ошибку при удалении участника
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryDelete_Failure() {
	participantID := uuid.New() // Генерируем новый UUID

	// Настраиваем мок для удаления участника с ошибкой
	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return errors.New("participant not found") // Возвращаем ошибку
	}

	// Пытаемся удалить участника
	err := suite.mockRepo.Delete(participantID)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "participant not found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestParticipantRepositoryUpdate_Success тестирует успешное обновление участника
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryUpdate_Success() {
	inputParticipant := builders.NewParticipantBuilder().
		WithID(uuid.New()).
		WithFIO("Jane Doe").
		WithCategory(2).
		WithGender(1).
		WithBirthday(time.Now()).
		WithCoach("Coach B").
		Build()

	// Настраиваем мок для обновления участника
	suite.mockRepo.UpdateFunc = func(participant *models.Participant) (*models.Participant, error) {
		return participant, nil // Возвращаем обновленного участника
	}

	// Пытаемся обновить участника
	updatedParticipant, err := suite.mockRepo.Update(inputParticipant)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputParticipant.ID, updatedParticipant.ID)
	require.Equal(suite.T(), inputParticipant.FIO, updatedParticipant.FIO)
	require.Equal(suite.T(), inputParticipant.Category, updatedParticipant.Category)
}

// TestParticipantRepositoryUpdate_Failure тестирует ошибку при обновлении участника
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryUpdate_Failure() {
	inputParticipant := builders.ParticipantMother.Default() // Используем значение по умолчанию

	// Настраиваем мок для обновления участника с ошибкой
	suite.mockRepo.UpdateFunc = func(participant *models.Participant) (*models.Participant, error) {
		return nil, errors.New("participant update failed") // Возвращаем ошибку
	}

	// Пытаемся обновить участника
	updatedParticipant, err := suite.mockRepo.Update(inputParticipant)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), updatedParticipant)                      // Убедимся, что обновленный участник равен nil
	require.EqualError(suite.T(), err, "participant update failed") // Проверяем, что ошибка соответствует ожиданиям
}

// TestParticipantRepositoryGetByID_Success тестирует успешное получение участника по ID
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetByID_Success() {
	participantID := uuid.New() // Генерируем новый UUID
	expectedParticipant := builders.NewParticipantBuilder().
		WithID(participantID).
		WithFIO("Alice Smith").
		WithCategory(1).
		WithGender(0).
		WithBirthday(time.Now()).
		WithCoach("Coach C").
		Build()

	// Настраиваем мок для получения участника
	suite.mockRepo.GetParticipantDataByIDFunc = func(id uuid.UUID) (*models.Participant, error) {
		return expectedParticipant, nil // Возвращаем ожидаемого участника
	}

	// Пытаемся получить участника
	participant, err := suite.mockRepo.GetParticipantDataByID(participantID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), expectedParticipant.ID, participant.ID)
	require.Equal(suite.T(), expectedParticipant.FIO, participant.FIO)
}

// TestParticipantRepositoryGetByID_Failure тестирует ошибку при получении участника по ID
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetByID_Failure() {
	participantID := uuid.New() // Генерируем новый UUID

	// Настраиваем мок для получения участника с ошибкой
	suite.mockRepo.GetParticipantDataByIDFunc = func(id uuid.UUID) (*models.Participant, error) {
		return nil, errors.New("participant not found") // Возвращаем ошибку
	}

	// Пытаемся получить участника
	participant, err := suite.mockRepo.GetParticipantDataByID(participantID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), participant)                         // Убедимся, что участник равен nil
	require.EqualError(suite.T(), err, "participant not found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestParticipantRepositoryGetByCrewID_Success тестирует успешное получение участников по Crew ID
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetByCrewID_Success() {
	crewID := uuid.New() // Генерируем новый UUID
	expectedParticipants := []models.Participant{
		*builders.NewParticipantBuilder().WithID(uuid.New()).WithFIO("Participant 1").Build(),
		*builders.NewParticipantBuilder().WithID(uuid.New()).WithFIO("Participant 2").Build(),
	}

	// Настраиваем мок для получения участников
	suite.mockRepo.GetParticipantsDataByCrewIDFunc = func(id uuid.UUID) ([]models.Participant, error) {
		return expectedParticipants, nil // Возвращаем ожидаемых участников
	}

	// Пытаемся получить участников
	participants, err := suite.mockRepo.GetParticipantsDataByCrewID(crewID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(expectedParticipants), len(participants))
}

// TestParticipantRepositoryGetByCrewID_Failure тестирует ошибку при получении участников по Crew ID
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetByCrewID_Failure() {
	crewID := uuid.New() // Генерируем новый UUID

	// Настраиваем мок для получения участников с ошибкой
	suite.mockRepo.GetParticipantsDataByCrewIDFunc = func(id uuid.UUID) ([]models.Participant, error) {
		return nil, errors.New("no participants found") // Возвращаем ошибку
	}

	// Пытаемся получить участников
	participants, err := suite.mockRepo.GetParticipantsDataByCrewID(crewID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), participants)                        // Убедимся, что участники равны nil
	require.EqualError(suite.T(), err, "no participants found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestParticipantRepositoryGetByProtestID_Success тестирует успешное получение участников по Protest ID
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetByProtestID_Success() {
	protestID := uuid.New() // Генерируем новый UUID
	expectedParticipants := []models.Participant{
		*builders.NewParticipantBuilder().WithID(uuid.New()).WithFIO("Participant A").Build(),
		*builders.NewParticipantBuilder().WithID(uuid.New()).WithFIO("Participant B").Build(),
	}

	// Настраиваем мок для получения участников
	suite.mockRepo.GetParticipantsDataByProtestIDFunc = func(id uuid.UUID) ([]models.Participant, error) {
		return expectedParticipants, nil // Возвращаем ожидаемых участников
	}

	// Пытаемся получить участников
	participants, err := suite.mockRepo.GetParticipantsDataByProtestID(protestID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(expectedParticipants), len(participants))
}

// TestParticipantRepositoryGetByProtestID_Failure тестирует ошибку при получении участников по Protest ID
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetByProtestID_Failure() {
	protestID := uuid.New() // Генерируем новый UUID

	// Настраиваем мок для получения участников с ошибкой
	suite.mockRepo.GetParticipantsDataByProtestIDFunc = func(id uuid.UUID) ([]models.Participant, error) {
		return nil, errors.New("no participants found for protest") // Возвращаем ошибку
	}

	// Пытаемся получить участников
	participants, err := suite.mockRepo.GetParticipantsDataByProtestID(protestID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), participants)                                    // Убедимся, что участники равны nil
	require.EqualError(suite.T(), err, "no participants found for protest") // Проверяем, что ошибка соответствует ожиданиям
}

// TestParticipantRepositoryGetAll_Success тестирует успешное получение всех участников
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetAll_Success() {
	expectedParticipants := []models.Participant{
		*builders.NewParticipantBuilder().WithID(uuid.New()).WithFIO("Participant X").Build(),
		*builders.NewParticipantBuilder().WithID(uuid.New()).WithFIO("Participant Y").Build(),
	}

	// Настраиваем мок для получения всех участников
	suite.mockRepo.GetAllParticipantsFunc = func() ([]models.Participant, error) {
		return expectedParticipants, nil // Возвращаем ожидаемых участников
	}

	// Пытаемся получить всех участников
	participants, err := suite.mockRepo.GetAllParticipants()
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(expectedParticipants), len(participants))
}

// TestParticipantRepositoryGetAll_Failure тестирует ошибку при получении всех участников
func (suite *ParticipantRepositoryTestSuite) TestParticipantRepositoryGetAll_Failure() {
	// Настраиваем мок для получения всех участников с ошибкой
	suite.mockRepo.GetAllParticipantsFunc = func() ([]models.Participant, error) {
		return nil, errors.New("no participants found") // Возвращаем ошибку
	}

	// Пытаемся получить всех участников
	participants, err := suite.mockRepo.GetAllParticipants()
	require.Error(suite.T(), err)
	require.Nil(suite.T(), participants)                        // Убедимся, что участники равны nil
	require.EqualError(suite.T(), err, "no participants found") // Проверяем, что ошибка соответствует ожиданиям
}
