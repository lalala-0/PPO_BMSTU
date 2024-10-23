package test_services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	services "PPO_BMSTU/internal/services"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	mock_password_hash "PPO_BMSTU/tests/hasher_mocks"
	mock_repository_interfaces "PPO_BMSTU/tests/repository_mocks"
	builders2 "PPO_BMSTU/tests/unit_tests/builders"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

type judgeServiceFields struct {
	judgeRepoMock *mock_repository_interfaces.MockIJudgeRepository
	logger        *log.Logger
	hash          *mock_password_hash.MockPasswordHash
}

func initJudgeServiceFields(ctrl *gomock.Controller) *judgeServiceFields {
	judgeRepoMock := mock_repository_interfaces.NewMockIJudgeRepository(ctrl)
	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f)

	return &judgeServiceFields{
		judgeRepoMock: judgeRepoMock,
		hash:          mock_password_hash.NewMockPasswordHash(ctrl),
		logger:        logger,
	}
}

func initJudgeService(fields *judgeServiceFields) service_interfaces.IJudgeService {
	return services.NewJudgeService(fields.judgeRepoMock, fields.hash, fields.logger)
}

var testJudgeGetByID = []struct {
	testName  string
	inputData struct {
		id uuid.UUID
	}
	prepare     func(fields *judgeServiceFields)
	checkOutput func(t *testing.T, judge *models.Judge, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			id uuid.UUID
		}{
			id: uuid.New(),
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(
				builders2.NewJudgeBuilder().
					WithID(uuid.New()).
					WithFIO("Test").
					WithLogin("Test").
					WithPassword("password123").
					WithRole(1).
					WithPost("Test").
					Build(), nil)
		},
		checkOutput: func(t *testing.T, judge *models.Judge, err error) {
			assert.NoError(t, err)
			assert.Equal(t, "Test", judge.FIO)
			assert.Equal(t, "Test", judge.Login)
			assert.Equal(t, "Test", judge.Post)
			assert.Equal(t, 1, judge.Role)
		},
	},
	{
		testName: "judgeView not found",
		inputData: struct {
			id uuid.UUID
		}{
			id: uuid.New(),
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
		},
	},
}

func TestJudgeService_GetJudgeByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initJudgeServiceFields(ctrl)
	service := initJudgeService(fields)

	for _, tt := range testJudgeGetByID {
		tt.prepare(fields)
		t.Run(tt.testName, func(t *testing.T) {
			judge, err := service.GetJudgeDataByID(tt.inputData.id)
			tt.checkOutput(t, judge, err)
		})
	}
}

var testJudgeGetJudgeDataByProtestID = []struct {
	testName  string
	inputData struct {
		id uuid.UUID
	}
	prepare     func(fields *judgeServiceFields)
	checkOutput func(t *testing.T, judge *models.Judge, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			id uuid.UUID
		}{
			id: uuid.New(),
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByProtestID(gomock.Any()).Return(
				builders2.NewJudgeBuilder().
					WithID(uuid.New()).
					WithFIO("Test").
					WithLogin("Test").
					WithPassword("password123").
					WithRole(1).
					WithPost("Test").
					Build(), nil)
		},
		checkOutput: func(t *testing.T, judge *models.Judge, err error) {
			assert.NoError(t, err)
			assert.Equal(t, "Test", judge.FIO)
			assert.Equal(t, "Test", judge.Login)
			assert.Equal(t, "Test", judge.Post)
			assert.Equal(t, 1, judge.Role)
		},
	},
	{
		testName: "judgeView not found",
		inputData: struct {
			id uuid.UUID
		}{
			id: uuid.New(),
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByProtestID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkOutput: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
		},
	},
}

func TestJudgeServiceGetJudgeDataByProtestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initJudgeServiceFields(ctrl)
	service := initJudgeService(fields)

	for _, tt := range testJudgeGetJudgeDataByProtestID {
		tt.prepare(fields)
		t.Run(tt.testName, func(t *testing.T) {
			judge, err := service.GetJudgeDataByProtestID(tt.inputData.id)
			tt.checkOutput(t, judge, err)
		})
	}
}

