package judge_repository_tests

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
)

// Определяем структуру для тестового набора
type JudgeRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockJudgeRepository
}

// setupTest и teardownTest для каждого теста
func (suite *JudgeRepositoryTestSuite) SetupTest() {
	suite.mockRepo = &mocks.MockJudgeRepository{} // Возврат нового экземпляра мока для каждого теста
}

// Пример теста для успешного создания профиля судьи с использованием Builder
func (suite *JudgeRepositoryTestSuite) TestCreateProfile_Success() {

	// Используем Builder для создания судьи
	judge := builders2.NewJudgeBuilder().
		WithID(uuid.New()).
		WithFIO("John Doe").
		WithLogin("johndoe").
		WithPassword("password123").
		WithRole(1).
		WithPost("Judge").
		Build()

	suite.mockRepo.CreateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return j, nil
	}

	createdJudge, err := suite.mockRepo.CreateProfile(judge)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), judge.FIO, createdJudge.FIO)
	assert.Equal(suite.T(), judge.Login, createdJudge.Login)
	assert.Equal(suite.T(), judge.Password, createdJudge.Password)
	assert.Equal(suite.T(), judge.Role, createdJudge.Role)
}

// Пример теста для неудачного создания профиля судьи с использованием Mother
func (suite *JudgeRepositoryTestSuite) TestCreateProfile_Failure() {

	// Используем Object Mother для создания судьи
	judge := builders2.JudgeMother.CustomJudge(uuid.New(), "John Doe", "johndoe", "password123", 1, "Judge")

	suite.mockRepo.CreateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return nil, errors.New("creation failed")
	}

	createdJudge, err := suite.mockRepo.CreateProfile(judge)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdJudge)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetByID_Success() {

	// Используем Builder для создания судьи
	judge := builders2.NewJudgeBuilder().
		WithID(uuid.New()).
		WithFIO("John Doe").
		WithLogin("johndoe").
		WithPassword("password123").
		WithPost("Senior Judge").
		Build()

	suite.mockRepo.CreateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return j, nil
	}

	suite.mockRepo.GetJudgeDataByIDFunc = func(id uuid.UUID) (*models.Judge, error) {
		return judge, nil
	}

	createdJudge, err := suite.mockRepo.CreateProfile(judge)
	require.NoError(suite.T(), err)

	receivedJudge, err := suite.mockRepo.GetJudgeDataByID(createdJudge.ID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createdJudge.ID, receivedJudge.ID)
	require.Equal(suite.T(), createdJudge.FIO, receivedJudge.FIO)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetByID_Failure() {

	// Используем Object Mother для создания судьи
	judge := builders2.JudgeMother.CustomJudge(uuid.New(), "John Doe", "johndoe", "password123", 1, "Senior Judge")

	suite.mockRepo.CreateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return j, nil
	}

	suite.mockRepo.GetJudgeDataByIDFunc = func(id uuid.UUID) (*models.Judge, error) {
		return nil, errors.New("judge not found")
	}

	_, err := suite.mockRepo.CreateProfile(judge)
	require.NoError(suite.T(), err)

	receivedJudge, err := suite.mockRepo.GetJudgeDataByID(judge.ID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedJudge)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetJudgeDataByProtestID_Success() {
	// Используем Builder для создания судьи
	judge := builders2.NewJudgeBuilder().
		WithID(uuid.New()).
		WithFIO("John Doe").
		WithLogin("johndoe").
		WithPassword("password123").
		WithPost("Senior Judge").
		Build()

	protestID := uuid.New() // Генерируем новый UUID для протеста

	suite.mockRepo.GetJudgeDataByProtestIDFunc = func(id uuid.UUID) (*models.Judge, error) {
		return judge, nil
	}

	receivedJudge, err := suite.mockRepo.GetJudgeDataByProtestID(protestID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), judge.ID, receivedJudge.ID)
	require.Equal(suite.T(), judge.FIO, receivedJudge.FIO)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetJudgeDataByProtestID_Failure() {

	protestID := uuid.New() // Генерируем новый UUID для протеста

	suite.mockRepo.GetJudgeDataByProtestIDFunc = func(id uuid.UUID) (*models.Judge, error) {
		return nil, errors.New("judge not found")
	}

	receivedJudge, err := suite.mockRepo.GetJudgeDataByProtestID(protestID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedJudge)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetJudgeDataByLogin_Success() {

	// Используем Builder для создания судьи
	judge := builders2.NewJudgeBuilder().
		WithID(uuid.New()).
		WithFIO("John Doe").
		WithLogin("johndoe").
		WithPassword("password123").
		WithPost("Senior Judge").
		Build()

	suite.mockRepo.GetJudgeDataByLoginFunc = func(login string) (*models.Judge, error) {
		return judge, nil
	}

	receivedJudge, err := suite.mockRepo.GetJudgeDataByLogin(judge.Login)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), judge.ID, receivedJudge.ID)
	require.Equal(suite.T(), judge.FIO, receivedJudge.FIO)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetJudgeDataByLogin_Failure() {

	// Используем Object Mother для создания судьи
	judge := builders2.JudgeMother.CustomJudge(uuid.New(), "John Doe", "johndoe", "password123", 1, "Senior Judge")

	suite.mockRepo.GetJudgeDataByLoginFunc = func(login string) (*models.Judge, error) {
		return nil, errors.New("judge not found")
	}

	receivedJudge, err := suite.mockRepo.GetJudgeDataByLogin(judge.Login)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedJudge)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryDelete_Success() {

	// Используем Builder для создания судьи
	judge := builders2.NewJudgeBuilder().
		WithID(uuid.New()).
		WithFIO("John Doe").
		WithLogin("johndoe").
		WithPassword("password123").
		WithPost("Senior Judge").
		Build()

	// Настраиваем мок на удаление судьи
	suite.mockRepo.DeleteProfileFunc = func(id uuid.UUID) error {
		return nil // Успешное удаление
	}

	err := suite.mockRepo.DeleteProfile(judge.ID)
	require.NoError(suite.T(), err)

	// Проверяем, что судья действительно удален
	suite.mockRepo.GetJudgeDataByIDFunc = func(id uuid.UUID) (*models.Judge, error) {
		return nil, errors.New("judge not found") // Судья не найден
	}
	_, err = suite.mockRepo.GetJudgeDataByID(judge.ID)
	require.Error(suite.T(), err)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryDelete_Failure() {

	// Используем Object Mother для создания судьи
	judge := builders2.JudgeMother.CustomJudge(uuid.New(), "John Doe", "johndoe", "password123", 1, "Senior Judge")

	// Настраиваем мок на удаление судьи с ошибкой
	suite.mockRepo.DeleteProfileFunc = func(id uuid.UUID) error {
		return errors.New("deletion failed") // Ошибка удаления
	}

	err := suite.mockRepo.DeleteProfile(judge.ID)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, "deletion failed") // Проверяем, что ошибка соответствует ожиданиям
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryUpdate_Success() {

	// Используем Builder для создания судьи
	judge := builders2.NewJudgeBuilder().
		WithID(uuid.New()).
		WithFIO("John Doe").
		WithLogin("johndoe").
		WithPassword("password123").
		WithPost("Senior Judge").
		Build()

	// Настраиваем мок на создание судьи
	suite.mockRepo.CreateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return judge, nil
	}

	// Создаем судью
	createdJudge, err := suite.mockRepo.CreateProfile(judge)
	require.NoError(suite.T(), err)

	// Настраиваем мок на обновление судьи
	suite.mockRepo.UpdateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return j, nil // Возвращаем обновленного судью
	}

	// Обновляем судью
	updatedJudge, err := suite.mockRepo.UpdateProfile(createdJudge)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createdJudge.ID, updatedJudge.ID)
	require.Equal(suite.T(), createdJudge.FIO, updatedJudge.FIO)
	require.Equal(suite.T(), createdJudge.Login, updatedJudge.Login)
	require.Equal(suite.T(), createdJudge.Password, updatedJudge.Password)
	require.Equal(suite.T(), createdJudge.Role, updatedJudge.Role)
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryUpdate_Failure() {

	// Используем Object Mother для создания судьи
	judge := builders2.JudgeMother.CustomJudge(uuid.New(), "John Doe", "johndoe", "password123", 1, "Senior Judge")

	// Настраиваем мок на создание судьи
	suite.mockRepo.CreateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return judge, nil
	}

	// Создаем судью
	createdJudge, err := suite.mockRepo.CreateProfile(judge)
	require.NoError(suite.T(), err)

	// Настраиваем мок на обновление судьи с ошибкой
	suite.mockRepo.UpdateProfileFunc = func(j *models.Judge) (*models.Judge, error) {
		return nil, errors.New("update failed") // Ошибка обновления
	}

	// Пытаемся обновить судью
	updatedJudge, err := suite.mockRepo.UpdateProfile(createdJudge)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), updatedJudge)                // Убедимся, что обновленный судья равен nil
	require.EqualError(suite.T(), err, "update failed") // Проверяем, что ошибка соответствует ожиданиям
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetJudgesDataByRatingID_Success() {

	// Создаем рейтинг
	rating := builders2.NewRatingBuilder().WithID(uuid.New()).Build()

	// Создаем судей с помощью Builder
	judges := []models.Judge{
		*builders2.NewJudgeBuilder().WithID(uuid.New()).WithFIO("Judge 1").WithLogin("judge1").WithPassword("password1").WithPost("Post 1").WithRole(1).Build(),
		*builders2.NewJudgeBuilder().WithID(uuid.New()).WithFIO("Judge 2").WithLogin("judge2").WithPassword("password2").WithPost("Post 2").WithRole(1).Build(),
	}

	// Настраиваем мок для получения судей по ID рейтинга
	suite.mockRepo.GetJudgesDataByRatingIDFunc = func(ratingID uuid.UUID) ([]models.Judge, error) {
		return judges, nil // Возвращаем созданных судей
	}

	// Получаем судей
	receivedJudges, err := suite.mockRepo.GetJudgesDataByRatingID(rating.ID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(judges), len(receivedJudges))
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetJudgesDataByRatingID_Failure() {

	// Создаем рейтинг
	rating := builders2.RatingMother.Default()

	// Настраиваем мок для получения судей по ID рейтинга с ошибкой
	suite.mockRepo.GetJudgesDataByRatingIDFunc = func(ratingID uuid.UUID) ([]models.Judge, error) {
		return nil, errors.New("failed to get judges") // Возвращаем ошибку
	}

	// Пытаемся получить судей
	receivedJudges, err := suite.mockRepo.GetJudgesDataByRatingID(rating.ID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedJudges)                     // Убедимся, что полученные судьи равны nil
	require.EqualError(suite.T(), err, "failed to get judges") // Проверяем, что ошибка соответствует ожиданиям
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetAllJudges_Success() {

	// Создаем судей с помощью Builder
	judges := []models.Judge{
		*builders2.NewJudgeBuilder().WithID(uuid.New()).WithFIO("Judge 1").WithLogin("judge1").WithPassword("password1").WithPost("Post 1").WithRole(1).Build(),
		*builders2.NewJudgeBuilder().WithID(uuid.New()).WithFIO("Judge 2").WithLogin("judge2").WithPassword("password2").WithPost("Post 2").WithRole(1).Build(),
	}

	// Настраиваем мок для получения всех судей
	suite.mockRepo.GetAllJudgesFunc = func() ([]models.Judge, error) {
		return judges, nil // Возвращаем созданных судей
	}

	// Получаем всех судей
	receivedJudges, err := suite.mockRepo.GetAllJudges()
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(judges), len(receivedJudges))
}

func (suite *JudgeRepositoryTestSuite) TestJudgeRepositoryGetAllJudges_Failure() {

	// Настраиваем мок для получения всех судей с ошибкой
	suite.mockRepo.GetAllJudgesFunc = func() ([]models.Judge, error) {
		return nil, errors.New("failed to get all judges") // Возвращаем ошибку
	}

	// Пытаемся получить всех судей
	receivedJudges, err := suite.mockRepo.GetAllJudges()
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedJudges)                         // Убедимся, что полученные судьи равны nil
	require.EqualError(suite.T(), err, "failed to get all judges") // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewResInRaceRepository запускает тесты в наборе
func TestJudgeRepository(t *testing.T) {
	suite.Run(t, new(JudgeRepositoryTestSuite))
}
