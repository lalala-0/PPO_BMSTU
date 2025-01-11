package postgres

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type RatingDB struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Class      int       `db:"class"`
	BlowoutCnt int       `db:"blowout_cnt"`
}

type RatingTableDB struct {
	CrewID               string    `db:"crew_id"`
	SailNum              int       `db:"sail_num"`
	ParticipantName      string    `db:"name"`
	ParticipantBirthDate time.Time `db:"birthdate"`
	ParticipantCategory  int       `db:"category"`
	Points               int       `db:"points"`
	RaceNumber           int       `db:"number"`
	PointsSum            int       `db:"points_summ"`
	Rank                 int       `db:"rang"`
	CoachName            string    `db:"coach_name"`
}

type RatingRepository struct {
	db *TracedDB
}

func NewRatingRepository(db *TracedDB) repository_interfaces.IRatingRepository {
	return &RatingRepository{db: db}
}

func copyRatingResultToModel(ratingDB *RatingDB) *models.Rating {
	return &models.Rating{
		ID:         ratingDB.ID,
		Name:       ratingDB.Name,
		Class:      ratingDB.Class,
		BlowoutCnt: ratingDB.BlowoutCnt,
	}
}

func copyRatingTableResultToModel(ratingTableDB []RatingTableDB) []models.RatingTableLine {
	var ratingLines []models.RatingTableLine

	if len(ratingTableDB) == 0 {
		return ratingLines
	}

	// Переменные для текущего `SailNum`
	var participantNames []string
	var participantBirthDates []time.Time
	var participantCategories []int
	var coachNames []string
	resInRace := make(map[int]int)

	// Текущий `SailNum`
	currentSailNum := ratingTableDB[0].SailNum
	currentCrewID := ratingTableDB[0].CrewID
	pointsSum := ratingTableDB[0].PointsSum
	rank := ratingTableDB[0].Rank

	for _, line := range ratingTableDB {
		// Если `SailNum` изменился, добавляем текущую строку в результат
		if line.SailNum != currentSailNum {
			ratingLines = append(ratingLines, models.RatingTableLine{
				CrewID:                currentCrewID,
				SailNum:               currentSailNum,
				ParticipantNames:      participantNames,
				ParticipantBirthDates: participantBirthDates,
				ParticipantCategories: participantCategories,
				ResInRace:             resInRace,
				PointsSum:             pointsSum,
				Rank:                  rank,
				CoachNames:            coachNames,
			})

			// Очищаем переменные для нового `SailNum`
			participantNames = []string{}
			participantBirthDates = []time.Time{}
			participantCategories = []int{}
			coachNames = []string{}
			resInRace = make(map[int]int)

			// Обновляем текущий `SailNum` и связанные данные
			currentSailNum = line.SailNum
			currentCrewID = line.CrewID
			pointsSum = line.PointsSum
			rank = line.Rank
		}

		// Добавляем данные текущей строки
		participantNames = append(participantNames, line.ParticipantName)
		participantBirthDates = append(participantBirthDates, line.ParticipantBirthDate)
		participantCategories = append(participantCategories, line.ParticipantCategory)
		coachNames = append(coachNames, line.CoachName)
		resInRace[line.RaceNumber] = line.Points
	}

	// Добавляем последнюю группу данных
	ratingLines = append(ratingLines, models.RatingTableLine{
		CrewID:                currentCrewID,
		SailNum:               currentSailNum,
		ParticipantNames:      participantNames,
		ParticipantBirthDates: participantBirthDates,
		ParticipantCategories: participantCategories,
		ResInRace:             resInRace,
		PointsSum:             pointsSum,
		Rank:                  rank,
		CoachNames:            coachNames,
	})

	return ratingLines
}

func (w RatingRepository) Create(ctx context.Context, rating *models.Rating) (*models.Rating, error) {
	query := `INSERT INTO ratings(name, class, blowout_cnt) VALUES ($1, $2, $3) RETURNING id;`

	var ratingID uuid.UUID
	err := w.db.QueryRowContext(ctx, query, rating.Name, rating.Class, rating.BlowoutCnt).Scan(&ratingID)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Rating{
		ID:         ratingID,
		Name:       rating.Name,
		Class:      rating.Class,
		BlowoutCnt: rating.BlowoutCnt,
	}, nil
}