var testJudgeGetJudgesDataByRatingID = []struct {
	testName  string
	prepare   func(fields *judgeServiceFields)
	checkFunc func(t *testing.T, workers []models.Judge, err error)
}{
	{
		testName: "Success",
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgesDataByRatingID(gomock.Any()).Return([]models.Judge{
				*builders2.NewJudgeBuilder().
					WithID(uuid.New()).
					WithFIO("Test").
					WithLogin("Test").
					WithPassword("password123").
					WithRole(1).
					WithPost("Test").
					Build(),
				*builders2.NewJudgeBuilder().
					WithID(uuid.New()).
					WithRole(2).
					Build(),
				*builders2.NewJudgeBuilder().
					WithID(uuid.New()).
					WithRole(3).
					Build(),
				*builders2.NewJudgeBuilder().
					WithID(uuid.New()).
					WithLogin("Test3").
					WithRole(2).
					Build(),
			}, nil)
		},
		checkFunc: func(t *testing.T, workers []models.Judge, err error) {
			assert.NoError(t, err)
			assert.Equal(t, 4, len(workers))
			assert.Equal(t, "Test", workers[0].FIO)
			assert.Equal(t, "Test", workers[0].Login)
			assert.Equal(t, 3, workers[2].Role)
			assert.Equal(t, "Test", workers[0].Post)
			assert.Equal(t, "Test3", workers[3].Login)
		},
	},
	{
		testName: "empty list",
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgesDataByRatingID(gomock.Any()).Return([]models.Judge{}, nil)
		},
		checkFunc: func(t *testing.T, workers []models.Judge, err error) {
			assert.NoError(t, err)
			assert.Equal(t, 0, len(workers))
		},
	},
}

func TestJudgeServiceGetJudgesDataByRatingID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initJudgeServiceFields(ctrl)
	service := initJudgeService(fields)

	for _, tt := range testJudgeGetJudgesDataByRatingID {
		tt.prepare(fields)
		t.Run(tt.testName, func(t *testing.T) {
			workers, err := service.GetJudgesDataByRatingID(uuid.New())
			tt.checkFunc(t, workers, err)
		})
	}
}

