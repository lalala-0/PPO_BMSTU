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
	"time"
)

//DetachCrewFromProtest(crewID uuid.UUID, protestID uuid.UUID) error

var testProtestRepositoryCreateSuccess = []struct {
	TestName    string
	InputData   *models.Protest
	CheckOutput func(t *testing.T, inputData *models.Protest, createdProtest *models.Protest, err error)
}{
	{
		TestName: "create success test",
		InputData: &models.Protest{
			ID:         uuid.New(),
			RuleNum:    23,
			ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			Status:     1,
			Comment:    "",
		},
		CheckOutput: func(t *testing.T, inputData *models.Protest, createdProtest *models.Protest, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.RaceID, createdProtest.RaceID)
			require.Equal(t, inputData.JudgeID, createdProtest.JudgeID)
			require.Equal(t, inputData.RatingID, createdProtest.RatingID)
			require.Equal(t, inputData.RuleNum, createdProtest.RuleNum)
			require.Equal(t, inputData.ReviewDate, createdProtest.ReviewDate)
			require.Equal(t, inputData.Status, createdProtest.Status)
			require.Equal(t, inputData.Comment, createdProtest.Comment)
		},
	},
}

func TestProtestRepositoryCreate(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}

	for _, test := range testProtestRepositoryCreateSuccess {
		protestRepository := postgres.CreateProtestRepository(&fields)
		t.Run(test.TestName, func(t *testing.T) {
			rating := postgres_init.CreateRating(&fields)
			race := postgres_init.CreateRace(&fields, rating.ID)
			judge := postgres_init.CreateJudge(&fields)

			test.InputData.RaceID = race.ID
			test.InputData.JudgeID = judge.ID
			test.InputData.RatingID = rating.ID

			createdProtest, err := protestRepository.Create(test.InputData)
			test.CheckOutput(t, test.InputData, createdProtest, err)
		})
	}
}

var testProtestRepositoryGetByIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdProtest *models.Protest, receivedProtest *models.Protest, err error)
}{
	{
		TestName: "get by id success test",
		CheckOutput: func(t *testing.T, createdProtest *models.Protest, receivedProtest *models.Protest, err error) {
			require.NoError(t, err)
			require.Equal(t, createdProtest.ID, receivedProtest.ID)
		},
	},
}

func TestProtestRepositoryGetByID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	protestRepository := postgres.CreateProtestRepository(&fields)

	for _, test := range testProtestRepositoryGetByIDSuccess {
		rating := postgres_init.CreateRating(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		judge := postgres_init.CreateJudge(&fields)
		protest := postgres_init.CreateProtest(&fields, race.ID, judge.ID, rating.ID)

		receivedProtest, err := protestRepository.GetProtestDataByID(protest.ID)
		test.CheckOutput(t, protest, receivedProtest, err)
	}
}

var testProtestRepositoryDeleteSuccess = []struct {
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

func TestProtestRepositoryDelete(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	protestRepository := postgres.CreateProtestRepository(&fields)

	for _, test := range testProtestRepositoryDeleteSuccess {
		rating := postgres_init.CreateRating(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		judge := postgres_init.CreateJudge(&fields)
		protest := postgres_init.CreateProtest(&fields, race.ID, judge.ID, rating.ID)

		err := protestRepository.Delete(protest.ID)
		test.CheckOutput(t, err)

		_, err = protestRepository.GetProtestDataByID(protest.ID)
		require.Error(t, err)
	}
}

var testProtestRepositoryUpdateSuccess = []struct {
	TestName    string
	InputData   *models.Protest
	CheckOutput func(t *testing.T, createdProtest *models.Protest, updatedProtest *models.Protest, err error)
}{
	{
		TestName: "update success test",
		InputData: &models.Protest{
			ID:         uuid.New(),
			RuleNum:    23,
			ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
			Status:     1,
			Comment:    "",
		},
		CheckOutput: func(t *testing.T, inputData *models.Protest, createdProtest *models.Protest, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.RaceID, createdProtest.RaceID)
			require.Equal(t, inputData.JudgeID, createdProtest.JudgeID)
			require.Equal(t, inputData.RatingID, createdProtest.RatingID)
			require.NotEqual(t, inputData.RuleNum, createdProtest.RuleNum)
			require.NotEqual(t, inputData.ReviewDate, createdProtest.ReviewDate)
			require.NotEqual(t, inputData.Status, createdProtest.Status)
			require.NotEqual(t, inputData.Comment, createdProtest.Comment)
		},
	},
}

func TestProtestRepositoryUpdate(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	protestRepository := postgres.CreateProtestRepository(&fields)

	for _, test := range testProtestRepositoryUpdateSuccess {
		rating := postgres_init.CreateRating(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		judge := postgres_init.CreateJudge(&fields)

		test.InputData.RaceID = race.ID
		test.InputData.JudgeID = judge.ID
		test.InputData.RatingID = rating.ID

		createdProtest, _ := protestRepository.Create(test.InputData)

		updatedProtest, err := protestRepository.Update(
			&models.Protest{
				ID:         createdProtest.ID,
				RaceID:     race.ID,
				RatingID:   rating.ID,
				JudgeID:    judge.ID,
				RuleNum:    22,
				ReviewDate: time.Date(2023, time.November, 10, 10, 0, 0, 0, time.UTC),
				Status:     2,
				Comment:    "Test",
			},
		)

		test.CheckOutput(t, createdProtest, updatedProtest, err)
	}
}

//GetProtestsDataByRaceID(raceID uuid.UUID) ([]models.Protest, error)

var testProtestRepositoryGetProtestsDataByRaceID = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdProtests []models.Protest, receivedProtests []models.Protest, err error)
}{
	{
		TestName: "Get Protests Data By Race ID success test",
		CheckOutput: func(t *testing.T, createdProtests []models.Protest, receivedProtests []models.Protest, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdProtests), len(receivedProtests))
		},
	},
}

