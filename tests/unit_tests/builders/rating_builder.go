package builders

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// RatingBuilder реализует паттерн Data Builder для Rating
type RatingBuilder struct {
	rating models.Rating
}

// NewRatingBuilder создает новый экземпляр RatingBuilder с настройками по умолчанию
func NewRatingBuilder() *RatingBuilder {
	return &RatingBuilder{
		rating: models.Rating{
			ID:         uuid.New(), // Генерация нового UUID по умолчанию
			Name:       "DefaultRating",
			Class:      1, // Класс по умолчанию
			BlowoutCnt: 0, // BlowoutCnt по умолчанию
		},
	}
}

// WithID устанавливает ID рейтинга
func (b *RatingBuilder) WithID(id uuid.UUID) *RatingBuilder {
	b.rating.ID = id
	return b
}

// WithName устанавливает имя рейтинга
func (b *RatingBuilder) WithName(name string) *RatingBuilder {
	b.rating.Name = name
	return b
}

// WithClass устанавливает класс рейтинга
func (b *RatingBuilder) WithClass(class int) *RatingBuilder {
	b.rating.Class = class
	return b
}

// WithBlowoutCnt устанавливает количество "blowouts"
func (b *RatingBuilder) WithBlowoutCnt(blowoutCnt int) *RatingBuilder {
	b.rating.BlowoutCnt = blowoutCnt
	return b
}

// Build возвращает готовый объект Rating
func (b *RatingBuilder) Build() *models.Rating {
	return &b.rating
}

// RatingMother реализует паттерн Object Mother для Rating
var RatingMother = struct {
	Default        func() *models.Rating
	WithID         func(id uuid.UUID) *models.Rating
	WithName       func(name string) *models.Rating
	WithClass      func(class int) *models.Rating
	WithBlowoutCnt func(blowoutCnt int) *models.Rating
	CustomRating   func(id uuid.UUID, name string, class int, blowoutCnt int) *models.Rating
}{
	Default: func() *models.Rating {
		return &models.Rating{
			ID:         uuid.New(),
			Name:       "DefaultRating",
			Class:      1,
			BlowoutCnt: 0,
		}
	},
	WithID: func(id uuid.UUID) *models.Rating {
		return &models.Rating{
			ID:         id,
			Name:       "RatingWithSpecificID",
			Class:      1,
			BlowoutCnt: 0,
		}
	},
	WithName: func(name string) *models.Rating {
		return &models.Rating{
			ID:         uuid.New(),
			Name:       name,
			Class:      1,
			BlowoutCnt: 0,
		}
	},
	WithClass: func(class int) *models.Rating {
		return &models.Rating{
			ID:         uuid.New(),
			Name:       "RatingWithCustomClass",
			Class:      class,
			BlowoutCnt: 0,
		}
	},
	WithBlowoutCnt: func(blowoutCnt int) *models.Rating {
		return &models.Rating{
			ID:         uuid.New(),
			Name:       "RatingWithBlowoutCnt",
			Class:      1,
			BlowoutCnt: blowoutCnt,
		}
	},
	CustomRating: func(id uuid.UUID, name string, class int, blowoutCnt int) *models.Rating {
		return &models.Rating{
			ID:         id,
			Name:       name,
			Class:      class,
			BlowoutCnt: blowoutCnt,
		}
	},
}
