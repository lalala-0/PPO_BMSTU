package test_repositories

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/postgres"
	"PPO_BMSTU/tests/test_repositories/postgres_init"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

var testJudgeRepositoryCreateProfileSuccess = []struct {
	TestName    string
	InputData   *models.Judge
	CheckOutput func(t *testing.T, inputData *models.Judge, createdJudge *models.Judge, err error)
}{
	{
		TestName: "create success test",
		InputData: &models.Judge{
			ID:       uuid.New(),
			FIO:      "Test",
			Login:    "Test",
			Password: "test123",
			Post:     "Test",
			Role:     1,
		},
		CheckOutput: func(t *testing.T, inputData *models.Judge, createdJudge *models.Judge, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.FIO, createdJudge.FIO)
			require.Equal(t, inputData.Login, createdJudge.Login)
			require.Equal(t, inputData.Post, createdJudge.Post)
			require.Equal(t, inputData.Role, createdJudge.Role)
		},
	},
}

func TestJudgeRepositoryCreateProfile(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}

	for _, test := range testJudgeRepositoryCreateProfileSuccess {
		judgeRepository := postgres.CreateJudgeRepository(&fields)
		t.Run(test.TestName, func(t *testing.T) {
			createdJudge, err := judgeRepository.CreateProfile(test.InputData)
			test.CheckOutput(t, test.InputData, createdJudge, err)
		})
	}
}

var testJudgeRepositoryGetByIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdJudge *models.Judge, receivedJudge *models.Judge, err error)
}{
	{
		TestName: "get by id success test",
		CheckOutput: func(t *testing.T, createdJudge *models.Judge, receivedJudge *models.Judge, err error) {
			require.NoError(t, err)
			require.Equal(t, createdJudge.ID, receivedJudge.ID)
		},
	},
}

func TestJudgeRepositoryGetByID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	judgeRepository := postgres.CreateJudgeRepository(&fields)

	for _, test := range testJudgeRepositoryGetByIDSuccess {
		createdJudge, err := judgeRepository.CreateProfile(
			&models.Judge{
				ID:       uuid.New(),
				FIO:      "Test",
				Login:    "Test",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
		)

		receivedJudge, err := judgeRepository.GetJudgeDataByID(createdJudge.ID)
		test.CheckOutput(t, createdJudge, receivedJudge, err)
	}
}

//GetJudgeDataByProtestID(protestID uuid.UUID) (*models.Judge, error)

var testJudgeRepositoryGetJudgeDataByProtestIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdJudge *models.Judge, receivedJudge *models.Judge, err error)
}{
	{
		TestName: "get by Protest ID success test",
		CheckOutput: func(t *testing.T, createdJudge *models.Judge, receivedJudge *models.Judge, err error) {
			require.NoError(t, err)
			require.Equal(t, createdJudge.ID, receivedJudge.ID)
		},
	},
}

func TestJudgeRepositoryGetJudgeDataByProtestID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	judgeRepository := postgres.CreateJudgeRepository(&fields)

	for _, test := range testJudgeRepositoryGetJudgeDataByProtestIDSuccess {
		rating := postgres_init.CreateRating(&fields)
		judge := postgres_init.CreateJudge(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		protest := postgres_init.CreateProtest(&fields, race.ID, judge.ID, rating.ID)

		receivedJudge, err := judgeRepository.GetJudgeDataByProtestID(protest.ID)
		test.CheckOutput(t, judge, receivedJudge, err)
	}
}

//GetJudgeDataByLogin(login string) (*models.Judge, error)

var testJudgeRepositoryGetJudgeDataByLoginSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdJudge *models.Judge, receivedJudge *models.Judge, err error)
}{
	{
		TestName: "get by Login success test",
		CheckOutput: func(t *testing.T, createdJudge *models.Judge, receivedJudge *models.Judge, err error) {
			require.NoError(t, err)
			require.Equal(t, createdJudge.ID, receivedJudge.ID)
		},
	},
}

func TestJudgeRepositoryGetJudgeDataByLogin(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	judgeRepository := postgres.CreateJudgeRepository(&fields)

	for _, test := range testJudgeRepositoryGetJudgeDataByLoginSuccess {
		judge := postgres_init.CreateJudge(&fields)

		receivedJudge, err := judgeRepository.GetJudgeDataByLogin(judge.Login)
		test.CheckOutput(t, judge, receivedJudge, err)
	}
}

var testJudgeRepositoryDeleteSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "delete success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestJudgeRepositoryDelete(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	judgeRepository := postgres.CreateJudgeRepository(&fields)

	for _, test := range testJudgeRepositoryDeleteSuccess {
		createdJudge, err := judgeRepository.CreateProfile(
			&models.Judge{
				ID:       uuid.New(),
				FIO:      "Test",
				Login:    "Test",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
		)

		err = judgeRepository.DeleteProfile(createdJudge.ID)
		test.CheckOutput(t, err)

		_, err = judgeRepository.GetJudgeDataByID(createdJudge.ID)
		require.Error(t, err)
	}
}

