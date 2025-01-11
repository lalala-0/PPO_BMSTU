package postgres

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type CrewResInRaceDB struct {
	ID               uuid.UUID `db:"id"`
	CrewID           uuid.UUID `db:"crew_id"`
	RaceID           uuid.UUID `db:"race_id"`
	Points           int       `db:"points"`
	SpecCircumstance int       `db:"spec_circumstance"`
}

type CrewResInRaceRepository struct {
	db *TracedDB
}

func NewCrewResInRaceRepository(db *TracedDB) repository_interfaces.ICrewResInRaceRepository {
	return &CrewResInRaceRepository{db: db}
}

func copyCrewResInRaceResultToModel(crewResInRaceDB *CrewResInRaceDB) *models.CrewResInRace {
	return &models.CrewResInRace{
		CrewID:           crewResInRaceDB.CrewID,
		RaceID:           crewResInRaceDB.RaceID,
		Points:           crewResInRaceDB.Points,
		SpecCircumstance: crewResInRaceDB.SpecCircumstance,
	}
}

func (w CrewResInRaceRepository) Create(ctx context.Context, crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	query := `INSERT INTO crew_race(crew_id, race_id, points, spec_circumstance) VALUES ($1, $2, $3, $4);`

	_, err := w.db.ExecContext(ctx, query, crewResInRace.CrewID, crewResInRace.RaceID, crewResInRace.Points, crewResInRace.SpecCircumstance)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.CrewResInRace{
		CrewID:           crewResInRace.CrewID,
		RaceID:           crewResInRace.RaceID,
		Points:           crewResInRace.Points,
		SpecCircumstance: crewResInRace.SpecCircumstance,
	}, nil
}

func (w CrewResInRaceRepository) Delete(ctx context.Context, raceID uuid.UUID, crewID uuid.UUID) error {
	query := `DELETE FROM crew_race WHERE race_id = $1 and crew_id = $2;`
	res, err := w.db.ExecContext(ctx, query, raceID, crewID)

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

func (w CrewResInRaceRepository) Update(ctx context.Context, crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	query := `UPDATE crew_race SET points = $3, spec_circumstance = $4 WHERE race_id = $1 and crew_id = $2 RETURNING race_id, crew_id, points, spec_circumstance;`

	var updatedCrewResInRace models.CrewResInRace
	err := w.db.QueryRowContext(ctx, query, crewResInRace.RaceID, crewResInRace.CrewID, crewResInRace.Points, crewResInRace.SpecCircumstance).Scan(&updatedCrewResInRace.RaceID, &updatedCrewResInRace.CrewID, &updatedCrewResInRace.Points, &updatedCrewResInRace.SpecCircumstance)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return &updatedCrewResInRace, nil
}

func (w CrewResInRaceRepository) GetCrewResByRaceIDAndCrewID(ctx context.Context, raceID uuid.UUID, crewID uuid.UUID) (*models.CrewResInRace, error) {
	query := `SELECT * FROM crew_race WHERE race_id = $1 and crew_id = $2;`
	crewResInRaceDB := &CrewResInRaceDB{}
	err := w.db.GetContext(ctx, crewResInRaceDB, query, raceID, crewID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	crewResInRaceModels := copyCrewResInRaceResultToModel(crewResInRaceDB)

	return crewResInRaceModels, nil
}

func (c CrewResInRaceRepository) GetAllCrewResInRace(ctx context.Context, raceID uuid.UUID) ([]models.CrewResInRace, error) {
	query := `SELECT * FROM crew_race WHERE race_id = $1;`
	var crewResInRaceDB []CrewResInRaceDB
	err := c.db.SelectContext(ctx, &crewResInRaceDB, query, raceID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	var crewResInRaceModels []models.CrewResInRace
	for i := range crewResInRaceDB {
		crewResInRace := copyCrewResInRaceResultToModel(&crewResInRaceDB[i])
		crewResInRaceModels = append(crewResInRaceModels, *crewResInRace)
	}
	return crewResInRaceModels, nil
}
