package repository

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CrewDB struct {
	ID       uuid.UUID `db:"id"`
	RatingID uuid.UUID `db:"rating_id"`
	Class    int       `db:"class"`
	SailNum  int       `db:"sail_num"`
}

type CrewRepository struct {
	db *sqlx.DB
}

func NewCrewRepository(db *sqlx.DB) repository_interfaces.ICrewRepository {
	return &CrewRepository{db: db}
}

func copyCrewResultToModel(crewDB *CrewDB) *models.Crew {
	return &models.Crew{
		ID:       crewDB.ID,
		RatingID: crewDB.RatingID,
		SailNum:  crewDB.SailNum,
		Class:    crewDB.Class,
	}
}

func (w CrewRepository) Create(crew *models.Crew) (*models.Crew, error) {
	query := `INSERT INTO crews(rating_id, class, sail_num) VALUES ($1, $2, $3) RETURNING id;`

	var crewID uuid.UUID
	err := w.db.QueryRow(query, crew.RatingID, crew.Class, crew.SailNum).Scan(&crewID)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Crew{
		ID:       crewID,
		RatingID: crew.RatingID,
		Class:    crew.Class,
		SailNum:  crew.SailNum,
	}, nil
}

func (w CrewRepository) Update(crew *models.Crew) (*models.Crew, error) {
	query := `UPDATE crews SET rating_id = $1, class = $2, sail_num = $3 WHERE id = $4 RETURNING id, rating_id, class, sail_num;`

	var updatedCrew models.Crew
	err := w.db.QueryRow(query, crew.RatingID, crew.Class, crew.SailNum, crew.ID).Scan(&updatedCrew.ID, &updatedCrew.RatingID, &updatedCrew.Class, &updatedCrew.SailNum)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return &updatedCrew, nil
}

func (w CrewRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM crews WHERE id = $1;`
	_, err := w.db.Exec(query, id)

	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w CrewRepository) GetCrewDataByID(id uuid.UUID) (*models.Crew, error) {
	query := `SELECT * FROM crews WHERE id = $1;`
	crewDB := &CrewDB{}
	err := w.db.Get(crewDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	crewModels := copyCrewResultToModel(crewDB)

	return crewModels, nil
}

func (w CrewRepository) AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error {
	query := `INSERT INTO participant_crew(participant_id, crew_id, helmsman, active) VALUES ($1, $2, $3, $4);`

	helmsmanBool := helmsman != 0
	_, err := w.db.Exec(query, participantID, crewID, helmsmanBool, true)

	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (w CrewRepository) DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error {
	query := `DELETE FROM participant_crew WHERE participant_id = $1 and crew_id = $2;`
	_, err := w.db.Exec(query, participantID, crewID)

	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w CrewRepository) ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error {
	query := `UPDATE participant_crew SET helmsman = $3, active = $4 WHERE participant_id = $1 and crew_id = $2;`
	helmsmanBool := helmsman != 0
	activeBool := active != 0

	_, err := w.db.Exec(query, participantID, crewID, helmsmanBool, activeBool)
	if err != nil {
		return repository_errors.UpdateError
	}
	return nil
}

func (w CrewRepository) GetCrewsDataByRatingID(id uuid.UUID) ([]models.Crew, error) {
	query := `SELECT * FROM crews WHERE rating_id = $1;`
	var crewDB []CrewDB
	err := w.db.Select(&crewDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	var crewModels []models.Crew
	for i := range crewDB {
		crew := copyCrewResultToModel(&crewDB[i])
		crewModels = append(crewModels, *crew)
	}

	return crewModels, nil
}

func (w CrewRepository) GetCrewsDataByProtestID(id uuid.UUID) ([]models.Crew, error) {
	query := `SELECT * FROM crews WHERE id in (SELECT crew_id from crew_protest where protest_id = $1);`
	var crewDB []CrewDB
	err := w.db.Select(&crewDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	var crewModels []models.Crew
	for i := range crewDB {
		crew := copyCrewResultToModel(&crewDB[i])
		crewModels = append(crewModels, *crew)
	}
	return crewModels, nil
}

func (w CrewRepository) GetCrewDataBySailNumAndRatingID(sailNum int, ratingID uuid.UUID) (*models.Crew, error) {
	query := `SELECT * FROM crews WHERE sail_num = $1 and rating_id = $2;`
	crewDB := &CrewDB{}
	err := w.db.Get(crewDB, query, sailNum, ratingID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	crewModels := copyCrewResultToModel(crewDB)

	return crewModels, nil
}
