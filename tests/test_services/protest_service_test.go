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

type protestServiceFields struct {
	protestRepoMock       *mock_repository_interfaces.MockIProtestRepository
	crewResInRaceRepoMock *mock_repository_interfaces.MockICrewResInRaceRepository
	crewRepoMock          *mock_repository_interfaces.MockICrewRepository
	logger                *log.Logger
}

func initProtestServiceFields(ctrl *gomock.Controller) *protestServiceFields {
	protestRepoMock := mock_repository_interfaces.NewMockIProtestRepository(ctrl)
	crewResInRaceRepoMock := mock_repository_interfaces.NewMockICrewResInRaceRepository(ctrl)
	crewRepoMock := mock_repository_interfaces.NewMockICrewRepository(ctrl)
	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f)
	return &protestServiceFields{
		protestRepoMock:       protestRepoMock,
		crewResInRaceRepoMock: crewResInRaceRepoMock,
		crewRepoMock:          crewRepoMock,
		logger:                logger,
	}
}

func initProtestService(fields *protestServiceFields) service_interfaces.IProtestService {
	return services.NewProtestService(fields.protestRepoMock, fields.crewResInRaceRepoMock, fields.crewRepoMock, fields.logger)
}

var testProtestCreateSuccess = []struct {
	testName  string
	inputData struct {
		protestID        uuid.UUID
		raceID           uuid.UUID
		ratingID         uuid.UUID
		judgeID          uuid.UUID
		ruleNum          int
		reviewDate       time.Time
		comment          string
		protesteeSailNum int
		protestorSailNum int
		witnessesSailNum []int
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName: "basic create",
		inputData: struct {
			protestID        uuid.UUID
			raceID           uuid.UUID
			ratingID         uuid.UUID
			judgeID          uuid.UUID
			ruleNum          int
			reviewDate       time.Time
			comment          string
			protesteeSailNum int
			protestorSailNum int
			witnessesSailNum []int
		}{
			protestID:        uuid.New(),
			raceID:           uuid.New(),
			ratingID:         uuid.New(),
			ruleNum:          31,
			reviewDate:       time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			comment:          "Test",
			protesteeSailNum: 123,
			protestorSailNum: 245,
			witnessesSailNum: []int{1, 2, 3, 4},
		},
		prepare: func(fields *protestServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)

			fields.protestRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)

			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.NoError(t, err)

		},
	},
}

