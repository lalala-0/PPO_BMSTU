package test_repositories

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type DBInterface interface {
	CreateParticipant() *models.Participant
	CreateRating() *models.Rating
	CreateCrew(ratingID uuid.UUID) *models.Crew
	CreateCrewResInRace(crewID uuid.UUID, raceID uuid.UUID) *models.CrewResInRace
	CreateRace(ratingID uuid.UUID) *models.Race
	CreateJudge() *models.Judge
	CreateProtest(raceID uuid.UUID, judgeID uuid.UUID, ratingID uuid.UUID) *models.Protest
	AttachCrewToProtest(crewID uuid.UUID, protestID uuid.UUID)
	AttachCrewToProtestStatus(crewID uuid.UUID, protestID uuid.UUID, status int)
	AttachJudgeToRating(judgeID uuid.UUID, ratingID uuid.UUID)
	AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID)
}
