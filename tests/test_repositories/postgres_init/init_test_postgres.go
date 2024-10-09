package postgres_init

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/postgres"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

// PostgresConnection хранит соединение с базой данных PostgreSQL
type PostgresConnection struct {
	DB *sql.DB
}

// NewPostgresConnection создает новый экземпляр PostgresConnection
func NewPostgresConnection(db *sql.DB) *PostgresConnection {
	return &PostgresConnection{DB: db}
}

func CreateParticipant(fields *postgres.PostgresConnection) *models.Participant {

	query := `INSERT INTO participants(name, category, gender, birthdate, coach_name) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var participantID uuid.UUID
	err := fields.DB.QueryRow(query, "test", models.Junior2category, true, time.Date(2003, time.November, 10, 23, 0, 0, 0, time.UTC), "Test").Scan(&participantID)

	if err != nil {
		return nil
	}

	return &models.Participant{
		ID:       participantID,
		FIO:      "test",
		Gender:   models.Female,
		Category: models.Junior2category,
		Coach:    "Test",
		Birthday: time.Date(2003, time.November, 10, 23, 0, 0, 0, time.UTC),
	}
}

func CreateRating(fields *postgres.PostgresConnection) *models.Rating {
	query := `INSERT INTO ratings(name, class, blowout_cnt) VALUES ($1, $2, $3) RETURNING id;`

	var ratingID uuid.UUID
	err := fields.DB.QueryRow(query, "Name", models.Laser, 1).Scan(&ratingID)

	if err != nil {
		return nil
	}

	return &models.Rating{
		ID:         ratingID,
		Name:       "Name",
		Class:      models.Laser,
		BlowoutCnt: 1,
	}
}

func CreateCrew(fields *postgres.PostgresConnection, ratingID uuid.UUID) *models.Crew {
	query := `INSERT INTO crews(rating_id, class, sail_num) VALUES ($1, $2, $3) RETURNING id;`

	var crewID uuid.UUID
	err := fields.DB.QueryRow(query, ratingID, 2, 123).Scan(&crewID)

	if err != nil {
		return nil
	}

	return &models.Crew{
		ID:       crewID,
		RatingID: ratingID,
		SailNum:  123,
		Class:    2,
	}
}

func CreateCrewResInRace(fields *postgres.PostgresConnection, crewID uuid.UUID, raceID uuid.UUID) *models.CrewResInRace {
	query := `INSERT INTO crew_race(crew_id, race_id, points, spec_circumstance) VALUES ($1, $2, $3, $4);`

	_, err := fields.DB.Exec(query, crewID, raceID, 12, 0)

	if err != nil {
		return nil
	}

	return &models.CrewResInRace{
		CrewID:           crewID,
		RaceID:           raceID,
		Points:           12,
		SpecCircumstance: 0,
	}
}

func CreateRace(fields *postgres.PostgresConnection, ratingID uuid.UUID) *models.Race {
	query := `INSERT INTO races(rating_id, date, number, class) VALUES ($1, $2, $3, $4) RETURNING id;`

	var raceID uuid.UUID
	err := fields.DB.QueryRow(query, ratingID, time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC), 1, 4).Scan(&raceID)

	if err != nil {
		return nil
	}

	return &models.Race{
		ID:       raceID,
		RatingID: ratingID,
		Date:     time.Date(2012, time.November, 10, 23, 0, 0, 0, time.UTC),
		Number:   1,
		Class:    4,
	}
}

func CreateJudge(fields *postgres.PostgresConnection) *models.Judge {
	query := `INSERT INTO judges(name, login, password, role, post) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var judgeID uuid.UUID
	err := fields.DB.QueryRow(query, "Test", "Test", "test123", 1, "Test").Scan(&judgeID)

	if err != nil {
		return nil
	}

	return &models.Judge{
		ID:       judgeID,
		FIO:      "Test",
		Login:    "Test",
		Password: "test123",
		Post:     "Test",
		Role:     1,
	}
}

func CreateProtest(fields *postgres.PostgresConnection, raceID uuid.UUID, judgeID uuid.UUID, ratingID uuid.UUID) *models.Protest {
	query := `INSERT INTO protests(race_id, rating_id, judge_id, rule_num, review_date, status, comment) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	var protestID uuid.UUID
	err := fields.DB.QueryRow(query, raceID, ratingID, judgeID, 23, time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC), 1, "").Scan(&protestID)

	if err != nil {
		return nil
	}

	return &models.Protest{
		ID:         protestID,
		RaceID:     raceID,
		JudgeID:    judgeID,
		RatingID:   ratingID,
		RuleNum:    23,
		ReviewDate: time.Date(2024, time.November, 10, 23, 0, 0, 0, time.UTC),
		Status:     1,
		Comment:    "",
	}
}

func AttachCrewToProtest(fields *postgres.PostgresConnection, crewID uuid.UUID, protestID uuid.UUID) {
	query := `INSERT INTO crew_protest(crew_id, protest_id, crew_status) VALUES ($1, $2, $3);`

	fields.DB.Exec(query, crewID, protestID, 1)
}

func AttachCrewToProtestStatus(fields *postgres.PostgresConnection, crewID uuid.UUID, protestID uuid.UUID, status int) {
	query := `INSERT INTO crew_protest(crew_id, protest_id, crew_status) VALUES ($1, $2, $3);`

	fields.DB.Exec(query, crewID, protestID, status)
}

func AttachJudgeToRating(fields *postgres.PostgresConnection, judgeID uuid.UUID, ratingID uuid.UUID) {
	query := `INSERT INTO judge_rating(judge_id, rating_id) VALUES ($1, $2);`

	fields.DB.Exec(query, judgeID, ratingID)
}

func AttachParticipantToCrew(fields *postgres.PostgresConnection, participantID uuid.UUID, crewID uuid.UUID) {
	query := `INSERT INTO participant_crew(participant_id, crew_id, helmsman, active) VALUES ($1, $2, $3, $4);`

	fields.DB.Exec(query, participantID, crewID, false, true)
}