var testProtestCreateFail = []struct {
	testName  string
	inputData struct {
		protestID        uuid.UUID
		raceID           uuid.UUID
		ratingID         uuid.UUID
		judgeID          uuid.UUID
		ruleNum          int
		reviewDate       time.Time
		comment          string
		protesteeSailNum int
		protestorSailNum int
		witnessesSailNum []int
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName: "invalid ruleNum",
		inputData: struct {
			protestID        uuid.UUID
			raceID           uuid.UUID
			ratingID         uuid.UUID
			judgeID          uuid.UUID
			ruleNum          int
			reviewDate       time.Time
			comment          string
			protesteeSailNum int
			protestorSailNum int
			witnessesSailNum []int
		}{
			protestID:  uuid.New(),
			raceID:     uuid.New(),
			ratingID:   uuid.New(),
			judgeID:    uuid.New(),
			ruleNum:    30,
			reviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			comment:    "Test"},
		prepare: func(fields *protestServiceFields) {},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
		},
	},
	{
		testName: "create failed",
		inputData: struct {
			protestID        uuid.UUID
			raceID           uuid.UUID
			ratingID         uuid.UUID
			judgeID          uuid.UUID
			ruleNum          int
			reviewDate       time.Time
			comment          string
			protesteeSailNum int
			protestorSailNum int
			witnessesSailNum []int
		}{
			protestID:        uuid.New(),
			raceID:           uuid.New(),
			ratingID:         uuid.New(),
			judgeID:          uuid.New(),
			ruleNum:          31,
			reviewDate:       time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			comment:          "Test",
			protesteeSailNum: 12,
			protestorSailNum: 45,
			witnessesSailNum: []int{1, 8, 4},
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().Create(gomock.Any()).Return(nil, repository_errors.InsertError)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{
				ID:       uuid.New(),
				RatingID: uuid.New(),
				SailNum:  12,
				Class:    models.Laser,
			}, nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
		},
	},
	{
		testName: "GetCrewDataBySailNumAndRatingID method failed",
		inputData: struct {
			protestID        uuid.UUID
			raceID           uuid.UUID
			ratingID         uuid.UUID
			judgeID          uuid.UUID
			ruleNum          int
			reviewDate       time.Time
			comment          string
			protesteeSailNum int
			protestorSailNum int
			witnessesSailNum []int
		}{
			protestID:        uuid.New(),
			raceID:           uuid.New(),
			ratingID:         uuid.New(),
			judgeID:          uuid.New(),
			ruleNum:          31,
			reviewDate:       time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			comment:          "Test",
			protesteeSailNum: 12,
			protestorSailNum: 45,
			witnessesSailNum: []int{1, 8, 4},
		},
		prepare: func(fields *protestServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{
				ID:       uuid.New(),
				RatingID: uuid.New(),
				SailNum:  12,
				Class:    models.Laser,
			}, repository_errors.SelectError)
		},

		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
		},
	},
	{
		testName: "AttachCrewToProtest method failed",
		inputData: struct {
			protestID        uuid.UUID
			raceID           uuid.UUID
			ratingID         uuid.UUID
			judgeID          uuid.UUID
			ruleNum          int
			reviewDate       time.Time
			comment          string
			protesteeSailNum int
			protestorSailNum int
			witnessesSailNum []int
		}{
			protestID:        uuid.New(),
			raceID:           uuid.New(),
			ratingID:         uuid.New(),
			judgeID:          uuid.New(),
			ruleNum:          31,
			reviewDate:       time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			comment:          "Test",
			protesteeSailNum: 12,
			protestorSailNum: 45,
			witnessesSailNum: []int{1, 8, 4},
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Protest{
				ID:         uuid.New(),
				RaceID:     uuid.New(),
				RatingID:   uuid.New(),
				JudgeID:    uuid.New(),
				RuleNum:    31,
				ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
				Status:     models.PendingReview,
				Comment:    "Test",
			}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{
				ID:       uuid.New(),
				RatingID: uuid.New(),
				SailNum:  12,
				Class:    models.Laser,
			}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{
				ID:       uuid.New(),
				RatingID: uuid.New(),
				SailNum:  12,
				Class:    models.Laser,
			}, nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.InsertError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
		},
	},
	{
		testName: "Attach witnesse Crew To Protest failed",
		inputData: struct {
			protestID        uuid.UUID
			raceID           uuid.UUID
			ratingID         uuid.UUID
			judgeID          uuid.UUID
			ruleNum          int
			reviewDate       time.Time
			comment          string
			protesteeSailNum int
			protestorSailNum int
			witnessesSailNum []int
		}{
			protestID:        uuid.New(),
			raceID:           uuid.New(),
			ratingID:         uuid.New(),
			judgeID:          uuid.New(),
			ruleNum:          31,
			reviewDate:       time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			comment:          "Test",
			protesteeSailNum: 12,
			protestorSailNum: 45,
			witnessesSailNum: []int{1, 8, 4},
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Protest{
				ID:         uuid.New(),
				RaceID:     uuid.New(),
				RatingID:   uuid.New(),
				JudgeID:    uuid.New(),
				RuleNum:    31,
				ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
				Status:     models.PendingReview,
				Comment:    "Test",
			}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{
				ID:       uuid.New(),
				RatingID: uuid.New(),
				SailNum:  12,
				Class:    models.Laser,
			}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{
				ID:       uuid.New(),
				RatingID: uuid.New(),
				SailNum:  12,
				Class:    models.Laser,
			}, nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.InsertError)
			fields.protestRepoMock.EXPECT().DetachCrewFromProtest(gomock.Any(), gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
		},
	},
	{
		testName: "Get witnesse Crew Data By SailNum And RatingID failed",
		inputData: struct {
			protestID        uuid.UUID
			raceID           uuid.UUID
			ratingID         uuid.UUID
			judgeID          uuid.UUID
			ruleNum          int
			reviewDate       time.Time
			comment          string
			protesteeSailNum int
			protestorSailNum int
			witnessesSailNum []int
		}{
			protestID:        uuid.New(),
			raceID:           uuid.New(),
			ratingID:         uuid.New(),
			judgeID:          uuid.New(),
			ruleNum:          31,
			reviewDate:       time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			comment:          "Test",
			protesteeSailNum: 12,
			protestorSailNum: 45,
			witnessesSailNum: []int{1, 8, 4},
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Protest{
				ID:         uuid.New(),
				RaceID:     uuid.New(),
				RatingID:   uuid.New(),
				JudgeID:    uuid.New(),
				RuleNum:    31,
				ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
				Status:     models.PendingReview,
				Comment:    "Test",
			}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)

			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.InsertError)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
		},
	},
}

