package test_repository

import (
	"PPO_BMSTU/internal/models"
	builders2 "PPO_BMSTU/tests/unit_tests/builders"
	"PPO_BMSTU/tests/unit_tests/test_repository/mocks"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type RaceRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockRaceRepository
}

func (suite *RaceRepositoryTestSuite) SetupTest() {
	suite.mockRepo = &mocks.MockRaceRepository{}
}

// Позитивный тест для создания гонки
func (suite *RaceRepositoryTestSuite) TestCreate_Success() {
	// Создаем гонку с помощью Builder
	inputRace := builders2.NewRaceBuilder().WithID(uuid.New()).WithRatingID(uuid.New()).WithDate(time.Now()).WithNumber(1).WithClass(4).Build()

	// Настраиваем мок для создания гонки
	suite.mockRepo.CreateFunc = func(race *models.Race) (*models.Race, error) {
		return race, nil // Возвращаем созданную гонку
	}

	// Пытаемся создать гонку
	createdRace, err := suite.mockRepo.Create(inputRace)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputRace.RatingID, createdRace.RatingID)
	require.Equal(suite.T(), inputRace.Date, createdRace.Date)
	require.Equal(suite.T(), inputRace.Number, createdRace.Number)
	require.Equal(suite.T(), inputRace.Class, createdRace.Class)
}

// Негативный тест для создания гонки
func (suite *RaceRepositoryTestSuite) TestCreate_Failure() {
	// Создаем гонку с помощью Builder
	inputRace := builders2.NewRaceBuilder().WithID(uuid.New()).WithRatingID(uuid.New()).WithDate(time.Now()).WithNumber(1).WithClass(4).Build()

	// Настраиваем мок для создания гонки с ошибкой
	suite.mockRepo.CreateFunc = func(race *models.Race) (*models.Race, error) {
		return nil, errors.New("failed to create race") // Возвращаем ошибку
	}

	// Пытаемся создать гонку
	createdRace, err := suite.mockRepo.Create(inputRace)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), createdRace)                         // Убедимся, что созданная гонка равна nil
	require.EqualError(suite.T(), err, "failed to create race") // Проверяем, что ошибка соответствует ожиданиям
}

// Позитивный тест для получения гонки по ID
func (suite *RaceRepositoryTestSuite) TestGetRaceDataByID_Success() {
	// Используем Object Mother для создания рейтинга и гонки
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	race := builders2.NewRaceBuilder().WithID(uuid.New()).WithRatingID(rating.ID).WithDate(time.Now()).WithNumber(1).WithClass(4).Build()

	// Моковая функция для получения гонки по ID
	suite.mockRepo.GetRaceDataByIDFunc = func(id uuid.UUID) (*models.Race, error) {
		return race, nil // Возвращаем созданную гонку
	}

	// Получаем гонку по ID
	receivedRace, err := suite.mockRepo.GetRaceDataByID(race.ID)

	// Проверяем, что получение гонки прошло успешно
	assert.NoError(suite.T(), err)
	require.Equal(suite.T(), race.RatingID, receivedRace.RatingID)
	require.Equal(suite.T(), race.Date, receivedRace.Date)
	require.Equal(suite.T(), race.Number, receivedRace.Number)
	require.Equal(suite.T(), race.Class, receivedRace.Class)
}

// Негативный тест для получения гонки по ID
func (suite *RaceRepositoryTestSuite) TestGetRaceDataByID_Failure() {
	// Моковая функция для получения гонки по ID с ошибкой
	suite.mockRepo.GetRaceDataByIDFunc = func(id uuid.UUID) (*models.Race, error) {
		return nil, errors.New("race not found") // Возвращаем ошибку
	}

	// Пытаемся получить гонку по несуществующему ID
	receivedRace, err := suite.mockRepo.GetRaceDataByID(uuid.New()) // Используем новый UUID

	// Проверяем, что произошла ошибка
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedRace)                 // Убедимся, что полученная гонка равна nil
	require.EqualError(suite.T(), err, "race not found") // Проверяем, что ошибка соответствует ожиданиям
}

// Позитивный тест для удаления гонки
func (suite *RaceRepositoryTestSuite) TestDelete_Success() {
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	race := builders2.NewRaceBuilder().WithID(uuid.New()).WithRatingID(rating.ID).WithDate(time.Now()).WithNumber(1).WithClass(4).Build()

	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return nil // Успешное удаление гонки
	}

	err := suite.mockRepo.Delete(race.ID)

	require.NoError(suite.T(), err)
}

// Негативный тест для удаления гонки
func (suite *RaceRepositoryTestSuite) TestDelete_Failure() {
	suite.mockRepo.DeleteFunc = func(id uuid.UUID) error {
		return errors.New("delete failed") // Возвращаем ошибку
	}

	err := suite.mockRepo.Delete(uuid.New()) // Используем новый UUID

	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "delete failed") // Проверяем, что ошибка соответствует ожиданиям
}

