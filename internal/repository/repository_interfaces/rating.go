package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type IRatingRepository interface {
	Create(rating *models.Rating) (*models.Rating, error)
	Delete(id uuid.UUID) error
	Update(rating *models.Rating) (*models.Rating, error)
	GetRatingDataByID(id uuid.UUID) (*models.Rating, error)
	AttachJudgeToRating(ratingID uuid.UUID, judgeID uuid.UUID) error
	DetachJudgeFromRating(ratingID uuid.UUID, judgeID uuid.UUID) error
}
