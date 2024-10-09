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
	"sort"
	"time"
)

type RatingDB struct {
	ID         string `bson:"_id"`
	Name       string `bson:"name"`
	Class      int    `bson:"class"`
	BlowoutCnt int    `bson:"blowout_cnt"`
}

type RatingTableDB struct {
	CrewID               uuid.UUID `bson:"_id"`
	SailNum              int       `bson:"sail_num"`
	ParticipantName      string    `bson:"name"`
	ParticipantBirthDate time.Time `bson:"birthdate"`
	ParticipantCategory  int       `bson:"category"`
	Points               int       `bson:"points"`
	RaceNumber           int       `bson:"number"`
	PointsSum            int       `bson:"points_summ"`
	Rank                 int       `bson:"rang"`
	CoachName            string    `bson:"coach_name"`
}

type RatingRepository struct {
	db *mongo.Database
}

func NewRatingRepository(db *mongo.Database) repository_interfaces.IRatingRepository {
	return &RatingRepository{db: db}
}

func copyRatingResultToModel(ratingDB *RatingDB) *models.Rating {
	id, err := uuid.Parse(ratingDB.ID)
	if err != nil {
		return nil
	}
	return &models.Rating{
		ID:         id,
		Name:       ratingDB.Name,
		Class:      ratingDB.Class,
		BlowoutCnt: ratingDB.BlowoutCnt,
	}
}

func copyRatingTableResultToModel(ratingTableDB []RatingTableDB) []models.RatingTableLine {
	var ratingLines []models.RatingTableLine
	var participantNames []string
	var participantBirthDates []time.Time
	var participantCategories []int
	var coachNames []string
	resInRace := make(map[int]int)
	for i, line := range ratingTableDB[:len(ratingTableDB)-1] {
		participantNames = append(participantNames, line.ParticipantName)
		participantBirthDates = append(participantBirthDates, line.ParticipantBirthDate)
		participantCategories = append(participantCategories, line.ParticipantCategory)
		coachNames = append(coachNames, line.CoachName)
		resInRace[line.RaceNumber] = line.Points

		if line.SailNum != ratingTableDB[i+1].SailNum {
			ratingLines = append(ratingLines, models.RatingTableLine{
				SailNum:               line.SailNum,
				ParticipantNames:      participantNames,
				ParticipantBirthDates: participantBirthDates,
				ParticipantCategories: participantCategories,
				ResInRace:             resInRace,
				PointsSum:             line.PointsSum,
				Rank:                  line.Rank,
				CoachNames:            coachNames,
			})
			participantNames = make([]string, 0)
			participantBirthDates = make([]time.Time, 0)
			participantCategories = make([]int, 0)
			coachNames = make([]string, 0)
			resInRace = make(map[int]int, 0)
		}
	}

	line := ratingTableDB[len(ratingTableDB)-1]
	participantNames = append(participantNames, line.ParticipantName)
	participantBirthDates = append(participantBirthDates, line.ParticipantBirthDate)
	participantCategories = append(participantCategories, line.ParticipantCategory)
	coachNames = append(coachNames, line.CoachName)
	resInRace[line.RaceNumber] = line.RaceNumber

	ratingLines = append(ratingLines, models.RatingTableLine{
		SailNum:               line.SailNum,
		ParticipantNames:      participantNames,
		ParticipantBirthDates: participantBirthDates,
		ParticipantCategories: participantCategories,
		ResInRace:             resInRace,
		PointsSum:             line.PointsSum,
		Rank:                  line.Rank,
		CoachNames:            coachNames,
	})

	// Сортировка по полю SailNum
	sort.Slice(ratingLines, func(i, j int) bool {
		return ratingLines[i].PointsSum < ratingLines[j].PointsSum
	})

	return ratingLines
}

func (w RatingRepository) Create(rating *models.Rating) (*models.Rating, error) {
	collection := w.db.Collection("ratings")
	ratingDB := &RatingDB{
		ID:         uuid.New().String(),
		Name:       rating.Name,
		Class:      rating.Class,
		BlowoutCnt: rating.BlowoutCnt,
	}

	_, err := collection.InsertOne(context.TODO(), ratingDB)
	if err != nil {
		return nil, repository_errors.InsertError
	}

	return copyRatingResultToModel(ratingDB), nil
}

