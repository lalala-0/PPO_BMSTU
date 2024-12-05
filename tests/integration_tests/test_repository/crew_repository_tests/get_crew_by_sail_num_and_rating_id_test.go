package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Success тестирует успешное получение записи о команде по номеру паруса и ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Success() {
	// arrange
	// чистим бд
	// добавляем все нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)
	crew, err := suite.initializer.CreateCrew(builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1))
	require.NoError(suite.T(), err)

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataBySailNumAndRatingID(crew.SailNum, crew.RatingID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), crew.ID, receivedCrew.ID)
	require.Equal(suite.T(), crew.SailNum, receivedCrew.SailNum)
	require.Equal(suite.T(), crew.RatingID, receivedCrew.RatingID)
}

// TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Failure тестирует ошибку при получении записи о команде по номеру паруса и ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Failure() {
	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataBySailNumAndRatingID(1, uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrew)                                       // Убедимся, что полученная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