func TestProtestServiceCreate(t *testing.T) {
	for _, tt := range testProtestCreateSuccess {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initProtestServiceFields(ctrl)
			protestService := initProtestService(fields)

			tt.prepare(fields)

			protest, err := protestService.AddNewProtest(tt.inputData.protestID, tt.inputData.raceID, tt.inputData.ratingID, tt.inputData.judgeID, tt.inputData.ruleNum, tt.inputData.reviewDate, tt.inputData.comment, tt.inputData.protesteeSailNum, tt.inputData.protestorSailNum, tt.inputData.witnessesSailNum)
			tt.checkOutput(t, protest, err)
		})
	}

	for _, tt := range testProtestCreateFail {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initProtestServiceFields(ctrl)
			protestService := initProtestService(fields)

			tt.prepare(fields)

			protest, err := protestService.AddNewProtest(tt.inputData.protestID, tt.inputData.raceID, tt.inputData.ratingID, tt.inputData.judgeID, tt.inputData.ruleNum, tt.inputData.reviewDate, tt.inputData.comment, tt.inputData.protesteeSailNum, tt.inputData.protestorSailNum, tt.inputData.witnessesSailNum)
			tt.checkOutput(t, protest, err)
		})
	}
}

var testProtestUpdateSuccess = []struct {
	testName  string
	inputData struct {
		protestID  uuid.UUID
		raceID     uuid.UUID
		judgeID    uuid.UUID
		ruleNum    int
		reviewDate time.Time
		status     int
		comment    string
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName: "basic update",
		inputData: struct {
			protestID  uuid.UUID
			raceID     uuid.UUID
			judgeID    uuid.UUID
			ruleNum    int
			reviewDate time.Time
			status     int
			comment    string
		}{
			protestID:  uuid.New(),
			raceID:     uuid.New(),
			ruleNum:    31,
			reviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			status:     models.Reviewed,
			comment:    "Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.UUID{}}, nil)
			fields.protestRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Protest{
				ID:         uuid.New(),
				RaceID:     uuid.New(),
				RatingID:   uuid.New(),
				RuleNum:    31,
				ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
				Status:     models.Reviewed,
				Comment:    "Test",
			}, nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.NoError(t, err)
			assert.Equal(t, 31, protest.RuleNum)
			assert.Equal(t, models.Reviewed, protest.Status)
			assert.Equal(t, "Test", protest.Comment)
		},
	},
}

var testProtestUpdateFail = []struct {
	testName  string
	inputData struct {
		protestID  uuid.UUID
		raceID     uuid.UUID
		judgeID    uuid.UUID
		ruleNum    int
		reviewDate time.Time
		status     int
		comment    string
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName: "invalid ruleNum",
		inputData: struct {
			protestID  uuid.UUID
			raceID     uuid.UUID
			judgeID    uuid.UUID
			ruleNum    int
			reviewDate time.Time
			status     int
			comment    string
		}{
			protestID:  uuid.New(),
			raceID:     uuid.New(),
			ruleNum:    30,
			reviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			status:     models.Reviewed,
			comment:    "Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.UUID{}}, nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
		},
	},
	{
		testName: "invalid status",
		inputData: struct {
			protestID  uuid.UUID
			raceID     uuid.UUID
			judgeID    uuid.UUID
			ruleNum    int
			reviewDate time.Time
			status     int
			comment    string
		}{
			protestID:  uuid.New(),
			raceID:     uuid.New(),
			ruleNum:    31,
			reviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			status:     10,
			comment:    "Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.UUID{}}, nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
		},
	},
	{
		testName: "update failed",
		inputData: struct {
			protestID  uuid.UUID
			raceID     uuid.UUID
			judgeID    uuid.UUID
			ruleNum    int
			reviewDate time.Time
			status     int
			comment    string
		}{
			protestID:  uuid.New(),
			raceID:     uuid.New(),
			ruleNum:    30,
			reviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			status:     models.Reviewed,
			comment:    "Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.UUID{}}, nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
		},
	},
}

