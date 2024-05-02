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

type RatingDB struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Class      int       `db:"class"`
	BlowoutCnt int       `db:"blowout_cnt"`
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
	query := `SELECT * FROM ratings;`
	ratingDB := []RatingDB{}
	err := w.db.Get(ratingDB, query)

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