func TestProtestRepositoryGetProtestsDataByRaceID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	protestRepository := postgres.CreateProtestRepository(&fields)

	for _, test := range testProtestRepositoryGetProtestsDataByRaceID {
		rating := postgres_init.CreateRating(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		judge := postgres_init.CreateJudge(&fields)

		createdProtests := []models.Protest{
			{
				RaceID:     race.ID,
				RatingID:   rating.ID,
				JudgeID:    judge.ID,
				RuleNum:    23,
				ReviewDate: time.Date(2024, time.November, 10, 9, 0, 0, 0, time.UTC),
				Status:     1,
				Comment:    "",
			},
			{
				RaceID:     race.ID,
				RatingID:   rating.ID,
				JudgeID:    judge.ID,
				RuleNum:    23,
				ReviewDate: time.Date(2024, time.November, 10, 10, 0, 0, 0, time.UTC),
				Status:     1,
				Comment:    "",
			},
			{
				RaceID:     race.ID,
				RatingID:   rating.ID,
				JudgeID:    judge.ID,
				RuleNum:    23,
				ReviewDate: time.Date(2024, time.November, 10, 11, 0, 0, 0, time.UTC),
				Status:     1,
				Comment:    "",
			},
		}

		for i, _ := range createdProtests {
			_, err := protestRepository.Create(&createdProtests[i])
			require.NoError(t, err)
		}

		receivedProtests, err := protestRepository.GetProtestsDataByRaceID(race.ID)
		test.CheckOutput(t, createdProtests, receivedProtests, err)
	}
}

// //GetProtestParticipantsIDByID(id uuid.UUID) (map[uuid.UUID]int, error)
var testProtestRepositoryGetProtestParticipantsIDByID = []struct {
	TestName    string
	InputData   map[int]int
	CheckOutput func(t *testing.T, createdProtestParticipants map[int]int, receivedProtestParticipants map[uuid.UUID]int, err error)
}{
	{
		TestName: "Get the crews participating in the protest By Protest ID success test",
		InputData: map[int]int{
			1: models.Protestor,
			2: models.Protestee,
			3: models.Witness,
			4: models.Witness,
			5: models.Witness,
		},
		CheckOutput: func(t *testing.T, createdProtestParticipants map[int]int, receivedProtestParticipants map[uuid.UUID]int, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdProtestParticipants), len(receivedProtestParticipants))
		},
	},
}

func TestProtestRepositoryGetProtestParticipantsIDByID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	protestRepository := postgres.CreateProtestRepository(&fields)
	crewRepository := postgres.CreateCrewRepository(&fields)

	for _, test := range testProtestRepositoryGetProtestParticipantsIDByID {
		rating := postgres_init.CreateRating(&fields)
		judge := postgres_init.CreateJudge(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		createdCrews := []models.Crew{
			{
				RatingID: rating.ID,
				SailNum:  1,
				Class:    models.Laser,
			},
			{
				RatingID: rating.ID,
				SailNum:  12,
				Class:    models.Laser,
			},
			{
				RatingID: rating.ID,
				SailNum:  2,
				Class:    models.Laser,
			},
			{
				RatingID: rating.ID,
				SailNum:  31,
				Class:    models.Laser,
			},
			{
				RatingID: rating.ID,
				SailNum:  13,
				Class:    models.Laser,
			},
		}
		protest := postgres_init.CreateProtest(&fields, race.ID, judge.ID, rating.ID)

		for i, _ := range createdCrews {
			c, err := crewRepository.Create(&createdCrews[i])
			require.NoError(t, err)
			postgres_init.AttachCrewToProtestStatus(&fields, c.ID, protest.ID, test.InputData[i])

		}

		receivedMap, err := protestRepository.GetProtestParticipantsIDByID(protest.ID)
		test.CheckOutput(t, test.InputData, receivedMap, err)
	}
}

//AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error

var testProtestRepositoryAttachCrewToProtest = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "Attach Crew To Protest success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestProtestRepositoryAttachCrewToProtest(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	protestRepository := postgres.CreateProtestRepository(&fields)

	for _, test := range testProtestRepositoryAttachCrewToProtest {
		rating := postgres_init.CreateRating(&fields)
		judge := postgres_init.CreateJudge(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		crew := postgres_init.CreateCrew(&fields, rating.ID)
		protest := postgres_init.CreateProtest(&fields, race.ID, judge.ID, rating.ID)

		err := protestRepository.AttachCrewToProtest(crew.ID, protest.ID, 1)
		test.CheckOutput(t, err)
	}
}

//AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error

var testProtestRepositoryDetachCrewFromProtest = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "Detach Crew from Protest success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestProtestRepositoryDetachCrewFromProtest(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	protestRepository := postgres.CreateProtestRepository(&fields)

	for _, test := range testProtestRepositoryDetachCrewFromProtest {
		rating := postgres_init.CreateRating(&fields)
		judge := postgres_init.CreateJudge(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		crew := postgres_init.CreateCrew(&fields, rating.ID)
		protest := postgres_init.CreateProtest(&fields, race.ID, judge.ID, rating.ID)

		err := protestRepository.AttachCrewToProtest(crew.ID, protest.ID, 1)
		require.NoError(t, err)
		err = protestRepository.DetachCrewFromProtest(crew.ID, protest.ID)
		test.CheckOutput(t, err)
	}
}
