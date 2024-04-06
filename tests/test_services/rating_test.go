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

type ratingServiceFields struct {
	ratingRepoMock *mock_repository_interfaces.MockIRatingRepository
	logger         *log.Logger
}

func initRatingServiceFields(ctrl *gomock.Controller) *ratingServiceFields {
	ratingRepoMock := mock_repository_interfaces.NewMockIRatingRepository(ctrl)

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f)

	return &ratingServiceFields{
		ratingRepoMock: ratingRepoMock,
		logger:         logger,
	}
}

func initRatingService(fields *ratingServiceFields) service_interfaces.IRatingService {
	return services.NewRatingService(fields.ratingRepoMock, fields.logger)
}

var testRatingServiceAddNewRating = []struct {
	testName  string
	inputData struct {
		id         uuid.UUID
		class      string
		blowoutCnt int
	}
	prepare     func(fields *ratingServiceFields)
	checkOutput func(t *testing.T, rating *models.Rating, err error)
}{
	{
		testName: "create rating success test",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"laser",
			0,
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().Create(gomock.Any()).Return(&models.Rating{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, rating)
		},
	},
	{
		testName: "invalid class",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"klkl",
			0,
		},
		prepare: func(fields *ratingServiceFields) {},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Nil(t, rating)
			assert.Equal(t, service_errors.InvalidClass, err)
		},
	},
	{
		testName: "invalid blowout count",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"laser",
			-90,
		},
		prepare: func(fields *ratingServiceFields) {},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Nil(t, rating)
			assert.Equal(t, service_errors.InvalidBlowoutCnt, err)
		},
	},
	{
		testName: "rating creation error",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"laser",
			9,
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().Create(gomock.Any()).Return(nil, repository_errors.InsertError)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Nil(t, rating)
			assert.Equal(t, repository_errors.InsertError, err)
		},
	},
}

func TestRatingService_CreateRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRatingServiceFields(ctrl)
	ratingService := initRatingService(fields)

	for _, tt := range testRatingServiceAddNewRating {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			rating, err := ratingService.AddNewRating(tt.inputData.id, tt.inputData.class, tt.inputData.blowoutCnt)
			tt.checkOutput(t, rating, err)
		})
	}
}

var testRatingServiceDelete = []struct {
	testName  string
	inputData struct {
		ratingID uuid.UUID
	}
	prepare     func(fields *ratingServiceFields)
	checkOutput func(t *testing.T, err error)
}{
	{
		testName: "delete rating success test",
		inputData: struct {
			ratingID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(&models.Rating{ID: uuid.UUID{}}, nil)
			fields.ratingRepoMock.EXPECT().Delete(gomock.Any()).Return(nil)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "rating not found",
		inputData: struct {
			ratingID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName: "delete rating error",
		inputData: struct {
			ratingID uuid.UUID
		}{
			uuid.New(),
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(&models.Rating{ID: uuid.UUID{}}, nil)
			fields.ratingRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DeleteError)
		},
		checkOutput: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DeleteError, err)
		},
	},
}

func TestRatingServiceDeleteRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRatingServiceFields(ctrl)
	ratingService := initRatingService(fields)

	for _, tt := range testRatingServiceDelete {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := ratingService.DeleteRatingByID(tt.inputData.ratingID)
			tt.checkOutput(t, err)
		})
	}
}

var testRatingServiceUpdateRatingByID = []struct {
	testName  string
	inputData struct {
		id         uuid.UUID
		class      string
		blowoutCnt int
	}
	prepare     func(fields *ratingServiceFields)
	checkOutput func(t *testing.T, rating *models.Rating, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"laser",
			0,
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(&models.Rating{
				ID:         uuid.New(),
				Class:      "jmkkkk",
				BlowoutCnt: 1,
			}, nil)
			fields.ratingRepoMock.EXPECT().Update(gomock.Any()).Return(&models.Rating{
				uuid.New(),
				"laser",
				0,
			}, nil)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "rating not found",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"laser",
			0,
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName: "invalid class",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"",
			0,
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(&models.Rating{
				ID:         uuid.New(),
				Class:      "jmkkkk",
				BlowoutCnt: 1,
			}, nil)

		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Equal(t, service_errors.InvalidClass, err)
		},
	},
	{
		testName: "invalid class blow count",
		inputData: struct {
			id         uuid.UUID
			class      string
			blowoutCnt int
		}{
			uuid.New(),
			"kmlm",
			-10,
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(&models.Rating{
				ID:         uuid.New(),
				Class:      "jmkkkk",
				BlowoutCnt: 1,
			}, nil)

			fields.ratingRepoMock.EXPECT().Update(gomock.Any()).Return(nil, repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Equal(t, service_errors.InvalidBlowoutCnt, err)
		},
	},
}

func TestRatingServiceUpdateRatingByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRatingServiceFields(ctrl)
	ratingService := initRatingService(fields)

	for _, tt := range testRatingServiceUpdateRatingByID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			rating, err := ratingService.UpdateRatingByID(tt.inputData.id, tt.inputData.class, tt.inputData.blowoutCnt)
			tt.checkOutput(t, rating, err)
		})
	}
}

var testRatingServiceAttachJudgeToRating = []struct {
	testName  string
	inputData struct {
		ratingID uuid.UUID
		judgeID  uuid.UUID
	}
	prepare     func(fields *ratingServiceFields)
	checkOutput func(t *testing.T, rating *models.Rating, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			ratingID uuid.UUID
			judgeID  uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().AttachJudgeToRating(gomock.Any(), gomock.Any()).Return(&models.Rating{}, nil)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "attach worker to rating error",
		inputData: struct {
			ratingID uuid.UUID
			judgeID  uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().AttachJudgeToRating(gomock.Any(), gomock.Any()).Return(nil, repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
}

func TestRatingServiceAttachJudgeToRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRatingServiceFields(ctrl)
	ratingService := initRatingService(fields)

	for _, tt := range testRatingServiceAttachJudgeToRating {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := ratingService.AttachJudgeToRating(tt.inputData.ratingID, tt.inputData.judgeID)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testRatingServiceDetachJudgeFromRating = []struct {
	testName  string
	inputData struct {
		ratingID uuid.UUID
		judgeID  uuid.UUID
	}
	prepare     func(fields *ratingServiceFields)
	checkOutput func(t *testing.T, rating *models.Rating, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			ratingID uuid.UUID
			judgeID  uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().DetachJudgeFromRating(gomock.Any(), gomock.Any()).Return(&models.Rating{}, nil)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "detach judge from rating error",
		inputData: struct {
			ratingID uuid.UUID
			judgeID  uuid.UUID
		}{
			uuid.New(),
			uuid.New(),
		},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().DetachJudgeFromRating(gomock.Any(), gomock.Any()).Return(nil, repository_errors.UpdateError)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.UpdateError, err)
		},
	},
}

func TestRatingServiceDetachJudgeFromRating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRatingServiceFields(ctrl)
	ratingService := initRatingService(fields)

	for _, tt := range testRatingServiceDetachJudgeFromRating {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			err := ratingService.DetachJudgeFromRating(tt.inputData.ratingID, tt.inputData.judgeID)
			tt.checkOutput(t, nil, err)
		})
	}
}

var testRatingServiceGetRatingDataByID = []struct {
	testName    string
	inputData   struct{ ratingID uuid.UUID }
	prepare     func(fields *ratingServiceFields)
	checkOutput func(t *testing.T, rating *models.Rating, err error)
}{
	{
		testName:  "get rating by id success test",
		inputData: struct{ ratingID uuid.UUID }{uuid.New()},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(&models.Rating{ID: uuid.New()}, nil)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, rating)
		},
	},
	{
		testName:  "rating not found",
		inputData: struct{ ratingID uuid.UUID }{uuid.New()},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Nil(t, rating)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName:  "get rating by id error",
		inputData: struct{ ratingID uuid.UUID }{uuid.New()},
		prepare: func(fields *ratingServiceFields) {
			fields.ratingRepoMock.EXPECT().GetRatingDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
		},
		checkOutput: func(t *testing.T, rating *models.Rating, err error) {
			assert.Error(t, err)
			assert.Nil(t, rating)
			assert.Equal(t, repository_errors.SelectError, err)
		},
	},
}

func TestRatingServiceGetRatingDataByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initRatingServiceFields(ctrl)
	ratingService := initRatingService(fields)

	for _, tt := range testRatingServiceGetRatingDataByID {
		t.Run(tt.testName, func(t *testing.T) {
			tt.prepare(fields)
			rating, err := ratingService.GetRatingDataByID(tt.inputData.ratingID)
			tt.checkOutput(t, rating, err)
		})
	}
}
