package builders

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// CrewResInRaceBuilder реализует паттерн Data Builder для CrewResInRace
type CrewResInRaceBuilder struct {
	crew models.CrewResInRace
}

// NewCrewResInRaceBuilder создает новый экземпляр CrewResInRaceBuilder с настройками по умолчанию
func NewCrewResInRaceBuilder() *CrewResInRaceBuilder {
	return &CrewResInRaceBuilder{
		crew: models.CrewResInRace{
			CrewID:           uuid.New(), // Генерация нового UUID по умолчанию
			RaceID:           uuid.New(), // Генерация нового UUID по умолчанию
			Points:           0,          // Значение Points по умолчанию
			SpecCircumstance: 0,          // Значение SpecCircumstance по умолчанию
		},
	}
}

// WithCrewID устанавливает ID экипажа
func (b *CrewResInRaceBuilder) WithCrewID(crewID uuid.UUID) *CrewResInRaceBuilder {
	b.crew.CrewID = crewID
	return b
}

// WithRaceID устанавливает ID гонки
func (b *CrewResInRaceBuilder) WithRaceID(raceID uuid.UUID) *CrewResInRaceBuilder {
	b.crew.RaceID = raceID
	return b
}

// WithPoints устанавливает количество очков
func (b *CrewResInRaceBuilder) WithPoints(points int) *CrewResInRaceBuilder {
	b.crew.Points = points
	return b
}

// WithSpecCircumstance устанавливает специальные обстоятельства
func (b *CrewResInRaceBuilder) WithSpecCircumstance(specCircumstance int) *CrewResInRaceBuilder {
	b.crew.SpecCircumstance = specCircumstance
	return b
}

// Build возвращает готовый объект CrewResInRace
func (b *CrewResInRaceBuilder) Build() *models.CrewResInRace {
	return &b.crew
}

// CrewMother реализует паттерн Object Mother для CrewResInRace
var CrewResInRaceMother = struct {
	Default              func() *models.CrewResInRace
	WithCrewID           func(crewID uuid.UUID) *models.CrewResInRace
	WithRaceID           func(raceID uuid.UUID) *models.CrewResInRace
	WithPoints           func(points int) *models.CrewResInRace
	WithSpecCircumstance func(specCircumstance int) *models.CrewResInRace
	CustomCrew           func(crewID, raceID uuid.UUID, points, specCircumstance int) *models.CrewResInRace
}{
	Default: func() *models.CrewResInRace {
		return &models.CrewResInRace{
			CrewID:           uuid.New(),
			RaceID:           uuid.New(),
			Points:           0,
			SpecCircumstance: 0,
		}
	},
	WithCrewID: func(crewID uuid.UUID) *models.CrewResInRace {
		return &models.CrewResInRace{
			CrewID:           crewID,
			RaceID:           uuid.New(),
			Points:           0,
			SpecCircumstance: 0,
		}
	},
	WithRaceID: func(raceID uuid.UUID) *models.CrewResInRace {
		return &models.CrewResInRace{
			CrewID:           uuid.New(),
			RaceID:           raceID,
			Points:           0,
			SpecCircumstance: 0,
		}
	},
	WithPoints: func(points int) *models.CrewResInRace {
		return &models.CrewResInRace{
			CrewID:           uuid.New(),
			RaceID:           uuid.New(),
			Points:           points,
			SpecCircumstance: 0,
		}
	},
	WithSpecCircumstance: func(specCircumstance int) *models.CrewResInRace {
		return &models.CrewResInRace{
			CrewID:           uuid.New(),
			RaceID:           uuid.New(),
			Points:           0,
			SpecCircumstance: specCircumstance,
		}
	},
	CustomCrew: func(crewID, raceID uuid.UUID, points, specCircumstance int) *models.CrewResInRace {
		return &models.CrewResInRace{
			CrewID:           crewID,
			RaceID:           raceID,
			Points:           points,
			SpecCircumstance: specCircumstance,
		}
	},
}
