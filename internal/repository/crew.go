package repository

//
//import (
//	"PPO_BMSTU/internal/repository/repository_errors"
//	"PPO_BMSTU/internal/repository/repository_interfaces"
//	"database/sql"
//	"errors"
//	"github.com/google/uuid"
//	"github.com/jmoiron/sqlx"
//	"PPO_BMSTU/internal/models"
//)
//
//type CrewDB struct {
//	ID          uuid.UUID `db:"id"`
//	Class        string    `db:"class"`
//	SailNum     int    `db:"sailNum"`
//	RatingID     uuid.UUID    `db:"ratingID"`
//}
//
//type CrewRepository struct {
//	db *sqlx.DB
//}
//
//func NewCrewRepository(db *sqlx.DB) repository_interfaces.ICrewRepository {
//	return &CrewRepository{db: db}
//}
//
//func copyCrewResultToModel(crewDB *CrewDB) *models.Crew {
//	return &models.Crew{
//		ID:          crewDB.ID,
//		RatingID: crewDB.RatingID,
//		SailNum:  crewDB.SailNum,
//		Class:    crewDB.Class,
//	}
//}
//
//Create(race *models.Crew) (*models.Crew, error)
//Delete(id uuid.UUID) error
//Update(crew *models.Crew) error
//GetCrewDataByID(id uuid.UUID) (*models.Crew, error)
//GetCrewsDataByRatingID(ratingID uuid.UUID) ([]models.Crew, error)
//GetCrewsDataByProtestID(protestID uuid.UUID) ([]models.Crew, error)
//GetCrewDataBySailNumAndRatingID(sailNum int, ratingID uuid.UUID) (*models.Crew, error)
//AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error
//DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error
//ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error
//
//func (w CrewRepository) Create(crew *models.Crew) (*models.Crew, error) {
//	query := `INSERT INTO crews(name, surname, address, phone_number, email, role, password) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
//
//	var crewID uuid.UUID
//	err := w.db.QueryRow(query, crew.Name, crew.Surname, crew.Address, crew.PhoneNumber, crew.Email, crew.Role, crew.Password).Scan(&crewID)
//
//	if err != nil {
//		return nil, repository_errors.InsertError
//	}
//
//	return &models.Crew{
//		ID:          crewID,
//		Name:        crew.Name,
//		Surname:     crew.Surname,
//		Address:     crew.Address,
//		PhoneNumber: crew.PhoneNumber,
//		Email:       crew.Email,
//		Role:        crew.Role,
//		Password:    crew.Password,
//	}, nil
//}
//
//func (w CrewRepository) Update(crew *models.Crew) (*models.Crew, error) {
//	query := `UPDATE crews SET name = $1, surname = $2, address = $3, phone_number = $4, email = $5, role = $6, password = $7 WHERE crews.id = $7 RETURNING id, name, surname, address, phone_number, email, role, password;`
//
//	var updatedCrew models.Crew
//	err := w.db.QueryRow(query, crew.Name, crew.Surname, crew.Address, crew.PhoneNumber, crew.Email, crew.Role, crew.Password, crew.ID).Scan(&updatedCrew.ID, &updatedCrew.Name, &updatedCrew.Surname, &updatedCrew.Address, &updatedCrew.PhoneNumber, &updatedCrew.Email, &updatedCrew.Role, &updatedCrew.Password)
//	if err != nil {
//		return nil, repository_errors.UpdateError
//	}
//	return &updatedCrew, nil
//}
//
//func (w CrewRepository) Delete(id uuid.UUID) error {
//	query := `DELETE FROM crews WHERE id = $1;`
//	_, err := w.db.Exec(query, id)
//
//	if err != nil {
//		return repository_errors.DeleteError
//	}
//
//	return nil
//}
//
//func (w CrewRepository) GetCrewByID(id uuid.UUID) (*models.Crew, error) {
//	query := `SELECT * FROM crews WHERE id = $1;`
//	crewDB := &CrewDB{}
//	err := w.db.Get(crewDB, query, id)
//
//	if errors.Is(err, sql.ErrNoRows) {
//		return nil, repository_errors.DoesNotExist
//	} else if err != nil {
//		return nil, repository_errors.SelectError
//	}
//
//	crewModels := copyCrewResultToModel(crewDB)
//
//	return crewModels, nil
//}
//
//func (w CrewRepository) GetAllCrews() ([]models.Crew, error) {
//	query := `SELECT name, surname, address, phone_number, email, role FROM crews;`
//	var crewDB []CrewDB
//
//	err := w.db.Select(&crewDB, query)
//
//	if err != nil {
//		return nil, repository_errors.SelectError
//	}
//
//	var crewModels []models.Crew
//	for i := range crewDB {
//		crew := copyCrewResultToModel(&crewDB[i])
//		crewModels = append(crewModels, *crew)
//	}
//
//	return crewModels, nil
//}
//
//func (w CrewRepository) GetCrewByEmail(email string) (*models.Crew, error) {
//	query := `SELECT * FROM crews WHERE email = $1;`
//	crewDB := &CrewDB{}
//	err := w.db.Get(crewDB, query, email)
//
//	if errors.Is(err, sql.ErrNoRows) {
//		return nil, repository_errors.DoesNotExist
//	} else if err != nil {
//		return nil, repository_errors.SelectError
//	}
//
//	crewModels := copyCrewResultToModel(crewDB)
//
//	return crewModels, nil
//}
