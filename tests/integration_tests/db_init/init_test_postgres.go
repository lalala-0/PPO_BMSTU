package db_init

//
//import (
//	"PPO_BMSTU/internal/models"
//	"database/sql"
//	"github.com/google/uuid"
//	"time"
//)
//
//// PostgresRepository - реализация интерфейса TestsRepository
//type PostgresRepository struct {
//	DB *sql.DB
//}
//
//// NewPostgresRepository создает новый экземпляр PostgresRepository
//func NewPostgresRepository(db *sql.DB) *PostgresRepository {
//	return &PostgresRepository{DB: db}
//}
//
//// Реализация методов интерфейса Repository
//
//func (r *PostgresRepository) CreateParticipant(name string, category models.Category, gender models.Gender, birthdate time.Time, coachName string) (*models.Participant, error) {
//	query := `INSERT INTO participants(name, category, gender, birthdate, coach_name) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
//	var participantID uuid.UUID
//	err := r.DB.QueryRow(query, name, category, gender, birthdate, coachName).Scan(&participantID)
//	if err != nil {
//		return nil, err
//	}
//	return &models.Participant{
//		ID:       participantID,
//		FIO:      name,
//		Gender:   gender,
//		Category: category,
//		Coach:    coachName,
//		Birthday: birthdate,
//	}, nil
//}
//
//func (r *PostgresRepository) CreateRating(name string, class models.Class, blowoutCount int) (*models.Rating, error) {
//	query := `INSERT INTO ratings(name, class, blowout_cnt) VALUES ($1, $2, $3) RETURNING id;`
//	var ratingID uuid.UUID
//	err := r.DB.QueryRow(query, name, class, blowoutCount).Scan(&ratingID)
//	if err != nil {
//		return nil, err
//	}
//	return &models.Rating{
//		ID:         ratingID,
//		Name:       name,
//		Class:      class,
//		BlowoutCnt: blowoutCount,
//	}, nil
//}
//
//func (r *PostgresRepository) CreateCrew(ratingID uuid.UUID, class int, sailNum int) (*models.Crew, error) {
//	query := `INSERT INTO crews(rating_id, class, sail_num) VALUES ($1, $2, $3) RETURNING id;`
//	var crewID uuid.UUID
//	err := r.DB.QueryRow(query, ratingID, class, sailNum).Scan(&crewID)
//	if err != nil {
//		return nil, err
//	}
//	return &models.Crew{
//		ID:       crewID,
//		RatingID: ratingID,
//		SailNum:  sailNum,
//		Class:    class,
//	}, nil
//}
//
//func (r *PostgresRepository) CreateCrewResInRace(crewID uuid.UUID, raceID uuid.UUID, points int, specCircumstance int) (*models.CrewResInRace, error) {
//	query := `INSERT INTO crew_race(crew_id, race_id, points, spec_circumstance) VALUES ($1, $2, $3, $4);`
//	_, err := r.DB.Exec(query, crewID, raceID, points, specCircumstance)
//	if err != nil {
//		return nil, err
//	}
//	return &models.CrewResInRace{
//		CrewID:           crewID,
//		RaceID:           raceID,
//		Points:           points,
//		SpecCircumstance: specCircumstance,
//	}, nil
//}
//
//func (r *PostgresRepository) CreateRace(ratingID uuid.UUID, date time.Time, number int, class int) (*models.Race, error) {
//	query := `INSERT INTO races(rating_id, date, number, class) VALUES ($1, $2, $3, $4) RETURNING id;`
//	var raceID uuid.UUID
//	err := r.DB.QueryRow(query, ratingID, date, number, class).Scan(&raceID)
//	if err != nil {
//		return nil, err
//	}
//	return &models.Race{
//		ID:       raceID,
//		RatingID: ratingID,
//		Date:     date,
//		Number:   number,
//		Class:    class,
//	}, nil
//}
//
//func (r *PostgresRepository) CreateJudge(name string, login string, password string, role int, post string) (*models.Judge, error) {
//	query := `INSERT INTO judges(name, login, password, role, post) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
//	var judgeID uuid.UUID
//	err := r.DB.QueryRow(query, name, login, password, role, post).Scan(&judgeID)
//	if err != nil {
//		return nil, err
//	}
//	return &models.Judge{
//		ID:       judgeID,
//		FIO:      name,
//		Login:    login,
//		Password: password,
//		Post:     post,
//		Role:     role,
//	}, nil
//}
//
//func (r *PostgresRepository) CreateProtest(raceID uuid.UUID, ratingID uuid.UUID, judgeID uuid.UUID, ruleNum int, reviewDate time.Time, status int, comment string) (*models.Protest, error) {
//	query := `INSERT INTO protests(race_id, rating_id, judge_id, rule_num, review_date, status, comment) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
//	var protestID uuid.UUID
//	err := r.DB.QueryRow(query, raceID, ratingID, judgeID, ruleNum, reviewDate, status, comment).Scan(&protestID)
//	if err != nil {
//		return nil, err
//	}
//	return &models.Protest{
//		ID:         protestID,
//		RaceID:     raceID,
//		JudgeID:    judgeID,
//		RatingID:   ratingID,
//		RuleNum:    ruleNum,
//		ReviewDate: reviewDate,
//		Status:     status,
//		Comment:    comment,
//	}, nil
//}
//
//func (r *PostgresRepository) AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID, status int) error {
//	query := `INSERT INTO crew_protest(crew_id, protest_id, crew_status) VALUES ($1, $2, $3);`
//	_, err := r.DB.Exec(query, crewID, protestID, status)
//	return err
//}
//
//func (r *PostgresRepository) AttachJudgeToRating(judgeID uuid.UUID, ratingID uuid.UUID) error {
//	query := `INSERT INTO judge_rating(judge_id, rating_id) VALUES ($1, $2);`
//	_, err := r.DB.Exec(query, judgeID, ratingID)
//	return err
//}
//
//func (r *PostgresRepository) AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman bool, active bool) error {
//	query := `INSERT INTO participant_crew(participant_id, crew_id, helmsman, active) VALUES ($1, $2, $3, $4);`
//	_, err := r.DB.Exec(query, participantID, crewID, helmsman, active)
//	return err
//}