func (w RatingRepository) Update(rating *models.Rating) (*models.Rating, error) {
	collection := w.db.Collection("ratings")
	filter := bson.M{"_id": rating.ID.String()}
	update := bson.M{
		"$set": bson.M{
			"name":        rating.Name,
			"class":       rating.Class,
			"blowout_cnt": rating.BlowoutCnt,
		},
	}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedRating RatingDB
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&updatedRating)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.UpdateError
	}

	return copyRatingResultToModel(&updatedRating), nil
}

func (w RatingRepository) Delete(id uuid.UUID) error {
	collection := w.db.Collection("ratings")
	filter := bson.M{"_id": id.String()}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w RatingRepository) GetRatingDataByID(id uuid.UUID) (*models.Rating, error) {
	collection := w.db.Collection("ratings")
	filter := bson.M{"_id": id.String()}
	var ratingDB RatingDB

	err := collection.FindOne(context.TODO(), filter).Decode(&ratingDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository_errors.DoesNotExist
		}
		return nil, repository_errors.SelectError
	}

	return copyRatingResultToModel(&ratingDB), nil
}

func (w RatingRepository) AttachJudgeToRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	collection := w.db.Collection("judge_rating")
	_, err := collection.InsertOne(context.TODO(), bson.M{"judge_id": judgeID.String(), "rating_id": ratingID.String()})

	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (w RatingRepository) DetachJudgeFromRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	collection := w.db.Collection("judge_rating")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"judge_id": judgeID.String(), "rating_id": ratingID.String()})

	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w RatingRepository) GetAllRatings() ([]models.Rating, error) {
	collection := w.db.Collection("ratings")
	cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find().SetSort(bson.M{"name": 1}))

	if err != nil {
		return nil, repository_errors.SelectError
	}
	defer cursor.Close(context.TODO())

	var ratingDB []RatingDB
	if err = cursor.All(context.TODO(), &ratingDB); err != nil {
		return nil, repository_errors.SelectError
	}

	var ratingModels []models.Rating
	for _, rating := range ratingDB {
		ratingModels = append(ratingModels, *copyRatingResultToModel(&rating))
	}

	return ratingModels, nil
}

func (r *RatingRepository) GetRatingTable(id uuid.UUID) ([]models.RatingTableLine, error) {

	var ratingTableDB []RatingTableDB

	crewCollection := r.db.Collection("crews")
	filter := bson.M{"rating_id": id.String()}
	cur, err := crewCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var crewModels []CrewDB
	for cur.Next(context.Background()) {
		var crew CrewDB
		err := cur.Decode(&crew)
		if err != nil {
			return nil, err
		}
		crewModels = append(crewModels, crew)
	}

	raceCollection := r.db.Collection("races")
	filter = bson.M{"rating_id": id.String()}

	cursor, err := raceCollection.Find(context.Background(), filter)
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

	for i, race := range racesDB {
		for _, crew := range crewModels {
			resCollection := r.db.Collection("crew_race")
			filter := bson.M{"race_id": race.ID, "crew_id": crew.ID}

			var crewResInRaceDB CrewResInRaceDB
			err := resCollection.FindOne(context.TODO(), filter).Decode(&crewResInRaceDB)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, repository_errors.DoesNotExist
				}
				return nil, repository_errors.SelectError
			}

			ratingTableDB = append(ratingTableDB, RatingTableDB{SailNum: crew.SailNum, Points: crewResInRaceDB.Points, RaceNumber: i + 1, PointsSum: crewResInRaceDB.Points * len(racesDB)})
		}
	}

	sort.Slice(ratingTableDB, func(i, j int) bool {
		return false
	})
	// Сортировка по полю SailNum
	sort.Slice(ratingTableDB, func(i, j int) bool {
		return ratingTableDB[i].SailNum < ratingTableDB[j].SailNum
	})
	calculatePointsSum(ratingTableDB)

	// Преобразуем результат в модель
	ratingTable := copyRatingTableResultToModel(ratingTableDB)

	return ratingTable, nil
}

