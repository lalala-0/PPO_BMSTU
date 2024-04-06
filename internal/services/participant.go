package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
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

func (p ParticipantService) AddNewParticipant(participantID uuid.UUID, fio string, category string, gender string, birthDay time.Time, coach string) (*models.Participant, error) {
	if !validFIO(fio) {
		p.logger.Error("SERVICE: Invalid fio", "fio", fio)
		return nil, service_errors.InvalidFIO
	}

	if !validCategory(category) {
		p.logger.Error("SERVICE: Invalid category", "category", category)
		return nil, service_errors.InvalidCategory
	}

	if !validGender(gender) {
		p.logger.Error("SERVICE: Invalid gender", "gender", gender)
		return nil, service_errors.InvalidGender
	}

	if !validBirthDay(birthDay) {
		p.logger.Error("SERVICE: Invalid birthDay", "birthDay", birthDay)
		return nil, service_errors.InvalidBirthDay
	}

	if !validFIO(coach) {
		p.logger.Error("SERVICE: Invalid coach", "coach", coach)
		return nil, service_errors.InvalidFIO
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

func (p ParticipantService) UpdateParticipantByID(participantID uuid.UUID, fio string, category string, birthDay time.Time, coach string) (*models.Participant, error) {
	participant, err := p.ParticipantRepository.GetParticipantDataByID(participantID)
	participantCopy := participant

	if err != nil {
		p.logger.Error("SERVICE: GetParticipantByID method failed", "id", participantID, "error", err)
		return participant, err
	}

	if !validFIO(fio) {
		p.logger.Error("SERVICE: Invalid fio", "fio", fio)
		return participant, service_errors.InvalidFIO
	}

	if !validCategory(category) {
		p.logger.Error("SERVICE: Invalid category", "category", category)
		return participant, service_errors.InvalidCategory
	}

	if !validBirthDay(birthDay) {
		p.logger.Error("SERVICE: Invalid birthDay", "birthDay", birthDay)
		return participant, service_errors.InvalidBirthDay
	}

	if !validFIO(coach) {
		p.logger.Error("SERVICE: Invalid coach", "coach", coach)
		return participant, service_errors.InvalidFIO
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
