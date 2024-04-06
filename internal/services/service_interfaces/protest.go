package service_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"time"
)

type IProtestService interface {
	AddNewProtest(protestID uuid.UUID, raceID uuid.UUID, ratingID uuid.UUID, ruleNum int, reviewDate time.Time, comment string, protesteeSailNum int, protestorSailNum int, witnessesSailNum []int) (*models.Protest, error)
	DeleteProtestByID(id uuid.UUID) error
	UpdateProtestByID(protestID uuid.UUID, ruleNum int, reviewDate time.Time, status int, comment string) (*models.Protest, error)
	GetProtestDataByID(id uuid.UUID) (*models.Protest, error)
	GetProtestsDataByRaceID(raceID uuid.UUID) ([]models.Protest, error)
	CompleteReview(protestID uuid.UUID, protesteePoints int, comment string)
	AttachJudgeToProtest(judgeID uuid.UUID, protestID uuid.UUID) error
	DetachJudgeFromProtest(judgeID uuid.UUID, protestID uuid.UUID) error
}
