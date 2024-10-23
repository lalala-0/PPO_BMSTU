package mongo

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ProtestDB struct {
	ID         string    `bson:"_id"`
	RaceID     string    `bson:"race_id"`
	JudgeID    string    `bson:"judge_id"`
	RatingID   string    `bson:"rating_id"`
	RuleNum    int       `bson:"rule_num"`
	ReviewDate time.Time `bson:"review_date"`
	Status     int       `bson:"status"`
	Comment    string    `bson:"comment"`
}

type ProtestRepository struct {
	db *mongo.Database
}

func NewProtestRepository(db *mongo.Database) repository_interfaces.IProtestRepository {
	return &ProtestRepository{db: db}
}

func copyProtestResultToModel(protestDB *ProtestDB) *models.Protest {
	id, _ := uuid.Parse(protestDB.ID)
	raceID, _ := uuid.Parse(protestDB.RaceID)
	judgeID, _ := uuid.Parse(protestDB.JudgeID)
	ratingID, _ := uuid.Parse(protestDB.RatingID)

	return &models.Protest{
		ID:         id,
		RaceID:     raceID,
		JudgeID:    judgeID,
		RatingID:   ratingID,
		RuleNum:    protestDB.RuleNum,
		ReviewDate: protestDB.ReviewDate,
		Status:     protestDB.Status,
		Comment:    protestDB.Comment,
	}
}

func (w *ProtestRepository) Create(protest *models.Protest) (*models.Protest, error) {
	collection := w.db.Collection("protests")

	protestDB := &ProtestDB{
		ID:         uuid.New().String(),
		RaceID:     protest.RaceID.String(),
		JudgeID:    protest.JudgeID.String(),
		RatingID:   protest.RatingID.String(),
		RuleNum:    protest.RuleNum,
		ReviewDate: protest.ReviewDate,
		Status:     protest.Status,
		Comment:    protest.Comment,
	}

	_, err := collection.InsertOne(context.TODO(), protestDB)
	if err != nil {
		return nil, repository_errors.InsertError
	}

	return copyProtestResultToModel(protestDB), nil
}

func (w *ProtestRepository) Update(protest *models.Protest) (*models.Protest, error) {
	collection := w.db.Collection("protests")
	filter := bson.M{"_id": protest.ID.String()}
	update := bson.M{
		"$set": bson.M{
			"race_id":     protest.RaceID.String(),
			"rating_id":   protest.RatingID.String(),
			"judge_id":    protest.JudgeID.String(),
			"rule_num":    protest.RuleNum,
			"review_date": protest.ReviewDate,
			"status":      protest.Status,
			"comment":     protest.Comment,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedProtest ProtestDB
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedProtest)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.UpdateError
	}

	return copyProtestResultToModel(&updatedProtest), nil
}

func (w *ProtestRepository) Delete(id uuid.UUID) error {
	collection := w.db.Collection("protests")
	filter := bson.M{"_id": id.String()}

	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}

	if res.DeletedCount == 0 {
		return repository_errors.DoesNotExist
	}

	return nil
}

func (w *ProtestRepository) GetProtestDataByID(id uuid.UUID) (*models.Protest, error) {
	collection := w.db.Collection("protests")
	filter := bson.M{"_id": id.String()}
	var protestDB ProtestDB

	err := collection.FindOne(context.TODO(), filter).Decode(&protestDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}

	return copyProtestResultToModel(&protestDB), nil
}

func (w *ProtestRepository) AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error {
	collection := w.db.Collection("crew_protest")
	document := bson.M{
		"crew_id":     crewID.String(),
		"protest_id":  protestID.String(),
		"crew_status": crewStatus,
	}

	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (w *ProtestRepository) DetachCrewFromProtest(protestID uuid.UUID, crewID uuid.UUID) error {
	collection := w.db.Collection("crew_protest")
	filter := bson.M{"crew_id": crewID.String(), "protest_id": protestID.String()}

	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}

	if res.DeletedCount == 0 {
		return repository_errors.DoesNotExist
	}

	return nil
}

func (w *ProtestRepository) GetProtestsDataByRaceID(id uuid.UUID) ([]models.Protest, error) {
	collection := w.db.Collection("protests")
	filter := bson.M{"race_id": id.String()}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.TODO())

	var protestsDB []ProtestDB
	for cursor.Next(context.TODO()) {
		var protestDB ProtestDB
		err := cursor.Decode(&protestDB)
		if err != nil {
			return nil, repository_errors.SelectError
		}
		protestsDB = append(protestsDB, protestDB)
	}

	var protests []models.Protest
	for _, protestDB := range protestsDB {
		protests = append(protests, *copyProtestResultToModel(&protestDB))
	}

	return protests, nil
}

func (w *ProtestRepository) GetProtestParticipantsIDByID(id uuid.UUID) (map[uuid.UUID]int, error) {
	collection := w.db.Collection("crew_protest")
	filter := bson.M{"protest_id": id.String()}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.TODO())

	ids := make(map[uuid.UUID]int)
	for cursor.Next(context.TODO()) {
		var crewProtest struct {
			CrewID     string `bson:"crew_id"`
			CrewStatus int    `bson:"crew_status"`
		}
		err := cursor.Decode(&crewProtest)
		if err != nil {
			return nil, repository_errors.SelectError
		}

		crewID, _ := uuid.Parse(crewProtest.CrewID)
		ids[crewID] = crewProtest.CrewStatus
	}

	return ids, nil
}
