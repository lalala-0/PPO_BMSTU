package service_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type IJudgeService interface {
	Login(login, password string) (*models.Judge, error)
	CreateProfile(judgeID uuid.UUID, fio string, login string, password string, role int) (*models.Judge, error)
	DeleteProfile(id uuid.UUID) error
	UpdateProfile(judgeID uuid.UUID, fio string, login string, password string, role int) (*models.Judge, error)
	GetJudgeDataByID(id uuid.UUID) (*models.Judge, error)
	GetJudgeDataByProtestID(protestID uuid.UUID) (*models.Judge, error)
	GetJudgesDataByRatingID(ratingID uuid.UUID) ([]models.Judge, error)
}
