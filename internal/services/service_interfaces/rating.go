package service_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type IRatingService interface {
	AddNewRating(ratingID uuid.UUID, class string, blowoutCnt int) (*models.Rating, error)
	DeleteRatingByID(id uuid.UUID) error
	UpdateRatingByID(ratingID uuid.UUID, class string, blowoutCnt int) (*models.Rating, error)
	GetRatingDataByID(id uuid.UUID) (*models.Rating, error)
	AttachJudgeToRating(ratingID uuid.UUID, judgeID uuid.UUID) error
	DetachJudgeFromRating(ratingID uuid.UUID, judgeID uuid.UUID) error
}
