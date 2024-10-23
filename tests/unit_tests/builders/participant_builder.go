package builders

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"time"
)

// ParticipantBuilder реализует паттерн Data Builder для Participant
type ParticipantBuilder struct {
	participant models.Participant
}

// NewParticipantBuilder создает новый экземпляр ParticipantBuilder с настройками по умолчанию
func NewParticipantBuilder() *ParticipantBuilder {
	return &ParticipantBuilder{
		participant: models.Participant{
			ID:       uuid.New(),      // Генерация нового UUID по умолчанию
			FIO:      "Default FIO",   // ФИО по умолчанию
			Category: 1,               // Категория по умолчанию
			Gender:   0,               // Пол по умолчанию (например, 0 — мужской, 1 — женский)
			Birthday: time.Now(),      // Дата рождения по умолчанию
			Coach:    "Default Coach", // Тренер по умолчанию
		},
	}
}

// WithID устанавливает ID участника
func (b *ParticipantBuilder) WithID(id uuid.UUID) *ParticipantBuilder {
	b.participant.ID = id
	return b
}

// WithFIO устанавливает ФИО участника
func (b *ParticipantBuilder) WithFIO(fio string) *ParticipantBuilder {
	b.participant.FIO = fio
	return b
}

// WithCategory устанавливает категорию участника
func (b *ParticipantBuilder) WithCategory(category int) *ParticipantBuilder {
	b.participant.Category = category
	return b
}

// WithGender устанавливает пол участника
func (b *ParticipantBuilder) WithGender(gender int) *ParticipantBuilder {
	b.participant.Gender = gender
	return b
}

// WithBirthday устанавливает дату рождения участника
func (b *ParticipantBuilder) WithBirthday(birthday time.Time) *ParticipantBuilder {
	b.participant.Birthday = birthday
	return b
}

// WithCoach устанавливает тренера участника
func (b *ParticipantBuilder) WithCoach(coach string) *ParticipantBuilder {
	b.participant.Coach = coach
	return b
}

// Build возвращает готовый объект Participant
func (b *ParticipantBuilder) Build() *models.Participant {
	return &b.participant
}

// ParticipantMother реализует паттерн Object Mother для Participant
var ParticipantMother = struct {
	Default           func() *models.Participant
	WithID            func(id uuid.UUID) *models.Participant
	WithFIO           func(fio string) *models.Participant
	WithCategory      func(category int) *models.Participant
	WithGender        func(gender int) *models.Participant
	WithBirthday      func(birthday time.Time) *models.Participant
	WithCoach         func(coach string) *models.Participant
	CustomParticipant func(id uuid.UUID, fio string, category, gender int, birthday time.Time, coach string) *models.Participant
}{
	Default: func() *models.Participant {
		return &models.Participant{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Category: 1,
			Gender:   0,
			Birthday: time.Now(),
			Coach:    "Default Coach",
		}
	},
	WithID: func(id uuid.UUID) *models.Participant {
		return &models.Participant{
			ID:       id,
			FIO:      "ParticipantWithSpecificID",
			Category: 1,
			Gender:   0,
			Birthday: time.Now(),
			Coach:    "Default Coach",
		}
	},
	WithFIO: func(fio string) *models.Participant {
		return &models.Participant{
			ID:       uuid.New(),
			FIO:      fio,
			Category: 1,
			Gender:   0,
			Birthday: time.Now(),
			Coach:    "Default Coach",
		}
	},
	WithCategory: func(category int) *models.Participant {
		return &models.Participant{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Category: category,
			Gender:   0,
			Birthday: time.Now(),
			Coach:    "Default Coach",
		}
	},
	WithGender: func(gender int) *models.Participant {
		return &models.Participant{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Category: 1,
			Gender:   gender,
			Birthday: time.Now(),
			Coach:    "Default Coach",
		}
	},
	WithBirthday: func(birthday time.Time) *models.Participant {
		return &models.Participant{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Category: 1,
			Gender:   0,
			Birthday: birthday,
			Coach:    "Default Coach",
		}
	},
	WithCoach: func(coach string) *models.Participant {
		return &models.Participant{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Category: 1,
			Gender:   0,
			Birthday: time.Now(),
			Coach:    coach,
		}
	},
	CustomParticipant: func(id uuid.UUID, fio string, category, gender int, birthday time.Time, coach string) *models.Participant {
		return &models.Participant{
			ID:       id,
			FIO:      fio,
			Category: category,
			Gender:   gender,
			Birthday: birthday,
			Coach:    coach,
		}
	},
}
