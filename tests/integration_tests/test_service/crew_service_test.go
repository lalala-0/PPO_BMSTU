package test_services

import (
	"PPO_BMSTU/internal/models"
	mongo_rep "PPO_BMSTU/internal/repository/mongo"
	postgres_rep "PPO_BMSTU/internal/repository/postgres"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services"
	"PPO_BMSTU/internal/services/service_interfaces"
	"PPO_BMSTU/tests/integration_tests/db_init"
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"runtime"
	"testing"
)

type crewServiceTestSuite struct {
	suite.Suite
	postgresClient *sqlx.DB
	mongoClient    *mongo.Database
	postgresC      testcontainers.Container
	mongoC         testcontainers.Container
	repo           repository_interfaces.ICrewRepository
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

var testCrewServiceAddNewCrew = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		ratingID uuid.UUID
		sailNum  int
		class    int
	}
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "create crew success test",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			models.Laser,
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, crew)
		},
	},
	{
		testName: "invalid input",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    int
		}{
			uuid.New(),
			uuid.New(),
			-89,
			90,
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, fmt.Errorf("SERVICE: Create method failed"), err)
		},
	},
	{
		testName: "crew creation error",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			models.Laser,
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, fmt.Errorf("SERVICE: Create method failed"), err)
		},
	},
}

func (suite *crewServiceTestSuite) TestCrewService_CreateCrew() {
	for _, tt := range testCrewServiceAddNewCrew {
		suite.T().Run(tt.testName, func(t *testing.T) {
			crew, err := suite.service.AddNewCrew(tt.inputData.id, tt.inputData.ratingID, tt.inputData.class, tt.inputData.sailNum)
			tt.checkOutput(t, crew, err)
		})
	}
}

//
//var testCrewServiceDelete = []struct {
//	testName  string
//	inputData struct {
//		crewID uuid.UUID
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, err error)
//}{
//	{
//		testName: "delete crew success test",
//		inputData: struct {
//			crewID uuid.UUID
//		}{
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(builders.CrewMother.Default(), nil)
//			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "crew not found",
//		inputData: struct {
//			crewID uuid.UUID
//		}{
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
//		},
//		checkOutput: func(t *testing.T, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: GetCrewDataByID method failed"), err)
//		},
//	},
//	{
//		testName: "delete crew error",
//		inputData: struct {
//			crewID uuid.UUID
//		}{
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(builders.CrewMother.Default(), nil)
//			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DeleteError)
//		},
//		checkOutput: func(t *testing.T, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: Delete method failed"), err)
//		},
//	},
//}
//
//func TestCrewServiceDeleteCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceDelete {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.DeleteCrewByID(tt.inputData.crewID)
//			tt.checkOutput(t, err)
//		})
//	}
//}
//
//var testCrewServiceUpdateCrewByID = []struct {
//	testName  string
//	inputData struct {
//		id       uuid.UUID
//		ratingID uuid.UUID
//		sailNum  int
//		class    int
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			id       uuid.UUID
//			ratingID uuid.UUID
//			sailNum  int
//			class    int
//		}{
//			uuid.New(),
//			uuid.New(),
//			89,
//			models.Cadet,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(
//				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 189, models.LaserRadial), nil)
//			fields.crewRepoMock.EXPECT().Update(gomock.Any()).Return(
//				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 89, models.LaserRadial), nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "crew not found",
//		inputData: struct {
//			id       uuid.UUID
//			ratingID uuid.UUID
//			sailNum  int
//			class    int
//		}{
//			uuid.New(),
//			uuid.New(),
//			89,
//			models.Laser,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: GetCrewByID method failed"), err)
//		},
//	},
//	{
//		testName: "invalid input",
//		inputData: struct {
//			id       uuid.UUID
//			ratingID uuid.UUID
//			sailNum  int
//			class    int
//		}{
//			uuid.New(),
//			uuid.New(),
//			-89,
//			90,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(
//				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 189, models.LaserRadial),
//				nil)
//
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
//		},
//	},
//}
//
//func TestCrewServiceUpdateCrewByID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceUpdateCrewByID {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			crew, err := crewService.UpdateCrewByID(tt.inputData.id, tt.inputData.ratingID, tt.inputData.class, tt.inputData.sailNum)
//			tt.checkOutput(t, crew, err)
//		})
//	}
//}
//
//var testCrewServiceAttachParticipantToCrew = []struct {
//	testName  string
//	inputData struct {
//		participantID uuid.UUID
//		crewID        uuid.UUID
//		helmsman      int
//		active        int
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "attach participant to crew error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, repository_errors.UpdateError, err)
//		},
//	},
//	{
//		testName: "input data error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			-1,
//			-1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//}
//
//func TestCrewServiceAttachParticipantToCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceAttachParticipantToCrew {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.AttachParticipantToCrew(tt.inputData.participantID, tt.inputData.crewID, tt.inputData.helmsman)
//			tt.checkOutput(t, nil, err)
//		})
//	}
//}
//
//var testCrewServiceDetachParticipantFromCrew = []struct {
//	testName  string
//	inputData struct {
//		crewID        uuid.UUID
//		participantID uuid.UUID
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			crewID        uuid.UUID
//			participantID uuid.UUID
//		}{
//			uuid.New(),
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "detach participant from crew error",
//		inputData: struct {
//			crewID        uuid.UUID
//			participantID uuid.UUID
//		}{
//			uuid.New(),
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, repository_errors.UpdateError, err)
//		},
//	},
//}
//
//func TestCrewServiceDetachParticipantFromCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceDetachParticipantFromCrew {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.DetachParticipantFromCrew(tt.inputData.participantID, tt.inputData.crewID)
//			tt.checkOutput(t, nil, err)
//		})
//	}
//}
//
//var testCrewServiceReplaceParticipantStatusInCrew = []struct {
//	testName  string
//	inputData struct {
//		participantID uuid.UUID
//		crewID        uuid.UUID
//		helmsman      int
//		active        int
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//			//fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//
//			//fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "replace participant status in crew error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, repository_errors.UpdateError, err)
//		},
//	},
//	{
//		testName: "input data error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			-1,
//		},
//		prepare: func(fields *crewServiceFields) {
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
//		},
//	},
//}
//
//func TestCrewServiceReplaceParticipantStatusInCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceReplaceParticipantStatusInCrew {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.ReplaceParticipantStatusInCrew(tt.inputData.participantID, tt.inputData.crewID, tt.inputData.helmsman, tt.inputData.active)
//			tt.checkOutput(t, nil, err)
//		})
//	}
//}
//
//var testCrewServiceGetCrewDataByID = []struct {
//	testName    string
//	inputData   struct{ crewID uuid.UUID }
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName:  "get crew by id success test",
//		inputData: struct{ crewID uuid.UUID }{uuid.New()},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(builders.CrewMother.Default(), nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//			assert.NotNil(t, crew)
//		},
//	},
//	{
//		testName:  "crew not found",
//		inputData: struct{ crewID uuid.UUID }{uuid.New()},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Nil(t, crew)
//			assert.Equal(t, repository_errors.DoesNotExist, err)
//		},
//	},
//	{
//		testName:  "get crew by id error",
//		inputData: struct{ crewID uuid.UUID }{uuid.New()},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Nil(t, crew)
//			assert.Equal(t, repository_errors.SelectError, err)
//		},
//	},
//}
//
//func TestCrewServiceGetCrewDataByID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceGetCrewDataByID {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			crew, err := crewService.GetCrewDataByID(tt.inputData.crewID)
//			tt.checkOutput(t, crew, err)
//		})
//	}
//}
