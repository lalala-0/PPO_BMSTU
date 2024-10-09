package mongo_init

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/mongo"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"
)

var testRaceRepositoryCreateSuccess = []struct {
	TestName    string
	InputData   *models.Race
	CheckOutput func(t *testing.T, inputData *models.Race, createdRace *models.Race, err error)
}{
	{
		TestName: "create success test",
		InputData: &models.Race{
			RatingID: uuid.New(),
			Date:     time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
			Number:   1,
			Class:    4,
		},
		CheckOutput: func(t *testing.T, inputData *models.Race, createdRace *models.Race, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.RatingID, createdRace.RatingID)
			require.Equal(t, inputData.Date, createdRace.Date)
			require.Equal(t, inputData.Number, createdRace.Number)
			require.Equal(t, inputData.Class, createdRace.Class)
		},
	},
}

func TestRaceRepositoryCreate(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}

	for _, test := range testRaceRepositoryCreateSuccess {
		raceRepository := mongo.CreateRaceRepository(&fields)
		t.Run(test.TestName, func(t *testing.T) {
			rating := CreateRating(&fields)
			test.InputData.RatingID = rating.ID
			createdRace, err := raceRepository.Create(test.InputData)
			test.CheckOutput(t, test.InputData, createdRace, err)
		})
	}
}

var testRaceRepositoryGetRaceByIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdRace *models.Race, receivedRace *models.Race, err error)
}{
	{
		TestName: "Get Race By ID success test",
		CheckOutput: func(t *testing.T, createdRace *models.Race, receivedRace *models.Race, err error) {
			require.NoError(t, err)
			require.Equal(t, receivedRace.RatingID, createdRace.RatingID)
			require.Equal(t, receivedRace.Date, createdRace.Date)
			require.Equal(t, receivedRace.Number, createdRace.Number)
			require.Equal(t, receivedRace.Class, createdRace.Class)
		},
	},
}

func TestRaceRepositoryGetByID(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	raceRepository := mongo.CreateRaceRepository(&fields)

	for _, test := range testRaceRepositoryGetRaceByIDSuccess {
		rating := CreateRating(&fields)
		race := CreateRace(&fields, rating.ID)

		receivedRace, err := raceRepository.GetRaceDataByID(race.ID)
		test.CheckOutput(t, race, receivedRace, err)
	}
}

var testRaceRepositoryDeleteSuccess = []struct {
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

func TestRaceRepositoryDelete(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	raceRepository := mongo.CreateRaceRepository(&fields)

	for _, test := range testRaceRepositoryDeleteSuccess {
		rating := CreateRating(&fields)
		race := CreateRace(&fields, rating.ID)

		err := raceRepository.Delete(race.ID)
		test.CheckOutput(t, err)

		_, err = raceRepository.GetRaceDataByID(race.ID)
		require.Error(t, err)
	}
}

var testRaceRepositoryUpdateSuccess = []struct {
	TestName  string
	InputData struct {
		Race *models.Race
	}
	CheckOutput func(t *testing.T, createdRace *models.Race, updatedRace *models.Race, err error)
}{
	{
		TestName: "update success test",
		InputData: struct {
			Race *models.Race
		}{
			&models.Race{
				RatingID: uuid.New(),
				Date:     time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
				Number:   1,
				Class:    4,
			},
		},
		CheckOutput: func(t *testing.T, createdRace *models.Race, updatedRace *models.Race, err error) {
			require.NoError(t, err)
			require.Equal(t, updatedRace.RatingID, createdRace.RatingID)
			require.NotEqual(t, updatedRace.Date, createdRace.Date)
			require.NotEqual(t, updatedRace.Number, createdRace.Number)
			require.NotEqual(t, updatedRace.Class, createdRace.Class)
		},
	},
}

func TestRaceRepositoryUpdate(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	raceRepository := mongo.CreateRaceRepository(&fields)

	for _, test := range testRaceRepositoryUpdateSuccess {
		rating := CreateRating(&fields)
		race := CreateRace(&fields, rating.ID)

		updatedRace, err := raceRepository.Update(
			&models.Race{
				ID:       race.ID,
				RatingID: rating.ID,
				Date:     time.Date(2012, time.November, 11, 23, 0, 0, 0, time.UTC),
				Number:   2,
				Class:    5,
			},
		)

		test.CheckOutput(t, race, updatedRace, err)
	}
}

//GetRacesDataByRatingID(rating_id uuid.UUID) ([]models.Race, error)

var testRaceRepositoryGetRacesDataByRatingID = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdRaces []models.Race, receivedRaces []models.Race, err error)
}{
	{
		TestName: "Get Races Data By Rating ID success test",
		CheckOutput: func(t *testing.T, createdRaces []models.Race, receivedRaces []models.Race, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdRaces), len(receivedRaces))
		},
	},
}

func TestRaceRepositoryGetRacesDataByRatingID(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	raceRepository := mongo.CreateRaceRepository(&fields)

	for _, test := range testRaceRepositoryGetRacesDataByRatingID {
		rating := CreateRating(&fields)

		createdRaces := []models.Race{
			{
				RatingID: rating.ID,
				Date:     time.Date(2012, time.November, 11, 10, 0, 0, 0, time.UTC),
				Number:   2,
				Class:    5,
			},
			{
				RatingID: rating.ID,
				Date:     time.Date(2012, time.November, 11, 13, 0, 0, 0, time.UTC),
				Number:   3,
				Class:    5,
			},
			{
				RatingID: rating.ID,
				Date:     time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
				Number:   1,
				Class:    5,
			},
		}

		for i, _ := range createdRaces {
			_, err := raceRepository.Create(&createdRaces[i])
			require.NoError(t, err)
		}

		receivedRaces, err := raceRepository.GetRacesDataByRatingID(rating.ID)
		test.CheckOutput(t, createdRaces, receivedRaces, err)
	}
}
