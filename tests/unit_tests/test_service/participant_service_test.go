package test_services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	services "PPO_BMSTU/internal/services"
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

type participantServiceFields struct {
	participantRepoMock *mock_repository_interfaces.MockIParticipantRepository
	logger              *log.Logger
}

func initParticipantServiceFields(ctrl *gomock.Controller) *participantServiceFields {
	participantRepoMock := mock_repository_interfaces.NewMockIParticipantRepository(ctrl)
	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f)

	return &participantServiceFields{
		participantRepoMock: participantRepoMock,
		logger:              logger,
	}
}

func initParticipantService(fields *participantServiceFields) service_interfaces.IParticipantService {
	return services.NewParticipantService(fields.participantRepoMock, fields.logger)
}

var testParticipantGetByIDSuccess = []struct {
	testName  string
	inputData struct {
		id uuid.UUID
	}
	prepare     func(fields *participantServiceFields)
	checkOutput func(t *testing.T, participant *models.Participant, err error)
}{
	{
		testName: "basic get by id",
		inputData: struct {
			id uuid.UUID
		}{
			id: uuid.New(),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(&models.Participant{}, nil)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, participant)
		},
	},
}

var testParticipantGetByIDFail = []struct {
	testName  string
	inputData struct {
		id uuid.UUID
	}
	prepare     func(fields *participantServiceFields)
	checkOutput func(t *testing.T, participant *models.Participant, err error)
}{
	{
		testName: "participant not found",
		inputData: struct {
			id uuid.UUID
		}{
			id: uuid.New(),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Nil(t, participant)
		},
	},
}

func TestParticipantServiceGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initParticipantServiceFields(ctrl)
	service := initParticipantService(fields)

	for _, test := range testParticipantGetByIDSuccess {
		t.Run(test.testName, func(t *testing.T) {
			test.prepare(fields)
			participant, err := service.GetParticipantDataByID(test.inputData.id)
			test.checkOutput(t, participant, err)
		})
	}

	for _, test := range testParticipantGetByIDFail {
		t.Run(test.testName, func(t *testing.T) {
			test.prepare(fields)
			participant, err := service.GetParticipantDataByID(test.inputData.id)
			test.checkOutput(t, participant, err)
		})
	}
}

