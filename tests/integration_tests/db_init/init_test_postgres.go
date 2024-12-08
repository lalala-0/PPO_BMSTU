package db_init

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// PostgresRepository - реализация интерфейса TestRepositoryInitializer
type PostgresRepository struct {
	DB *sqlx.DB
}

// NewPostgresRepository создает новый экземпляр PostgresRepository
func NewPostgresRepository(db *sqlx.DB) TestRepositoryInitializer {
	return &PostgresRepository{DB: db}
}

// Реализация методов интерфейса TestRepositoryInitializer
func (r *PostgresRepository) CreateParticipant(participant *models.Participant) (*models.Participant, error) {
	var gender bool
	if participant.Gender == 0 {
		gender = false
	} else {
		gender = true
	}
	query := `INSERT INTO participants(id, name, category, gender, birthdate, coach_name) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	var participantID uuid.UUID
	err := r.DB.QueryRow(query, participant.ID, participant.FIO, participant.Category, gender, participant.Birthday, participant.Coach).Scan(&participantID)
	if err != nil {
		return nil, err
	}
	participant.ID = participantID
	return participant, nil
}

func (r *PostgresRepository) CreateRating(rating *models.Rating) (*models.Rating, error) {
	query := `INSERT INTO ratings(id, name, class, blowout_cnt) VALUES ($1, $2, $3, $4) RETURNING id;`
	var ratingID uuid.UUID
	err := r.DB.QueryRow(query, rating.ID, rating.Name, rating.Class, rating.BlowoutCnt).Scan(&ratingID)
	if err != nil {
		return nil, err
	}
	rating.ID = ratingID
	return rating, nil
}

func (r *PostgresRepository) CreateCrew(crew *models.Crew) (*models.Crew, error) {
	query := `INSERT INTO crews(id, rating_id, class, sail_num) VALUES ($1, $2, $3, $4) RETURNING id;`
	var crewID uuid.UUID
	err := r.DB.QueryRow(query, crew.ID, crew.RatingID, crew.Class, crew.SailNum).Scan(&crewID)
	if err != nil {
		return nil, err
	}
	crew.ID = crewID
	return crew, nil
}

func (r *PostgresRepository) CreateCrewResInRace(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	query := `INSERT INTO crew_race(crew_id, race_id, points, spec_circumstance) VALUES ($1, $2, $3, $4);`
	_, err := r.DB.Exec(query, crewResInRace.CrewID, crewResInRace.RaceID, crewResInRace.Points, crewResInRace.SpecCircumstance)
	if err != nil {
		return nil, err
	}
	return crewResInRace, nil
}

func (r *PostgresRepository) CreateRace(race *models.Race) (*models.Race, error) {
	query := `INSERT INTO races(id, rating_id, date, number, class) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	var raceID uuid.UUID
	err := r.DB.QueryRow(query, race.ID, race.RatingID, race.Date, race.Number, race.Class).Scan(&raceID)
	if err != nil {
		return nil, err
	}
	race.ID = raceID
	return race, nil
}

func (r *PostgresRepository) CreateJudge(judge *models.Judge) (*models.Judge, error) {
	query := `INSERT INTO judges(id, name, login, password, role, post) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	var judgeID uuid.UUID
	err := r.DB.QueryRow(query, judge.ID, judge.FIO, judge.Login, judge.Password, judge.Role, judge.Post).Scan(&judgeID)
	if err != nil {
		return nil, err
	}
	judge.ID = judgeID
	return judge, nil
}

func (r *PostgresRepository) CreateProtest(protest *models.Protest) (*models.Protest, error) {
	query := `INSERT INTO protests(id, race_id, rating_id, judge_id, rule_num, review_date, status, comment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	var protestID uuid.UUID
	err := r.DB.QueryRow(query, protest.ID, protest.RaceID, protest.RatingID, protest.JudgeID, protest.RuleNum, protest.ReviewDate, protest.Status, protest.Comment).Scan(&protestID)
	if err != nil {
		return nil, err
	}
	protest.ID = protestID
	return protest, nil
}

func (r *PostgresRepository) AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID) error {
	query := `INSERT INTO crew_protest(crew_id, protest_id) VALUES ($1, $2);`
	_, err := r.DB.Exec(query, crewID, protestID)
	return err
}

func (r *PostgresRepository) AttachCrewToProtestStatus(crewID uuid.UUID, protestID uuid.UUID, status int) error {
	query := `UPDATE crew_protest SET crew_status = $3 WHERE crew_id = $1 AND protest_id = $2;`
	_, err := r.DB.Exec(query, crewID, protestID, status)
	return err
}

func (r *PostgresRepository) AttachJudgeToRating(judgeID uuid.UUID, ratingID uuid.UUID) error {
	query := `INSERT INTO judge_rating(judge_id, rating_id) VALUES ($1, $2);`
	_, err := r.DB.Exec(query, judgeID, ratingID)
	return err
}

func (r *PostgresRepository) AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID) error {
	query := `INSERT INTO participant_crew(participant_id, crew_id) VALUES ($1, $2);`
	_, err := r.DB.Exec(query, participantID, crewID)
	return err
}
func (r *PostgresRepository) ClearAll() error {
	// Начинаем транзакцию
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	// Список таблиц для очистки
	tables := []string{
		"participant_crew",
		"crew_protest",
		"judge_rating",
		"crew_race",
		"protests",
		"judges",
		"crews",
		"ratings",
		"participants",
		"races",
	}

	// Удаление данных из каждой таблицы
	for _, table := range tables {
		if _, err = tx.Exec("DELETE FROM " + table + ";"); err != nil {
			return err
		}
	}

	return nil
}
