package builders

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// JudgeBuilder реализует паттерн Data Builder для Judge
type JudgeBuilder struct {
	judge models.Judge
}

// NewJudgeBuilder создает новый экземпляр JudgeBuilder с настройками по умолчанию
func NewJudgeBuilder() *JudgeBuilder {
	return &JudgeBuilder{
		judge: models.Judge{
			ID:       uuid.New(),         // Генерация нового UUID по умолчанию
			FIO:      "Default FIO",      // ФИО по умолчанию
			Login:    "default_login",    // Логин по умолчанию
			Password: "default_password", // Пароль по умолчанию
			Role:     1,                  // Роль по умолчанию
			Post:     "Default Post",     // Должность по умолчанию
		},
	}
}

// WithID устанавливает ID судьи
func (b *JudgeBuilder) WithID(id uuid.UUID) *JudgeBuilder {
	b.judge.ID = id
	return b
}

// WithFIO устанавливает ФИО судьи
func (b *JudgeBuilder) WithFIO(fio string) *JudgeBuilder {
	b.judge.FIO = fio
	return b
}

// WithLogin устанавливает логин судьи
func (b *JudgeBuilder) WithLogin(login string) *JudgeBuilder {
	b.judge.Login = login
	return b
}

// WithPassword устанавливает пароль судьи
func (b *JudgeBuilder) WithPassword(password string) *JudgeBuilder {
	b.judge.Password = password
	return b
}

// WithRole устанавливает роль судьи
func (b *JudgeBuilder) WithRole(role int) *JudgeBuilder {
	b.judge.Role = role
	return b
}

// WithPost устанавливает должность судьи
func (b *JudgeBuilder) WithPost(post string) *JudgeBuilder {
	b.judge.Post = post
	return b
}

// Build возвращает готовый объект Judge
func (b *JudgeBuilder) Build() *models.Judge {
	return &b.judge
}

// JudgeMother реализует паттерн Object Mother для Judge
var JudgeMother = struct {
	Default      func() *models.Judge
	WithID       func(id uuid.UUID) *models.Judge
	WithFIO      func(fio string) *models.Judge
	WithLogin    func(login string) *models.Judge
	WithPassword func(password string) *models.Judge
	WithRole     func(role int) *models.Judge
	WithPost     func(post string) *models.Judge
	CustomJudge  func(id uuid.UUID, fio, login, password string, role int, post string) *models.Judge
}{
	Default: func() *models.Judge {
		return &models.Judge{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Login:    "default_login",
			Password: "default_password",
			Role:     1,
			Post:     "Default Post",
		}
	},
	WithID: func(id uuid.UUID) *models.Judge {
		return &models.Judge{
			ID:       id,
			FIO:      "JudgeWithSpecificID",
			Login:    "default_login",
			Password: "default_password",
			Role:     1,
			Post:     "Default Post",
		}
	},
	WithFIO: func(fio string) *models.Judge {
		return &models.Judge{
			ID:       uuid.New(),
			FIO:      fio,
			Login:    "default_login",
			Password: "default_password",
			Role:     1,
			Post:     "Default Post",
		}
	},
	WithLogin: func(login string) *models.Judge {
		return &models.Judge{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Login:    login,
			Password: "default_password",
			Role:     1,
			Post:     "Default Post",
		}
	},
	WithPassword: func(password string) *models.Judge {
		return &models.Judge{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Login:    "default_login",
			Password: password,
			Role:     1,
			Post:     "Default Post",
		}
	},
	WithRole: func(role int) *models.Judge {
		return &models.Judge{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Login:    "default_login",
			Password: "default_password",
			Role:     role,
			Post:     "Default Post",
		}
	},
	WithPost: func(post string) *models.Judge {
		return &models.Judge{
			ID:       uuid.New(),
			FIO:      "Default FIO",
			Login:    "default_login",
			Password: "default_password",
			Role:     1,
			Post:     post,
		}
	},
	CustomJudge: func(id uuid.UUID, fio, login, password string, role int, post string) *models.Judge {
		return &models.Judge{
			ID:       id,
			FIO:      fio,
			Login:    login,
			Password: password,
			Role:     role,
			Post:     post,
		}
	},
}
