package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryGetCrewsDataByRatingID_Success тестирует успешное получение списка команд по ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByRatingID_Success() {
	// arrange
	// чистим бд
	suite.initializer.ClearAll()
	// добавляем все нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)

	for _ = range 3 {
		_, err := suite.initializer.CreateCrew(builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1))
		require.NoError(suite.T(), err)
	}

	// act
	// Пытаемся получить список команд по ID рейтинга
	receivedCrews, err := suite.repo.GetCrewsDataByRatingID(rating.ID)
	require.NoError(suite.T(), err)
	require.Len(suite.T(), receivedCrews, 3)
}

// TestCrewRepositoryGetCrewsDataByRatingID_Failure тестирует ошибку при получении списка команд по ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByRatingID_Failure() {
	// Пытаемся получить список команд по ID рейтинга
	receivedCrews, err := suite.repo.GetCrewsDataByRatingID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrews)                                      // Убедимся, что полученный список команд равен nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