// Негативный тест для обновления гонки
func (suite *RaceRepositoryTestSuite) TestUpdate_Failure() {
	// Моковая функция для обновления гонки с ошибкой
	suite.mockRepo.UpdateFunc = func(race *models.Race) (*models.Race, error) {
		return nil, errors.New("update failed") // Возвращаем ошибку
	}

	// Создаем гонку с использованием Builder
	race := builders2.RaceMother.Default()

	// Пытаемся обновить гонку
	updatedRace, err := suite.mockRepo.Update(race)

	// Проверяем, что произошла ошибка
	require.Error(suite.T(), err)
	require.Nil(suite.T(), updatedRace)                 // Убедимся, что обновленная гонка равна nil
	require.EqualError(suite.T(), err, "update failed") // Проверяем, что ошибка соответствует ожиданиям
}

// Позитивный тест для обновления гонки
func (suite *RaceRepositoryTestSuite) TestUpdate_Success() {
	// Используем Object Mother для создания рейтинга и гонки
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	race := builders2.NewRaceBuilder().
		WithID(uuid.New()).
		WithRatingID(rating.ID).
		WithDate(time.Now()).
		WithNumber(1).
		WithClass(4).
		Build()

	// Моковая функция для обновления гонки
	suite.mockRepo.UpdateFunc = func(race *models.Race) (*models.Race, error) {
		// Возвращаем обновленную гонку с новыми значениями
		return &models.Race{
			ID:       race.ID,
			RatingID: race.RatingID,
			Date:     time.Date(2012, time.November, 11, 23, 0, 0, 0, time.UTC),
			Number:   2,
			Class:    5,
		}, nil
	}

	// Обновляем гонку с использованием Builder
	updatedRace, err := suite.mockRepo.Update(builders2.NewRaceBuilder().
		WithID(race.ID).
		WithRatingID(rating.ID).
		WithDate(time.Date(2012, time.November, 11, 23, 0, 0, 0, time.UTC)).
		WithNumber(2).
		WithClass(5).
		Build(),
	)

	// Проверяем, что обновление прошло успешно
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), updatedRace.RatingID, race.RatingID)
	require.NotEqual(suite.T(), updatedRace.Date, race.Date)
	require.NotEqual(suite.T(), updatedRace.Number, race.Number)
	require.NotEqual(suite.T(), updatedRace.Class, race.Class)
}

// Негативный тест для получения гонок по ID рейтинга
func (suite *RaceRepositoryTestSuite) TestGetRacesDataByRatingID_Failure() {
	// Моковая функция для получения гонок по ID рейтинга с ошибкой
	suite.mockRepo.GetRacesDataByRatingIDFunc = func(ratingID uuid.UUID) ([]models.Race, error) {
		return nil, errors.New("no races found") // Возвращаем ошибку
	}

	// Пытаемся получить гонки по несуществующему ID рейтинга
	races, err := suite.mockRepo.GetRacesDataByRatingID(uuid.New()) // Используем новый UUID

	// Проверяем, что произошла ошибка
	require.Error(suite.T(), err)
	require.Nil(suite.T(), races)                        // Убедимся, что список гонок равен nil
	require.EqualError(suite.T(), err, "no races found") // Проверяем, что ошибка соответствует ожиданиям
}

// Позитивный тест для получения гонок по ID рейтинга
func (suite *RaceRepositoryTestSuite) TestGetRacesDataByRatingID_Success() {
	rating := builders2.RatingMother.CustomRating(uuid.New(), "RatingName", models.Laser, 1)
	race := builders2.NewRaceBuilder().WithID(uuid.New()).WithRatingID(rating.ID).WithDate(time.Now()).WithNumber(1).WithClass(4).Build()

	// Настраиваем мок для получения гонок по ID рейтинга
	suite.mockRepo.GetRacesDataByRatingIDFunc = func(ratingID uuid.UUID) ([]models.Race, error) {
		return []models.Race{*race}, nil // Возвращаем список гонок
	}

	// Получаем гонки по ID рейтинга
	races, err := suite.mockRepo.GetRacesDataByRatingID(rating.ID)

	// Проверяем, что получение гонок прошло успешно
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(races))
	require.Equal(suite.T(), race.RatingID, races[0].RatingID)
	require.Equal(suite.T(), race.Date, races[0].Date)
	require.Equal(suite.T(), race.Number, races[0].Number)
	require.Equal(suite.T(), race.Class, races[0].Class)
}

// Запуск тестового сьюта
func TestRaceRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RaceRepositoryTestSuite))
}
