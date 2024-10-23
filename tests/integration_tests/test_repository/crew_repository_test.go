package test_repository

import (
	"PPO_BMSTU/internal/models"
	mongo_rep "PPO_BMSTU/internal/repository/mongo"
	postgres_rep "PPO_BMSTU/internal/repository/postgres"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/tests/integration_tests/db_init"
	"PPO_BMSTU/tests/unit_tests/builders"
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"runtime"
	"testing"
)

// CrewRepositoryTestSuite описывает тестовый набор для ICrewRepository
type CrewRepositoryTestSuite struct {
	suite.Suite
	postgresClient *sqlx.DB
	mongoClient    *mongo.Database
	postgresC      testcontainers.Container
	mongoC         testcontainers.Container
	repo           repository_interfaces.ICrewRepository // Изменено на общий интерфейс
}

// TestICrewRepository запускает тесты в наборе
func TestICrewRepository(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Тестирование PostgreSQL
	os.Setenv("DB_TYPE", "postgres")
	suite.Run(t, new(CrewRepositoryTestSuite))

	//// Тестирование MongoDB
	//os.Setenv("DB_TYPE", "mongo")
	//suite.Run(t, new(CrewRepositoryTestSuite))
}

func (suite *CrewRepositoryTestSuite) SetupSuite() {
	dbType := os.Getenv("DB_TYPE") // Получаем значение переменной окружения DB_TYPE

	var err error
	if dbType == "postgres" || dbType == "" { // По умолчанию тестируем PostgreSQL, если не указано
		suite.postgresC, suite.postgresClient, err = db_init.SetupTestDatabasePostgres()
		require.NoError(suite.T(), err)

		suite.repo = postgres_rep.NewCrewRepository(suite.postgresClient)
	}

	if dbType == "mongo" { // Тестируем MongoDB, если указано
		suite.mongoC, suite.mongoClient = db_init.SetupTestDatabaseMongo()

		// Проверяем, что клиент MongoDB и контейнер не nil
		require.NotNil(suite.T(), suite.mongoClient, "Mongo client is nil")
		require.NotNil(suite.T(), suite.mongoC, "Mongo container is nil")

		suite.repo = mongo_rep.NewCrewRepository(suite.mongoClient)
	}
}

// TearDownSuite выполняется один раз после завершения тестов
func (suite *CrewRepositoryTestSuite) TearDownSuite() {
	// Завершаем работу с PostgreSQL контейнером
	if suite.postgresC != nil {
		err := suite.postgresC.Terminate(context.Background())
		require.NoError(suite.T(), err)
	}

	// Завершаем работу с MongoDB контейнером
	if suite.mongoC != nil {
		err := suite.mongoC.Terminate(context.Background())
		require.NoError(suite.T(), err)
	}

	// Закрываем подключение к PostgreSQL
	if suite.postgresClient != nil {
		err := suite.postgresClient.Close()
		require.NoError(suite.T(), err)
	}

	// Закрываем подключение к MongoDB
	if suite.mongoClient != nil {
		err := suite.mongoClient.Client().Disconnect(context.Background())
		require.NoError(suite.T(), err)
	}
}

// TestCrewRepositoryCreate_Success тестирует успешное создание записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryCreate_Success() {
	// Создаем запись о команде с помощью Builder
	inputCrew := builders.NewCrewBuilder().
		WithSailNum(2).
		WithClass(1).
		Build()

	// Пытаемся создать запись о команде
	createdCrew, err := suite.repo.Create(inputCrew)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), inputCrew.RatingID, createdCrew.RatingID)
	require.Equal(suite.T(), inputCrew.SailNum, createdCrew.SailNum)
	require.Equal(suite.T(), inputCrew.Class, createdCrew.Class)
}

// TestCrewRepositoryCreate_Failure тестирует ошибку при создании записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryCreate_Failure() {
	// Создаем запись о команде с помощью Object Mother
	inputCrew := builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 0, 0) // Некорректное значение для номера паруса

	// Пытаемся создать запись о команде
	createdCrew, err := suite.repo.Create(inputCrew)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), createdCrew)                                       // Убедимся, что созданная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.InsertError.Error()) // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryDelete_Success тестирует успешное удаление записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDelete_Success() {
	// Создаем запись о команде перед удалением
	crew := builders.NewCrewBuilder().WithSailNum(3).WithClass(2).Build()
	_, err := suite.repo.Create(crew)
	require.NoError(suite.T(), err)

	// Пытаемся удалить запись о команде
	err = suite.repo.Delete(uuid.New())
	require.NoError(suite.T(), err)

	// Проверяем, что запись была удалена
	deletedCrew, err := suite.repo.GetCrewDataByID(crew.ID)
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