var testJudgeDelete = []struct {
	testName  string
	inputData struct {
		id uuid.UUID
	}
	prepare   func(fields *judgeServiceFields)
	checkFunc func(t *testing.T, err error)
}{
	{
		testName:  "Success",
		inputData: struct{ id uuid.UUID }{id: uuid.New()},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(builders2.JudgeMother.Default(), nil)
			fields.judgeRepoMock.EXPECT().DeleteProfile(gomock.Any()).Return(nil)
		},
		checkFunc: func(t *testing.T, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName:  "judgeView not found",
		inputData: struct{ id uuid.UUID }{id: uuid.New()},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkFunc: func(t *testing.T, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
}

func TestJudgeService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initJudgeServiceFields(ctrl)
	service := initJudgeService(fields)

	for _, tt := range testJudgeDelete {
		tt.prepare(fields)
		t.Run(tt.testName, func(t *testing.T) {
			err := service.DeleteProfile(tt.inputData.id)
			tt.checkFunc(t, err)
		})
	}
}

// ----------------------------------------
var testJudgeUpdateProfile = []struct {
	testName  string
	inputData struct {
		id       uuid.UUID
		fio      string
		login    string
		post     string
		role     int
		password string
	}
	prepare   func(fields *judgeServiceFields)
	checkFunc func(t *testing.T, judge *models.Judge, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			id       uuid.UUID
			fio      string
			login    string
			post     string
			role     int
			password string
		}{
			id:       uuid.New(),
			fio:      "Test",
			login:    "Test",
			post:     "Test",
			role:     1,
			password: "password123", //change password from
		},
		prepare: func(fields *judgeServiceFields) {
			var judge = builders2.JudgeMother.Default()
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(judge, nil)
			fields.hash.EXPECT().GetHash(gomock.Any()).Return("hash", nil)
			fields.judgeRepoMock.EXPECT().UpdateProfile(judge).Return(judge, nil)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.NoError(t, err)
		},
	},
	{
		testName: "crew not found",
		inputData: struct {
			id       uuid.UUID
			fio      string
			login    string
			post     string
			role     int
			password string
		}{
			id:       uuid.New(),
			fio:      "Test",
			login:    "Test",
			post:     "Test",
			role:     1,
			password: "password123", //change password from
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName: "invalid password",
		inputData: struct {
			id       uuid.UUID
			fio      string
			login    string
			post     string
			role     int
			password string
		}{
			id:       uuid.New(),
			fio:      "Test",
			login:    "Test",
			post:     "Test",
			role:     1,
			password: "123", //change password from
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(builders2.JudgeMother.Default(), nil)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Equal(t, service_errors.InvalidPassword, err)
		},
	},
	{
		testName: "invalid fio",
		inputData: struct {
			id       uuid.UUID
			fio      string
			login    string
			post     string
			role     int
			password string
		}{
			id:       uuid.New(),
			fio:      "",
			login:    "Test",
			post:     "Test",
			role:     1,
			password: "password123", //change password from
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(builders2.JudgeMother.Default(), nil)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Equal(t, service_errors.InvalidFIO, err)
		},
	},
	{
		testName: "invalid role",
		inputData: struct {
			id       uuid.UUID
			fio      string
			login    string
			post     string
			role     int
			password string
		}{
			id:       uuid.New(),
			fio:      "ol",
			login:    "Test",
			post:     "Test",
			role:     10,
			password: "password123", //change password from
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(builders2.JudgeMother.Default(), nil)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Equal(t, service_errors.InvalidRole, err)
		},
	},
	{
		testName: "invalid login",
		inputData: struct {
			id       uuid.UUID
			fio      string
			login    string
			post     string
			role     int
			password string
		}{
			id:       uuid.New(),
			fio:      "ol",
			login:    "",
			post:     "Test",
			role:     1,
			password: "password123", //change password from
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByID(gomock.Any()).Return(builders2.JudgeMother.Default(), nil)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Equal(t, service_errors.InvalidLogin, err)
		},
	},
}

func TestJudgeServiceUpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initJudgeServiceFields(ctrl)
	service := initJudgeService(fields)

	for _, tt := range testJudgeUpdateProfile {
		tt.prepare(fields)
		t.Run(tt.testName, func(t *testing.T) {
			judge, err := service.UpdateProfile(tt.inputData.id, tt.inputData.fio, tt.inputData.login, tt.inputData.password, tt.inputData.role)
			tt.checkFunc(t, judge, err)
		})
	}
}

