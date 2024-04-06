package test_services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	services "PPO_BMSTU/internal/services"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	mock_repository_interfaces "PPO_BMSTU/tests/repository_mocks"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

type crewServiceFields struct {
	crewRepoMock *mock_repository_interfaces.MockICrewRepository
	logger         *log.Logger
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
		logger:         logger,
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
		class    string
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
			class    string
		}{
			uuid.New(),
			uuid.New(),
			89,
			"laser",
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
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
			class    string
		}{
			uuid.New(),
			uuid.New(),
			-89,
			"",
		},
		prepare: func(fields *crewServiceFields) {},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, service_errors.InvalidClass, err)
		},
	},
	{
		testName: "crew creation error",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    string
		}{
			uuid.New(),
			uuid.New(),
			89,
			"laser",
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().Create(gomock.Any()).Return(nil, repository_errors.InsertError)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, repository_errors.InsertError, err)
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
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(&models.Crew{ID: uuid.UUID{}}, nil)
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
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
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
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(&models.Crew{ID: uuid.UUID{}}, nil)
			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DeleteError)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DeleteError, err)
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
		class    string
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
			class    string
		}{
			uuid.New(),
			uuid.New(),
			89,
			"laser",
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(&models.Crew{
				uuid.New(),
				uuid.New(),
				189,
				"laser radial",
			}, nil)
			fields.crewRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Crew{
				uuid.New(),
				uuid.New(),
				89,
				"laser",
			}, nil)
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
			class    string
		}{
			uuid.New(),
			uuid.New(),
			89,
			"laser",
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName: "invalid input",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    string
		}{
			uuid.New(),
			uuid.New(),
			-89,
			"",
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(&models.Crew{
				uuid.New(),
				uuid.New(),
				189,
				"laser radial",
			}, nil)

		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, service_errors.InvalidClass, err)
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


AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error
DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error
ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error

var testCrewServiceAttachParticipantToCrew = []struct {
	testName  string
	inputData struct {
		participantID  uuid.UUID
		crewID uuid.UUID
		helmsman int
		active int
	}
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			participantID  uuid.UUID
			crewID uuid.UUID
			helmsman int
			active int
		}{
			uuid.New(),
			uuid.New(),
			1,
			1,

		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Crew{}, nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "attach worker to crew error",
		inputData: struct {
			participantID  uuid.UUID
			crewID uuid.UUID
			helmsman int
			active int
		}{
			uuid.New(),
			uuid.New(),
			1,
			1,
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
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
			err := crewService.AttachParticipantToCrew(tt.inputData.crewID, tt.inputData.judgeID)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testCrewServiceDetachParticipantFromCrew = []struct {
	testName  string
	inputData struct {
		crewID uuid.UUID
		judgeID  uuid.UUID
	}
	prepare     func(fields *crewServiceFields)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			crewID uuid.UUID
			judgeID  uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(&models.Crew{}, nil)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "detach judge from crew error",
		inputData: struct {
			crewID uuid.UUID
			judgeID  uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *crewServiceFields) {
			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(nil, repository_errors.UpdateError)
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
			err := crewService.DetachParticipantFromCrew(tt.inputData.crewID, tt.inputData.judgeID)
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
			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
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

