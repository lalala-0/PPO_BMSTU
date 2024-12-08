package test_services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	services "PPO_BMSTU/internal/services"
	_ "PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	mock_repository_interfaces "PPO_BMSTU/tests/repository_mocks"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
	"time"
)

type raceServiceFields struct {
	raceRepoMock          *mock_repository_interfaces.MockIRaceRepository
	crewResInRaceRepoMock *mock_repository_interfaces.MockICrewResInRaceRepository
	crewRepoMock          *mock_repository_interfaces.MockICrewRepository
	logger                *log.Logger
}

func initRaceServiceFields(ctrl *gomock.Controller) *raceServiceFields {
	raceRepoMock := mock_repository_interfaces.NewMockIRaceRepository(ctrl)
	crewResInRaceRepoMock := mock_repository_interfaces.NewMockICrewResInRaceRepository(ctrl)
	crewRepoMock := mock_repository_interfaces.NewMockICrewRepository(ctrl)

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f)

	return &raceServiceFields{
		raceRepoMock:          raceRepoMock,
		crewResInRaceRepoMock: crewResInRaceRepoMock,
		crewRepoMock:          crewRepoMock,
		logger:                logger,
	}
}

func initRaceService(fields *raceServiceFields) service_interfaces.IRaceService {
	return services.NewRaceService(fields.raceRepoMock, fields.crewRepoMock, fields.crewResInRaceRepoMock, fields.logger)
}

var testRaceServiceAddNewRace = []struct {
	testName  string
	inputData struct {
		raceID   uuid.UUID
		ratingID uuid.UUID
		number   int
		date     time.Time
		class    int
	}
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, race *models.Race, err error)
}{
	{
		testName: "create race success test",
		inputData: struct {
			raceID   uuid.UUID
			ratingID uuid.UUID
			number   int
			date     time.Time
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
			models.Laser,
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, race)
		},
	},
	{
		testName: "invalid input",
		inputData: struct {
			raceID   uuid.UUID
			ratingID uuid.UUID
			number   int
			date     time.Time
			class    int
		}{
			uuid.New(),
			uuid.New(),
			-89,
			time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
			models.Laser,
		},
		prepare: func(fields *raceServiceFields) {},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.Error(t, err)
			assert.Nil(t, race)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "race creation error",
		inputData: struct {
			raceID   uuid.UUID
			ratingID uuid.UUID
			number   int
			date     time.Time
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
			models.Laser,
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().Create(gomock.Any()).Return(nil, repository_errors.InsertError)
		},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.Error(t, err)
			assert.Nil(t, race)
			assert.Equal(t, fmt.Errorf("DB ERROR: Insert operation was not successful"), err)
		},
	},
}

func TestRaceService_CreateRace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRaceServiceFields(ctrl)
	raceService := initRaceService(fields)

	for _, tt := range testRaceServiceAddNewRace {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			race, err := raceService.AddNewRace(tt.inputData.raceID, tt.inputData.ratingID, tt.inputData.number, tt.inputData.date, tt.inputData.class)
			tt.checkOutput(t, race, err)
		})
	}
}

var testRaceServiceDelete = []struct {
	testName  string
	inputData struct {
		raceID uuid.UUID
	}
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, err error)
}{
	{
		testName: "delete race success test",
		inputData: struct {
			raceID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().Delete(gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "race not found",
		inputData: struct {
			raceID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("GET operation has failed. Such row does not exist"), err)
		},
	},
	{
		testName: "delete race error",
		inputData: struct {
			raceID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DeleteError)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("DB ERROR: Delete operation was not successful"), err)
		},
	},
}

func TestRaceServiceDeleteRace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRaceServiceFields(ctrl)
	raceService := initRaceService(fields)

	for _, tt := range testRaceServiceDelete {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := raceService.DeleteRaceByID(tt.inputData.raceID)
			tt.checkOutput(t, err)
		})
	}
}

var testRaceServiceUpdateRace = []struct {
	testName  string
	inputData struct {
		raceID   uuid.UUID
		ratingID uuid.UUID
		number   int
		date     time.Time
		class    int
	}
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, race *models.Race, err error)
}{
	{
		testName: "update race success test",
		inputData: struct {
			raceID   uuid.UUID
			ratingID uuid.UUID
			number   int
			date     time.Time
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
			models.Laser,
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.UUID{}}, nil)
			fields.raceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, race)
		},
	},
	{
		testName: "invalid input",
		inputData: struct {
			raceID   uuid.UUID
			ratingID uuid.UUID
			number   int
			date     time.Time
			class    int
		}{
			uuid.New(),
			uuid.New(),
			-89,
			time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
			models.Laser,
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.UUID{}}, nil)
		},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "race update error",
		inputData: struct {
			raceID   uuid.UUID
			ratingID uuid.UUID
			number   int
			date     time.Time
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
			models.Laser,
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestRaceService_UpdateRace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRaceServiceFields(ctrl)
	raceService := initRaceService(fields)

	for _, tt := range testRaceServiceUpdateRace {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			race, err := raceService.UpdateRaceByID(tt.inputData.raceID, tt.inputData.ratingID, tt.inputData.number, tt.inputData.date, tt.inputData.class)
			tt.checkOutput(t, race, err)
		})
	}
}

