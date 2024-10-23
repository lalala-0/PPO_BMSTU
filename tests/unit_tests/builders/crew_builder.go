package builders

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// CrewBuilder реализует паттерн Data Builder для Crew
type CrewBuilder struct {
	crew models.Crew
}

// NewCrewBuilder создает новый экземпляр CrewBuilder с настройками по умолчанию
func NewCrewBuilder() *CrewBuilder {
	return &CrewBuilder{
		crew: models.Crew{
			ID:       uuid.New(), // Генерация нового UUID по умолчанию
			RatingID: uuid.New(), // Генерация нового RatingID по умолчанию
			SailNum:  1,          // Номер паруса по умолчанию
			Class:    1,          // Класс по умолчанию
		},
	}
}

// WithID устанавливает ID экипажа
func (b *CrewBuilder) WithID(id uuid.UUID) *CrewBuilder {
	b.crew.ID = id
	return b
}

// WithRatingID устанавливает RatingID экипажа
func (b *CrewBuilder) WithRatingID(ratingID uuid.UUID) *CrewBuilder {
	b.crew.RatingID = ratingID
	return b
}

// WithSailNum устанавливает номер паруса экипажа
func (b *CrewBuilder) WithSailNum(sailNum int) *CrewBuilder {
	b.crew.SailNum = sailNum
	return b
}

// WithClass устанавливает класс экипажа
func (b *CrewBuilder) WithClass(class int) *CrewBuilder {
	b.crew.Class = class
	return b
}

// Build возвращает готовый объект Crew
func (b *CrewBuilder) Build() *models.Crew {
	return &b.crew
}

// CrewMother реализует паттерн Object Mother для Crew
var CrewMother = struct {
	Default      func() *models.Crew
	WithID       func(id uuid.UUID) *models.Crew
	WithRatingID func(ratingID uuid.UUID) *models.Crew
	WithSailNum  func(sailNum int) *models.Crew
	WithClass    func(class int) *models.Crew
	CustomCrew   func(id, ratingID uuid.UUID, sailNum, class int) *models.Crew
}{
	Default: func() *models.Crew {
		return &models.Crew{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			SailNum:  1,
			Class:    1,
		}
	},
	WithID: func(id uuid.UUID) *models.Crew {
		return &models.Crew{
			ID:       id,
			RatingID: uuid.New(),
			SailNum:  1,
			Class:    1,
		}
	},
	WithRatingID: func(ratingID uuid.UUID) *models.Crew {
		return &models.Crew{
			ID:       uuid.New(),
			RatingID: ratingID,
			SailNum:  1,
			Class:    1,
		}
	},
	WithSailNum: func(sailNum int) *models.Crew {
		return &models.Crew{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			SailNum:  sailNum,
			Class:    1,
		}
	},
	WithClass: func(class int) *models.Crew {
		return &models.Crew{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			SailNum:  1,
			Class:    class,
		}
	},
	CustomCrew: func(id, ratingID uuid.UUID, sailNum, class int) *models.Crew {
		return &models.Crew{
			ID:       id,
			RatingID: ratingID,
			SailNum:  sailNum,
			Class:    class,
		}
	},
}
