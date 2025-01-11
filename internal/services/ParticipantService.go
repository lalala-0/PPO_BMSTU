package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_interfaces"
	"PPO_BMSTU/logger"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"time"
)

type ParticipantService struct {
	ParticipantRepository repository_interfaces.IParticipantRepository
	logger                *logger.CustomLogger
}

func NewParticipantService(ParticipantRepository repository_interfaces.IParticipantRepository, logger *logger.CustomLogger) service_interfaces.IParticipantService {
	return &ParticipantService{
		ParticipantRepository: ParticipantRepository,
		logger:                logger,
	}
}

func (p ParticipantService) AddNewParticipant(participantID uuid.UUID, fio string, category int, gender int, birthDay time.Time, coach string) (*models.Participant, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "AddNewParticipant")
	defer span.End()

	if !validFIO(fio) || !validCategory(category) || !validGender(gender) || !validBirthDay(birthDay) || !validFIO(coach) {
		p.logger.Error("SERVICE: Invalid input data", "fio", fio, "category", category, "gender", gender, "birthDate", birthDay, "coach", coach)
		span.SetStatus(codes.Error, "Invalid input data")
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

	participant, err := p.ParticipantRepository.Create(ctx, participant)
	if err != nil {
		p.logger.Error("SERVICE: CreateNewParticipant method failed", "error", err)
		span.SetStatus(codes.Error, "CreateNewParticipant failed")
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully created new participant", "participant", participant)
	span.SetStatus(codes.Ok, "Successfully created participant")
	return participant, nil
}

func (p ParticipantService) DeleteParticipantByID(id uuid.UUID) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DeleteParticipantByID")
	defer span.End()

	_, err := p.ParticipantRepository.GetParticipantDataByID(ctx, id)
	if err != nil {
		p.logger.Error("SERVICE: GetParticipantDataByID method failed", "id", id, "error", err)
		span.SetStatus(codes.Error, "GetParticipantDataByID failed")
		return err
	}

	err = p.ParticipantRepository.Delete(ctx, id)
	if err != nil {
		p.logger.Error("SERVICE: DeleteParticipantByID method failed", "error", err)
		span.SetStatus(codes.Error, "DeleteParticipantByID failed")
		return err
	}

	p.logger.Info("SERVICE: Successfully deleted participant", "participant", id)
	span.SetStatus(codes.Ok, "Successfully deleted participant")
	return nil
}

func (p ParticipantService) UpdateParticipantByID(participantID uuid.UUID, fio string, category int, birthDay time.Time, coach string) (*models.Participant, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "UpdateParticipantByID")
	defer span.End()

	participant, err := p.ParticipantRepository.GetParticipantDataByID(ctx, participantID)
	participantCopy := participant

	if err != nil {
		p.logger.Error("SERVICE: GetParticipantByID method failed", "id", participantID, "error", err)
		span.SetStatus(codes.Error, "GetParticipantByID failed")
		return participant, err
	}

	if !validFIO(fio) || !validCategory(category) || !validBirthDay(birthDay) || !validFIO(coach) {
		p.logger.Error("SERVICE: Invalid input data", "fio", fio, "category", category, "birthDate", birthDay, "coach", coach)
		span.SetStatus(codes.Error, "Invalid input data")
		return participant, fmt.Errorf("SERVICE: Invalid input data")
	}

	participant.FIO = fio
	participant.Category = category
	participant.Birthday = birthDay
	participant.Coach = coach

	participant, err = p.ParticipantRepository.Update(ctx, participant)
	if err != nil {
		participant = participantCopy
		p.logger.Error("SERVICE: UpdateParticipant method failed", "error", err)
		span.SetStatus(codes.Error, "UpdateParticipant failed")
		return participant, err
	}

	p.logger.Info("SERVICE: Successfully updated participant", "participant", participant)
	span.SetStatus(codes.Ok, "Successfully updated participant")
	return participant, nil
}

func (p ParticipantService) GetParticipantDataByID(id uuid.UUID) (*models.Participant, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetParticipantDataByID")
	defer span.End()

	participant, err := p.ParticipantRepository.GetParticipantDataByID(ctx, id)

	if err != nil {
		p.logger.Error("SERVICE: GetParticipantByID method failed", "id", id, "error", err)
		span.SetStatus(codes.Error, "GetParticipantByID failed")
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got participant with GetParticipantByID", "id", id)
	span.SetStatus(codes.Ok, "Successfully retrieved participant")
	return participant, nil
}

func (p ParticipantService) GetParticipantsDataByCrewID(crewID uuid.UUID) ([]models.Participant, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetParticipantsDataByCrewID")
	defer span.End()

	participants, err := p.ParticipantRepository.GetParticipantsDataByCrewID(ctx, crewID)
	if err != nil {
		p.logger.Error("SERVICE: GetParticipantsDataByCrewID method failed", "error", err)
		span.SetStatus(codes.Error, "GetParticipantsDataByCrewID failed")
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got participants by crew id", "participants", participants)
	span.SetStatus(codes.Ok, "Successfully retrieved participants by crew id")
	return participants, nil
}

func (p ParticipantService) GetParticipantsDataByProtestID(protestID uuid.UUID) ([]models.Participant, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetParticipantsDataByProtestID")
	defer span.End()

	participants, err := p.ParticipantRepository.GetParticipantsDataByCrewID(ctx, protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetParticipantsDataByProtestID method failed", "error", err)
		span.SetStatus(codes.Error, "GetParticipantsDataByProtestID failed")
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got participants by protest id", "participants", participants)
	span.SetStatus(codes.Ok, "Successfully retrieved participants by protest id")
	return participants, nil
}

func (p ParticipantService) GetAllParticipants() ([]models.Participant, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetAllParticipants")
	defer span.End()

	participants, err := p.ParticipantRepository.GetAllParticipants(ctx)
	if err != nil {
		p.logger.Error("SERVICE: GetAllParticipants method failed", "error", err)
		span.SetStatus(codes.Error, "GetAllParticipants failed")
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got all participants", "participants", participants)
	span.SetStatus(codes.Ok, "Successfully retrieved all participants")
	return participants, nil
}
