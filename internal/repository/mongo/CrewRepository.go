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
	"log"
)

type CrewDB struct {
	ID       string `bson:"_id"`
	RatingID string `bson:"rating_id"`
	Class    int    `bson:"class"`
	SailNum  int    `bson:"sail_num"`
}

type CrewRepository struct {
	db *mongo.Database
}

func NewCrewRepository(db *mongo.Database) repository_interfaces.ICrewRepository {
	return &CrewRepository{db: db}
}

func copyCrewResultToModel(crewDB *CrewDB) *models.Crew {
	id, _ := uuid.Parse(crewDB.ID)
	ratingID, _ := uuid.Parse(crewDB.RatingID)
	return &models.Crew{
		ID:       id,
		RatingID: ratingID,
		SailNum:  crewDB.SailNum,
		Class:    crewDB.Class,
	}
}

func (c *CrewRepository) Create(crew *models.Crew) (*models.Crew, error) {
	collection := c.db.Collection("crews")

	id := uuid.New().String()
	ratingID := crew.RatingID.String()
	createdCrew := CrewDB{
		ID:       id,
		RatingID: ratingID,
		SailNum:  crew.SailNum,
		Class:    crew.Class,
	}

	_, err := collection.InsertOne(context.Background(), createdCrew)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return copyCrewResultToModel(&createdCrew), nil
}

func (c *CrewRepository) Update(crew *models.Crew) (*models.Crew, error) {
	collection := c.db.Collection("crews")
	filter := bson.M{"_id": crew.ID.String()}
	update := bson.M{"$set": bson.M{
		"rating_id": crew.RatingID.String(),
		"class":     crew.Class,
		"sail_num":  crew.SailNum,
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return crew, nil
}

func (c *CrewRepository) Delete(id uuid.UUID) error {
	collection := c.db.Collection("crews")
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

func (c *CrewRepository) GetCrewDataByID(id uuid.UUID) (*models.Crew, error) {
	collection := c.db.Collection("crews")
	filter := bson.M{"_id": id.String()}
	var crewDB CrewDB

	err := collection.FindOne(context.Background(), filter).Decode(&crewDB)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}
	return copyCrewResultToModel(&crewDB), nil
}

func (c *CrewRepository) AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error {
	m2mCollection := c.db.Collection("participant_crew")

	_, err := m2mCollection.InsertOne(context.Background(), bson.M{
		"crew_id":        crewID.String(),
		"participant_id": participantID.String(),
		"helmsman":       helmsman,
	})
	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (c *CrewRepository) DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error {
	m2mCollection := c.db.Collection("participant_crew")

	res, err := m2mCollection.DeleteOne(context.Background(), bson.M{
		"crew_id":        crewID.String(),
		"participant_id": participantID.String(),
	})
	if err != nil {
		return repository_errors.DeleteError
	}
	if res.DeletedCount == 0 {
		return repository_errors.DoesNotExist
	}

	return nil
}

func (c *CrewRepository) ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error {
	collection := c.db.Collection("participant_crew")
	filter := bson.M{"participant_id": participantID.String(), "crew_id": crewID.String()}
	update := bson.M{"$set": bson.M{
		"helmsman": helmsman,
		"active":   active,
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return repository_errors.UpdateError
	}
	return nil
}

func (c *CrewRepository) GetCrewsDataByRatingID(id uuid.UUID) ([]models.Crew, error) {
	collection := c.db.Collection("crews")
	filter := bson.M{"rating_id": id.String()}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var crewModels []models.Crew
	for cur.Next(context.Background()) {
		var crew CrewDB
		err := cur.Decode(&crew)
		if err != nil {
			return nil, err
		}
		crewModels = append(crewModels, *copyCrewResultToModel(&crew))
	}
	return crewModels, nil
}

func (c *CrewRepository) GetCrewsDataByProtestID(id uuid.UUID) ([]models.Crew, error) {
	m2mCollection := c.db.Collection("crew_protest")
	crewsCollection := c.db.Collection("crews")

	var crews []models.Crew

	filter := bson.M{"protest_id": id.String()}

	// Опции для получения только поля crew_id
	projection := bson.M{"crew_id": 1, "_id": 0} // _id исключаем

	// Выполнение запроса
	cursor, err := m2mCollection.Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result struct {
			CrewID string `bson:"crew_id"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		var crewDB CrewDB
		err = crewsCollection.FindOne(context.Background(), bson.M{"_id": result.CrewID}).Decode(&crewDB)
		if err != nil {
			return nil, repository_errors.DoesNotExist
		}
		crews = append(crews, *copyCrewResultToModel(&crewDB))
	}

	return crews, nil
}

func (c *CrewRepository) GetCrewDataBySailNumAndRatingID(sailNum int, ratingID uuid.UUID) (*models.Crew, error) {
	collection := c.db.Collection("crews")
	filter := bson.M{"sail_num": sailNum, "rating_id": ratingID.String()}
	var crew CrewDB

	err := collection.FindOne(context.Background(), filter).Decode(&crew)
	if err != nil {
		return nil, repository_errors.DoesNotExist
	}

	return copyCrewResultToModel(&crew), nil
}
