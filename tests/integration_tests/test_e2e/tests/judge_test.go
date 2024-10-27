package tests

// TestJudgeCreatesRatingTeamsAndParticipants: основной тестовый сценарий: Судья создает рейтинг, добавляет три команды, по одному участнику в каждую команду, запускает стартовую процедуру гонки и завершает её финишной процедурой.
func (suite *e2eTestSuite) TestJudgeCreatesRatingTeamsAndParticipants() {
	// Создать рейтинг, используя сервис RatingService
	//rating, err := suite.ratingService.AddNewRating(uuid.New(), "e2eTest rating", 1, 1)
	//assert.NoError(suite.T(), err)
	//
	//// Создать три команды для рейтинга по одному участнику
	//for i := range 3 {
	//	crew, err := suite.crewService.AddNewCrew(uuid.New(), rating.ID, 1, i+2)
	//	assert.NoError(suite.T(), err)
	//	participant, err := suite.participantService.AddNewParticipant(uuid.New(), "test", 1, 1, time.Now(), "test")
	//	assert.NoError(suite.T(), err)
	//	// Добавить участников в каждую из команд, используя ParticipantService
	//	err = suite.crewService.AttachParticipantToCrew(participant.ID, crew.ID, 1)
	//	assert.NoError(suite.T(), err)
	//}

	// Создать гонку
	//_, err = suite.raceService.AddNewRace(uuid.New(), rating.ID, 1, time.Now(), 1)
	//assert.NoError(suite.T(), err)

	// Провести стартовую процедуру гонки с помощью RaceService

	// Провести финишную процедуру для гонки, используя RaceService
	// Пример: err := suite.raceService.FinishRaceProcedure(...)

}