var testRaceServiceGetRaceDataByID = []struct {
	testName    string
	inputData   struct{ raceID uuid.UUID }
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, race *models.Race, err error)
}{
	{
		testName:  "get race by id success test",
		inputData: struct{ raceID uuid.UUID }{uuid.New()},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, race)
		},
	},
	{
		testName:  "get race by id error",
		inputData: struct{ raceID uuid.UUID }{uuid.New()},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, race *models.Race, err error) {
			assert.Error(t, err)
			assert.Nil(t, race)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestRaceServiceGetRaceDataByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRaceServiceFields(ctrl)
	raceService := initRaceService(fields)

	for _, tt := range testRaceServiceGetRaceDataByID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			race, err := raceService.GetRaceDataByID(tt.inputData.raceID)
			tt.checkOutput(t, race, err)
		})
	}
}

var testRaceServiceGetRacesDataByRatingID = []struct {
	testName  string
	inputData struct{ ratingID uuid.UUID }
	prepare   func(fields *raceServiceFields)
	checkFunc func(t *testing.T, races []models.Race, err error)
}{
	{
		testName:  "get races by rating id success test",
		inputData: struct{ ratingID uuid.UUID }{uuid.New()},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRacesDataByRatingID(gomock.Any()).Return([]models.Race{{ID: uuid.New()}, {ID: uuid.New()}}, nil)
		},
		checkFunc: func(t *testing.T, race []models.Race, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, race)
		},
	},
	{
		testName:  "races not found",
		inputData: struct{ ratingID uuid.UUID }{uuid.New()},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRacesDataByRatingID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkFunc: func(t *testing.T, race []models.Race, err error) {
			assert.Error(t, err)
			assert.Nil(t, race)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName:  "get races by rating id error",
		inputData: struct{ ratingID uuid.UUID }{uuid.New()},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRacesDataByRatingID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkFunc: func(t *testing.T, race []models.Race, err error) {
			assert.Error(t, err)
			assert.Nil(t, race)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestRaceServiceGetRacesDataByRatingID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRaceServiceFields(ctrl)
	raceService := initRaceService(fields)

	for _, tt := range testRaceServiceGetRacesDataByRatingID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			race, err := raceService.GetRacesDataByRatingID(tt.inputData.ratingID)
			tt.checkFunc(t, race, err)
		})
	}
}

