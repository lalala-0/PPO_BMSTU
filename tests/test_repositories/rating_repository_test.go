package test_repositories

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

var testRatingRepositoryCreateSuccess = []struct {
	TestName    string
	InputData   *models.Rating
	CheckOutput func(t *testing.T, inputData *models.Rating, createdRating *models.Rating, err error)
}{
	{
		TestName: "create success test",
		InputData: &models.Rating{
			ID:         uuid.New(),
			Name:       "Name",
			Class:      models.Laser,
			BlowoutCnt: 1,
		},
		CheckOutput: func(t *testing.T, inputData *models.Rating, createdRating *models.Rating, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.Name, createdRating.Name)
			require.Equal(t, inputData.Class, createdRating.Class)
			require.Equal(t, inputData.BlowoutCnt, createdRating.BlowoutCnt)
		},
	},
}

func TestRatingRepositoryCreate(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := repository.PostgresConnection{DB: db}

	for _, test := range testRatingRepositoryCreateSuccess {
		ratingRepository := repository.CreateRatingRepository(&fields)
		t.Run(test.TestName, func(t *testing.T) {

			createdRating, err := ratingRepository.Create(test.InputData)
			test.CheckOutput(t, test.InputData, createdRating, err)
		})
	}
}

var testRatingRepositoryGetByIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdRating *models.Rating, receivedRating *models.Rating, err error)
}{
	{
		TestName: "get by id success test",
		CheckOutput: func(t *testing.T, createdRating *models.Rating, receivedRating *models.Rating, err error) {
			require.NoError(t, err)
			require.Equal(t, createdRating.ID, receivedRating.ID)
		},
	},
}

func TestRatingRepositoryGetByID(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := repository.PostgresConnection{DB: db}
	ratingRepository := repository.CreateRatingRepository(&fields)

	for _, test := range testRatingRepositoryGetByIDSuccess {
		rating := createRating(&fields)

		receivedRating, err := ratingRepository.GetRatingDataByID(rating.ID)
		test.CheckOutput(t, rating, receivedRating, err)
	}
}

var testRatingRepositoryDeleteSuccess = []struct {
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

func TestRatingRepositoryDelete(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := repository.PostgresConnection{DB: db}
	ratingRepository := repository.CreateRatingRepository(&fields)

	for _, test := range testRatingRepositoryDeleteSuccess {
		rating := createRating(&fields)

		err := ratingRepository.Delete(rating.ID)
		test.CheckOutput(t, err)

		_, err = ratingRepository.GetRatingDataByID(rating.ID)
		require.Error(t, err)
	}
}

var testRatingRepositoryUpdateSuccess = []struct {
	TestName    string
	InputData   *models.Rating
	CheckOutput func(t *testing.T, createdRating *models.Rating, updatedRating *models.Rating, err error)
}{
	{
		TestName: "update success test",
		InputData: &models.Rating{
			ID:         uuid.New(),
			Name:       "Name",
			Class:      models.Laser,
			BlowoutCnt: 1,
		},
		CheckOutput: func(t *testing.T, inputData *models.Rating, createdRating *models.Rating, err error) {
			require.NoError(t, err)
			require.NotEqual(t, inputData.Name, createdRating.Name)
			require.NotEqual(t, inputData.Class, createdRating.Class)
			require.NotEqual(t, inputData.BlowoutCnt, createdRating.BlowoutCnt)
		},
	},
}

func TestRatingRepositoryUpdate(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := repository.PostgresConnection{DB: db}
	ratingRepository := repository.CreateRatingRepository(&fields)

	for _, test := range testRatingRepositoryUpdateSuccess {
		rating := createRating(&fields)

		createdRating, err := ratingRepository.Create(test.InputData)

		updatedRating, err := ratingRepository.Update(
			&models.Rating{
				ID:         rating.ID,
				Name:       "Name1",
				Class:      models.LaserRadial,
				BlowoutCnt: 2,
			},
		)

		test.CheckOutput(t, createdRating, updatedRating, err)
	}
}

//AttachCrewToRating(crewID uuid.UUID, ratingID uuid.UUID, crewStatus int) error

var testRatingRepositoryAttachJudgeToRating = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "Attach Judge To Rating success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestRatingRepositoryAttachJudgeToRating(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := repository.PostgresConnection{DB: db}
	ratingRepository := repository.CreateRatingRepository(&fields)

	for _, test := range testRatingRepositoryAttachJudgeToRating {
		rating := createRating(&fields)
		judge := createJudge(&fields)

		err := ratingRepository.AttachJudgeToRating(rating.ID, judge.ID)
		test.CheckOutput(t, err)
	}
}

var testRatingRepositoryDetachJudgeFromRating = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "Detach Judge from Rating success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestRatingRepositoryDetachJudgeFromRating(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := repository.PostgresConnection{DB: db}
	ratingRepository := repository.CreateRatingRepository(&fields)

	for _, test := range testRatingRepositoryDetachJudgeFromRating {
		rating := createRating(&fields)
		judge := createJudge(&fields)

		err := ratingRepository.AttachJudgeToRating(rating.ID, judge.ID)
		require.NoError(t, err)
		err = ratingRepository.DetachJudgeFromRating(rating.ID, judge.ID)
		test.CheckOutput(t, err)
	}
}

var testRatingRepositoryGetAllRatings = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdRatings []models.Rating, receivedRatings []models.Rating, err error)
}{
	{
		TestName: "Get all Ratings success test",
		CheckOutput: func(t *testing.T, createdRatings []models.Rating, receivedRatings []models.Rating, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdRatings), len(receivedRatings))
		},
	},
}

func TestRatingRepositoryGetAllRatings(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := repository.PostgresConnection{DB: db}
	ratingRepository := repository.CreateRatingRepository(&fields)

	for _, test := range testRatingRepositoryGetAllRatings {
		createdRatings := []models.Rating{
			{
				Name:       "Name1",
				Class:      models.LaserRadial,
				BlowoutCnt: 2,
			},
			{
				Name:       "Name2",
				Class:      models.Laser,
				BlowoutCnt: 2,
			},
			{
				Name:       "Name3",
				Class:      models.Zoom8,
				BlowoutCnt: 2,
			},
		}

		for i, _ := range createdRatings {
			_, err := ratingRepository.Create(&createdRatings[i])
			require.NoError(t, err)
		}

		receivedRatings, err := ratingRepository.GetAllRatings()
		test.CheckOutput(t, createdRatings, receivedRatings, err)
	}
}
