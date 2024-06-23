package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_interfaces"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"time"
)

type ParticipantService struct {
	ParticipantRepository repository_interfaces.IParticipantRepository
	logger                *log.Logger
}

func NewParticipantService(ParticipantRepository repository_interfaces.IParticipantRepository, logger *log.Logger) service_interfaces.IParticipantService {
	return &ParticipantService{
		ParticipantRepository: ParticipantRepository,
		logger:                logger,
	}
}

func (p ParticipantService) AddNewParticipant(participantID uuid.UUID, fio string, category int, gender int, birthDay time.Time, coach string) (*models.Participant, error) {
	if !validFIO(fio) || !validCategory(category) || !validGender(gender) || !validBirthDay(birthDay) || !validFIO(coach) {
		p.logger.Error("SERVICE: Invalid input data", "fio", fio, "category", category, "gender", gender, "birthDate", birthDay, "coach", coach)
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}

	participant := &models.Participant{
		ID:       participantID,
		FIO:      fio,
		Category: category,
		Gender:   gender,
		Birthday: birthDay,
		Coach:    coach,
	}

	participant, err := p.ParticipantRepository.Create(participant)
	if err != nil {
		p.logger.Error("SERVICE: CreateNewParticipant method failed", "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully created new participant", "participant", participant)
	return participant, nil
}

func (p ParticipantService) DeleteParticipantByID(id uuid.UUID) error {
	_, err := p.ParticipantRepository.GetParticipantDataByID(id)
	if err != nil {
		p.logger.Error("SERVICE: GetParticipantDataByID method failed", "id", id, "error", err)
		return err
	}

	err = p.ParticipantRepository.Delete(id)
	if err != nil {
		p.logger.Error("SERVICE: DeleteParticipantByID method failed", "error", err)
		return err
	}

	p.logger.Info("SERVICE: Successfully deleted participant", "participant", id)
	return nil
}

func (p ParticipantService) UpdateParticipantByID(participantID uuid.UUID, fio string, category int, birthDay time.Time, coach string) (*models.Participant, error) {
	participant, err := p.ParticipantRepository.GetParticipantDataByID(participantID)
	participantCopy := participant

	if err != nil {
		p.logger.Error("SERVICE: GetParticipantByID method failed", "id", participantID, "error", err)
		return participant, err
	}

	if !validFIO(fio) || !validCategory(category) || !validBirthDay(birthDay) || !validFIO(coach) {
		p.logger.Error("SERVICE: Invalid input data", "fio", fio, "category", category, "birthDate", birthDay, "coach", coach)
		return participant, fmt.Errorf("SERVICE: Invalid input data")
	}

	participant.FIO = fio
	participant.Category = category
	participant.Birthday = birthDay
	participant.Coach = coach

	participant, err = p.ParticipantRepository.Update(participant)
	if err != nil {
		participant = participantCopy
		p.logger.Error("SERVICE: UpdateParticipant method failed", "error", err)
		return participant, err
	}

	p.logger.Info("SERVICE: Successfully created new participant", "participant", participant)
	return participant, nil
}

func (p ParticipantService) GetParticipantDataByID(id uuid.UUID) (*models.Participant, error) {
	participant, err := p.ParticipantRepository.GetParticipantDataByID(id)

	if err != nil {
		p.logger.Error("SERVICE: GetParticipantByID method failed", "id", id, "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got participant with GetParticipantByID", "id", id)
	return participant, nil
}

func (p ParticipantService) GetParticipantsDataByCrewID(crewID uuid.UUID) ([]models.Participant, error) {
	participants, err := p.ParticipantRepository.GetParticipantsDataByCrewID(crewID)
	if err != nil {
		p.logger.Error("SERVICE: GetParticipantsDataByCrewID method failed", "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got participants by crew id", "participants", participants)
	return participants, nil
}

func (p ParticipantService) GetParticipantsDataByProtestID(protestID uuid.UUID) ([]models.Participant, error) {
	participants, err := p.ParticipantRepository.GetParticipantsDataByCrewID(protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetParticipantsDataByProtestID method failed", "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got participants by protest id", "participants", participants)
	return participants, nil
}

func (p ParticipantService) GetAllParticipants() ([]models.Participant, error) {
	participants, err := p.ParticipantRepository.GetAllParticipants()
	if err != nil {
		p.logger.Error("SERVICE: GetAllParticipants method failed", "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got all participants", "participants", participants)
	return participants, nil
}
