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

var testCrewResInRaceRepositoryCreateSuccess = []struct {
	TestName    string
	InputData   *models.CrewResInRace
	CheckOutput func(t *testing.T, inputData *models.CrewResInRace, createdCrewResInRace *models.CrewResInRace, err error)
}{
	{
		TestName: "create success test",
		InputData: &models.CrewResInRace{
			CrewID:           uuid.New(),
			RaceID:           uuid.New(),
			Points:           12,
			SpecCircumstance: 0,
		},
		CheckOutput: func(t *testing.T, inputData *models.CrewResInRace, createdCrewResInRace *models.CrewResInRace, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.CrewID, createdCrewResInRace.CrewID)
			require.Equal(t, inputData.RaceID, createdCrewResInRace.RaceID)
			require.Equal(t, inputData.Points, createdCrewResInRace.Points)
			require.Equal(t, inputData.SpecCircumstance, createdCrewResInRace.SpecCircumstance)
		},
	},
}

func TestCrewResInRaceRepositoryCreate(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}

	for _, test := range testCrewResInRaceRepositoryCreateSuccess {
		crewResInRaceRepository := postgres.CreateCrewResInRaceRepository(&fields)
		t.Run(test.TestName, func(t *testing.T) {
			rating := postgres_init.CreateRating(&fields)
			crew := postgres_init.CreateCrew(&fields, rating.ID)
			race := postgres_init.CreateRace(&fields, rating.ID)
			test.InputData.CrewID = crew.ID
			test.InputData.RaceID = race.ID
			createdCrewResInRace, err := crewResInRaceRepository.Create(test.InputData)
			test.CheckOutput(t, test.InputData, createdCrewResInRace, err)
		})
	}
}

var testCrewResInRaceRepositoryGetCrewResByRaceIDAndCrewIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdCrewResInRace *models.CrewResInRace, receivedCrewResInRace *models.CrewResInRace, err error)
}{
	{
		TestName: "Get Crew Res By Race ID And Crew ID success test",
		CheckOutput: func(t *testing.T, createdCrewResInRace *models.CrewResInRace, receivedCrewResInRace *models.CrewResInRace, err error) {
			require.NoError(t, err)
			require.Equal(t, createdCrewResInRace.RaceID, receivedCrewResInRace.RaceID)
			require.Equal(t, createdCrewResInRace.CrewID, receivedCrewResInRace.CrewID)
		},
	},
}

func TestCrewResInRaceRepositoryGetByID(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	crewResInRaceRepository := postgres.CreateCrewResInRaceRepository(&fields)

	for _, test := range testCrewResInRaceRepositoryGetCrewResByRaceIDAndCrewIDSuccess {
		rating := postgres_init.CreateRating(&fields)
		crew := postgres_init.CreateCrew(&fields, rating.ID)
		race := postgres_init.CreateRace(&fields, rating.ID)
		crewResInRace := postgres_init.CreateCrewResInRace(&fields, crew.ID, race.ID)

		receivedCrewResInRace, err := crewResInRaceRepository.GetCrewResByRaceIDAndCrewID(race.ID, crew.ID)
		test.CheckOutput(t, crewResInRace, receivedCrewResInRace, err)
	}
}

var testCrewResInRaceRepositoryDeleteSuccess = []struct {
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

func TestCrewResInRaceRepositoryDelete(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	crewResInRaceRepository := postgres.CreateCrewResInRaceRepository(&fields)

	for _, test := range testCrewResInRaceRepositoryDeleteSuccess {
		rating := postgres_init.CreateRating(&fields)
		crew := postgres_init.CreateCrew(&fields, rating.ID)
		race := postgres_init.CreateRace(&fields, rating.ID)
		crewResInRace := postgres_init.CreateCrewResInRace(&fields, crew.ID, race.ID)

		err := crewResInRaceRepository.Delete(crewResInRace.RaceID, crewResInRace.CrewID)
		test.CheckOutput(t, err)

		_, err = crewResInRaceRepository.GetCrewResByRaceIDAndCrewID(crewResInRace.RaceID, crewResInRace.CrewID)
		require.Error(t, err)
	}
}

// Update(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error)

var testCrewResInRaceRepositoryUpdateSuccess = []struct {
	TestName  string
	InputData struct {
		CrewResInRace *models.CrewResInRace
	}
	CheckOutput func(t *testing.T, createdCrewResInRace *models.CrewResInRace, updatedCrewResInRace *models.CrewResInRace, err error)
}{
	{
		TestName: "update success test",
		InputData: struct {
			CrewResInRace *models.CrewResInRace
		}{
			&models.CrewResInRace{
				CrewID:           uuid.New(),
				RaceID:           uuid.New(),
				Points:           12,
				SpecCircumstance: 0,
			},
		},
		CheckOutput: func(t *testing.T, createdCrewResInRace *models.CrewResInRace, updatedCrewResInRace *models.CrewResInRace, err error) {
			require.NoError(t, err)
			require.Equal(t, createdCrewResInRace.CrewID, updatedCrewResInRace.CrewID)
			require.Equal(t, updatedCrewResInRace.RaceID, createdCrewResInRace.RaceID)
			require.NotEqual(t, updatedCrewResInRace.Points, createdCrewResInRace.Points)
			require.NotEqual(t, updatedCrewResInRace.SpecCircumstance, createdCrewResInRace.SpecCircumstance)
		},
	},
}

func TestCrewResInRaceRepositoryUpdate(t *testing.T) {
	dbContainer, db := postgres_init.SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := postgres.PostgresConnection{DB: db}
	crewResInRaceRepository := postgres.CreateCrewResInRaceRepository(&fields)

	for _, test := range testCrewResInRaceRepositoryUpdateSuccess {
		rating := postgres_init.CreateRating(&fields)
		crew := postgres_init.CreateCrew(&fields, rating.ID)
		race := postgres_init.CreateRace(&fields, rating.ID)
		crewResInRace := postgres_init.CreateCrewResInRace(&fields, crew.ID, race.ID)

		updatedCrewResInRace, err := crewResInRaceRepository.Update(
			&models.CrewResInRace{
				CrewID:           crew.ID,
				RaceID:           race.ID,
				Points:           135,
				SpecCircumstance: 1,
			},
		)

		test.CheckOutput(t, crewResInRace, updatedCrewResInRace, err)
	}
}
