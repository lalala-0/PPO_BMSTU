package mongo_init

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/mongo"
	"PPO_BMSTU/internal/repository/repository_errors"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

var testCrewRepositoryCreateSuccess = []struct {
	TestName    string
	InputData   *models.Crew
	CheckOutput func(t *testing.T, inputData *models.Crew, createdCrew *models.Crew, err error)
}{
	{
		TestName: "create success test",
		InputData: &models.Crew{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			SailNum:  123,
			Class:    2,
		},
		CheckOutput: func(t *testing.T, inputData *models.Crew, createdCrew *models.Crew, err error) {
			require.NoError(t, err)
			require.Equal(t, inputData.RatingID, createdCrew.RatingID)
			require.Equal(t, inputData.SailNum, createdCrew.SailNum)
			require.Equal(t, inputData.Class, createdCrew.Class)
		},
	},
}

func TestCrewRepositoryCreate(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}

	for _, test := range testCrewRepositoryCreateSuccess {
		crewRepository := mongo.CreateCrewRepository(&fields)
		t.Run(test.TestName, func(t *testing.T) {
			rating := CreateRating(&fields)

			test.InputData.RatingID = rating.ID

			createdCrew, err := crewRepository.Create(test.InputData)
			test.CheckOutput(t, test.InputData, createdCrew, err)
		})
	}
}

var testCrewRepositoryGetByIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdCrew *models.Crew, receivedCrew *models.Crew, err error)
}{
	{
		TestName: "get by id success test",
		CheckOutput: func(t *testing.T, createdCrew *models.Crew, receivedCrew *models.Crew, err error) {
			require.NoError(t, err)
			require.Equal(t, createdCrew.ID, receivedCrew.ID)
		},
	},
}

var testCrewRepositoryGetByIDFailed = []struct {
	TestName    string
	InputData   uuid.UUID
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName:  "get by id DoesNotExist test",
		InputData: uuid.New(),
		CheckOutput: func(t *testing.T, err error) {
			require.Error(t, err)
			require.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
}

func TestCrewRepositoryGetByID(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryGetByIDSuccess {
		rating := CreateRating(&fields)
		crew := CreateCrew(&fields, rating.ID)

		receivedCrew, err := crewRepository.GetCrewDataByID(crew.ID)
		test.CheckOutput(t, crew, receivedCrew, err)
	}

	for _, test := range testCrewRepositoryGetByIDFailed {
		_, err := crewRepository.GetCrewDataByID(uuid.New())
		test.CheckOutput(t, err)
	}
}

var testCrewRepositoryGetCrewDataBySailNumAndRatingIDSuccess = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdCrew *models.Crew, receivedCrew *models.Crew, err error)
}{
	{
		TestName: "get by Sail Number And Rating ID success test",
		CheckOutput: func(t *testing.T, createdCrew *models.Crew, receivedCrew *models.Crew, err error) {
			require.NoError(t, err)
			require.Equal(t, createdCrew.ID, receivedCrew.ID)
		},
	},
}

var testCrewRepositoryGetCrewDataBySailNumAndRatingIDFailed = []struct {
	TestName    string
	InputData   uuid.UUID
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName:  "get by Sail Number And Rating ID DoesNotExist test",
		InputData: uuid.New(),
		CheckOutput: func(t *testing.T, err error) {
			require.Error(t, err)
			require.Equal(t, repository_errors.DoesNotExist, err)
		},
	},
}

func TestCrewRepositoryGetCrewDataBySailNumAndRatingID(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryGetCrewDataBySailNumAndRatingIDSuccess {
		rating := CreateRating(&fields)
		crew := CreateCrew(&fields, rating.ID)

		receivedCrew, err := crewRepository.GetCrewDataBySailNumAndRatingID(crew.SailNum, crew.RatingID)
		test.CheckOutput(t, crew, receivedCrew, err)
	}

	for _, test := range testCrewRepositoryGetCrewDataBySailNumAndRatingIDFailed {
		rating := CreateRating(&fields)

		_, err := crewRepository.GetCrewDataBySailNumAndRatingID(123, rating.ID)
		test.CheckOutput(t, err)
	}
}

var testCrewRepositoryDeleteSuccess = []struct {
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

func TestCrewRepositoryDelete(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryDeleteSuccess {
		rating := CreateRating(&fields)
		createdCrew := CreateCrew(&fields, rating.ID)

		err := crewRepository.Delete(createdCrew.ID)
		test.CheckOutput(t, err)

		_, err = crewRepository.GetCrewDataByID(createdCrew.ID)
		require.Error(t, err)
	}
}

var testCrewRepositoryUpdateSuccess = []struct {
	TestName  string
	InputData struct {
		Crew *models.Crew
	}
	CheckOutput func(t *testing.T, createdCrew *models.Crew, updatedCrew *models.Crew, err error)
}{
	{
		TestName: "update success test",
		InputData: struct {
			Crew *models.Crew
		}{
			&models.Crew{
				SailNum: 123,
				Class:   2,
			},
		},
		CheckOutput: func(t *testing.T, createdCrew *models.Crew, updatedCrew *models.Crew, err error) {
			require.NoError(t, err)
			require.Equal(t, createdCrew.ID, updatedCrew.ID)
			require.Equal(t, createdCrew.RatingID, updatedCrew.RatingID)
			require.Equal(t, createdCrew.SailNum, updatedCrew.SailNum)
			require.Equal(t, createdCrew.Class, updatedCrew.Class)
		},
	},
}

func TestCrewRepositoryUpdate(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryUpdateSuccess {
		rating := CreateRating(&fields)
		test.InputData.Crew.RatingID = rating.ID
		createdCrew, _ := crewRepository.Create(test.InputData.Crew)

		updatedCrew, err := crewRepository.Update(
			&models.Crew{
				ID:       createdCrew.ID,
				RatingID: rating.ID,
				SailNum:  123,
				Class:    2,
			},
		)

		test.CheckOutput(t, createdCrew, updatedCrew, err)
	}
}

var testCrewRepositoryGetCrewsDataByRatingID = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdCrews []models.Crew, receivedCrews []models.Crew, err error)
}{
	{
		TestName: "Get Crews Data By Rating ID success test",
		CheckOutput: func(t *testing.T, createdCrews []models.Crew, receivedCrews []models.Crew, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdCrews), len(receivedCrews))
		},
	},
}

