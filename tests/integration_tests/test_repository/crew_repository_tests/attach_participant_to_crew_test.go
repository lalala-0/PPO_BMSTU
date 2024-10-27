package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryAttachParticipantToCrew_Success тестирует успешное добавление участника в команду
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryAttachParticipantToCrew_Success() {
	// arrange
	// чистим бд
	err := suite.initializer.ClearAll()
	assert.NoError(suite.T(), err)
	inputRating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	assert.NoError(suite.T(), err)
	crew, err := suite.initializer.CreateCrew(builders.CrewMother.WithRatingID(inputRating.ID))
	assert.NoError(suite.T(), err)
	participant, err := suite.initializer.CreateParticipant(builders.ParticipantMother.Default())
	require.NoError(suite.T(), err)

	// act
	err = suite.repo.AttachParticipantToCrew(participant.ID, crew.ID, 1)

	// assert
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryAttachParticipantToCrew_Failure тестирует ошибку при добавлении участника в команду
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryAttachParticipantToCrew_Failure() {
	// Пытаемся добавить участника в команду
	err := suite.repo.AttachParticipantToCrew(uuid.New(), uuid.New(), 1)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, repository_errors.InsertError.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