func TestProtestServiceUpdate(t *testing.T) {
	for _, tt := range testProtestUpdateSuccess {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initProtestServiceFields(ctrl)
			protestService := initProtestService(fields)

			tt.prepare(fields)
			protest, err := protestService.UpdateProtestByID(tt.inputData.protestID, tt.inputData.raceID, tt.inputData.judgeID, tt.inputData.ruleNum, tt.inputData.reviewDate, tt.inputData.status, tt.inputData.comment)
			tt.checkOutput(t, protest, err)
		})
	}

	for _, tt := range testProtestUpdateFail {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := initProtestServiceFields(ctrl)
			protestService := initProtestService(fields)

			tt.prepare(fields)

			protest, err := protestService.UpdateProtestByID(tt.inputData.protestID, tt.inputData.raceID, tt.inputData.judgeID, tt.inputData.ruleNum, tt.inputData.reviewDate, tt.inputData.status, tt.inputData.comment)
			tt.checkOutput(t, protest, err)
		})
	}
}

var testProtestServiceDelete = []struct {
	testName  string
	inputData struct {
		protestID uuid.UUID
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, err error)
}{
	{
		testName: "delete protest success test",
		inputData: struct {
			protestID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.UUID{}}, nil)
			fields.protestRepoMock.EXPECT().Delete(gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "protest not found",
		inputData: struct {
			protestID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName: "delete protest error",
		inputData: struct {
			protestID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.UUID{}}, nil)
			fields.protestRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DeleteError)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DeleteError, err)
		},
	},
}

func TestProtestServiceDeleteProtest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initProtestServiceFields(ctrl)
	protestService := initProtestService(fields)

	for _, tt := range testProtestServiceDelete {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := protestService.DeleteProtestByID(tt.inputData.protestID)
			tt.checkOutput(t, err)
		})
	}
}

var testProtestServiceGetProtestDataByID = []struct {
	testName    string
	inputData   struct{ protestID uuid.UUID }
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName:  "get protest by id success test",
		inputData: struct{ protestID uuid.UUID }{uuid.New()},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, protest)
		},
	},
	{
		testName:  "protest not found",
		inputData: struct{ protestID uuid.UUID }{uuid.New()},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName:  "get protest by id error",
		inputData: struct{ protestID uuid.UUID }{uuid.New()},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestProtestServiceGetProtestDataByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initProtestServiceFields(ctrl)
	protestService := initProtestService(fields)

	for _, tt := range testProtestServiceGetProtestDataByID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			protest, err := protestService.GetProtestDataByID(tt.inputData.protestID)
			tt.checkOutput(t, protest, err)
		})
	}
}

var testProtestServiceGetProtestsDataByRaceID = []struct {
	testName  string
	inputData struct{ raceID uuid.UUID }
	prepare   func(fields *protestServiceFields)
	checkFunc func(t *testing.T, protests []models.Protest, err error)
}{
	{
		testName:  "get protests by race id success test",
		inputData: struct{ raceID uuid.UUID }{uuid.New()},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestsDataByRaceID(gomock.Any()).Return([]models.Protest{{ID: uuid.New()}, {ID: uuid.New()}}, nil)
		},
		checkFunc: func(t *testing.T, protest []models.Protest, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, protest)
		},
	},
	{
		testName:  "protests not found",
		inputData: struct{ raceID uuid.UUID }{uuid.New()},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestsDataByRaceID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkFunc: func(t *testing.T, protest []models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName:  "get protests by race id error",
		inputData: struct{ raceID uuid.UUID }{uuid.New()},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestsDataByRaceID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkFunc: func(t *testing.T, protest []models.Protest, err error) {
			assert.Error(t, err)
			assert.Nil(t, protest)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestProtestServiceGetProtestsDataByRaceID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initProtestServiceFields(ctrl)
	protestService := initProtestService(fields)

	for _, tt := range testProtestServiceGetProtestsDataByRaceID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			protest, err := protestService.GetProtestsDataByRaceID(tt.inputData.raceID)
			tt.checkFunc(t, protest, err)
		})
	}
}

var testProtestServiceAttachCrewToProtest = []struct {
	testName  string
	inputData struct {
		protestID uuid.UUID
		sailNum   int
		role      int
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			protestID uuid.UUID
			sailNum   int
			role      int
		}{
			uuid.New(),
			121,
			models.Protestee,
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "attach crew to protest error",
		inputData: struct {
			protestID uuid.UUID
			sailNum   int
			role      int
		}{
			uuid.New(),
			121,
			models.Protestee,
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.InsertError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.InsertError, err)
		},
	},
	{
		testName: "invalid role error",
		inputData: struct {
			protestID uuid.UUID
			sailNum   int
			role      int
		}{
			uuid.New(),
			121,
			100,
		},
		prepare: func(fields *protestServiceFields) {
			//fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid protest role"), err)
		},
	},
	{
		testName: "crew not found error",
		inputData: struct {
			protestID uuid.UUID
			sailNum   int
			role      int
		}{
			uuid.New(),
			121,
			100,
		},
		prepare: func(fields *protestServiceFields) {
			//fields.protestRepoMock.EXPECT().AttachCrewToProtest(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid protest role"), err)
		},
	},
}

