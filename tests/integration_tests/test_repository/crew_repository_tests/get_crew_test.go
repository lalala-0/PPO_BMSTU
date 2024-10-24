package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryGetCrewDataByID_Success тестирует успешное получение записи о команде по ID
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataByID_Success() {
	// arrange
	// чистим бд
	suite.initializer.ClearAll()
	// добавляем все нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)
	crew, err := suite.initializer.CreateCrew(builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1))
	require.NoError(suite.T(), err)

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataByID(crew.ID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), crew.ID, receivedCrew.ID)
	require.Equal(suite.T(), crew.RatingID, receivedCrew.RatingID)
	require.Equal(suite.T(), crew.SailNum, receivedCrew.SailNum)
	require.Equal(suite.T(), crew.Class, receivedCrew.Class)
}

// TestCrewRepositoryGetCrewDataByID_Failure тестирует ошибку при получении записи о команде по ID
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataByID_Failure() {
	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataByID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrew)                                       // Убедимся, что полученная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
