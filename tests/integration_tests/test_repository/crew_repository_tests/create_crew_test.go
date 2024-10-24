package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryCreate_Success тестирует успешное создание записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryCreate_Success() {
	// чистим бд
	suite.initializer.ClearAll()
	// добавляем все нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1)

	// Пытаемся создать запись о команде
	createdCrew, err := suite.repo.Create(inputCrew)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputCrew.RatingID, createdCrew.RatingID)
	require.Equal(suite.T(), inputCrew.SailNum, createdCrew.SailNum)
	require.Equal(suite.T(), inputCrew.Class, createdCrew.Class)
}

// TestCrewRepositoryCreate_Failure тестирует ошибку при создании записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryCreate_Failure() {
	// Создаем запись о команде с помощью Object Mother
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 0, 0) // Некорректное значение для номера паруса

	// Пытаемся создать запись о команде
	createdCrew, err := suite.repo.Create(inputCrew)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), createdCrew)                                       // Убедимся, что созданная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.InsertError.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
