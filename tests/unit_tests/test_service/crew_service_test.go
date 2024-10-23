package test_services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	services "PPO_BMSTU/internal/services"
	"PPO_BMSTU/internal/services/service_interfaces"
	mock_repository_interfaces "PPO_BMSTU/tests/repository_mocks"
	"PPO_BMSTU/tests/unit_tests/builders"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

type crewServiceFields struct {
	crewRepoMock *mock_repository_interfaces.MockICrewRepository
	logger       *log.Logger
}

func initCrewServiceFields(ctrl *gomock.Controller) *crewServiceFields {
	crewRepoMock := mock_repository_interfaces.NewMockICrewRepository(ctrl)

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f)

	return &crewServiceFields{
		crewRepoMock: crewRepoMock,
		logger:       logger,
	}
}

func initCrewService(fields *crewServiceFields) service_interfaces.ICrewService {
	return services.NewCrewService(fields.crewRepoMock, fields.logger)
}

var testCrewServiceAddNewCrew = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		ratingID uuid.UUID
		sailNum  int
		class    int
	}
	prepare     func(fields *crewServiceFields)
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
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().Create(gomock.Any()).Return(builders.CrewMother.Default(), nil)
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
		prepare: func(fields *crewServiceFields) {},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
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
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().Create(gomock.Any()).Return(nil, repository_errors.InsertError)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, fmt.Errorf("SERVICE: Create method failed"), err)
		},
	},
}

func TestCrewService_CreateCrew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initCrewServiceFields(ctrl)
	crewService := initCrewService(fields)

	for _, tt := range testCrewServiceAddNewCrew {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			crew, err := crewService.AddNewCrew(tt.inputData.id, tt.inputData.ratingID, tt.inputData.class, tt.inputData.sailNum)
			tt.checkOutput(t, crew, err)
		})
	}
}

var testCrewServiceDelete = []struct {
	testName  string
	inputData struct {
		crewID uuid.UUID
	}
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, err error)
}{
	{
		testName: "delete crew success test",
		inputData: struct {
			crewID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "crew not found",
		inputData: struct {
			crewID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName: "delete crew error",
		inputData: struct {
			crewID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
}

func TestCrewServiceDeleteCrew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initCrewServiceFields(ctrl)
	crewService := initCrewService(fields)

	for _, tt := range testCrewServiceDelete {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := crewService.DeleteCrewByID(tt.inputData.crewID)
			tt.checkOutput(t, err)
		})
	}
}

var testCrewServiceUpdateCrewByID = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		ratingID uuid.UUID
		sailNum  int
		class    int
	}
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			models.Cadet,
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(
				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 189, models.LaserRadial), nil)
			fields.crewRepoMock.EXPECT().Update(gomock.Any()).Return(
				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 89, models.LaserRadial), nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "crew not found",
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
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: GetCrewByID method failed"), err)
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
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(
				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 189, models.LaserRadial),
				nil)

		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
		},
	},
}

func TestCrewServiceUpdateCrewByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initCrewServiceFields(ctrl)
	crewService := initCrewService(fields)

	for _, tt := range testCrewServiceUpdateCrewByID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			crew, err := crewService.UpdateCrewByID(tt.inputData.id, tt.inputData.ratingID, tt.inputData.class, tt.inputData.sailNum)
			tt.checkOutput(t, crew, err)
		})
	}
}

var testCrewServiceAttachParticipantToCrew = []struct {
	testName  string
	inputData struct {
		participantID uuid.UUID
		crewID        uuid.UUID
		helmsman      int
		active        int
	}
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			participantID uuid.UUID
			crewID        uuid.UUID
			helmsman      int
			active        int
		}{
			uuid.New(),
			uuid.New(),
			1,
			1,
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "attach participant to crew error",
		inputData: struct {
			participantID uuid.UUID
			crewID        uuid.UUID
			helmsman      int
			active        int
		}{
			uuid.New(),
			uuid.New(),
			1,
			1,
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
	{
		testName: "input data error",
		inputData: struct {
			participantID uuid.UUID
			crewID        uuid.UUID
			helmsman      int
			active        int
		}{
			uuid.New(),
			uuid.New(),
			-1,
			-1,
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
		},
	},
}

func TestCrewServiceAttachParticipantToCrew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initCrewServiceFields(ctrl)
	crewService := initCrewService(fields)

	for _, tt := range testCrewServiceAttachParticipantToCrew {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := crewService.AttachParticipantToCrew(tt.inputData.participantID, tt.inputData.crewID, tt.inputData.helmsman)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testCrewServiceDetachParticipantFromCrew = []struct {
	testName  string
	inputData struct {
		crewID        uuid.UUID
		participantID uuid.UUID
	}
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			crewID        uuid.UUID
			participantID uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "detach participant from crew error",
		inputData: struct {
			crewID        uuid.UUID
			participantID uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
}

func TestCrewServiceDetachParticipantFromCrew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initCrewServiceFields(ctrl)
	crewService := initCrewService(fields)

	for _, tt := range testCrewServiceDetachParticipantFromCrew {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := crewService.DetachParticipantFromCrew(tt.inputData.participantID, tt.inputData.crewID)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testCrewServiceReplaceParticipantStatusInCrew = []struct {
	testName  string
	inputData struct {
		participantID uuid.UUID
		crewID        uuid.UUID
		helmsman      int
		active        int
	}
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			participantID uuid.UUID
			crewID        uuid.UUID
			helmsman      int
			active        int
		}{
			uuid.New(),
			uuid.New(),
			1,
			1,
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			//fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			//fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "replace participant status in crew error",
		inputData: struct {
			participantID uuid.UUID
			crewID        uuid.UUID
			helmsman      int
			active        int
		}{
			uuid.New(),
			uuid.New(),
			1,
			1,
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
	{
		testName: "input data error",
		inputData: struct {
			participantID uuid.UUID
			crewID        uuid.UUID
			helmsman      int
			active        int
		}{
			uuid.New(),
			uuid.New(),
			1,
			-1,
		},
		prepare: func(fields *crewServiceFields) {
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
		},
	},
}

func TestCrewServiceReplaceParticipantStatusInCrew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initCrewServiceFields(ctrl)
	crewService := initCrewService(fields)

	for _, tt := range testCrewServiceReplaceParticipantStatusInCrew {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := crewService.ReplaceParticipantStatusInCrew(tt.inputData.participantID, tt.inputData.crewID, tt.inputData.helmsman, tt.inputData.active)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testCrewServiceGetCrewDataByID = []struct {
	testName    string
	inputData   struct{ crewID uuid.UUID }
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName:  "get crew by id success test",
		inputData: struct{ crewID uuid.UUID }{uuid.New()},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(builders.CrewMother.Default(), nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, crew)
		},
	},
	{
		testName:  "crew not found",
		inputData: struct{ crewID uuid.UUID }{uuid.New()},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName:  "get crew by id error",
		inputData: struct{ crewID uuid.UUID }{uuid.New()},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestCrewServiceGetCrewDataByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initCrewServiceFields(ctrl)
	crewService := initCrewService(fields)

	for _, tt := range testCrewServiceGetCrewDataByID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			crew, err := crewService.GetCrewDataByID(tt.inputData.crewID)
			tt.checkOutput(t, crew, err)
		})
	}
}