func (w RatingRepository) Update(ctx context.Context, rating *models.Rating) (*models.Rating, error) {
	query := `UPDATE ratings SET name = $1, class = $2, blowout_cnt = $3 WHERE ratings.id = $4 RETURNING id, name, class, blowout_cnt;`

	var updatedRating models.Rating
	err := w.db.QueryRowContext(ctx, query, rating.Name, rating.Class, rating.BlowoutCnt, rating.ID).Scan(&updatedRating.ID, &updatedRating.Name, &updatedRating.Class, &updatedRating.BlowoutCnt)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return &updatedRating, nil
}

func (w RatingRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM ratings WHERE id = $1;`
	res, err := w.db.ExecContext(ctx, query, id)

	if err != nil {
		return repository_errors.DeleteError
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository_errors.DoesNotExist
	}
	return nil
}

func (w RatingRepository) GetRatingDataByID(ctx context.Context, id uuid.UUID) (*models.Rating, error) {
	query := `SELECT * FROM ratings WHERE id = $1;`
	ratingDB := &RatingDB{}

	// Используем GetContext для выполнения запроса с учетом контекста
	err := w.db.GetContext(ctx, ratingDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	ratingModel := copyRatingResultToModel(ratingDB)

	return ratingModel, nil
}

func (w RatingRepository) AttachJudgeToRating(ctx context.Context, ratingID uuid.UUID, judgeID uuid.UUID) error {
	query := `INSERT INTO judge_rating(judge_id, rating_id) VALUES ($1, $2);`

	_, err := w.db.ExecContext(ctx, query, judgeID, ratingID)
	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (w RatingRepository) DetachJudgeFromRating(ctx context.Context, ratingID uuid.UUID, judgeID uuid.UUID) error {
	query := `DELETE FROM judge_rating WHERE judge_id = $1 AND rating_id = $2;`
	res, err := w.db.ExecContext(ctx, query, judgeID, ratingID)

	if err != nil {
		return repository_errors.DeleteError
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository_errors.DoesNotExist
	}

	return nil
}

func (w RatingRepository) GetAllRatings(ctx context.Context) ([]models.Rating, error) {
	query := `SELECT * FROM ratings ORDER BY name;`
	ratingDB := []RatingDB{}

	// Используем контекст в запросе
	err := w.db.SelectContext(ctx, &ratingDB, query)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	var ratingModels []models.Rating
	for i := range ratingDB {
		rating := copyRatingResultToModel(&ratingDB[i])
		ratingModels = append(ratingModels, *rating)
	}
	return ratingModels, nil
}

func (r RatingRepository) GetRatingTable(ctx context.Context, id uuid.UUID) ([]models.RatingTableLine, error) {
	query :=
		`SELECT c.id AS crew_id, c.sail_num, cp.name, cp.birthdate, category, points, number, points_summ, 
			DENSE_RANK() OVER (ORDER BY points_summ, id) AS rang, coach_name
		FROM crews AS c
		JOIN (
			SELECT crew_id, pc.name, pc.birthdate, pc.category, pc.gender, pc.coach_name
			FROM participant_crew
			JOIN (
				SELECT id, name, birthdate, category, gender, coach_name
				FROM participants
			) AS pc ON pc.id = participant_crew.participant_id AND participant_crew.helmsman
		) AS cp ON c.id = cp.crew_id
		JOIN (
			SELECT points, crew_id, race_id, rc.number
			FROM crew_race
			JOIN (
				SELECT id, number
				FROM races AS r
				GROUP BY r.id
			) AS rc ON crew_race.race_id = rc.id
			GROUP BY crew_id, points, crew_id, race_id, rc.number
			ORDER BY rc.number
		) AS cr ON c.id = cr.crew_id
		JOIN (
			SELECT SUM(points) AS points_summ, crew_id
			FROM crew_race
			GROUP BY crew_id
		) AS ps ON c.id = ps.crew_id
		WHERE rating_id = $1
		ORDER BY rang, sail_num, number;
	`

	ratingTableDB := []RatingTableDB{}
	err := r.db.SelectContext(ctx, &ratingTableDB, query, id)
	if err != nil {
		return nil, repository_errors.SelectError
	} else if len(ratingTableDB) == 0 {
		return nil, repository_errors.DoesNotExist
	}

	ratingTable := copyRatingTableResultToModel(ratingTableDB)
	return ratingTable, nil
}
