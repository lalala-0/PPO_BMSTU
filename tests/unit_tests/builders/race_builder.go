package builders

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"time"
)

// RaceBuilder реализует паттерн Data Builder для Race
type RaceBuilder struct {
	race models.Race
}

// NewRaceBuilder создает новый экземпляр RaceBuilder с настройками по умолчанию
func NewRaceBuilder() *RaceBuilder {
	return &RaceBuilder{
		race: models.Race{
			ID:       uuid.New(), // Генерация нового UUID по умолчанию
			RatingID: uuid.New(), // Генерация нового RatingID по умолчанию
			Date:     time.Now(), // Установка текущей даты по умолчанию
			Number:   1,          // Номер по умолчанию
			Class:    1,          // Класс по умолчанию
		},
	}
}

// WithID устанавливает ID гонки
func (b *RaceBuilder) WithID(id uuid.UUID) *RaceBuilder {
	b.race.ID = id
	return b
}

// WithRatingID устанавливает RatingID гонки
func (b *RaceBuilder) WithRatingID(ratingID uuid.UUID) *RaceBuilder {
	b.race.RatingID = ratingID
	return b
}

// WithDate устанавливает дату гонки
func (b *RaceBuilder) WithDate(date time.Time) *RaceBuilder {
	b.race.Date = date
	return b
}

// WithNumber устанавливает номер гонки
func (b *RaceBuilder) WithNumber(number int) *RaceBuilder {
	b.race.Number = number
	return b
}

// WithClass устанавливает класс гонки
func (b *RaceBuilder) WithClass(class int) *RaceBuilder {
	b.race.Class = class
	return b
}

// Build возвращает готовый объект Race
func (b *RaceBuilder) Build() *models.Race {
	return &b.race
}

// RaceMother реализует паттерн Object Mother для Race
var RaceMother = struct {
	Default      func() *models.Race
	WithID       func(id uuid.UUID) *models.Race
	WithRatingID func(ratingID uuid.UUID) *models.Race
	WithDate     func(date time.Time) *models.Race
	WithNumber   func(number int) *models.Race
	WithClass    func(class int) *models.Race
	CustomRace   func(id uuid.UUID, ratingID uuid.UUID, date time.Time, number int, class int) *models.Race
}{
	Default: func() *models.Race {
		return &models.Race{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			Date:     time.Now(),
			Number:   1,
			Class:    1,
		}
	},
	WithID: func(id uuid.UUID) *models.Race {
		return &models.Race{
			ID:       id,
			RatingID: uuid.New(),
			Date:     time.Now(),
			Number:   1,
			Class:    1,
		}
	},
	WithRatingID: func(ratingID uuid.UUID) *models.Race {
		return &models.Race{
			ID:       uuid.New(),
			RatingID: ratingID,
			Date:     time.Now(),
			Number:   1,
			Class:    1,
		}
	},
	WithDate: func(date time.Time) *models.Race {
		return &models.Race{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			Date:     date,
			Number:   1,
			Class:    1,
		}
	},
	WithNumber: func(number int) *models.Race {
		return &models.Race{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			Date:     time.Now(),
			Number:   number,
			Class:    1,
		}
	},
	WithClass: func(class int) *models.Race {
		return &models.Race{
			ID:       uuid.New(),
			RatingID: uuid.New(),
			Date:     time.Now(),
			Number:   1,
			Class:    class,
		}
	},
	CustomRace: func(id uuid.UUID, ratingID uuid.UUID, date time.Time, number int, class int) *models.Race {
		return &models.Race{
			ID:       id,
			RatingID: ratingID,
			Date:     date,
			Number:   number,
			Class:    class,
		}
	},
}
