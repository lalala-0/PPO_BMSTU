package crew_repository_tests

import (
	mongo_rep "PPO_BMSTU/internal/repository/mongo"
	postgres_rep "PPO_BMSTU/internal/repository/postgres"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/tests/integration_tests/db_init"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
	repo           repository_interfaces.ICrewRepository
	initializer    db_init.TestRepositoryInitializer
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
		suite.postgresClient, err = db_init.ConnectTestDatabasePostgres()
		require.NoError(suite.T(), err)

		// тестируемый репозиторий
		suite.repo = postgres_rep.NewCrewRepository(suite.postgresClient)

		// инициализатор элементов бд
		suite.initializer = db_init.NewPostgresRepository(suite.postgresClient)
		suite.initializer.ClearAll()
	}

	if dbType == "mongo" { // Тестируем MongoDB, если указано
		suite.mongoClient, err = db_init.ConnectTestDatabaseMongo()
		require.NoError(suite.T(), err)

		// тестируемый репозиторий
		suite.repo = mongo_rep.NewCrewRepository(suite.mongoClient)

		// инициализатор элементов бд
		suite.initializer = db_init.NewMongoRepository(suite.mongoClient)
		suite.initializer.ClearAll()
	}

}

// TearDownSuite выполняется один раз после завершения тестов
func (suite *CrewRepositoryTestSuite) TearDownSuite() {
	suite.initializer.ClearAll()
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
