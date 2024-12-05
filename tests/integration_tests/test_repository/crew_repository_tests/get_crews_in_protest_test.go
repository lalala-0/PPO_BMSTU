package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryGetCrewsDataByProtestID_Success тестирует успешное получение списка команд по ID протеста
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByProtestID_Success() {
	// arrange
	// чистим бд
	// добавляем все нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)
	judge, err := suite.initializer.CreateJudge(builders.JudgeMother.Default())
	require.NoError(suite.T(), err)
	race, err := suite.initializer.CreateRace(builders.RaceMother.WithRatingID(rating.ID))
	require.NoError(suite.T(), err)
	protest, err := suite.initializer.CreateProtest(builders.NewProtestBuilder().WithJudgeID(judge.ID).WithRatingID(rating.ID).WithRaceID(race.ID).Build())
	require.NoError(suite.T(), err)

	for _ = range 3 {
		crew, err := suite.initializer.CreateCrew(builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1))
		require.NoError(suite.T(), err)

		err = suite.initializer.AttachCrewToProtest(crew.ID, protest.ID)
		require.NoError(suite.T(), err)
	}

	// act
	// Пытаемся получить список команд по ID протеста
	receivedCrews, err := suite.repo.GetCrewsDataByProtestID(protest.ID)
	require.NoError(suite.T(), err)
	require.Len(suite.T(), receivedCrews, 3)
}

// TestCrewRepositoryGetCrewsDataByProtestID_Failure тестирует ошибку при получении списка команд по ID протеста
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByProtestID_Failure() {
	// Пытаемся получить список команд по ID протеста
	receivedCrews, err := suite.repo.GetCrewsDataByProtestID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrews)                                      // Убедимся, что полученный список команд равен nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
