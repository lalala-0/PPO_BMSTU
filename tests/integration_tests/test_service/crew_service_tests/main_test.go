package crew_service_tests

import (
	mongo_rep "PPO_BMSTU/internal/repository/mongo"
	postgres_rep "PPO_BMSTU/internal/repository/postgres"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services"
	"PPO_BMSTU/internal/services/service_interfaces"
	"PPO_BMSTU/tests/integration_tests/db_init"
	"context"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"runtime"
	"testing"
)

type crewServiceTestSuite struct {
	suite.Suite
	postgresClient *sqlx.DB
	mongoClient    *mongo.Database
	repo           repository_interfaces.ICrewRepository
	initializer    db_init.TestRepositoryInitializer
	service        service_interfaces.ICrewService
	logger         *log.Logger
}

// TestICrewRepository запускает тесты в наборе
func TestICrewRepository(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Тестирование PostgreSQL
	os.Setenv("DB_TYPE", "postgres")
	suite.Run(t, new(crewServiceTestSuite))

	//// Тестирование MongoDB
	//os.Setenv("DB_TYPE", "mongo")
	//suite.Run(t, new(crewServiceTestSuite))
}

func (suite *crewServiceTestSuite) SetupSuite() {
	// Создание контейнера для бд
	dbType := os.Getenv("DB_TYPE") // Получаем значение переменной окружения DB_TYPE
	var err error
	if dbType == "postgres" || dbType == "" { // По умолчанию тестируем PostgreSQL, если не указано
		suite.postgresClient, err = db_init.ConnectTestDatabasePostgres()
		require.NoError(suite.T(), err)

		suite.repo = postgres_rep.NewCrewRepository(suite.postgresClient)

		// инициализатор элементов бд
		suite.initializer = db_init.NewPostgresRepository(suite.postgresClient)
	}

	if dbType == "mongo" { // Тестируем MongoDB, если указано
		suite.mongoClient, err = db_init.ConnectTestDatabaseMongo()
		require.NoError(suite.T(), err)

		suite.repo = mongo_rep.NewCrewRepository(suite.mongoClient)

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
	suite.service = services.NewCrewService(suite.repo, suite.logger)

}

// TearDownSuite выполняется один раз после завершения тестов
func (suite *crewServiceTestSuite) TearDownSuite() {
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
