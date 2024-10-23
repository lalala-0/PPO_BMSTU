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

type RaceDB struct {
	ID       uuid.UUID `db:"id"`
	RatingID uuid.UUID `db:"rating_id"`
	Date     time.Time `db:"date"`
	Number   int       `db:"number"`
	Class    int       `db:"class"`
}

type RaceRepository struct {
	db *sqlx.DB
}

func NewRaceRepository(db *sqlx.DB) repository_interfaces.IRaceRepository {
	return &RaceRepository{db: db}
}

func copyRaceResultToModel(raceDB *RaceDB) *models.Race {
	return &models.Race{
		ID:       raceDB.ID,
		RatingID: raceDB.RatingID,
		Date:     raceDB.Date,
		Number:   raceDB.Number,
		Class:    raceDB.Class,
	}
}

func (w RaceRepository) Create(race *models.Race) (*models.Race, error) {
	query := `INSERT INTO races(rating_id, date, number, class) VALUES ($1, $2, $3, $4) RETURNING id;`

	var raceID uuid.UUID
	err := w.db.QueryRow(query, race.RatingID, race.Date, race.Number, race.Class).Scan(&raceID)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Race{
		ID:       raceID,
		RatingID: race.RatingID,
		Date:     race.Date,
		Number:   race.Number,
		Class:    race.Class,
	}, nil
}

func (w RaceRepository) Update(race *models.Race) (*models.Race, error) {
	query := `UPDATE races SET rating_id = $1, date = $2, number = $3, class = $4 WHERE id = $5 RETURNING id, rating_id, date, number, class;`

	var updatedRace models.Race
	err := w.db.QueryRow(query, race.RatingID, race.Date, race.Number, race.Class, race.ID).Scan(&updatedRace.ID, &updatedRace.RatingID, &updatedRace.Date, &updatedRace.Number, &updatedRace.Class)

	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return &updatedRace, nil
}

func (w RaceRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM races WHERE id = $1;`
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

func (w RaceRepository) GetRaceDataByID(id uuid.UUID) (*models.Race, error) {
	query := `SELECT * FROM races WHERE id = $1;`
	raceDB := &RaceDB{}
	err := w.db.Get(raceDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	raceModels := copyRaceResultToModel(raceDB)

	return raceModels, nil
}

func (w RaceRepository) GetRacesDataByRatingID(id uuid.UUID) ([]models.Race, error) {
	query := `SELECT * FROM races WHERE rating_id = $1;`
	var raceDB []RaceDB
	err := w.db.Select(&raceDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	var raceModels []models.Race
	for i := range raceDB {
		race := copyRaceResultToModel(&raceDB[i])
		raceModels = append(raceModels, *race)
	}

	return raceModels, nil
}