// TestCrewRepositoryUpdate_Success тестирует успешное обновление записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryUpdate_Success() {
	// Создаем запись о команде перед обновлением
	crew := builders.NewCrewBuilder().WithSailNum(4).WithClass(2).Build()
	_, err := suite.repo.Create(crew)
	require.NoError(suite.T(), err)

	// Обновляем номер паруса
	crew.SailNum = 5
	updatedCrew, err := suite.repo.Update(crew)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), crew.ID, updatedCrew.ID)
	require.Equal(suite.T(), crew.SailNum, updatedCrew.SailNum)
}

// TestCrewRepositoryUpdate_Failure тестирует ошибку при обновлении записи о команде
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryUpdate_Failure() {
	// Создаем запись о команде перед обновлением
	crew := builders.NewCrewBuilder().WithSailNum(4).WithClass(2).Build()
	_, err := suite.repo.Create(crew)
	require.NoError(suite.T(), err)

	// Пытаемся обновить запись о команде
	crew.SailNum = -5
	updatedCrew, err := suite.repo.Update(crew)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), updatedCrew)                                        // Убедимся, что обновленная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryGetCrewDataByID_Success тестирует успешное получение записи о команде по ID
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataByID_Success() {
	// Создаем запись о команде с помощью Builder
	crew := builders.NewCrewBuilder().
		WithSailNum(4).
		WithClass(3).
		Build()

	// Сохраняем запись в базе данных
	_, err := suite.repo.Create(crew)
	require.NoError(suite.T(), err)

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataByID(crew.ID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), crew.ID, receivedCrew.ID)
	require.Equal(suite.T(), crew.RatingID, receivedCrew.RatingID)
	require.Equal(suite.T(), crew.SailNum, receivedCrew.SailNum)
	require.Equal(suite.T(), crew.Class, receivedCrew.Class)
}

// TestCrewRepositoryGetCrewDataByID_Failure тестирует ошибку при получении записи о команде по ID
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataByID_Failure() {
	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataByID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrew)                                       // Убедимся, что полученная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryGetCrewsDataByRatingID_Success тестирует успешное получение списка команд по ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByRatingID_Success() {
	// Создаем записи о командах
	ratingID := uuid.New()
	crews := []models.Crew{
		*builders.NewCrewBuilder().WithSailNum(5).WithClass(1).WithRatingID(ratingID).Build(),
		*builders.NewCrewBuilder().WithSailNum(6).WithClass(2).WithRatingID(ratingID).Build(),
	}

	// Сохраняем записи в базе данных
	for _, crew := range crews {
		_, err := suite.repo.Create(&crew)
		require.NoError(suite.T(), err)
	}

	// Пытаемся получить список команд по ID рейтинга
	receivedCrews, err := suite.repo.GetCrewsDataByRatingID(ratingID)
	require.NoError(suite.T(), err)
	require.Len(suite.T(), receivedCrews, len(crews))
	require.Equal(suite.T(), crews[0].SailNum, receivedCrews[0].SailNum)
	require.Equal(suite.T(), crews[1].SailNum, receivedCrews[1].SailNum)
}

// TestCrewRepositoryGetCrewsDataByRatingID_Failure тестирует ошибку при получении списка команд по ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByRatingID_Failure() {
	// Пытаемся получить список команд по ID рейтинга
	receivedCrews, err := suite.repo.GetCrewsDataByRatingID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrews)                                      // Убедимся, что полученный список команд равен nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}

