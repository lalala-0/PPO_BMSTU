package db_init

import (
	"PPO_BMSTU/internal/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongoRepository реализует интерфейс TestRepositoryInitializer и использует соединение с MongoDB.
type mongoRepository struct {
	mongoDB *mongo.Database
}

// NewMongoRepository создает новый экземпляр mongoRepository.
func NewMongoRepository(m *mongo.Database) TestRepositoryInitializer {
	return &mongoRepository{mongoDB: m}
}

// CreateParticipant создает нового участника и возвращает его.
func (r *mongoRepository) CreateParticipant(participant *models.Participant) (*models.Participant, error) {
	collection := r.mongoDB.Collection("participants")

	_, err := collection.InsertOne(context.Background(), participant)
	if err != nil {
		fmt.Println("Error creating participant:", err)
		return nil, fmt.Errorf("error creating participant: %v", err)
	}

	return participant, nil
}

// CreateRating создает новый рейтинг и возвращает его.
func (r *mongoRepository) CreateRating(rating *models.Rating) (*models.Rating, error) {
	collection := r.mongoDB.Collection("ratings")

	_, err := collection.InsertOne(context.Background(), rating)
	if err != nil {
		fmt.Println("Error creating rating:", err)
		return nil, fmt.Errorf("error creating rating: %v", err)
	}

	return rating, nil
}

// CreateCrew создает новую команду и возвращает ее.
func (r *mongoRepository) CreateCrew(crew *models.Crew) (*models.Crew, error) {
	collection := r.mongoDB.Collection("crews")

	_, err := collection.InsertOne(context.Background(), crew)
	if err != nil {
		fmt.Println("Error creating crew:", err)
		return nil, fmt.Errorf("error creating crew: %v", err)
	}

	return crew, nil
}

// CreateCrewResInRace создает результат команды в гонке и возвращает его.
func (r *mongoRepository) CreateCrewResInRace(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	collection := r.mongoDB.Collection("crew_race")

	_, err := collection.InsertOne(context.Background(), crewResInRace)
	if err != nil {
		fmt.Println("Error creating crew result in race:", err)
		return nil, fmt.Errorf("error creating crew result in race: %v", err)
	}

	return crewResInRace, nil
}

// CreateRace создает новую гонку и возвращает ее.
func (r *mongoRepository) CreateRace(race *models.Race) (*models.Race, error) {
	collection := r.mongoDB.Collection("races")

	_, err := collection.InsertOne(context.Background(), race)
	if err != nil {
		fmt.Println("Error creating race:", err)
		return nil, fmt.Errorf("error creating race: %v", err)
	}

	return race, nil
}

// CreateJudge создает нового судью и возвращает его.
func (r *mongoRepository) CreateJudge(judge *models.Judge) (*models.Judge, error) {
	collection := r.mongoDB.Collection("judges")

	_, err := collection.InsertOne(context.Background(), judge)
	if err != nil {
		fmt.Println("Error creating judge:", err)
		return nil, fmt.Errorf("error creating judge: %v", err)
	}

	return judge, nil
}

// CreateProtest создает новый протест и возвращает его.
func (r *mongoRepository) CreateProtest(protest *models.Protest) (*models.Protest, error) {
	collection := r.mongoDB.Collection("protests")

	_, err := collection.InsertOne(context.Background(), protest)
	if err != nil {
		fmt.Println("Error creating protest:", err)
		return nil, fmt.Errorf("error creating protest: %v", err)
	}

	return protest, nil
}

// AttachCrewToProtest прикрепляет команду к протесту.
func (r *mongoRepository) AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID) error {
	collection := r.mongoDB.Collection("crew_protest")

	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":         uuid.New().String(),
		"crew_id":     crewID.String(),
		"protest_id":  protestID.String(),
		"crew_status": 1,
	})
	if err != nil {
		fmt.Println("Error attaching crew to protest:", err)
		return fmt.Errorf("error attaching crew to protest: %v", err)
	}
	return nil
}

// AttachCrewToProtestStatus обновляет статус команды в протесте.
func (r *mongoRepository) AttachCrewToProtestStatus(crewID uuid.UUID, protestID uuid.UUID, status int) error {
	collection := r.mongoDB.Collection("crew_protest")

	_, err := collection.UpdateOne(context.Background(),
		bson.M{"crew_id": crewID.String(), "protest_id": protestID.String()},
		bson.M{"$set": bson.M{"crew_status": status}})
	if err != nil {
		fmt.Println("Error updating crew status in protest:", err)
		return fmt.Errorf("error updating crew status in protest: %v", err)
	}
	return nil
}

// AttachJudgeToRating прикрепляет судью к рейтингу.
func (r *mongoRepository) AttachJudgeToRating(judgeID uuid.UUID, ratingID uuid.UUID) error {
	collection := r.mongoDB.Collection("judge_rating")

	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":       uuid.New().String(),
		"judge_id":  judgeID.String(),
		"rating_id": ratingID.String(),
	})
	if err != nil {
		fmt.Println("Error attaching judge to rating:", err)
		return fmt.Errorf("error attaching judge to rating: %v", err)
	}
	return nil
}

// AttachParticipantToCrew прикрепляет участника к команде.
func (r *mongoRepository) AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID) error {
	collection := r.mongoDB.Collection("participant_crew")

	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":            uuid.New().String(),
		"participant_id": participantID.String(),
		"crew_id":        crewID.String(),
	})
	if err != nil {
		fmt.Println("Error attaching participant to crew:", err)
		return fmt.Errorf("error attaching participant to crew: %v", err)
	}
	return nil
}
func (r *mongoRepository) ClearAll() error {
	// Список коллекций для очистки
	collections := []string{
		"crew_protest",
		"crew_race",
		"crews",
		"judge_rating",
		"judges",
		"participant_crew",
		"participants",
		"protests",
		"races",
		"rating",
	}

	// Очистка каждой коллекции
	for _, collectionName := range collections {
		collection := r.mongoDB.Collection(collectionName)
		_, err := collection.DeleteMany(context.Background(), bson.M{})
		if err != nil {
			fmt.Printf("Error clearing %s: %v\n", collectionName, err)
			return err // Завершаем функцию при ошибке
		}
	}

	return nil // Все коллекции успешно очищены
}