func TestProtestServiceAttachCrewToProtest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initProtestServiceFields(ctrl)
	protestService := initProtestService(fields)

	for _, tt := range testProtestServiceAttachCrewToProtest {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := protestService.AttachCrewToProtest(tt.inputData.protestID, tt.inputData.sailNum, tt.inputData.role)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testProtestServiceDetachCrewFromProtest = []struct {
	testName  string
	inputData struct {
		protestID uuid.UUID
		sailNum   int
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			protestID uuid.UUID
			sailNum   int
		}{
			uuid.New(),
			123,
		},
		prepare: func(fields *protestServiceFields) {
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().DetachCrewFromProtest(gomock.Any(), gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "detach crew from protest error",
		inputData: struct {
			protestID uuid.UUID
			sailNum   int
		}{
			uuid.New(),
			234,
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().DetachCrewFromProtest(gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
	{
		testName: "crew not found error",
		inputData: struct {
			protestID uuid.UUID
			sailNum   int
		}{
			uuid.New(),
			234,
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.crewRepoMock.EXPECT().GetCrewDataBySailNumAndRatingID(gomock.Any(), gomock.Any()).Return(&models.Crew{ID: uuid.New()}, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestProtestServiceDetachCrewFromProtest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initProtestServiceFields(ctrl)
	protestService := initProtestService(fields)

	for _, tt := range testProtestServiceDetachCrewFromProtest {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := protestService.DetachCrewFromProtest(tt.inputData.protestID, tt.inputData.sailNum)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testProtestServiceCompleteReview = []struct {
	testName  string
	inputData struct {
		protestID       uuid.UUID
		protesteePoints int
		comment         string
	}
	prepare     func(fields *protestServiceFields)
	checkOutput func(t *testing.T, protest *models.Protest, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			protestID       uuid.UUID
			protesteePoints int
			comment         string
		}{
			uuid.New(),
			123,
			"Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().GetProtestParticipantsIDByID(gomock.Any()).Return(make(map[int]uuid.UUID), nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			//fields.crewResInRaceRepoMock.EXPECT().GetCrewResInRaceDataByID(gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)

		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "protest not found error",
		inputData: struct {
			protestID       uuid.UUID
			protesteePoints int
			comment         string
		}{
			uuid.New(),
			123,
			"Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
	{
		testName: "protest update error",
		inputData: struct {
			protestID       uuid.UUID
			protesteePoints int
			comment         string
		}{
			uuid.New(),
			123,
			"Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().GetProtestParticipantsIDByID(gomock.Any()).Return(make(map[int]uuid.UUID), nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().Update(gomock.Any()).Return(nil, repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
	{
		testName: "crew res in race update error",
		inputData: struct {
			protestID       uuid.UUID
			protesteePoints int
			comment         string
		}{
			uuid.New(),
			123,
			"Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			//fields.protestRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, repository_errors.UpdateError)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().GetProtestParticipantsIDByID(gomock.Any()).Return(make(map[int]uuid.UUID), nil)

		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
	{
		testName: "GetCrewResByRaceIDAndCrewID error",
		inputData: struct {
			protestID       uuid.UUID
			protesteePoints int
			comment         string
		}{
			uuid.New(),
			123,
			"Test",
		},
		prepare: func(fields *protestServiceFields) {
			fields.protestRepoMock.EXPECT().GetProtestDataByID(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			//fields.protestRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Protest{ID: uuid.New()}, nil)
			//fields.crewResInRaceRepoMock.EXPECT().Update(gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, nil)
			fields.protestRepoMock.EXPECT().GetProtestParticipantsIDByID(gomock.Any()).Return(make(map[int]uuid.UUID), nil)
			fields.crewResInRaceRepoMock.EXPECT().GetCrewResByRaceIDAndCrewID(gomock.Any(), gomock.Any()).Return(&models.CrewResInRace{CrewID: uuid.New()}, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, protest *models.Protest, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestProtestServiceCompleteReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initProtestServiceFields(ctrl)
	protestService := initProtestService(fields)

	for _, tt := range testProtestServiceCompleteReview {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := protestService.CompleteReview(tt.inputData.protestID, tt.inputData.protesteePoints, tt.inputData.comment)
			tt.checkOutput(t, nil, err)
		})
	}
}
