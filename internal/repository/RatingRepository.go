package repository

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type RatingDB struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Class      int       `db:"class"`
	BlowoutCnt int       `db:"blowout_cnt"`
}

type RatingTableDB struct {
	CrewID               uuid.UUID `db:"id"`
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
	db *sqlx.DB
}

func NewRatingRepository(db *sqlx.DB) repository_interfaces.IRatingRepository {
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

	var participantNames []string
	var participantBirthDates []time.Time
	var participantCategories []int
	var coachNames []string
	resInRace := make(map[int]int)
	for i, line := range ratingTableDB[0 : len(ratingTableDB)-1] {
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
	return ratingLines
}

func (w RatingRepository) Create(rating *models.Rating) (*models.Rating, error) {
	query := `INSERT INTO ratings(name, class, blowout_cnt) VALUES ($1, $2, $3) RETURNING id;`

	var ratingID uuid.UUID
	err := w.db.QueryRow(query, rating.Name, rating.Class, rating.BlowoutCnt).Scan(&ratingID)

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

func (w RatingRepository) Update(rating *models.Rating) (*models.Rating, error) {
	query := `UPDATE ratings SET name = $1, class = $2, blowout_cnt = $3 WHERE ratings.id = $4 RETURNING id, name, class, blowout_cnt;`

	var updatedRating models.Rating
	err := w.db.QueryRow(query, rating.Name, rating.Class, rating.BlowoutCnt, rating.ID).Scan(&updatedRating.ID, &updatedRating.Name, &updatedRating.Class, &updatedRating.BlowoutCnt)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return &updatedRating, nil
}

func (w RatingRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM ratings WHERE id = $1;`
	_, err := w.db.Exec(query, id)

	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w RatingRepository) GetRatingDataByID(id uuid.UUID) (*models.Rating, error) {
	query := `SELECT * FROM ratings WHERE id = $1;`
	ratingDB := &RatingDB{}
	err := w.db.Get(ratingDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	ratingModels := copyRatingResultToModel(ratingDB)

	return ratingModels, nil
}

func (w RatingRepository) AttachJudgeToRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	query := `INSERT INTO judge_rating(judge_id, rating_id) VALUES ($1, $2);`

	_, err := w.db.Exec(query, judgeID, ratingID)

	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (w RatingRepository) DetachJudgeFromRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	query := `DELETE FROM judge_rating WHERE judge_id = $1 and rating_id = $2;`
	_, err := w.db.Exec(query, judgeID, ratingID)

	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w RatingRepository) GetAllRatings() ([]models.Rating, error) {
	query := `SELECT * FROM ratings ORDER BY name;`
	ratingDB := []RatingDB{}
	err := w.db.Select(&ratingDB, query)

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

func (r RatingRepository) GetRatingTable(id uuid.UUID) ([]models.RatingTableLine, error) {
	query :=
		`select c.id, c.sail_num, cp.name, cp.birthdate, category, points, number, points_summ, dense_rank() OVER (order BY points_summ, id) as rang, coach_name
		from crews as c
		join(
		select crew_id, pc.name, pc.birthdate, pc.category, pc.gender, pc.coach_name
			from participant_crew
			join(select id, name, birthdate, category, gender, coach_name
				from participants
				)as pc on pc.id = participant_crew.participant_id and participant_crew.helmsman
				) as cp on c.id = cp.crew_id
				join(
				select points, crew_id, race_id, rc.number
					from crew_race
					join (select id, number
						from races as r
						group by r.id
						)as rc on crew_race.race_id = rc.id
						GROUP BY crew_id, points, crew_id, race_id, rc.number
						order by rc.number 
						)as cr on c.id = cr.crew_id
						join( select sum(points) as points_summ , crew_id
							from crew_race
							GROUP BY crew_id
							)as ps on c.id = ps.crew_id		
		where rating_id = $1
		order by rang, sail_num, number;
	`
	ratingTableDB := []RatingTableDB{}
	err := r.db.Select(&ratingTableDB, query, id)
	if err != nil {
		return nil, repository_errors.SelectError
	} else if len(ratingTableDB) == 0 {
		return nil, repository_errors.DoesNotExist
	}

	ratingTable := copyRatingTableResultToModel(ratingTableDB)
	return ratingTable, nil
}
