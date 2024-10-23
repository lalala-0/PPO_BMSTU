package test_repository

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/tests/unit_tests/builders"
	"PPO_BMSTU/tests/unit_tests/test_repository/mocks"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

// CrewResInRaceRepositoryTestSuite описывает тестовый набор для CrewResInRaceRepository
type CrewResInRaceRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockCrewResInRaceRepository
}

// SetupSuite выполняется один раз перед запуском тестов в этом наборе
func (suite *CrewResInRaceRepositoryTestSuite) SetupSuite() {
	suite.mockRepo = &mocks.MockCrewResInRaceRepository{}
}

// TestCrewResInRaceRepositoryCreate_Success тестирует успешное создание записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryCreate_Success() {
	// Создаем запись о команде в гонке с помощью Builder
	inputCrewResInRace := builders.NewCrewResInRaceBuilder().
		WithPoints(12).
		WithSpecCircumstance(0).
		Build()

	// Настраиваем мок для создания записи о команде в гонке
	suite.mockRepo.CreateFunc = func(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
		return crewResInRace, nil // Возвращаем созданную запись
	}

	// Пытаемся создать запись о команде в гонке
	createdCrewResInRace, err := suite.mockRepo.Create(inputCrewResInRace)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputCrewResInRace.CrewID, createdCrewResInRace.CrewID)
	require.Equal(suite.T(), inputCrewResInRace.RaceID, createdCrewResInRace.RaceID)
	require.Equal(suite.T(), inputCrewResInRace.Points, createdCrewResInRace.Points)
	require.Equal(suite.T(), inputCrewResInRace.SpecCircumstance, createdCrewResInRace.SpecCircumstance)
}

// TestCrewResInRaceRepositoryCreate_Failure тестирует ошибку при создании записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryCreate_Failure() {
	// Создаем запись о команде в гонке с помощью Object Mother
	inputCrewResInRace := builders.CrewResInRaceMother.CustomCrew(uuid.New(), uuid.New(), 1, 0) // Некорректное значение для очков

	// Настраиваем мок для создания записи о команде в гонке с ошибкой
	suite.mockRepo.CreateFunc = func(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
		return nil, errors.New("invalid points value") // Возвращаем ошибку
	}

	// Пытаемся создать запись о команде в гонке
	createdCrewResInRace, err := suite.mockRepo.Create(inputCrewResInRace)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), createdCrewResInRace)               // Убедимся, что созданная запись равна nil
	require.EqualError(suite.T(), err, "invalid points value") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewResInRaceRepositoryGetCrewResByRaceIDAndCrewID_Success тестирует успешное получение записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryGetCrewResByRaceIDAndCrewID_Success() {
	// Создаем запись о команде в гонке с помощью Builder
	inputCrewResInRace := builders.NewCrewResInRaceBuilder().
		WithPoints(12).
		WithSpecCircumstance(0).
		Build()

	// Настраиваем мок для получения записи о команде в гонке
	suite.mockRepo.GetCrewResByRaceIDAndCrewIDFunc = func(raceID, crewID uuid.UUID) (*models.CrewResInRace, error) {
		return inputCrewResInRace, nil // Возвращаем созданную запись
	}

	// Пытаемся получить запись о команде в гонке
	receivedCrewResInRace, err := suite.mockRepo.GetCrewResByRaceIDAndCrewID(inputCrewResInRace.RaceID, inputCrewResInRace.CrewID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputCrewResInRace.CrewID, receivedCrewResInRace.CrewID)
	require.Equal(suite.T(), inputCrewResInRace.RaceID, receivedCrewResInRace.RaceID)
	require.Equal(suite.T(), inputCrewResInRace.Points, receivedCrewResInRace.Points)
	require.Equal(suite.T(), inputCrewResInRace.SpecCircumstance, receivedCrewResInRace.SpecCircumstance)
}