var testRaceMakeStartProcedureSuccess = []struct {
	testName  string
	inputData struct {
		raceID              uuid.UUID
		falseStartYachtList map[int]int
	}
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, crewResInRace *models.CrewResInRace, err error)
	checkFunc   func(t *testing.T, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			raceID              uuid.UUID
			falseStartYachtList map[int]int
		}{
			raceID:              uuid.New(),
			falseStartYachtList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewsDataByRatingID(gomock.Any()).Return([]models.Crew{{ID: uuid.New()}, {ID: uuid.New()}}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Create(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Create(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.NoError(t, err)
		},
	},
}

var testRaceMakeStartProcedureFail = []struct {
	testName  string
	inputData struct {
		raceID              uuid.UUID
		falseStartYachtList map[int]int
	}
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, crewResInRace *models.CrewResInRace, err error)
	checkFunc   func(t *testing.T, err error)
}{
	{
		testName: "GetRaceDataByID error",
		inputData: struct {
			raceID              uuid.UUID
			falseStartYachtList map[int]int
		}{
			raceID:              uuid.New(),
			falseStartYachtList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, repository_errors.SelectError)

		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
	{
		testName: "GetCrewsDataByRatingID error",
		inputData: struct {
			raceID              uuid.UUID
			falseStartYachtList map[int]int
		}{
			raceID:              uuid.New(),
			falseStartYachtList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewsDataByRatingID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
	{
		testName: "GetCrewDataBySailNumAndRatingID error",
		inputData: struct {
			raceID              uuid.UUID
			falseStartYachtList map[int]int
		}{
			raceID:              uuid.New(),
			falseStartYachtList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewsDataByRatingID(gomock.Any()).Return([]models.Crew{{ID: uuid.New(), SailNum: 1}, {ID: uuid.New(), SailNum: 345}, {ID: uuid.New(), SailNum: 654}}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Create(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Create(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Create(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, repository_errors.SelectError)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)

			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestRaceServiceMakeStartProcedure(t *testing.T) {
	for _, tt := range testRaceMakeStartProcedureSuccess {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initRaceServiceFields(ctrl)
			raceService := initRaceService(fields)

			tt.prepare(fields)
			err := raceService.MakeStartProcedure(tt.inputData.raceID, tt.inputData.falseStartYachtList)
			tt.checkOutput(t, nil, err)
		})
	}

	for _, tt := range testRaceMakeStartProcedureFail {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initRaceServiceFields(ctrl)
			raceService := initRaceService(fields)

			tt.prepare(fields)

			err := raceService.MakeStartProcedure(tt.inputData.raceID, tt.inputData.falseStartYachtList)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testRaceMakeFinishProcedureSuccess = []struct {
	testName  string
	inputData struct {
		raceID           uuid.UUID
		finishersList    map[int]int
		nonFinishersList map[int]int
	}
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, crewResInRace *models.CrewResInRace, err error)
	checkFunc   func(t *testing.T, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			raceID           uuid.UUID
			finishersList    map[int]int
			nonFinishersList map[int]int
		}{
			raceID:           uuid.New(),
			finishersList:    map[int]int{10: 87, 344: 65, 659: 1},
			nonFinishersList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)

			fields.crewResInRaceRepoMock.EXPECT().GetAllCrewResInRace(gomock.Any()).Return([]models.CrewResInRace{{RaceID: uuid.New()}}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.NoError(t, err)
		},
	},
}

var testRaceMakeFinishProcedureFail = []struct {
	testName  string
	inputData struct {
		raceID           uuid.UUID
		finishersList    map[int]int
		nonFinishersList map[int]int
	}
	prepare     func(fields *raceServiceFields)
	checkOutput func(t *testing.T, crewResInRace *models.CrewResInRace, err error)
	checkFunc   func(t *testing.T, err error)
}{
	{
		testName: "GetRaceDataByID error",
		inputData: struct {
			raceID           uuid.UUID
			finishersList    map[int]int
			nonFinishersList map[int]int
		}{
			raceID:           uuid.New(),
			finishersList:    map[int]int{10: 87, 344: 65, 659: 1},
			nonFinishersList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, repository_errors.SelectError)

		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
	{
		testName: "GetCrewDataBySailNumAndRatingID error",
		inputData: struct {
			raceID           uuid.UUID
			finishersList    map[int]int
			nonFinishersList map[int]int
		}{
			raceID:           uuid.New(),
			finishersList:    map[int]int{10: 87, 344: 65, 659: 1},
			nonFinishersList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(nil, repository_errors.SelectError)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)

			fields.crewResInRaceRepoMock.EXPECT().GetAllCrewResInRace(gomock.Any()).Return([]models.CrewResInRace{{RaceID: uuid.New()}}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
	{
		testName: "GetCrewResByRaceIDAndCrewID error",
		inputData: struct {
			raceID           uuid.UUID
			finishersList    map[int]int
			nonFinishersList map[int]int
		}{
			raceID:           uuid.New(),
			finishersList:    map[int]int{10: 87, 344: 65, 659: 1},
			nonFinishersList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(nil, repository_errors.SelectError)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)

			fields.crewResInRaceRepoMock.EXPECT().GetAllCrewResInRace(gomock.Any()).Return([]models.CrewResInRace{{RaceID: uuid.New()}}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
	{
		testName: "Update error",
		inputData: struct {
			raceID           uuid.UUID
			finishersList    map[int]int
			nonFinishersList map[int]int
		}{
			raceID:           uuid.New(),
			finishersList:    map[int]int{10: 87, 344: 65, 659: 1},
			nonFinishersList: map[int]int{1: models.DNF, 345: models.DNC, 654: models.DNS},
		},
		prepare: func(fields *raceServiceFields) {
			fields.raceRepoMock.EXPECT().GetRaceDataByID(gomock.Any()).Return(&models.Race{ID: uuid.New()}, nil)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, repository_errors.UpdateError)

			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(nil, repository_errors.UpdateError)

			fields.crewResInRaceRepoMock.EXPECT().GetAllCrewResInRace(gomock.Any()).Return([]models.CrewResInRace{{RaceID: uuid.New()}}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{RaceID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, crewResInRace *models.CrewResInRace, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
}

func TestRaceServiceMakeFinishProcedure(t *testing.T) {
	for _, tt := range testRaceMakeFinishProcedureSuccess {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initRaceServiceFields(ctrl)
			raceService := initRaceService(fields)

			tt.prepare(fields)
			err := raceService.MakeFinishProcedure(tt.inputData.raceID, tt.inputData.finishersList, tt.inputData.nonFinishersList)
			tt.checkOutput(t, nil, err)
		})
	}

	for _, tt := range testRaceMakeFinishProcedureFail {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initRaceServiceFields(ctrl)
			raceService := initRaceService(fields)

			tt.prepare(fields)
			err := raceService.MakeFinishProcedure(tt.inputData.raceID, tt.inputData.finishersList, tt.inputData.nonFinishersList)
			tt.checkOutput(t, nil, err)
		})
	}
}
