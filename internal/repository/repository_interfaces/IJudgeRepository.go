package repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type IJudgeRepository interface {
	CreateProfile(judge *models.Judge) (*models.Judge, error)
	DeleteProfile(id uuid.UUID) error
	UpdateProfile(judge *models.Judge) (*models.Judge, error)
	GetJudgeDataByID(id uuid.UUID) (*models.Judge, error)
	GetJudgeDataByProtestID(protestID uuid.UUID) (*models.Judge, error)
	GetJudgesDataByRatingID(ratingID uuid.UUID) ([]models.Judge, error)
	GetJudgeDataByLogin(login string) (*models.Judge, error)
}
