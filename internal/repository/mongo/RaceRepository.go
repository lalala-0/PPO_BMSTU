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

type RaceDB struct {
	ID       string    `bson:"_id"`
	RatingID string    `bson:"rating_id"`
	Date     time.Time `bson:"date"`
	Number   int       `bson:"number"`
	Class    int       `bson:"class"`
}

type RaceRepository struct {
	db *mongo.Database
}

func NewRaceRepository(db *mongo.Database) repository_interfaces.IRaceRepository {
	return &RaceRepository{db: db}
}

func copyRaceResultToModel(raceDB *RaceDB) *models.Race {
	id, _ := uuid.Parse(raceDB.ID)
	ratingID, _ := uuid.Parse(raceDB.RatingID)

	return &models.Race{
		ID:       id,
		RatingID: ratingID,
		Date:     raceDB.Date,
		Number:   raceDB.Number,
		Class:    raceDB.Class,
	}
}

func (w *RaceRepository) Create(race *models.Race) (*models.Race, error) {
	collection := w.db.Collection("races")

	raceDB := &RaceDB{
		ID:       uuid.New().String(),
		RatingID: race.RatingID.String(),
		Date:     race.Date,
		Number:   race.Number,
		Class:    race.Class,
	}

	_, err := collection.InsertOne(context.Background(), raceDB)
	if err != nil {
		return nil, repository_errors.InsertError
	}

	return copyRaceResultToModel(raceDB), nil
}

func (w *RaceRepository) Update(race *models.Race) (*models.Race, error) {
	collection := w.db.Collection("races")
	filter := bson.M{"_id": race.ID.String()}
	update := bson.M{
		"$set": bson.M{
			"rating_id": race.RatingID.String(),
			"date":      race.Date,
			"number":    race.Number,
			"class":     race.Class,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedRace RaceDB
	err := collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedRace)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.UpdateError
	}

	return copyRaceResultToModel(&updatedRace), nil
}

func (w *RaceRepository) Delete(id uuid.UUID) error {
	collection := w.db.Collection("races")
	filter := bson.M{"_id": id.String()}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}

	if res.DeletedCount == 0 {
		return repository_errors.DoesNotExist
	}

	return nil
}

func (w *RaceRepository) GetRaceDataByID(id uuid.UUID) (*models.Race, error) {
	collection := w.db.Collection("races")
	filter := bson.M{"_id": id.String()}
	var raceDB RaceDB

	err := collection.FindOne(context.Background(), filter).Decode(&raceDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}

	return copyRaceResultToModel(&raceDB), nil
}

func (w *RaceRepository) GetRacesDataByRatingID(id uuid.UUID) ([]models.Race, error) {
	collection := w.db.Collection("races")
	filter := bson.M{"rating_id": id.String()}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.Background())

	var racesDB []RaceDB
	for cursor.Next(context.Background()) {
		var raceDB RaceDB
		err := cursor.Decode(&raceDB)
		if err != nil {
			return nil, err
		}
		racesDB = append(racesDB, raceDB)
	}

	var races []models.Race
	for _, raceDB := range racesDB {
		races = append(races, *copyRaceResultToModel(&raceDB))
	}

	return races, nil
}
