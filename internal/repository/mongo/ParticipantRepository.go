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

type ParticipantDB struct {
	ID       string    `bson:"_id"`
	FIO      string    `bson:"name"`
	Category int       `bson:"category"`
	Gender   bool      `bson:"gender"`
	Birthday time.Time `bson:"birthdate"`
	Coach    string    `bson:"coach_name"`
}

type ParticipantRepository struct {
	db *mongo.Database
}

func NewParticipantRepository(db *mongo.Database) repository_interfaces.IParticipantRepository {
	return &ParticipantRepository{db: db}
}

func copyParticipantResultToModel(participantDB *ParticipantDB) *models.Participant {
	gender := 0
	if participantDB.Gender {
		gender = 1
	}

	id, _ := uuid.Parse(participantDB.ID)

	return &models.Participant{
		ID:       id,
		FIO:      participantDB.FIO,
		Category: participantDB.Category,
		Gender:   gender,
		Birthday: participantDB.Birthday,
		Coach:    participantDB.Coach,
	}
}

func (p *ParticipantRepository) Create(participant *models.Participant) (*models.Participant, error) {
	collection := p.db.Collection("participants")

	participantDB := &ParticipantDB{
		ID:       uuid.New().String(),
		FIO:      participant.FIO,
		Category: participant.Category,
		Gender:   participant.Gender != 0,
		Birthday: participant.Birthday,
		Coach:    participant.Coach,
	}

	_, err := collection.InsertOne(context.TODO(), participantDB)
	if err != nil {
		return nil, repository_errors.InsertError
	}

	return copyParticipantResultToModel(participantDB), nil
}

func (p *ParticipantRepository) Update(participant *models.Participant) (*models.Participant, error) {
	collection := p.db.Collection("participants")
	filter := bson.M{"_id": participant.ID.String()}
	update := bson.M{
		"$set": bson.M{
			"name":       participant.FIO,
			"category":   participant.Category,
			"gender":     participant.Gender != 0,
			"birthdate":  participant.Birthday,
			"coach_name": participant.Coach,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedParticipant ParticipantDB
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedParticipant)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.UpdateError
	}

	return copyParticipantResultToModel(&updatedParticipant), nil
}

func (p *ParticipantRepository) Delete(id uuid.UUID) error {
	collection := p.db.Collection("participants")
	filter := bson.M{"_id": id.String()}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (p *ParticipantRepository) GetParticipantDataByID(id uuid.UUID) (*models.Participant, error) {
	collection := p.db.Collection("participants")
	filter := bson.M{"_id": id.String()}
	var participantDB ParticipantDB

	err := collection.FindOne(context.TODO(), filter).Decode(&participantDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}

	return copyParticipantResultToModel(&participantDB), nil
}

func (p *ParticipantRepository) GetParticipantsDataByCrewID(id uuid.UUID) ([]models.Participant, error) {
	var m2mCollection = p.db.Collection("participant_crew")
	var participantsCollection = p.db.Collection("participants")

	var participants []models.Participant

	filter := bson.M{"crew_id": id.String()}

	// Опции для получения только поля participant_id
	projection := bson.M{"participant_id": 1, "_id": 0}

	// Выполнение запроса
	cursor, err := m2mCollection.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result struct {
			ParticipantID string `bson:"participant_id"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		var participantDB ParticipantDB
		err = participantsCollection.FindOne(context.Background(), bson.M{"_id": result.ParticipantID}).Decode(&participantDB)
		if err != nil {
			return nil, repository_errors.DoesNotExist
		}
		participants = append(participants, *copyParticipantResultToModel(&participantDB))
	}

	return participants, nil
}

func (p *ParticipantRepository) GetParticipantsDataByProtestID(id uuid.UUID) ([]models.Participant, error) {
	// Коллекция participant_crew, через которую будем искать участников
	participantCrewCollection := p.db.Collection("participant_crew")

	// Агрегационный pipeline
	pipeline := mongo.Pipeline{
		// Ищем записи в коллекции crew_protest по protest_id
		{{"$lookup", bson.D{
			{"from", "crew_protest"},
			{"localField", "crew_id"},
			{"foreignField", "crew_id"},
			{"as", "crew_protests"},
		}}},
		{{"$unwind", "$crew_protests"}},                                 // Разворачиваем массив
		{{"$match", bson.D{{"crew_protests.protest_id", id.String()}}}}, // Фильтрация по protest_id

		// Подключаем участников через participant_crew
		{{"$lookup", bson.D{
			{"from", "participants"},
			{"localField", "participant_id"},
			{"foreignField", "_id"},
			{"as", "participant_info"},
		}}},
		{{"$unwind", "$participant_info"}}, // Разворачиваем массив с участниками

		// Оставляем только нужную информацию о участниках
		{{"$replaceRoot", bson.D{{"newRoot", "$participant_info"}}}},
	}

	cursor, err := participantCrewCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.TODO())

	var participants []models.Participant
	if err := cursor.All(context.TODO(), &participants); err != nil {
		return nil, repository_errors.SelectError
	}

	if len(participants) == 0 {
		return nil, repository_errors.DoesNotExist
	}

	return participants, nil
}

func (p *ParticipantRepository) GetAllParticipants() ([]models.Participant, error) {
	collection := p.db.Collection("participants")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.TODO())

	var participantsDB []ParticipantDB
	for cursor.Next(context.TODO()) {
		var participantDB ParticipantDB
		err := cursor.Decode(&participantDB)
		if err != nil {
			return nil, repository_errors.SelectError
		}
		participantsDB = append(participantsDB, participantDB)
	}

	var participants []models.Participant
	for _, participantDB := range participantsDB {
		participants = append(participants, *copyParticipantResultToModel(&participantDB))
	}

	return participants, nil
}
