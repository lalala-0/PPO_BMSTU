package postgres

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

type ProtestDB struct {
	ID         uuid.UUID `db:"id"`
	RaceID     uuid.UUID `db:"race_id"`
	JudgeID    uuid.UUID `db:"judge_id"`
	RatingID   uuid.UUID `db:"rating_id"`
	RuleNum    int       `db:"rule_num"`
	ReviewDate time.Time `db:"review_date"`
	Status     int       `db:"status"`
	Comment    string    `db:"comment"`
}

type ProtestRepository struct {
	db *sqlx.DB
}

func NewProtestRepository(db *sqlx.DB) repository_interfaces.IProtestRepository {
	return &ProtestRepository{db: db}
}

func copyProtestResultToModel(protestDB *ProtestDB) *models.Protest {
	return &models.Protest{
		ID:         protestDB.ID,
		RaceID:     protestDB.RaceID,
		JudgeID:    protestDB.JudgeID,
		RatingID:   protestDB.RatingID,
		RuleNum:    protestDB.RuleNum,
		ReviewDate: protestDB.ReviewDate,
		Status:     protestDB.Status,
		Comment:    protestDB.Comment,
	}
}

func (w ProtestRepository) Create(protest *models.Protest) (*models.Protest, error) {
	query := `INSERT INTO protests(race_id, rating_id, judge_id, rule_num, review_date, status, comment) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	var protestID uuid.UUID
	err := w.db.QueryRow(query, protest.RaceID, protest.RatingID, protest.JudgeID, protest.RuleNum, protest.ReviewDate, protest.Status, protest.Comment).Scan(&protestID)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Protest{
		ID:         protestID,
		RaceID:     protest.RaceID,
		JudgeID:    protest.JudgeID,
		RatingID:   protest.RatingID,
		RuleNum:    protest.RuleNum,
		ReviewDate: protest.ReviewDate,
		Status:     protest.Status,
		Comment:    protest.Comment,
	}, nil
}

func (w ProtestRepository) Update(protest *models.Protest) (*models.Protest, error) {
	query := `UPDATE protests SET race_id = $1, rating_id = $2, judge_id = $3, rule_num = $4, review_date = $5, status = $6, comment = $7 WHERE id = $8 RETURNING id, race_id, rating_id, judge_id, rule_num, review_date, status, comment;`

	var updatedProtest models.Protest
	err := w.db.QueryRow(query, protest.RaceID, protest.RatingID, protest.JudgeID, protest.RuleNum, protest.ReviewDate, protest.Status, protest.Comment, protest.ID).Scan(&updatedProtest.ID, &updatedProtest.RaceID, &updatedProtest.RatingID, &updatedProtest.JudgeID, &updatedProtest.RuleNum, &updatedProtest.ReviewDate, &updatedProtest.Status, &updatedProtest.Comment)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return &updatedProtest, nil
}

func (w ProtestRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM protests WHERE id = $1;`
	res, err := w.db.Exec(query, id)

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

func (w ProtestRepository) GetProtestDataByID(id uuid.UUID) (*models.Protest, error) {
	query := `SELECT * FROM protests WHERE id = $1;`
	protestDB := &ProtestDB{}
	err := w.db.Get(protestDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	protestModels := copyProtestResultToModel(protestDB)

	return protestModels, nil
}

func (w ProtestRepository) AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID, crewStatus int) error {
	query := `INSERT INTO crew_protest(crew_id, protest_id, crew_status) VALUES ($1, $2, $3);`

	_, err := w.db.Exec(query, crewID, protestID, crewStatus)

	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (w ProtestRepository) DetachCrewFromProtest(protestID uuid.UUID, crewID uuid.UUID) error {
	query := `DELETE FROM crew_protest WHERE crew_id = $1 and protest_id = $2;`
	res, err := w.db.Exec(query, crewID, protestID)

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

func (w ProtestRepository) GetProtestsDataByRaceID(id uuid.UUID) ([]models.Protest, error) {
	query := `SELECT * FROM protests WHERE race_id = $1;`
	var protestDB []ProtestDB
	err := w.db.Select(&protestDB, query, id)

	if err != nil {
		return nil, repository_errors.SelectError
	}

	var protestModels []models.Protest
	for i := range protestDB {
		protest := copyProtestResultToModel(&protestDB[i])
		protestModels = append(protestModels, *protest)
	}

	return protestModels, nil
}

func (w ProtestRepository) GetProtestParticipantsIDByID(id uuid.UUID) (map[uuid.UUID]int, error) {
	query := `SELECT crew_id, crew_status FROM crew_protest WHERE protest_id = $1;`
	var ids = make(map[uuid.UUID]int)
	rows, err := w.db.Query(query, id)

	if err != nil {
		return nil, repository_errors.SelectError
	}
	for rows.Next() {
		var status int
		var uuid uuid.UUID
		_ = rows.Scan(&uuid, &status)
		ids[uuid] = status
	}
	return ids, nil
}
