package mongo_init

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/mongo"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	DBNAME   = "testdb"
	USER     = "user"
	PASSWORD = "password"
)

func CreateParticipant(m *mongo.MongoConnection) *models.Participant {
	collection := m.DB.Collection("participants")

	id := uuid.New()
	participant := &mongo.ParticipantDB{
		ID:       id.String(),
		FIO:      "test",
		Gender:   false,
		Category: models.Junior2category,
		Coach:    "Test",
		Birthday: time.Date(2003, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	_, err := collection.InsertOne(context.Background(), participant)
	if err != nil {
		fmt.Println("Error creating participant:", err)
		return nil
	}

	return &models.Participant{
		ID:       id,
		FIO:      "test",
		Gender:   models.Female,
		Category: models.Junior2category,
		Coach:    "Test",
		Birthday: time.Date(2003, time.November, 10, 23, 0, 0, 0, time.UTC),
	}
}

func CreateRating(m *mongo.MongoConnection) *models.Rating {
	collection := m.DB.Collection("ratings")
	id := uuid.New()

	rating := &mongo.RatingDB{
		ID:         id.String(),
		Name:       "Name",
		Class:      models.Laser,
		BlowoutCnt: 1,
	}

	_, err := collection.InsertOne(context.Background(), rating)
	if err != nil {
		fmt.Println("Error creating rating:", err)
		return nil
	}
	return &models.Rating{
		ID:         id,
		Name:       rating.Name,
		Class:      rating.Class,
		BlowoutCnt: rating.BlowoutCnt,
	}
}

func CreateCrew(m *mongo.MongoConnection, ratingID uuid.UUID) *models.Crew {
	collection := m.DB.Collection("crews")

	crew := mongo.CrewDB{
		ID:       uuid.New().String(),
		RatingID: ratingID.String(),
		Class:    2,
		SailNum:  123,
	}

	_, err := collection.InsertOne(context.Background(), crew)
	if err != nil {
		fmt.Println("Error creating crew:", err)
		return nil
	}
	id, _ := uuid.Parse(crew.ID)

	return &models.Crew{
		ID:       id,
		RatingID: ratingID,
		Class:    crew.Class,
		SailNum:  crew.SailNum,
	}
}

func CreateCrewResInRace(m *mongo.MongoConnection, crewID uuid.UUID, raceID uuid.UUID) *models.CrewResInRace {
	collection := m.DB.Collection("crew_race")

	crewRes := &mongo.CrewResInRaceDB{
		CrewID:           crewID.String(),
		RaceID:           raceID.String(),
		Points:           12,
		SpecCircumstance: 0,
	}

	_, err := collection.InsertOne(context.Background(), crewRes)
	if err != nil {
		fmt.Println("Error creating crew result in race:", err)
		return nil
	}

	return &models.CrewResInRace{
		CrewID:           crewID,
		RaceID:           raceID,
		Points:           crewRes.Points,
		SpecCircumstance: crewRes.SpecCircumstance,
	}
}

func CreateRace(m *mongo.MongoConnection, ratingID uuid.UUID) *models.Race {
	collection := m.DB.Collection("races")
	id := uuid.New()

	race := &mongo.RaceDB{
		ID:       id.String(),
		RatingID: ratingID.String(),
		Date:     time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
		Number:   1,
		Class:    4,
	}

	_, err := collection.InsertOne(context.Background(), race)
	if err != nil {
		fmt.Println("Error creating race:", err)
		return nil
	}

	return &models.Race{
		ID:       id,
		RatingID: ratingID,
		Date:     race.Date,
		Number:   race.Number,
		Class:    race.Class,
	}
}

func CreateJudge(m *mongo.MongoConnection) *models.Judge {
	collection := m.DB.Collection("judges")

	id := uuid.New()
	judge := &mongo.JudgeDB{
		ID:       id.String(),
		FIO:      "Test",
		Login:    "Test",
		Password: "test123",
		Post:     "Test",
		Role:     1,
	}

	_, err := collection.InsertOne(context.Background(), judge)
	if err != nil {
		fmt.Println("Error creating judge:", err)
		return nil
	}

	return &models.Judge{
		ID:       id,
		FIO:      judge.FIO,
		Login:    judge.Login,
		Password: judge.Password,
		Post:     judge.Post,
		Role:     judge.Role,
	}
}

func CreateProtest(m *mongo.MongoConnection, raceID uuid.UUID, judgeID uuid.UUID, ratingID uuid.UUID) *models.Protest {
	collection := m.DB.Collection("protests")
	id := uuid.New()
	protest := &mongo.ProtestDB{
		ID:         id.String(),
		RaceID:     raceID.String(),
		JudgeID:    judgeID.String(),
		RatingID:   ratingID.String(),
		RuleNum:    23,
		ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
		Status:     1,
		Comment:    "",
	}

	_, err := collection.InsertOne(context.Background(), protest)
	if err != nil {
		fmt.Println("Error creating protest:", err)
		return nil
	}

	return &models.Protest{
		ID:         id,
		RaceID:     raceID,
		JudgeID:    judgeID,
		RatingID:   ratingID,
		RuleNum:    protest.RuleNum,
		ReviewDate: protest.ReviewDate,
		Status:     protest.Status,
		Comment:    protest.Comment,
	}
}

func AttachCrewToProtest(m *mongo.MongoConnection, crewID uuid.UUID, protestID uuid.UUID) {
	collection := m.DB.Collection("crew_protest")

	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":         uuid.New().String(),
		"crew_id":     crewID.String(),
		"protest_id":  protestID.String(),
		"crew_status": 1,
	})
	if err != nil {
		fmt.Println("Error attaching crew to protest:", err)
	}
}

func AttachCrewToProtestStatus(m *mongo.MongoConnection, crewID uuid.UUID, protestID uuid.UUID, status int) {
	collection := m.DB.Collection("crew_protest")

	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":         uuid.New().String(),
		"crew_id":     crewID.String(),
		"protest_id":  protestID.String(),
		"crew_status": status})
	if err != nil {
		fmt.Println("Error updating crew status in protest:", err)
	}
}

func AttachJudgeToRating(m *mongo.MongoConnection, judgeID uuid.UUID, ratingID uuid.UUID) {
	collection := m.DB.Collection("judge_rating")

	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":       uuid.New().String(),
		"judge_id":  judgeID.String(),
		"rating_id": ratingID.String(),
	})
	if err != nil {
		fmt.Println("Error attaching judge to rating:", err)
	}
}

func AttachParticipantToCrew(m *mongo.MongoConnection, participantID uuid.UUID, crewID uuid.UUID) {
	collection := m.DB.Collection("participant_crew")

	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":            uuid.New().String(),
		"participant_id": participantID.String(),
		"crew_id":        crewID.String(),
		"helmsman":       false,
		"active":         true,
	})
	if err != nil {
		fmt.Println("Error attaching participant to crew:", err)
	}
}
