package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"context"
	"github.com/google/uuid"
)

type IRatingRepository interface {
	Create(ctx context.Context, rating *models.Rating) (*models.Rating, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, rating *models.Rating) (*models.Rating, error)
	GetRatingDataByID(ctx context.Context, id uuid.UUID) (*models.Rating, error)
	AttachJudgeToRating(ctx context.Context, ratingID uuid.UUID, judgeID uuid.UUID) error
	DetachJudgeFromRating(ctx context.Context, ratingID uuid.UUID, judgeID uuid.UUID) error
	GetAllRatings(ctx context.Context) ([]models.Rating, error)
	GetRatingTable(ctx context.Context, id uuid.UUID) ([]models.RatingTableLine, error)
}