//func (r *RatingRepository) GetRatingTable(id uuid.UUID) ([]models.RatingTableLine, error) {
//	collection := r.db.Collection("crews")
//
//	// Преобразуем UUID в строку для использования в MongoDB
//	ratingID := id.String()
//
//	// Шаги агрегации
//	pipeline := mongo.Pipeline{
//		// Шаг 1: Соединяем экипажи с участниками (только рулевые)
//		{
//			{"$lookup", bson.D{
//				{"from", "participant_crew"},
//				{"localField", "_id"},
//				{"foreignField", "crew_id"},
//				{"as", "crew_participants"},
//				{"pipeline", bson.A{
//					bson.D{{"$match", bson.D{{"helmsman", true}}}}, // Фильтрация по рулевым
//					bson.D{
//						{"$lookup", bson.D{
//							{"from", "participants"},
//							{"localField", "participant_id"},
//							{"foreignField", "_id"},
//							{"as", "participant_info"},
//						}},
//					},
//					bson.D{{"$unwind", "$participant_info"}}, // Разворачиваем массив участника
//				}},
//			}},
//		},
//		// Шаг 2: Соединяем экипажи с результатами гонок (crew_race)
//		{
//			{"$lookup", bson.D{
//				{"from", "crew_race"},
//				{"localField", "_id"},
//				{"foreignField", "crew_id"},
//				{"as", "crew_races"},
//			}},
//		},
//		// Шаг 3: Соединяем результаты гонок с номерами гонок из races
//		{
//			{"$lookup", bson.D{
//				{"from", "races"},
//				{"localField", "crew_races.race_id"},
//				{"foreignField", "_id"},
//				{"as", "race_info"},
//			}},
//		},
//		// Шаг 4: Считаем сумму очков экипажа
//		{
//			{"$addFields", bson.D{
//				{"points_summ", bson.D{{"$sum", "$crew_races.points"}}},
//			}},
//		},
//		// Шаг 5: Фильтруем экипажи по конкретному рейтингу
//		{
//			{"$match", bson.D{
//				{"rating_id", ratingID},
//			}},
//		},
//		// Шаг 6: Сортируем экипажи по очкам и номерам гонок
//		{
//			{"$sort", bson.D{
//				{"points_summ", -1},     // Сортировка по сумме очков (по убыванию)
//				{"_id", 1},              // Вторичная сортировка по ID
//				{"race_info.number", 1}, // Сортировка по номерам гонок
//			}},
//		},
//	}
//
//	// Выполняем запрос агрегации
//	cursor, err := collection.Aggregate(context.TODO(), pipeline)
//	if err != nil {
//		return nil, repository_errors.SelectError
//	}
//	defer cursor.Close(context.TODO())
//
//	// Обрабатываем результаты
//	var ratingTableDB []RatingTableDB
//	if err = cursor.All(context.TODO(), &ratingTableDB); err != nil {
//		return nil, repository_errors.SelectError
//	} else if len(ratingTableDB) == 0 {
//		return nil, repository_errors.DoesNotExist
//	}
//
//	// Преобразуем результат в модель
//	ratingTable := copyRatingTableResultToModel(ratingTableDB)
//
//	// Ранжирование (присваиваем ранги вручную)
//	for i := range ratingTable {
//		if i > 0 && ratingTable[i].PointsSum == ratingTable[i-1].PointsSum {
//			ratingTable[i].Rank = ratingTable[i-1].Rank
//		} else {
//			ratingTable[i].Rank = i + 1
//		}
//	}
//
//	return ratingTable, nil
//}

func calculatePointsSum(ratingTable []RatingTableDB) {
	// Создаем карту для хранения суммы очков по каждому CrewID
	pointsMap := make(map[int]int)

	// Сначала проходим по массиву и суммируем очки для каждой команды
	for _, item := range ratingTable {
		pointsMap[item.SailNum] += item.Points
	}

	// Далее обновляем PointsSum для каждой записи
	for i, item := range ratingTable {
		ratingTable[i].PointsSum = pointsMap[item.SailNum]
	}
}
