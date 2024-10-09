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

var testParticipantRepositoryCreateSuccess = []struct {
	TestName    string
	InputData   *models.Participant
	CheckOutput func(t *testing.T, inputData *models.Participant, createdParticipant *models.Participant, err error)
}{
	{
		TestName: "create success test",
		InputData: &models.Participant{
			ID:       uuid.New(),
			FIO:      "test",
			Gender:   models.Female,
			Category: models.Junior2category,
			Coach:    "Test",
			Birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		CheckOutput: func(t *testing.T, inputData *models.Participant, createdParticipant *models.Participant, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.FIO, createdParticipant.FIO)
			require.Equal(t, inputData.Gender, createdParticipant.Gender)
			require.Equal(t, inputData.Category, createdParticipant.Category)
			require.Equal(t, inputData.Coach, createdParticipant.Coach)
			require.Equal(t, inputData.Birthday, createdParticipant.Birthday)
		},
	},
}

func TestParticipantRepositoryCreate(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}

	for _, test := range testParticipantRepositoryCreateSuccess {
		participantRepository := postgres.CreateParticipantRepository(&fields)
		t.Run(test.TestName, func(t *testing.T) {

			createdParticipant, err := participantRepository.Create(test.InputData)
			test.CheckOutput(t, test.InputData, createdParticipant, err)
		})
	}
}

var testParticipantRepositoryGetByIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdParticipant *models.Participant, receivedParticipant *models.Participant, err error)
}{
	{
		TestName: "get by id success test",
		CheckOutput: func(t *testing.T, createdParticipant *models.Participant, receivedParticipant *models.Participant, err error) {
			require.NoError(t, err)
			require.Equal(t, createdParticipant.ID, receivedParticipant.ID)
		},
	},
}

func TestParticipantRepositoryGetByID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	participantRepository := postgres.CreateParticipantRepository(&fields)

	for _, test := range testParticipantRepositoryGetByIDSuccess {
		participant := postgres_init.CreateParticipant(&fields)

		receivedParticipant, err := participantRepository.GetParticipantDataByID(participant.ID)
		test.CheckOutput(t, participant, receivedParticipant, err)
	}
}

var testParticipantRepositoryDeleteSuccess = []struct {
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

func TestParticipantRepositoryDelete(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	participantRepository := postgres.CreateParticipantRepository(&fields)

	for _, test := range testParticipantRepositoryDeleteSuccess {
		participant := postgres_init.CreateParticipant(&fields)

		err := participantRepository.Delete(participant.ID)
		test.CheckOutput(t, err)

		_, err = participantRepository.GetParticipantDataByID(participant.ID)
		require.Error(t, err)
	}
}

var testParticipantRepositoryUpdateSuccess = []struct {
	TestName    string
	InputData   *models.Participant
	CheckOutput func(t *testing.T, createdParticipant *models.Participant, updatedParticipant *models.Participant, err error)
}{
	{
		TestName: "update success test",
		InputData: &models.Participant{
			ID:       uuid.New(),
			FIO:      "test1",
			Gender:   models.Male,
			Category: models.Junior1category,
			Coach:    "Test1",
			Birthday: time.Date(2008, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		CheckOutput: func(t *testing.T, createdParticipant *models.Participant, updatedParticipant *models.Participant, err error) {
			require.NoError(t, err)
			require.NotEqual(t, updatedParticipant.FIO, createdParticipant.FIO)
			require.NotEqual(t, updatedParticipant.Gender, createdParticipant.Gender)
			require.NotEqual(t, updatedParticipant.Category, createdParticipant.Category)
			require.NotEqual(t, updatedParticipant.Coach, createdParticipant.Coach)
			require.NotEqual(t, updatedParticipant.Birthday, createdParticipant.Birthday)
		},
	},
}

func TestParticipantRepositoryUpdate(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	participantRepository := postgres.CreateParticipantRepository(&fields)

	for _, test := range testParticipantRepositoryUpdateSuccess {
		createdParticipant, err := participantRepository.Create(test.InputData)

		updatedParticipant, err := participantRepository.Update(
			postgres_init.CreateParticipant(&fields),
		)

		test.CheckOutput(t, createdParticipant, updatedParticipant, err)
	}
}

//GetParticipantsDataByCrewID(crewID uuid.UUID) ([]models.Participant, error)

var testParticipantRepositoryGetParticipantsDataByCrewID = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdParticipants []models.Participant, receivedParticipants []models.Participant, err error)
}{
	{
		TestName: "Get Participants Data By Crew ID success test",
		CheckOutput: func(t *testing.T, createdParticipants []models.Participant, receivedParticipants []models.Participant, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdParticipants), len(receivedParticipants))
		},
	},
}

