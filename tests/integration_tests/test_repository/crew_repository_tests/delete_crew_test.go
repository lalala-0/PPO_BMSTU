package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryDelete_Success тестирует успешное удаление записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDelete_Success() {
	// Чистим БД
	// Добавляем нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)

	// Создаем команду для обновления
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1)
	createdCrew, err := suite.repo.Create(inputCrew)
	require.NoError(suite.T(), err)

	// Пытаемся удалить запись о команде
	err = suite.repo.Delete(createdCrew.ID)
	require.NoError(suite.T(), err)

	// Проверяем, что запись была удалена
	deletedCrew, err := suite.repo.GetCrewDataByID(createdCrew.ID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), deletedCrew)
}

// TestCrewRepositoryDelete_Failure тестирует ошибку при удалении записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDelete_Failure() {
	// Пытаемся удалить запись о команде
	err := suite.repo.Delete(uuid.New())
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