func TestCrewRepositoryGetCrewsDataByRatingID(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryGetCrewsDataByRatingID {
		rating := CreateRating(&fields)

		createdCrews := []models.Crew{
			{
				RatingID: rating.ID,
				SailNum:  123,
				Class:    2,
			},
			{
				RatingID: rating.ID,
				SailNum:  124,
				Class:    2,
			},
			{
				RatingID: rating.ID,
				SailNum:  125,
				Class:    2,
			},
		}

		for i := range createdCrews {
			_, err := crewRepository.Create(&createdCrews[i])
			require.NoError(t, err)
		}

		receivedCrews, err := crewRepository.GetCrewsDataByRatingID(rating.ID)
		test.CheckOutput(t, createdCrews, receivedCrews, err)
	}
}

// GetCrewsDataByProtestID(id uuid.UUID) ([]models.Crew, error)

var testCrewRepositoryGetCrewsDataByProtestID = []struct {
	TestName    string
	CheckOutput func(t *testing.T, createdCrews []models.Crew, receivedCrews []models.Crew, err error)
}{
	{
		TestName: "Get Crews Data By Protest ID success test",
		CheckOutput: func(t *testing.T, createdCrews []models.Crew, receivedCrews []models.Crew, err error) {
			require.NoError(t, err)
			require.Equal(t, len(createdCrews), len(receivedCrews))
		},
	},
}

func TestCrewRepositoryGetCrewsDataByProtestID(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryGetCrewsDataByProtestID {
		rating := CreateRating(&fields)
		judge := CreateJudge(&fields)
		race := CreateRace(&fields, rating.ID)
		protest := CreateProtest(&fields, race.ID, judge.ID, rating.ID)

		createdCrews := []models.Crew{
			{
				RatingID: rating.ID,
				SailNum:  123,
				Class:    2,
			},
			{
				RatingID: rating.ID,
				SailNum:  124,
				Class:    2,
			},
			{
				RatingID: rating.ID,
				SailNum:  125,
				Class:    2,
			},
		}

		for i := range createdCrews {
			c, err := crewRepository.Create(&createdCrews[i])
			require.NoError(t, err)
			AttachCrewToProtest(&fields, c.ID, protest.ID)
		}

		receivedCrews, err := crewRepository.GetCrewsDataByProtestID(protest.ID)
		test.CheckOutput(t, createdCrews, receivedCrews, err)
	}
}

var testCrewRepositoryAttachParticipantToCrew = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "Attach Participant To Crew success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestCrewRepositoryAttachParticipantToCrew(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryAttachParticipantToCrew {
		rating := CreateRating(&fields)
		participant := CreateParticipant(&fields)

		crew, err := crewRepository.Create(
			&models.Crew{
				RatingID: rating.ID,
				SailNum:  123,
				Class:    2,
			},
		)
		require.NoError(t, err)

		err = crewRepository.AttachParticipantToCrew(participant.ID, crew.ID, 1)
		test.CheckOutput(t, err)
	}
}

var testCrewRepositoryDetachParticipantFromCrew = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "Detach Participant From Crew success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestCrewRepositoryDetachParticipantFromCrew(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryDetachParticipantFromCrew {
		rating := CreateRating(&fields)
		participant := CreateParticipant(&fields)

		crew, err := crewRepository.Create(
			&models.Crew{
				RatingID: rating.ID,
				SailNum:  123,
				Class:    2,
			},
		)
		require.NoError(t, err)

		err = crewRepository.DetachParticipantFromCrew(participant.ID, crew.ID)
		test.CheckOutput(t, err)
	}
}

var testCrewRepositoryReplaceParticipantStatusInCrew = []struct {
	TestName    string
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "Replace Participant Status In Crew success test",
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestCrewRepositoryReplaceParticipantStatusInCrew(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := mongo.MongoConnection{DB: db}
	crewRepository := mongo.CreateCrewRepository(&fields)

	for _, test := range testCrewRepositoryReplaceParticipantStatusInCrew {
		rating := CreateRating(&fields)
		participant := CreateParticipant(&fields)

		crew, err := crewRepository.Create(
			&models.Crew{
				RatingID: rating.ID,
				SailNum:  123,
				Class:    2,
			},
		)
		require.NoError(t, err)

		err = crewRepository.ReplaceParticipantStatusInCrew(participant.ID, crew.ID, 1, 0)
		test.CheckOutput(t, err)
	}
}