func TestParticipantRepositoryGetParticipantsDataByCrewID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	participantRepository := postgres.CreateParticipantRepository(&fields)

	for _, test := range testParticipantRepositoryGetParticipantsDataByCrewID {
		createdParticipants := []models.Participant{
			{
				FIO:      "test1",
				Gender:   models.Male,
				Category: models.Junior1category,
				Coach:    "Test1",
				Birthday: time.Date(2008, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			{
				FIO:      "test12",
				Gender:   models.Female,
				Category: models.Junior2category,
				Coach:    "Test12",
				Birthday: time.Date(2006, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			{
				FIO:      "test13",
				Gender:   models.Male,
				Category: models.Junior2category,
				Coach:    "Test13",
				Birthday: time.Date(2007, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
		}
		rating := postgres_init.CreateRating(&fields)
		crew := postgres_init.CreateCrew(&fields, rating.ID)
		for i, _ := range createdParticipants {
			participant, err := participantRepository.Create(&createdParticipants[i])
			require.NoError(t, err)
			postgres_init.AttachParticipantToCrew(&fields, participant.ID, crew.ID)
		}

		receivedParticipants, err := participantRepository.GetParticipantsDataByCrewID(crew.ID)
		test.CheckOutput(t, createdParticipants, receivedParticipants, err)
	}
}

//GetParticipantsDataByProtestID(crewID uuid.UUID) ([]models.Participant, error)

var testParticipantRepositoryGetParticipantsDataByProtestID = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdParticipants []models.Participant, receivedParticipants []models.Participant, err error)
}{
	{
		TestName: "Get Participants Data By Crew ID success test",
		CheckOutput: func(t *testing.T, createdParticipants []models.Participant, receivedParticipants []models.Participant, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdParticipants), len(receivedParticipants))
		},
	},
}

func TestParticipantRepositoryGetParticipantsDataByProtestID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	participantRepository := postgres.CreateParticipantRepository(&fields)

	for _, test := range testParticipantRepositoryGetParticipantsDataByProtestID {
		rating := postgres_init.CreateRating(&fields)
		race := postgres_init.CreateRace(&fields, rating.ID)
		judge := postgres_init.CreateJudge(&fields)
		protest := postgres_init.CreateProtest(&fields, race.ID, judge.ID, rating.ID)
		crew := postgres_init.CreateCrew(&fields, rating.ID)
		postgres_init.AttachCrewToProtest(&fields, crew.ID, protest.ID)

		createdParticipants := []models.Participant{
			{
				FIO:      "test1",
				Gender:   models.Male,
				Category: models.Junior1category,
				Coach:    "Test1",
				Birthday: time.Date(2008, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			{
				FIO:      "test12",
				Gender:   models.Female,
				Category: models.Junior2category,
				Coach:    "Test12",
				Birthday: time.Date(2006, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			{
				FIO:      "test13",
				Gender:   models.Male,
				Category: models.Junior2category,
				Coach:    "Test13",
				Birthday: time.Date(2007, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
		}
		for i, _ := range createdParticipants {
			p, err := participantRepository.Create(&createdParticipants[i])
			require.NoError(t, err)
			postgres_init.AttachParticipantToCrew(&fields, p.ID, crew.ID)
		}

		receivedParticipants, err := participantRepository.GetParticipantsDataByProtestID(protest.ID)
		test.CheckOutput(t, createdParticipants, receivedParticipants, err)
	}
}
