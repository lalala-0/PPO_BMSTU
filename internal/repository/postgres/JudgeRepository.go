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

type JudgeDB struct {
	ID       uuid.UUID `db:"id"`
	FIO      string    `db:"name"`
	Login    string    `db:"login"`
	Password string    `db:"password"`
	Role     int       `db:"role"`
	Post     string    `db:"post"`
}

type JudgeRepository struct {
	db *TracedDB
}

func NewJudgeRepository(db *TracedDB) repository_interfaces.IJudgeRepository {
	return &JudgeRepository{db: db}
}

func copyJudgeResultToModel(judgeDB *JudgeDB) *models.Judge {
	return &models.Judge{
		ID:       judgeDB.ID,
		FIO:      judgeDB.FIO,
		Login:    judgeDB.Login,
		Password: judgeDB.Password,
		Role:     judgeDB.Role,
		Post:     judgeDB.Post,
	}
}

func (w JudgeRepository) CreateProfile(ctx context.Context, judge *models.Judge) (*models.Judge, error) {
	query := `INSERT INTO judges(name, login, password, role, post) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var judgeID uuid.UUID
	err := w.db.QueryRowContext(ctx, query, judge.FIO, judge.Login, judge.Password, judge.Role, judge.Post).Scan(&judgeID)

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Judge{
		ID:       judgeID,
		FIO:      judge.FIO,
		Login:    judge.Login,
		Password: judge.Password,
		Role:     judge.Role,
		Post:     judge.Post,
	}, nil
}

func (w JudgeRepository) UpdateProfile(ctx context.Context, judge *models.Judge) (*models.Judge, error) {
	query := `UPDATE judges SET name = $1, login = $2, password = $3, role = $4, post = $5 WHERE judges.id = $6 RETURNING id, name, login, password, role, post;`

	var updatedJudge models.Judge
	err := w.db.QueryRowContext(ctx, query, judge.FIO, judge.Login, judge.Password, judge.Role, judge.Post, judge.ID).
		Scan(&updatedJudge.ID, &updatedJudge.FIO, &updatedJudge.Login, &updatedJudge.Password, &updatedJudge.Role, &updatedJudge.Post)
	if err != nil {
		return nil, repository_errors.UpdateError
	}
	return &updatedJudge, nil
}

func (w JudgeRepository) DeleteProfile(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM judges WHERE id = $1;`
	res, err := w.db.ExecContext(ctx, query, id)

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

func (w JudgeRepository) GetJudgeDataByID(ctx context.Context, id uuid.UUID) (*models.Judge, error) {
	query := `SELECT * FROM judges WHERE id = $1;`
	judgeDB := &JudgeDB{}
	err := w.db.GetContext(ctx, judgeDB, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	judgeModels := copyJudgeResultToModel(judgeDB)

	return judgeModels, nil
}

func (w JudgeRepository) GetJudgeDataByLogin(ctx context.Context, login string) (*models.Judge, error) {
	query := `SELECT * FROM judges WHERE login like $1;`
	judgeDB := &JudgeDB{}
	err := w.db.GetContext(ctx, judgeDB, query, login)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	judgeModels := copyJudgeResultToModel(judgeDB)

	return judgeModels, nil
}

func (w JudgeRepository) GetJudgesDataByRatingID(ctx context.Context, ratingID uuid.UUID) ([]models.Judge, error) {
	query := `SELECT * FROM judges WHERE id IN (SELECT judge_id FROM judge_rating WHERE rating_id = $1);`
	var judgeDB []JudgeDB
	err := w.db.SelectContext(ctx, &judgeDB, query, ratingID)

	if err != nil {
		return nil, repository_errors.SelectError
	}

	var judgeModels []models.Judge
	for i := range judgeDB {
		judge := copyJudgeResultToModel(&judgeDB[i])
		judgeModels = append(judgeModels, *judge)
	}

	return judgeModels, nil
}

func (w JudgeRepository) GetAllJudges(ctx context.Context) ([]models.Judge, error) {
	query := `SELECT * FROM judges;`
	var judgeDB []JudgeDB
	err := w.db.SelectContext(ctx, &judgeDB, query)

	if err != nil {
		return nil, repository_errors.SelectError
	}

	var judgeModels []models.Judge
	for i := range judgeDB {
		judge := copyJudgeResultToModel(&judgeDB[i])
		judgeModels = append(judgeModels, *judge)
	}

	return judgeModels, nil
}

func (w JudgeRepository) GetJudgeDataByProtestID(ctx context.Context, protestID uuid.UUID) (*models.Judge, error) {
	query := `SELECT * FROM judges WHERE id IN (SELECT judge_id FROM protests WHERE id = $1);`
	judgeDB := &JudgeDB{}
	err := w.db.GetContext(ctx, judgeDB, query, protestID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	judgeModels := copyJudgeResultToModel(judgeDB)

	return judgeModels, nil
}
