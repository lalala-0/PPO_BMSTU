package test_e2e

import (
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestE2EScenario_OpenRatingRaceProtestParticipants тестирует сценарий:
// 1. Пользователь открывает рейтинг
// 2. Пользователь открывает гонку
// 3. Пользователь переходит к протестам
// 4. Пользователь просматривает участников протеста
func (suite *e2eTestSuite) TestE2EScenario_OpenRatingRaceProtestParticipants() {
	// Инициализируем необходимые данные
	err := suite.initializer.ClearAll()
	assert.NoError(suite.T(), err)
	inputRating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	assert.NoError(suite.T(), err)
	inputRace, err := suite.initializer.CreateRace(builders.RaceMother.WithRatingID(inputRating.ID))
	assert.NoError(suite.T(), err)
	inputJudge, err := suite.initializer.CreateJudge(builders.JudgeMother.Default())
	assert.NoError(suite.T(), err)
	inputProtest, err := suite.initializer.CreateProtest(builders.NewProtestBuilder().WithRaceID(inputRace.ID).WithRatingID(inputRating.ID).WithJudgeID(inputJudge.ID).Build())
	assert.NoError(suite.T(), err)

	for _ = range 3 {
		crew, err := suite.initializer.CreateCrew(builders.CrewMother.CustomCrew(uuid.New(), inputRating.ID, 1, 1))
		require.NoError(suite.T(), err)

		err = suite.initializer.AttachCrewToProtest(crew.ID, inputProtest.ID)
		require.NoError(suite.T(), err)
	}

	// Шаг 1: Пользователь открывает рейтинг
	rating, err := suite.ratingService.GetRatingDataByID(inputRating.ID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), rating)

	// Шаг 2: Пользователь выбирает гонку
	race, err := suite.raceService.GetRaceDataByID(inputRace.ID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), race)

	// Шаг 3: Пользователь переходит на страницу протестов этой гонки
	protests, err := suite.protestService.GetProtestsDataByRaceID(inputRace.ID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), protests)
	require.Greater(suite.T(), len(protests), 0) // Убедимся, что есть хотя бы один протест

	// Шаг 4: Пользователь просматривает участников протеста
	protest := protests[0] // Берем первый протест
	participantIDs, err := suite.protestService.GetProtestParticipantsIDByID(protest.ID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), participantIDs)
	require.Equal(suite.T(), len(participantIDs), 3) // Убедимся, что есть участники протеста

	// Шаг 5: Пользователь просматривает подробную информацию о каждом команде-участнике протеста
	for participantID, _ := range participantIDs {
		protestParticipant, err := suite.crewService.GetCrewDataByID(participantID)
		require.NoError(suite.T(), err)
		require.NotNil(suite.T(), protestParticipant)
		require.NotEmpty(suite.T(), protestParticipant.SailNum)
		require.NotEmpty(suite.T(), protestParticipant.ID)
	}
}
