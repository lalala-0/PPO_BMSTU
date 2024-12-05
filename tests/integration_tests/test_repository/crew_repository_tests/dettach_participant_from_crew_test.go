package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryDetachParticipantFromCrew_Success тестирует успешное удаление участника из команды
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDetachParticipantFromCrew_Success() {
	// arrange
	// чистим бд
	// добавляем все нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)
	crew, err := suite.initializer.CreateCrew(builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1))
	require.NoError(suite.T(), err)
	participant, err := suite.initializer.CreateParticipant(builders.ParticipantMother.Default())
	require.NoError(suite.T(), err)
	err = suite.initializer.AttachParticipantToCrew(participant.ID, crew.ID)
	require.NoError(suite.T(), err)

	// act
	err = suite.repo.DetachParticipantFromCrew(participant.ID, crew.ID)

	// assert
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryDetachParticipantFromCrew_Failure тестирует ошибку при удалении участника из команды
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDetachParticipantFromCrew_Failure() {
	// Пытаемся удалить участника из команды
	err := suite.repo.DetachParticipantFromCrew(uuid.New(), uuid.New())
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
