package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"context"
	"github.com/google/uuid"
)

type IJudgeRepository interface {
	CreateProfile(ctx context.Context, judge *models.Judge) (*models.Judge, error)
	DeleteProfile(ctx context.Context, id uuid.UUID) error
	UpdateProfile(ctx context.Context, judge *models.Judge) (*models.Judge, error)
	GetJudgeDataByID(ctx context.Context, id uuid.UUID) (*models.Judge, error)
	GetJudgeDataByProtestID(ctx context.Context, protestID uuid.UUID) (*models.Judge, error)
	GetJudgesDataByRatingID(ctx context.Context, ratingID uuid.UUID) ([]models.Judge, error)
	GetAllJudges(ctx context.Context) ([]models.Judge, error)
	GetJudgeDataByLogin(ctx context.Context, login string) (*models.Judge, error)
}