// TestCrewResInRaceRepositoryGetCrewResByRaceIDAndCrewID_Failure тестирует ошибку при получении записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryGetCrewResByRaceIDAndCrewID_Failure() {
	// Настраиваем мок для получения записи о команде в гонке с ошибкой
	suite.mockRepo.GetCrewResByRaceIDAndCrewIDFunc = func(raceID, crewID uuid.UUID) (*models.CrewResInRace, error) {
		return nil, errors.New("record not found") // Возвращаем ошибку
	}

	// Пытаемся получить запись о команде в гонке
	receivedCrewResInRace, err := suite.mockRepo.GetCrewResByRaceIDAndCrewID(uuid.New(), uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrewResInRace)          // Убедимся, что полученная запись равна nil
	require.EqualError(suite.T(), err, "record not found") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewResInRaceRepositoryDelete_Success тестирует успешное удаление записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryDelete_Success() {
	// Создаем запись о команде в гонке с помощью Builder
	inputCrewResInRace := builders.NewCrewResInRaceBuilder().
		WithPoints(12).
		WithSpecCircumstance(0).
		Build()

	// Настраиваем мок для удаления записи о команде в гонке
	suite.mockRepo.DeleteFunc = func(raceID, crewID uuid.UUID) error {
		return nil // Возвращаем nil для успешного удаления
	}

	// Пытаемся удалить запись о команде в гонке
	err := suite.mockRepo.Delete(inputCrewResInRace.RaceID, inputCrewResInRace.CrewID)
	require.NoError(suite.T(), err)
}

// TestCrewResInRaceRepositoryDelete_Failure тестирует ошибку при удалении записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryDelete_Failure() {
	// Настраиваем мок для удаления записи о команде в гонке с ошибкой
	suite.mockRepo.DeleteFunc = func(raceID, crewID uuid.UUID) error {
		return errors.New("failed to delete record") // Возвращаем ошибку
	}

	// Пытаемся удалить запись о команде в гонке
	err := suite.mockRepo.Delete(uuid.New(), uuid.New())
	require.Error(suite.T(), err)                                 // Проверяем, что возникает ошибка
	require.EqualError(suite.T(), err, "failed to delete record") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewResInRaceRepositoryUpdate_Success тестирует успешное обновление записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryUpdate_Success() {
	// Создаем запись о команде в гонке с помощью Builder
	inputCrewResInRace := builders.NewCrewResInRaceBuilder().
		WithPoints(135).
		WithSpecCircumstance(1).
		Build()

	// Настраиваем мок для обновления записи о команде в гонке
	suite.mockRepo.UpdateFunc = func(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
		// Возвращаем обновлённую запись
		crewResInRace.Points = 135
		crewResInRace.SpecCircumstance = 1
		return crewResInRace, nil
	}

	// Пытаемся обновить запись о команде в гонке
	updatedCrewResInRace, err := suite.mockRepo.Update(inputCrewResInRace) // Используем созданную запись
	require.NoError(suite.T(), err)

	// Проверяем, что обновлённая запись соответствует ожиданиям
	require.Equal(suite.T(), inputCrewResInRace.CrewID, updatedCrewResInRace.CrewID)
	require.Equal(suite.T(), inputCrewResInRace.RaceID, updatedCrewResInRace.RaceID)
	require.Equal(suite.T(), inputCrewResInRace.Points, updatedCrewResInRace.Points)
	require.Equal(suite.T(), inputCrewResInRace.SpecCircumstance, updatedCrewResInRace.SpecCircumstance)
}

// TestCrewResInRaceRepositoryUpdate_Failure тестирует ошибку при обновлении записи о команде в гонке
func (suite *CrewResInRaceRepositoryTestSuite) TestCrewResInRaceRepositoryUpdate_Failure() {
	// Создаем запись о команде в гонке с помощью Object Mother
	inputCrewResInRace := builders.CrewResInRaceMother.CustomCrew(uuid.New(), uuid.New(), 10, 0) // Некорректное значение для очков

	// Настраиваем мок для обновления записи о команде в гонке с ошибкой
	suite.mockRepo.UpdateFunc = func(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
		return nil, errors.New("failed to update record") // Возвращаем ошибку
	}

	// Пытаемся обновить запись о команде в гонке
	updatedCrewResInRace, err := suite.mockRepo.Update(inputCrewResInRace)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), updatedCrewResInRace)                  // Убедимся, что обновлённая запись равна nil
	require.EqualError(suite.T(), err, "failed to update record") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewResInRaceRepository запускает тесты в наборе
func TestCrewResInRaceRepository(t *testing.T) {
	suite.Run(t, new(CrewResInRaceRepositoryTestSuite))
}
