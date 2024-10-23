package builders

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"time"
)

// ProtestBuilder реализует паттерн Data Builder для Protest
type ProtestBuilder struct {
	protest models.Protest
}

// NewProtestBuilder создает новый экземпляр ProtestBuilder с настройками по умолчанию
func NewProtestBuilder() *ProtestBuilder {
	return &ProtestBuilder{
		protest: models.Protest{
			ID:         uuid.New(),        // Генерация нового UUID по умолчанию
			RaceID:     uuid.New(),        // Генерация нового RaceID по умолчанию
			JudgeID:    uuid.New(),        // Генерация нового JudgeID по умолчанию
			RuleNum:    1,                 // Номер правила по умолчанию
			ReviewDate: time.Now(),        // Дата рассмотрения по умолчанию
			Status:     0,                 // Статус по умолчанию
			Comment:    "Default Comment", // Комментарий по умолчанию
			RatingID:   uuid.New(),        // Генерация нового RatingID по умолчанию
		},
	}
}

// WithID устанавливает ID протеста
func (b *ProtestBuilder) WithID(id uuid.UUID) *ProtestBuilder {
	b.protest.ID = id
	return b
}

// WithRaceID устанавливает RaceID протеста
func (b *ProtestBuilder) WithRaceID(raceID uuid.UUID) *ProtestBuilder {
	b.protest.RaceID = raceID
	return b
}

// WithJudgeID устанавливает JudgeID протеста
func (b *ProtestBuilder) WithJudgeID(judgeID uuid.UUID) *ProtestBuilder {
	b.protest.JudgeID = judgeID
	return b
}

// WithRuleNum устанавливает номер правила протеста
func (b *ProtestBuilder) WithRuleNum(ruleNum int) *ProtestBuilder {
	b.protest.RuleNum = ruleNum
	return b
}

// WithReviewDate устанавливает дату рассмотрения протеста
func (b *ProtestBuilder) WithReviewDate(reviewDate time.Time) *ProtestBuilder {
	b.protest.ReviewDate = reviewDate
	return b
}

// WithStatus устанавливает статус протеста
func (b *ProtestBuilder) WithStatus(status int) *ProtestBuilder {
	b.protest.Status = status
	return b
}

// WithComment устанавливает комментарий протеста
func (b *ProtestBuilder) WithComment(comment string) *ProtestBuilder {
	b.protest.Comment = comment
	return b
}

// WithRatingID устанавливает RatingID протеста
func (b *ProtestBuilder) WithRatingID(ratingID uuid.UUID) *ProtestBuilder {
	b.protest.RatingID = ratingID
	return b
}

// Build возвращает готовый объект Protest
func (b *ProtestBuilder) Build() *models.Protest {
	return &b.protest
}

// ProtestMother реализует паттерн Object Mother для Protest
var ProtestMother = struct {
	Default        func() *models.Protest
	WithID         func(id uuid.UUID) *models.Protest
	WithRaceID     func(raceID uuid.UUID) *models.Protest
	WithJudgeID    func(judgeID uuid.UUID) *models.Protest
	WithRuleNum    func(ruleNum int) *models.Protest
	WithReviewDate func(reviewDate time.Time) *models.Protest
	WithStatus     func(status int) *models.Protest
	WithComment    func(comment string) *models.Protest
	WithRatingID   func(ratingID uuid.UUID) *models.Protest
	CustomProtest  func(id, raceID, judgeID, ratingID uuid.UUID, ruleNum, status int, reviewDate time.Time, comment string) *models.Protest
}{
	Default: func() *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     uuid.New(),
			JudgeID:    uuid.New(),
			RuleNum:    1,
			ReviewDate: time.Now(),
			Status:     0,
			Comment:    "Default Comment",
			RatingID:   uuid.New(),
		}
	},
	WithID: func(id uuid.UUID) *models.Protest {
		return &models.Protest{
			ID:         id,
			RaceID:     uuid.New(),
			JudgeID:    uuid.New(),
			RuleNum:    1,
			ReviewDate: time.Now(),
			Status:     0,
			Comment:    "ProtestWithSpecificID",
			RatingID:   uuid.New(),
		}
	},
	WithRaceID: func(raceID uuid.UUID) *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     raceID,
			JudgeID:    uuid.New(),
			RuleNum:    1,
			ReviewDate: time.Now(),
			Status:     0,
			Comment:    "Default Comment",
			RatingID:   uuid.New(),
		}
	},
	WithJudgeID: func(judgeID uuid.UUID) *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     uuid.New(),
			JudgeID:    judgeID,
			RuleNum:    1,
			ReviewDate: time.Now(),
			Status:     0,
			Comment:    "Default Comment",
			RatingID:   uuid.New(),
		}
	},
	WithRuleNum: func(ruleNum int) *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     uuid.New(),
			JudgeID:    uuid.New(),
			RuleNum:    ruleNum,
			ReviewDate: time.Now(),
			Status:     0,
			Comment:    "Default Comment",
			RatingID:   uuid.New(),
		}
	},
	WithReviewDate: func(reviewDate time.Time) *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     uuid.New(),
			JudgeID:    uuid.New(),
			RuleNum:    1,
			ReviewDate: reviewDate,
			Status:     0,
			Comment:    "Default Comment",
			RatingID:   uuid.New(),
		}
	},
	WithStatus: func(status int) *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     uuid.New(),
			JudgeID:    uuid.New(),
			RuleNum:    1,
			ReviewDate: time.Now(),
			Status:     status,
			Comment:    "Default Comment",
			RatingID:   uuid.New(),
		}
	},
	WithComment: func(comment string) *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     uuid.New(),
			JudgeID:    uuid.New(),
			RuleNum:    1,
			ReviewDate: time.Now(),
			Status:     0,
			Comment:    comment,
			RatingID:   uuid.New(),
		}
	},
	WithRatingID: func(ratingID uuid.UUID) *models.Protest {
		return &models.Protest{
			ID:         uuid.New(),
			RaceID:     uuid.New(),
			JudgeID:    uuid.New(),
			RuleNum:    1,
			ReviewDate: time.Now(),
			Status:     0,
			Comment:    "Default Comment",
			RatingID:   ratingID,
		}
	},
	CustomProtest: func(id, raceID, judgeID, ratingID uuid.UUID, ruleNum, status int, reviewDate time.Time, comment string) *models.Protest {
		return &models.Protest{
			ID:         id,
			RaceID:     raceID,
			JudgeID:    judgeID,
			RuleNum:    ruleNum,
			ReviewDate: reviewDate,
			Status:     status,
			Comment:    comment,
			RatingID:   ratingID,
		}
	},
}
