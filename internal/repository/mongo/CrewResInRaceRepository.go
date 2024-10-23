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
)

type CrewResInRaceDB struct {
	ID               string `bson:"_id,omitempty"`
	CrewID           string `bson:"crew_id"`
	RaceID           string `bson:"race_id"`
	Points           int    `bson:"points"`
	SpecCircumstance int    `bson:"spec_circumstance"`
}

type CrewResInRaceRepository struct {
	db *mongo.Database
}

func NewCrewResInRaceRepository(db *mongo.Database) repository_interfaces.ICrewResInRaceRepository {
	return &CrewResInRaceRepository{db: db}
}

func copyCrewResInRaceResultToModel(crewResInRaceDB *CrewResInRaceDB) *models.CrewResInRace {
	crewID, _ := uuid.Parse(crewResInRaceDB.CrewID)
	raceID, _ := uuid.Parse(crewResInRaceDB.RaceID)
	return &models.CrewResInRace{
		CrewID:           crewID,
		RaceID:           raceID,
		Points:           crewResInRaceDB.Points,
		SpecCircumstance: crewResInRaceDB.SpecCircumstance,
	}
}

// Create - создание новой записи
func (w *CrewResInRaceRepository) Create(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	collection := w.db.Collection("crew_race")
	crewResInRaceDB := &CrewResInRaceDB{
		ID:               uuid.NewString(),
		CrewID:           crewResInRace.CrewID.String(),
		RaceID:           crewResInRace.RaceID.String(),
		Points:           crewResInRace.Points,
		SpecCircumstance: crewResInRace.SpecCircumstance,
	}
	_, err := collection.InsertOne(context.TODO(), crewResInRaceDB)
	if err != nil {
		return nil, repository_errors.InsertError
	}
	return copyCrewResInRaceResultToModel(crewResInRaceDB), nil
}

// Delete - удаление записи по raceID и crewID
func (w *CrewResInRaceRepository) Delete(raceID uuid.UUID, crewID uuid.UUID) error {
	collection := w.db.Collection("crew_race")
	filter := bson.M{"race_id": raceID.String(), "crew_id": crewID.String()}

	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}

	if res.DeletedCount == 0 {
		return repository_errors.DoesNotExist
	}

	return nil
}

// Update - обновление записи
func (w *CrewResInRaceRepository) Update(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	collection := w.db.Collection("crew_race")
	filter := bson.M{"race_id": crewResInRace.RaceID.String(), "crew_id": crewResInRace.CrewID.String()}
	update := bson.M{
		"$set": bson.M{
			"points":            crewResInRace.Points,
			"spec_circumstance": crewResInRace.SpecCircumstance,
		},
	}

	var updatedCrewResInRace CrewResInRaceDB
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedCrewResInRace)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.UpdateError
	}

	return copyCrewResInRaceResultToModel(&updatedCrewResInRace), nil
}

// GetCrewResByRaceIDAndCrewID - получение записи по raceID и crewID
func (w *CrewResInRaceRepository) GetCrewResByRaceIDAndCrewID(raceID uuid.UUID, crewID uuid.UUID) (*models.CrewResInRace, error) {
	collection := w.db.Collection("crew_race")
	filter := bson.M{"race_id": raceID.String(), "crew_id": crewID.String()}

	var crewResInRaceDB CrewResInRaceDB
	err := collection.FindOne(context.TODO(), filter).Decode(&crewResInRaceDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}

	return copyCrewResInRaceResultToModel(&crewResInRaceDB), nil
}

// GetAllCrewResInRace - получение всех записей для конкретной гонки
func (w *CrewResInRaceRepository) GetAllCrewResInRace(raceID uuid.UUID) ([]models.CrewResInRace, error) {
	collection := w.db.Collection("crew_race")
	filter := bson.M{"race_id": raceID.String()}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.TODO())

	var crewResInRaceDBs []CrewResInRaceDB
	if err = cursor.All(context.TODO(), &crewResInRaceDBs); err != nil {
		return nil, repository_errors.SelectError
	}

	var crewResInRaceModels []models.CrewResInRace
	for _, crewResInRaceDB := range crewResInRaceDBs {
		crewResInRaceModels = append(crewResInRaceModels, *copyCrewResInRaceResultToModel(&crewResInRaceDB))
	}

	return crewResInRaceModels, nil
}
