package db_init

//
//import (
//	"PPO_BMSTU/internal/models"
//	"PPO_BMSTU/internal/repository/mongo"
//	"context"
//	"fmt"
//	"github.com/google/uuid"
//	"time"
//
//	"go.mongodb.org/mongo-driver/bson"
//)
//
//// mongoRepository реализует интерфейс MongoRepository и использует соединение с MongoDB.
//type mongoRepository struct {
//	mongoConnection *mongo.MongoConnection
//}
//
//func NewMongoRepository(m *mongo.MongoConnection) TestRepository {
//	return &mongoRepository{mongoConnection: m}
//}
//
//func (r *mongoRepository) 	CreateParticipant(participant *models.Participant) (*models.Participant, error) {
//	collection := r.mongoConnection.DB.Collection("participants")
//
//	participantDB := participantToDB(participant)
//
//	_, err := collection.InsertOne(context.Background(), participantDB)
//	if err != nil {
//		fmt.Println("Error creating participant:", err)
//		return nil, fmt.Errorf("error creating participant: %v", err)
//	}
//
//	res := participantDB(participantDB)
//
//	return res
//}
//
//func (r *mongoRepository) CreateRating() *models.Rating {
//	collection := r.mongoConnection.DB.Collection("ratings")
//	id := uuid.New()
//
//	rating := &mongo.RatingDB{
//		ID:         id.String(),
//		Name:       "Name",
//		Class:      models.Laser,
//		BlowoutCnt: 1,
//	}
//
//	_, err := collection.InsertOne(context.Background(), rating)
//	if err != nil {
//		fmt.Println("Error creating rating:", err)
//		return nil
//	}
//	return &models.Rating{
//		ID:         id,
//		Name:       rating.Name,
//		Class:      rating.Class,
//		BlowoutCnt: rating.BlowoutCnt,
//	}
//}
//
//func (r *mongoRepository) CreateCrew() *models.Crew {
//	collection := r.mongoConnection.DB.Collection("crews")
//
//	crew := mongo.CrewDB{
//		ID:       uuid.New().String(),
//		RatingID: ratingID.String(),
//		Class:    2,
//		SailNum:  123,
//	}
//
//	_, err := collection.InsertOne(context.Background(), crew)
//	if err != nil {
//		fmt.Println("Error creating crew:", err)
//		return nil
//	}
//	id, _ := uuid.Parse(crew.ID)
//
//	return &models.Crew{
//		ID:       id,
//		RatingID: ratingID,
//		Class:    crew.Class,
//		SailNum:  crew.SailNum,
//	}
//}
//
//func (r *mongoRepository) CreateCrewResInRace() *models.CrewResInRace {
//	collection := r.mongoConnection.DB.Collection("crew_race")
//
//	crewRes := &mongo.CrewResInRaceDB{
//		CrewID:           crewID.String(),
//		RaceID:           raceID.String(),
//		Points:           12,
//		SpecCircumstance: 0,
//	}
//
//	_, err := collection.InsertOne(context.Background(), crewRes)
//	if err != nil {
//		fmt.Println("Error creating crew result in race:", err)
//		return nil
//	}
//
//	return &models.CrewResInRace{
//		CrewID:           crewID,
//		RaceID:           raceID,
//		Points:           crewRes.Points,
//		SpecCircumstance: crewRes.SpecCircumstance,
//	}
//}
//
//func (r *mongoRepository) CreateRace() *models.Race {
//	collection := r.mongoConnection.DB.Collection("races")
//	id := uuid.New()
//
//	race := &mongo.RaceDB{
//		ID:       id.String(),
//		RatingID: ratingID.String(),
//		Date:     time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
//		Number:   1,
//		Class:    4,
//	}
//
//	_, err := collection.InsertOne(context.Background(), race)
//	if err != nil {
//		fmt.Println("Error creating race:", err)
//		return nil
//	}
//
//	return &models.Race{
//		ID:       id,
//		RatingID: ratingID,
//		Date:     race.Date,
//		Number:   race.Number,
//		Class:    race.Class,
//	}
//}
//
//func (r *mongoRepository) CreateJudge() *models.Judge {
//	collection := r.mongoConnection.DB.Collection("judges")
//
//	id := uuid.New()
//	judge := &mongo.JudgeDB{
//		ID:       id.String(),
//		FIO:      "Test",
//		Login:    "Test",
//		Password: "test123",
//		Post:     "Test",
//		Role:     1,
//	}
//
//	_, err := collection.InsertOne(context.Background(), judge)
//	if err != nil {
//		fmt.Println("Error creating judge:", err)
//		return nil
//	}
//
//	return &models.Judge{
//		ID:       id,
//		FIO:      judge.FIO,
//		Login:    judge.Login,
//		Password: judge.Password,
//		Post:     judge.Post,
//		Role:     judge.Role,
//	}
//}
//
//func (r *mongoRepository) CreateProtest() *models.Protest {
//	collection := r.mongoConnection.DB.Collection("protests")
//	id := uuid.New()
//	protest := &mongo.ProtestDB{
//		ID:         id.String(),
//		RaceID:     raceID.String(),
//		JudgeID:    judgeID.String(),
//		RatingID:   ratingID.String(),
//		RuleNum:    23,
//		ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
//		Status:     1,
//		Comment:    "",
//	}
//
//	_, err := collection.InsertOne(context.Background(), protest)
//	if err != nil {
//		fmt.Println("Error creating protest:", err)
//		return nil
//	}
//
//	return &models.Protest{
//		ID:         id,
//		RaceID:     raceID,
//		JudgeID:    judgeID,
//		RatingID:   ratingID,
//		RuleNum:    protest.RuleNum,
//		ReviewDate: protest.ReviewDate,
//		Status:     protest.Status,
//		Comment:    protest.Comment,
//	}
//}
//
//func (r *mongoRepository) AttachCrewToProtest() {
//	collection := r.mongoConnection.DB.Collection("crew_protest")
//
//	_, err := collection.InsertOne(context.Background(), bson.M{
//		"_id":         uuid.New().String(),
//		"crew_id":     crewID.String(),
//		"protest_id":  protestID.String(),
//		"crew_status": 1,
//	})
//	if err != nil {
//		fmt.Println("Error attaching crew to protest:", err)
//	}
//}
//
//func (r *mongoRepository) AttachCrewToProtestStatus() {
//	collection := r.mongoConnection.DB.Collection("crew_protest")
//
//	_, err := collection.InsertOne(context.Background(), bson.M{
//		"_id":         uuid.New().String(),
//		"crew_id":     crewID.String(),
//		"protest_id":  protestID.String(),
//		"crew_status": status})
//	if err != nil {
//		fmt.Println("Error updating crew status in protest:", err)
//	}
//}
//
//func (r *mongoRepository) AttachJudgeToRating() {
//	collection := r.mongoConnection.DB.Collection("judge_rating")
//
//	_, err := collection.InsertOne(context.Background(), bson.M{
//		"_id":       uuid.New().String(),
//		"judge_id":  judgeID.String(),
//		"rating_id": ratingID.String(),
//	})
//	if err != nil {
//		fmt.Println("Error attaching judge to rating:", err)
//	}
//}
//
//func (r *mongoRepository) AttachParticipantToCrew() {
//	collection := r.mongoConnection.DB.Collection("participant_crew")
//
//	_, err := collection.InsertOne(context.Background(), bson.M{
//		"_id":            uuid.New().String(),
//		"participant_id": participantID.String(),
//		"crew_id":        crewID.String(),
//		"helmsman":       false,
//		"active":         true,
//	})
//	if err != nil {
//		fmt.Println("Error attaching participant to crew:", err)
//	}
//}
//
//
//
//func participantToDB(participant *models.Participant) *mongo.ParticipantDB {
//	var gender bool
//	if participant.Gender == 1{
//		gender = true
//	} else{
//		gender= false
//	}
//	return &mongo.ParticipantDB{
//		ID:       participant.ID.String(),
//		FIO:      participant.FIO,
//		Gender:   gender,
//		Category: participant.Category,
//		Coach:    participant.Coach,
//		Birthday: participant.Birthday,
//	}
//}
//
//func participantFromDB(participantDB *mongo.ParticipantDB) *models.Participant {
//	id, _ := uuid.Parse(participantDB.ID) // Обработка ошибки при необходимости
//	var gender int
//	if participantDB.Gender{
//		gender = 1
//	} else{
//		gender= 0
//	}
//	return &models.Participant{
//		ID:       id,
//		FIO:      participantDB.FIO,
//		Gender:   gender,
//		Category: participantDB.Category,
//		Coach:    participantDB.Coach,
//		Birthday: participantDB.Birthday,
//	}
//}
//
//func ratingToDB(rating *models.Rating) *mongo.RatingDB {
//	return &mongo.RatingDB{
//		ID:         rating.ID.String(),
//		Name:       rating.Name,
//		Class:      rating.Class,
//		BlowoutCnt: rating.BlowoutCnt,
//	}
//}
//
//func ratingFromDB(ratingDB *mongo.RatingDB) *models.Rating {
//	id, _ := uuid.Parse(ratingDB.ID) // Обработка ошибки при необходимости
//	return &models.Rating{
//		ID:         id,
//		Name:       ratingDB.Name,
//		Class:      ratingDB.Class,
//		BlowoutCnt: ratingDB.BlowoutCnt,
//	}
//}
//
//func crewToDB(crew *models.Crew) *mongo.CrewDB {
//	return &mongo.CrewDB{
//		ID:       crew.ID.String(),
//		RatingID: crew.RatingID.String(),
//		Class:    crew.Class,
//		SailNum:  crew.SailNum,
//	}
//}
//
//func crewFromDB(crewDB *mongo.CrewDB) *models.Crew {
//	id, _ := uuid.Parse(crewDB.ID) // Обработка ошибки при необходимости
//	ratingID, _ := uuid.Parse(crewDB.RatingID) // Обработка ошибки при необходимости
//	return &models.Crew{
//		ID:       id,
//		RatingID: ratingID,
//		Class:    crewDB.Class,
//		SailNum:  crewDB.SailNum,
//	}
//}
//
//func raceToDB(race *models.Race) *mongo.RaceDB {
//	return &mongo.RaceDB{
//		ID:       race.ID.String(),
//		RatingID: race.RatingID.String(),
//		Date:     race.Date,
//		Number:   race.Number,
//		Class:    race.Class,
//	}
//}
//
//func raceFromDB(raceDB *mongo.RaceDB) *models.Race {
//	id, _ := uuid.Parse(raceDB.ID) // Обработка ошибки при необходимости
//	ratingID, _ := uuid.Parse(raceDB.RatingID) // Обработка ошибки при необходимости
//	return &models.Race{
//		ID:       id,
//		RatingID: ratingID,
//		Date:     raceDB.Date,
//		Number:   raceDB.Number,
//		Class:    raceDB.Class,
//	}
//}
//
//func judgeToDB(judge *models.Judge) *mongo.JudgeDB {
//	return &mongo.JudgeDB{
//		ID:       judge.ID.String(),
//		FIO:      judge.FIO,
//		Login:    judge.Login,
//		Password: judge.Password,
//		Post:     judge.Post,
//		Role:     judge.Role,
//	}
//}
//
//func judgeFromDB(judgeDB *mongo.JudgeDB) *models.Judge {
//	id, _ := uuid.Parse(judgeDB.ID) // Обработка ошибки при необходимости
//	return &models.Judge{
//		ID:       id,
//		FIO:      judgeDB.FIO,
//		Login:    judgeDB.Login,
//		Password: judgeDB.Password,
//		Post:     judgeDB.Post,
//		Role:     judgeDB.Role,
//	}
//}
//
//func protestToDB(protest *models.Protest) *mongo.ProtestDB {
//	return &mongo.ProtestDB{
//		ID:         protest.ID.String(),
//		RaceID:     protest.RaceID.String(),
//		JudgeID:    protest.JudgeID.String(),
//		RatingID:   protest.RatingID.String(),
//		RuleNum:    protest.RuleNum,
//		ReviewDate: protest.ReviewDate,
//		Status:     protest.Status,
//		Comment:    protest.Comment,
//	}
//}
//
//func protestFromDB(protestDB *mongo.ProtestDB) *models.Protest {
//	id, _ := uuid.Parse(protestDB.ID) // Обработка ошибки при необходимости
//	raceID, _ := uuid.Parse(protestDB.RaceID) // Обработка ошибки при необходимости
//	judgeID, _ := uuid.Parse(protestDB.JudgeID) // Обработка ошибки при необходимости
//	ratingID, _ := uuid.Parse(protestDB.RatingID) // Обработка ошибки при необходимости
//	return &models.Protest{
//		ID:         id,
//		RaceID:     raceID,
//		JudgeID:    judgeID,
//		RatingID:   ratingID,
//		RuleNum:    protestDB.RuleNum,
//		ReviewDate: protestDB.ReviewDate,
//		Status:     protestDB.Status,
//		Comment:    protestDB.Comment,
//	}
//}
