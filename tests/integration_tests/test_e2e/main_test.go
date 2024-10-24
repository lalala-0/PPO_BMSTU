package test_e2e

import (
	mongo_rep "PPO_BMSTU/internal/repository/mongo"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services"
	"PPO_BMSTU/internal/services/service_interfaces"
	mock_password_hash "PPO_BMSTU/tests/hasher_mocks"
	"github.com/charmbracelet/log"

	postgres_rep "PPO_BMSTU/internal/repository/postgres"
	"PPO_BMSTU/tests/integration_tests/db_init"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"runtime"
	"testing"
)

// e2eTestSuite описывает тестовый набор для Ie2e
type e2eTestSuite struct {
	suite.Suite
	postgresClient *sqlx.DB
	mongoClient    *mongo.Database
	postgresC      testcontainers.Container
	mongoC         testcontainers.Container
	initializer    db_init.TestRepositoryInitializer
	logger         *log.Logger
	hash           *mock_password_hash.MockPasswordHash

	crewService        service_interfaces.ICrewService
	judgeService       service_interfaces.IJudgeService
	participantService service_interfaces.IParticipantService
	protestService     service_interfaces.IProtestService
	raceService        service_interfaces.IRaceService
	ratingService      service_interfaces.IRatingService

	crewRepository          repository_interfaces.ICrewRepository
	crewResInRaceRepository repository_interfaces.ICrewResInRaceRepository
	judgeRepository         repository_interfaces.IJudgeRepository
	participantRepository   repository_interfaces.IParticipantRepository
	protestRepository       repository_interfaces.IProtestRepository
	raceRepository          repository_interfaces.IRaceRepository
	ratingRepository        repository_interfaces.IRatingRepository
}

// TestIe2e запускает тесты в наборе
func TestIe2e(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Тестирование PostgreSQL
	os.Setenv("DB_TYPE", "postgres")
	suite.Run(t, new(e2eTestSuite))

	//// Тестирование MongoDB
	//os.Setenv("DB_TYPE", "mongo")
	//suite.Run(t, new(e2eTestSuite))
}

func (suite *e2eTestSuite) SetupSuite() {
	dbType := os.Getenv("DB_TYPE") // Получаем значение переменной окружения DB_TYPE

	var err error
	if dbType == "postgres" || dbType == "" { // По умолчанию тестируем PostgreSQL, если не указано
		suite.postgresC, suite.postgresClient, err = db_init.SetupTestDatabasePostgres()
		require.NoError(suite.T(), err)

		// все репозитории
		suite.crewRepository = postgres_rep.NewCrewRepository(suite.postgresClient)
		suite.crewResInRaceRepository = postgres_rep.NewCrewResInRaceRepository(suite.postgresClient)
		suite.judgeRepository = postgres_rep.NewJudgeRepository(suite.postgresClient)
		suite.participantRepository = postgres_rep.NewParticipantRepository(suite.postgresClient)
		suite.protestRepository = postgres_rep.NewProtestRepository(suite.postgresClient)
		suite.raceRepository = postgres_rep.NewRaceRepository(suite.postgresClient)
		suite.ratingRepository = postgres_rep.NewRatingRepository(suite.postgresClient)

		// инициализатор элементов бд
		suite.initializer = db_init.NewPostgresRepository(suite.postgresClient)
	}

	if dbType == "mongo" { // Тестируем MongoDB, если указано
		suite.mongoC, suite.mongoClient = db_init.SetupTestDatabaseMongo()

		// Проверяем, что клиент MongoDB и контейнер не nil
		require.NotNil(suite.T(), suite.mongoClient, "Mongo client is nil")
		require.NotNil(suite.T(), suite.mongoC, "Mongo container is nil")

		// тестируемый репозиторий
		suite.crewRepository = mongo_rep.NewCrewRepository(suite.mongoClient)
		suite.crewResInRaceRepository = mongo_rep.NewCrewResInRaceRepository(suite.mongoClient)
		suite.judgeRepository = mongo_rep.NewJudgeRepository(suite.mongoClient)
		suite.participantRepository = mongo_rep.NewParticipantRepository(suite.mongoClient)
		suite.protestRepository = mongo_rep.NewProtestRepository(suite.mongoClient)
		suite.raceRepository = mongo_rep.NewRaceRepository(suite.mongoClient)
		suite.ratingRepository = mongo_rep.NewRatingRepository(suite.mongoClient)

		// инициализатор элементов бд
		suite.initializer = db_init.NewMongoRepository(suite.mongoClient)
	}

	// Инициализация лога
	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	suite.logger = log.New(f)

	// Инициализация тестируемого сервиса
	suite.crewService = services.NewCrewService(suite.crewRepository, suite.logger)
	suite.judgeService = services.NewJudgeService(suite.judgeRepository, suite.hash, suite.logger)
	suite.participantService = services.NewParticipantService(suite.participantRepository, suite.logger)
	suite.protestService = services.NewProtestService(suite.protestRepository, suite.crewResInRaceRepository, suite.crewRepository, suite.logger)
	suite.raceService = services.NewRaceService(suite.raceRepository, suite.crewRepository, suite.crewResInRaceRepository, suite.logger)
	suite.ratingService = services.NewRatingService(suite.ratingRepository, suite.judgeRepository, suite.logger)

}

// TearDownSuite выполняется один раз после завершения тестов
func (suite *e2eTestSuite) TearDownSuite() {
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