var testJudgeCreateProfile = []struct {
	testName  string
	inputData struct {
		judge    *models.Judge
		password string
	}
	prepare   func(fields *judgeServiceFields)
	checkFunc func(t *testing.T, judge *models.Judge, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			judge    *models.Judge
			password string
		}{
			judge: &models.Judge{
				FIO:   "Test",
				Login: "Test",
				Post:  "Test",
				Role:  1,
			},
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByLogin(gomock.Any()).Return(nil, nil)
			fields.hash.EXPECT().GetHash(gomock.Any()).Return("hash", nil)
			fields.judgeRepoMock.EXPECT().CreateProfile(gomock.Any()).Return(&models.Judge{
				ID:       uuid.New(),
				FIO:      "Test",
				Login:    "Test",
				Post:     "Test",
				Role:     1,
				Password: "hash",
			}, nil)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.NoError(t, err)
			assert.Equal(t, "Test", judge.FIO)
			assert.Equal(t, "Test", judge.Login)
			assert.Equal(t, "Test", judge.Post)
		},
	},
	{
		testName: "judgeView already exists",
		inputData: struct {
			judge    *models.Judge
			password string
		}{
			judge: &models.Judge{
				FIO:   "Test",
				Login: "Test",
				Post:  "Test",
				Role:  1,
			},
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByLogin(gomock.Any()).Return(builders2.JudgeMother.Default(), nil)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
			assert.Equal(t, service_errors.NotUnique, err)
		},
	},
	{
		testName: "invalid fio",
		inputData: struct {
			judge    *models.Judge
			password string
		}{
			judge: &models.Judge{
				FIO:   "",
				Login: "Test",
				Post:  "Test",
				Role:  1,
			},
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid Login",
		inputData: struct {
			judge    *models.Judge
			password string
		}{
			judge: &models.Judge{
				FIO:   "aa",
				Login: "",
				Post:  "Test",
				Role:  1,
			},
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid role",
		inputData: struct {
			judge    *models.Judge
			password string
		}{
			judge: &models.Judge{
				FIO:   "hgrtunju",
				Login: "Test",
				Post:  "Test",
				Role:  10,
			},
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
	{
		testName: "invalid password",
		inputData: struct {
			judge    *models.Judge
			password string
		}{
			judge: &models.Judge{
				FIO:   "Test",
				Login: "Test",
				Post:  "Test",
				Role:  1,
			},
			password: "123",
		},
		prepare: func(fields *judgeServiceFields) {},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input data"), err)
		},
	},
}

func TestJudgeServiceCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initJudgeServiceFields(ctrl)
	service := initJudgeService(fields)

	for _, tt := range testJudgeCreateProfile {
		tt.prepare(fields)
		t.Run(tt.testName, func(t *testing.T) {
			judge, err := service.CreateProfile(tt.inputData.judge.ID, tt.inputData.judge.FIO, tt.inputData.judge.Login, tt.inputData.password, tt.inputData.judge.Role, tt.inputData.judge.Post)
			tt.checkFunc(t, judge, err)
		})
	}
}

var testJudgeLogin = []struct {
	testName  string
	inputData struct {
		login    string
		password string
	}
	prepare   func(fields *judgeServiceFields)
	checkFunc func(t *testing.T, judge *models.Judge, err error)
}{
	{
		testName: "Success",
		inputData: struct {
			login    string
			password string
		}{
			login:    "abcdef",
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByLogin(gomock.Any()).Return(&models.Judge{
				FIO:      "Test",
				Login:    "abcdef",
				Post:     "Test",
				Role:     1,
				Password: "hash",
			}, nil)
			fields.hash.EXPECT().CompareHashAndPassword(gomock.Any(), gomock.Any()).Return(true)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.NoError(t, err)
			assert.Equal(t, "Test", judge.FIO)
			assert.Equal(t, "abcdef", judge.Login)
			assert.Equal(t, "hash", judge.Password)
		},
	},
	{
		testName: "judgeView not found",
		inputData: struct {
			login    string
			password string
		}{
			login:    "notFound",
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByLogin(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
			assert.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
	{
		testName: "invalid password",
		inputData: struct {
			login    string
			password string
		}{
			login:    "abcdef",
			password: "password123",
		},
		prepare: func(fields *judgeServiceFields) {
			fields.judgeRepoMock.EXPECT().GetJudgeDataByLogin(gomock.Any()).Return(&models.Judge{
				FIO:   "Test",
				Login: "qwerty",
			}, nil)
			fields.hash.EXPECT().CompareHashAndPassword(gomock.Any(), gomock.Any()).Return(false)
		},
		checkFunc: func(t *testing.T, judge *models.Judge, err error) {
			assert.Error(t, err)
			assert.Nil(t, judge)
			assert.Equal(t, service_errors.MismatchedPassword, err)
		},
	},
}

func TestJudgeServiceLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields := initJudgeServiceFields(ctrl)
	service := initJudgeService(fields)

	for _, tt := range testJudgeLogin {
		tt.prepare(fields)
		t.Run(tt.testName, func(t *testing.T) {
			judge, err := service.Login(tt.inputData.login, tt.inputData.password)
			tt.checkFunc(t, judge, err)
		})
	}
}
