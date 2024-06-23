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

type ParticipantDB struct {
	ID       uuid.UUID `db:"id"`
	FIO      string    `db:"name"`
	Category int       `db:"category"`
	Gender   bool      `db:"gender"`
	Birthday time.Time `db:"birthdate"`
	Coach    string    `db:"coach_name"`
}

type ParticipantRepository struct {
	db *sqlx.DB
}

func NewParticipantRepository(db *sqlx.DB) repository_interfaces.IParticipantRepository {
	return &ParticipantRepository{db: db}
}

func copyParticipantResultToModel(participantDB *ParticipantDB) *models.Participant {
	gender := 0
	if participantDB.Gender {
		gender = 1
	}

	return &models.Participant{

		ID:       participantDB.ID,
		FIO:      participantDB.FIO,
		Category: participantDB.Category,
		Gender:   gender,
		Birthday: participantDB.Birthday,
		Coach:    participantDB.Coach,
	}
}

func (w ParticipantRepository) Create(participant *models.Participant) (*models.Participant, error) {
	query := `INSERT INTO participants(name, category, gender, birthdate, coach_name) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var participantID uuid.UUID
	gender := participant.Gender != 0
	err := w.db.QueryRow(query, participant.FIO, participant.Category, gender, participant.Birthday, participant.Coach).Scan(&participantID)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Participant{
		ID:       participantID,
		FIO:      participant.FIO,
		Category: participant.Category,
		Gender:   participant.Gender,
		Birthday: participant.Birthday,
		Coach:    participant.Coach,
	}, nil
}

func (w ParticipantRepository) Update(participant *models.Participant) (*models.Participant, error) {
	query := `UPDATE participants SET name = $2, category = $3, gender = $4, birthdate = $5, coach_name = $6 WHERE id = $1 RETURNING id, name, category, gender, birthdate, coach_name;`
	gender := participant.Gender != 0

	participantDB := &ParticipantDB{}
	err := w.db.QueryRow(query, participant.ID, participant.FIO, participant.Category, gender, participant.Birthday, participant.Coach).Scan(&participantDB.ID, &participantDB.FIO, &participantDB.Category, &participantDB.Gender, &participantDB.Birthday, &participantDB.Coach)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	participantModels := copyParticipantResultToModel(participantDB)

	return participantModels, nil
}

func (w ParticipantRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM participants WHERE id = $1;`
	_, err := w.db.Exec(query, id)

	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w ParticipantRepository) GetParticipantDataByID(id uuid.UUID) (*models.Participant, error) {
	query := `SELECT * FROM participants WHERE id = $1;`
	participantDB := &ParticipantDB{}
	err := w.db.Get(participantDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	participantModels := copyParticipantResultToModel(participantDB)

	return participantModels, nil
}

func (w ParticipantRepository) GetParticipantsDataByCrewID(id uuid.UUID) ([]models.Participant, error) {
	query := `SELECT * FROM participants WHERE id IN (SELECT participant_id FROM participant_crew WHERE crew_id = $1);`
	var participantDB []ParticipantDB
	err := w.db.Select(&participantDB, query, id)

	if err != nil {
		return nil, repository_errors.SelectError
	}

	var participantModels []models.Participant
	for i := range participantDB {
		participant := copyParticipantResultToModel(&participantDB[i])
		participantModels = append(participantModels, *participant)
	}

	return participantModels, nil
}

func (w ParticipantRepository) GetParticipantsDataByProtestID(id uuid.UUID) ([]models.Participant, error) {
	query := `SELECT * FROM participants WHERE id in (SELECT participant_id FROM participant_crew WHERE crew_id IN (SELECT crew_id from crew_protest where protest_id = $1));`
	var participantDB []ParticipantDB
	err := w.db.Select(&participantDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	var participantModels []models.Participant
	for i := range participantDB {
		participant := copyParticipantResultToModel(&participantDB[i])
		participantModels = append(participantModels, *participant)
	}
	return participantModels, nil
}

func (w ParticipantRepository) GetAllParticipants() ([]models.Participant, error) {
	query := `SELECT * FROM participants;`
	var participantDB []ParticipantDB
	err := w.db.Select(&participantDB, query)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	var participantModels []models.Participant
	for i := range participantDB {
		participant := copyParticipantResultToModel(&participantDB[i])
		participantModels = append(participantModels, *participant)
	}
	return participantModels, nil
}
