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

type JudgeDB struct {
	ID       string `bson:"_id"`
	FIO      string `bson:"name"`
	Login    string `bson:"login"`
	Password string `bson:"password"`
	Role     int    `bson:"role"`
	Post     string `bson:"post"`
}

type JudgeRepository struct {
	db *mongo.Database
}

func NewJudgeRepository(db *mongo.Database) repository_interfaces.IJudgeRepository {
	return &JudgeRepository{db: db}
}

func copyJudgeResultToModel(judgeDB *JudgeDB) *models.Judge {
	id, _ := uuid.Parse(judgeDB.ID)
	return &models.Judge{
		ID:       id,
		FIO:      judgeDB.FIO,
		Login:    judgeDB.Login,
		Password: judgeDB.Password,
		Role:     judgeDB.Role,
		Post:     judgeDB.Post,
	}
}

func (w *JudgeRepository) CreateProfile(judge *models.Judge) (*models.Judge, error) {
	collection := w.db.Collection("judges")
	judgeDB := &JudgeDB{
		ID:       judge.ID.String(),
		FIO:      judge.FIO,
		Login:    judge.Login,
		Password: judge.Password,
		Role:     judge.Role,
		Post:     judge.Post,
	}

	_, err := collection.InsertOne(context.TODO(), judgeDB)
	if err != nil {
		return nil, repository_errors.InsertError
	}

	return copyJudgeResultToModel(judgeDB), nil
}

func (w *JudgeRepository) UpdateProfile(judge *models.Judge) (*models.Judge, error) {
	collection := w.db.Collection("judges")
	filter := bson.M{"_id": judge.ID.String()}
	update := bson.M{
		"$set": bson.M{
			"name":     judge.FIO,
			"login":    judge.Login,
			"password": judge.Password,
			"role":     judge.Role,
			"post":     judge.Post,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedJudge JudgeDB
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedJudge)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.UpdateError
	}

	return copyJudgeResultToModel(&updatedJudge), nil
}

func (w *JudgeRepository) DeleteProfile(id uuid.UUID) error {
	collection := w.db.Collection("judges")
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

func (w *JudgeRepository) GetJudgeDataByID(id uuid.UUID) (*models.Judge, error) {
	collection := w.db.Collection("judges")
	filter := bson.M{"_id": id.String()}
	var judgeDB JudgeDB

	err := collection.FindOne(context.TODO(), filter).Decode(&judgeDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}

	return copyJudgeResultToModel(&judgeDB), nil
}

func (w *JudgeRepository) GetJudgeDataByLogin(login string) (*models.Judge, error) {
	collection := w.db.Collection("judges")
	filter := bson.M{"login": login}
	var judgeDB JudgeDB

	err := collection.FindOne(context.TODO(), filter).Decode(&judgeDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}

	return copyJudgeResultToModel(&judgeDB), nil
}

func (w *JudgeRepository) GetJudgesDataByRatingID(ratingID uuid.UUID) ([]models.Judge, error) {
	var m2mCollection = w.db.Collection("judge_rating")
	var judgesCollection = w.db.Collection("judges")

	var judges []models.Judge

	// Фильтр для выборки по rating_id
	filter := bson.M{"rating_id": ratingID.String()}

	// Опции для получения только поля judge_id
	projection := bson.M{"judge_id": 1, "_id": 0} // _id исключаем

	// Выполнение запроса
	cursor, err := m2mCollection.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result struct {
			JudgeID string `bson:"judge_id"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		var judgeDB JudgeDB
		err = judgesCollection.FindOne(context.Background(), bson.M{"_id": result.JudgeID}).Decode(&judgeDB)
		if err != nil {
			return nil, repository_errors.DoesNotExist
		}
		judges = append(judges, *copyJudgeResultToModel(&judgeDB))
	}

	return judges, nil
}

func (w *JudgeRepository) GetAllJudges() ([]models.Judge, error) {
	collection := w.db.Collection("judges")
	var judgesDB []JudgeDB

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var judgeDB JudgeDB
		err := cursor.Decode(&judgeDB)
		if err != nil {
			return nil, repository_errors.SelectError
		}
		judgesDB = append(judgesDB, judgeDB)
	}

	var judgeModels []models.Judge
	for _, judge := range judgesDB {
		judgeModels = append(judgeModels, *copyJudgeResultToModel(&judge))
	}

	return judgeModels, nil
}

func (w *JudgeRepository) GetJudgeDataByProtestID(protestID uuid.UUID) (*models.Judge, error) {
	var m2mCollection = w.db.Collection("protests")
	var judgesCollection = w.db.Collection("judges")

	var protest struct {
		JudgeID string `bson:"judge_id"`
	}
	err := m2mCollection.FindOne(context.Background(), bson.M{"_id": protestID.String()}).Decode(&protest)
	if err != nil {
		return nil, repository_errors.DoesNotExist
	}

	var judgeDB JudgeDB
	err = judgesCollection.FindOne(context.Background(), bson.M{"_id": protest.JudgeID}).Decode(&judgeDB)
	if err != nil {
		return nil, repository_errors.SelectError
	}

	return copyJudgeResultToModel(&judgeDB), nil
}