//UpdateProfile(judgeView *models.Judge) (*models.Judge, error)

var testJudgeRepositoryUpdateSuccess = []struct {
	TestName  string
	InputData struct {
		Judge *models.Judge
	}
	CheckOutput func(t *testing.T, createdJudge *models.Judge, updatedJudge *models.Judge, err error)
}{
	{
		TestName: "update success test",
		InputData: struct {
			Judge *models.Judge
		}{
			&models.Judge{
				ID:       uuid.New(),
				FIO:      "Test",
				Login:    "Test",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
		},
		CheckOutput: func(t *testing.T, createdJudge *models.Judge, updatedJudge *models.Judge, err error) {
			require.NoError(t, err)
			require.Equal(t, createdJudge.ID, updatedJudge.ID)
			require.Equal(t, updatedJudge.ID, createdJudge.ID)
			require.Equal(t, updatedJudge.FIO, createdJudge.FIO)
			require.Equal(t, updatedJudge.Login, createdJudge.Login)
			require.Equal(t, updatedJudge.Post, createdJudge.Post)
			require.Equal(t, updatedJudge.Role, createdJudge.Role)
		},
	},
}

func TestJudgeRepositoryUpdate(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	judgeRepository := postgres.CreateJudgeRepository(&fields)

	for _, test := range testJudgeRepositoryUpdateSuccess {
		createdJudge, err := judgeRepository.CreateProfile(test.InputData.Judge)

		updatedJudge, err := judgeRepository.UpdateProfile(
			&models.Judge{
				ID:       createdJudge.ID,
				FIO:      "Test",
				Login:    "Test",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
		)

		test.CheckOutput(t, createdJudge, updatedJudge, err)
	}
}

//GetJudgesDataByRatingID(ratingID uuid.UUID) ([]models.Judge, error)

var testJudgeRepositoryGetJudgesDataByRatingID = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdJudges []models.Judge, receivedJudges []models.Judge, err error)
}{
	{
		TestName: "Get Judges Data By Rating ID success test",
		CheckOutput: func(t *testing.T, createdJudges []models.Judge, receivedJudges []models.Judge, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdJudges), len(receivedJudges))
		},
	},
}

func TestJudgeRepositoryGetJudgesDataByRatingID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	judgeRepository := postgres.CreateJudgeRepository(&fields)

	for _, test := range testJudgeRepositoryGetJudgesDataByRatingID {
		rating := postgres_init.CreateRating(&fields)

		createdJudges := []models.Judge{
			{
				ID:       uuid.New(),
				FIO:      "Test",
				Login:    "Te1st",
				Password: "tes2t123",
				Post:     "Tes1	t",
				Role:     1,
			},
			{
				ID:       uuid.New(),
				FIO:      "Test1",
				Login:    "Test",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
			{
				ID:       uuid.New(),
				FIO:      "Test2",
				Login:    "Tes222t",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
		}

		for i, _ := range createdJudges {
			j, err := judgeRepository.CreateProfile(&createdJudges[i])
			require.NoError(t, err)
			postgres_init.AttachJudgeToRating(&fields, j.ID, rating.ID)
		}

		receivedJudges, err := judgeRepository.GetJudgesDataByRatingID(rating.ID)
		test.CheckOutput(t, createdJudges, receivedJudges, err)
	}
}

var testJudgeRepositoryGetAllJudges = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdJudges []models.Judge, receivedJudges []models.Judge, err error)
}{
	{
		TestName: "Get all Judges success test",
		CheckOutput: func(t *testing.T, createdJudges []models.Judge, receivedJudges []models.Judge, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdJudges), len(receivedJudges))
		},
	},
}

func TestJudgeRepositoryGetAllJudges(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	judgeRepository := postgres.CreateJudgeRepository(&fields)

	for _, test := range testJudgeRepositoryGetAllJudges {
		rating := postgres_init.CreateRating(&fields)

		createdJudges := []models.Judge{
			{
				ID:       uuid.New(),
				FIO:      "Test",
				Login:    "Te1st",
				Password: "tes2t123",
				Post:     "Tes1	t",
				Role:     1,
			},
			{
				ID:       uuid.New(),
				FIO:      "Test1",
				Login:    "Test",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
			{
				ID:       uuid.New(),
				FIO:      "Test2",
				Login:    "Tes222t",
				Password: "test123",
				Post:     "Test",
				Role:     1,
			},
		}

		for i, _ := range createdJudges {
			j, err := judgeRepository.CreateProfile(&createdJudges[i])
			require.NoError(t, err)
			postgres_init.AttachJudgeToRating(&fields, j.ID, rating.ID)
		}

		receivedJudges, err := judgeRepository.GetAllJudges()
		test.CheckOutput(t, createdJudges, receivedJudges, err)
	}
}