var testParticipantUpdateFail = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		fio      string
		category int
		gender   int
		birthday time.Time
		coach    string
	}
	prepare     func(fields *participantServiceFields)
	checkOutput func(t *testing.T, participant *models.Participant, err error)
}{
	{
		testName: "participant not found",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "Test",
			gender:   models.Male,
			category: models.MasterRussia,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Nil(t, participant)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
	{
		testName: "invalid fio",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "",
			gender:   models.Male,
			category: models.Junior2category,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(&models.Participant{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid category",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "test",
			gender:   models.Female,
			category: 90,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(&models.Participant{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid coach",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "test",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(&models.Participant{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid birthdate",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "jet",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2023, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(&models.Participant{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "update failed",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "Test",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(&models.Participant{ID: uuid.New()}, nil)
			fields.participantRepoMock.EXPECT().Update(gomock.Any()).Return(nil, repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
}

var testParticipantUpdateSuccess = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		fio      string
		category int
		gender   int
		birthday time.Time
		coach    string
	}
	prepare     func(fields *participantServiceFields)
	checkOutput func(t *testing.T, participant *models.Participant, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "Test",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2002, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().GetParticipantDataByID(gomock.Any()).Return(&models.Participant{ID: uuid.New()}, nil)
			fields.participantRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Participant{
				FIO:      "Test",
				Gender:   models.Female,
				Category: models.Sport1category,
				Coach:    "Test",
				Birthday: time.Date(2002, time.November, 10, 23, 0, 0, 0, time.UTC),
			}, nil)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, participant)
			assert.Equal(t, "Test", participant.FIO)
			assert.Equal(t, models.Female, participant.Gender)
			assert.Equal(t, models.Sport1category, participant.Category)
			assert.Equal(t, "Test", participant.Coach)
			assert.Equal(t, time.Date(2002, time.November, 10, 23, 0, 0, 0, time.UTC), participant.Birthday)
		},
	},
}

func TestParticipantServiceUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initParticipantServiceFields(ctrl)
	service := initParticipantService(fields)

	for _, test := range testParticipantUpdateSuccess {
		t.Run(test.testName, func(t *testing.T) {
			test.prepare(fields)
			participant, err := service.UpdateParticipantByID(test.inputData.id, test.inputData.fio, test.inputData.category, test.inputData.birthday, test.inputData.coach)
			test.checkOutput(t, participant, err)

		})
	}

	for _, test := range testParticipantUpdateFail {
		t.Run(test.testName, func(t *testing.T) {
			test.prepare(fields)
			participant, err := service.UpdateParticipantByID(test.inputData.id, test.inputData.fio, test.inputData.category, test.inputData.birthday, test.inputData.coach)
			test.checkOutput(t, participant, err)
		})
	}
}

var testParticipantCreateFail = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		fio      string
		category int
		gender   int
		birthday time.Time
		coach    string
	}
	prepare     func(fields *participantServiceFields)
	checkOutput func(t *testing.T, participant *models.Participant, err error)
}{
	{
		testName: "invalid fio",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Nil(t, participant)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid category",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "test",
			gender:   models.Female,
			category: 99,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Nil(t, participant)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid gender",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "test",
			gender:   0,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Nil(t, participant)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid coach",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "test",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Nil(t, participant)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid birthdate",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "test",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2023, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Nil(t, participant)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
}

var testParticipantCreateSuccess = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		fio      string
		category int
		gender   int
		birthday time.Time
		coach    string
	}
	prepare     func(fields *participantServiceFields)
	checkOutput func(t *testing.T, participant *models.Participant, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "Test",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2002, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Participant{
				FIO:      "Test",
				Gender:   models.Female,
				Category: models.Sport2category,
				Coach:    "Test",
				Birthday: time.Date(2002, time.November, 10, 23, 0, 0, 0, time.UTC),
			}, nil)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, participant)
			assert.Equal(t, "Test", participant.FIO)
			assert.Equal(t, models.Female, participant.Gender)
			assert.Equal(t, models.Sport2category, participant.Category)
			assert.Equal(t, "Test", participant.Coach)
			assert.Equal(t, time.Date(2002, time.November, 10, 23, 0, 0, 0, time.UTC), participant.Birthday)
		},
	},
	{
		testName: "create failed",
		inputData: struct {
			id       uuid.UUID
			fio      string
			category int
			gender   int
			birthday time.Time
			coach    string
		}{
			id:       uuid.New(),
			fio:      "Test",
			gender:   models.Female,
			category: models.Sport2category,
			coach:    "Test",
			birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		prepare: func(fields *participantServiceFields) {
			fields.participantRepoMock.EXPECT().Create(gomock.Any()).Return(nil, repository_errors.InsertError)
		},
		checkOutput: func(t *testing.T, participant *models.Participant, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.InsertError, err)
		},
	},
}

func TestParticipantServiceCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initParticipantServiceFields(ctrl)
	service := initParticipantService(fields)

	for _, test := range testParticipantCreateSuccess {
		t.Run(test.testName, func(t *testing.T) {
			test.prepare(fields)
			participant, err := service.AddNewParticipant(test.inputData.id, test.inputData.fio, test.inputData.category, test.inputData.gender, test.inputData.birthday, test.inputData.coach)
			test.checkOutput(t, participant, err)
		})
	}

	for _, test := range testParticipantCreateFail {
		t.Run(test.testName, func(t *testing.T) {
			test.prepare(fields)
			participant, err := service.AddNewParticipant(test.inputData.id, test.inputData.fio, test.inputData.category, test.inputData.gender, test.inputData.birthday, test.inputData.coach)
			test.checkOutput(t, participant, err)
		})
	}
}
