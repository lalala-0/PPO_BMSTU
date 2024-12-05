package crew_repository_tests

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/tests/unit_tests/builders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCrewRepositoryUpdate_Success тестирует успешное обновление записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryUpdate_Success() {
	// Чистим БД
	// Добавляем нужные записи
	rating, err := suite.initializer.CreateRating(builders.RatingMother.Default())
	require.NoError(suite.T(), err)

	// Создаем команду для обновления
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), rating.ID, 1, 1)
	createdCrew, err := suite.repo.Create(inputCrew)
	require.NoError(suite.T(), err)

	// Изменяем данные команды
	updatedCrew := *createdCrew
	updatedCrew.Class = 4
	updatedCrew.SailNum = 2 // Обновляем номер паруса

	// Пытаемся обновить запись о команде
	crewAfterUpdate, err := suite.repo.Update(&updatedCrew)
	require.NoError(suite.T(), err)

	// Проверяем, что обновленная запись совпадает с ожидаемыми данными
	require.Equal(suite.T(), updatedCrew.ID, crewAfterUpdate.ID)
	require.Equal(suite.T(), updatedCrew.RatingID, crewAfterUpdate.RatingID)
	require.Equal(suite.T(), updatedCrew.SailNum, crewAfterUpdate.SailNum)
	require.Equal(suite.T(), updatedCrew.Class, crewAfterUpdate.Class)
}

// TestCrewRepositoryUpdate_Failure тестирует ошибку при обновлении записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryUpdate_Failure() {
	// Чистим БД
	// Создаем запись о команде с помощью Object Mother
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 0, 0) // Некорректное значение для номера паруса

	// Пытаемся обновить несуществующую команду
	crewAfterUpdate, err := suite.repo.Update(inputCrew)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), crewAfterUpdate)                                   // Убедимся, что обновленная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.UpdateError.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
