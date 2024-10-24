package db_init

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

// TestsRepository определяет методы для работы с MongoDB.
type TestRepositoryInitializer interface {
	CreateParticipant(participant *models.Participant) (*models.Participant, error)
	CreateRating(rating *models.Rating) (*models.Rating, error)
	CreateCrew(crew *models.Crew) (*models.Crew, error)
	CreateCrewResInRace(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error)
	CreateRace(race *models.Race) (*models.Race, error)
	CreateJudge(judge *models.Judge) (*models.Judge, error)
	CreateProtest(protest *models.Protest) (*models.Protest, error)
	AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID) error
	AttachCrewToProtestStatus(crewID uuid.UUID, protestID uuid.UUID, status int) error
	AttachJudgeToRating(judgeID uuid.UUID, ratingID uuid.UUID) error
	AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID) error
	ClearAll() error
}