//// TestCrewRepositoryGetCrewsDataByProtestID_Success тестирует успешное получение списка команд по ID протеста
//func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByProtestID_Success() {
//	// Создаем записи о командах
//	protestID := uuid.New()
//	crews := []models.Crew{
//		*builders.NewCrewBuilder().WithSailNum(1).WithClass(1).WithProtestID(protestID).Build(),
//		*builders.NewCrewBuilder().WithSailNum(2).WithClass(2).WithProtestID(protestID).Build(),
//	}
//
//	// Сохраняем записи в базе данных
//	for _, crew := range crews {
//		_, err := suite.repo.Create(crew)
//		require.NoError(suite.T(), err)
//	}
//
//	// Пытаемся получить список команд по ID протеста
//	receivedCrews, err := suite.repo.GetCrewsDataByProtestID(protestID)
//	require.NoError(suite.T(), err)
//	require.Len(suite.T(), receivedCrews, len(crews))
//	require.Equal(suite.T(), crews[0].SailNum, receivedCrews[0].SailNum)
//	require.Equal(suite.T(), crews[1].SailNum, receivedCrews[1].SailNum)
//}

// TestCrewRepositoryGetCrewsDataByProtestID_Failure тестирует ошибку при получении списка команд по ID протеста
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewsDataByProtestID_Failure() {
	// Пытаемся получить список команд по ID протеста
	receivedCrews, err := suite.repo.GetCrewsDataByProtestID(uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrews)                                      // Убедимся, что полученный список команд равен nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Success тестирует успешное получение записи о команде по номеру паруса и ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Success() {
	crew := builders.NewCrewBuilder().
		WithSailNum(1).
		WithRatingID(uuid.New()).
		Build()

	// Сохраняем запись в базе данных
	_, err := suite.repo.Create(crew)
	require.NoError(suite.T(), err)

	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataBySailNumAndRatingID(1, crew.RatingID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), crew.ID, receivedCrew.ID)
	require.Equal(suite.T(), crew.SailNum, receivedCrew.SailNum)
	require.Equal(suite.T(), crew.RatingID, receivedCrew.RatingID)
}

// TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Failure тестирует ошибку при получении записи о команде по номеру паруса и ID рейтинга
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryGetCrewDataBySailNumAndRatingID_Failure() {
	// Пытаемся получить запись о команде
	receivedCrew, err := suite.repo.GetCrewDataBySailNumAndRatingID(1, uuid.New())
	require.Error(suite.T(), err)
	require.Nil(suite.T(), receivedCrew)                                       // Убедимся, что полученная запись равна nil
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryAttachParticipantToCrew_Success тестирует успешное добавление участника в команду
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryAttachParticipantToCrew_Success() {
	crew := builders.NewCrewBuilder().
		WithSailNum(1).
		Build()

	// Сохраняем запись о команде
	crewRecord, err := suite.repo.Create(crew)
	require.NoError(suite.T(), err)

	// Пытаемся добавить участника в команду
	err = suite.repo.AttachParticipantToCrew(crewRecord.ID, uuid.New(), 1)
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryAttachParticipantToCrew_Failure тестирует ошибку при добавлении участника в команду
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryAttachParticipantToCrew_Failure() {
	// Пытаемся добавить участника в команду
	err := suite.repo.AttachParticipantToCrew(uuid.New(), uuid.New(), 1)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, repository_errors.InsertError.Error()) // Проверяем, что ошибка соответствует ожиданиям
}

// TestCrewRepositoryDetachParticipantFromCrew_Success тестирует успешное удаление участника из команды
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDetachParticipantFromCrew_Success() {
	crew := builders.NewCrewBuilder().
		WithSailNum(1).
		Build()

	// Сохраняем запись о команде
	crewRecord, err := suite.repo.Create(crew)
	require.NoError(suite.T(), err)

	// Пытаемся удалить участника из команды
	err = suite.repo.DetachParticipantFromCrew(crewRecord.ID, uuid.New())
	require.NoError(suite.T(), err)
}

// TestCrewRepositoryDetachParticipantFromCrew_Failure тестирует ошибку при удалении участника из команды
func (suite *CrewRepositoryTestSuite) TestCrewRepositoryDetachParticipantFromCrew_Failure() {
	// Пытаемся удалить участника из команды
	err := suite.repo.DetachParticipantFromCrew(uuid.New(), uuid.New())
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, repository_errors.DoesNotExist.Error()) // Проверяем, что ошибка соответствует ожиданиям
}
