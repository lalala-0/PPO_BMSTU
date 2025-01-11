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
)

type CrewService struct {
	CrewRepository repository_interfaces.ICrewRepository
	logger         *logger.CustomLogger
}

func NewCrewService(CrewRepository repository_interfaces.ICrewRepository, logger *logger.CustomLogger) service_interfaces.ICrewService {
	return &CrewService{
		CrewRepository: CrewRepository,
		logger:         logger,
	}
}
func (c CrewService) AddNewCrew(crewID uuid.UUID, ratingID uuid.UUID, class int, sailNum int) (*models.Crew, error) {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "AddNewCrew") // Спан для добавления нового экипажа
	defer span.End()

	c.logger.Info("SERVICE: Validating data")

	if !validClass(class) || !(validSailNum(sailNum)) {
		c.logger.Error("SERVICE: Invalid input", "class", class, "sail number", sailNum)
		return nil, fmt.Errorf("SERVICE: Invalid input")
	}

	c.logger.Info("SERVICE: Creating new crew %d", crewID)

	var crew = &models.Crew{
		ID:       crewID,
		RatingID: ratingID,
		Class:    class,
		SailNum:  sailNum,
	}

	// Создание экипажа в репозитории с трассировкой
	createdCrew, err := c.CrewRepository.Create(ctx, crew)
	if err != nil {
		c.logger.Error("SERVICE: Create method failed", "error", err)
		return nil, fmt.Errorf("SERVICE: Create method failed")
	}
	c.logger.Info("SERVICE: Successfully created new crew with ", "id", createdCrew.ID)
	return createdCrew, nil
}

func (c CrewService) DeleteCrewByID(id uuid.UUID) error {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DeleteCrewByID") // Спан для удаления экипажа
	defer span.End()

	err := c.CrewRepository.Delete(ctx, id)
	if err != nil {
		c.logger.Error("SERVICE: Delete method failed", "error", err)
		span.SetStatus(codes.Error, "Delete method failed")
		return err
	}

	c.logger.Info("SERVICE: Successfully deleted crew", "id", id)
	span.SetStatus(codes.Ok, "Success")
	return nil
}

func (c CrewService) UpdateCrewByID(crewID uuid.UUID, ratingID uuid.UUID, class int, sailNum int) (*models.Crew, error) {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "UpdateCrewByID") // Спан для обновления экипажа
	defer span.End()

	crew, err := c.CrewRepository.GetCrewDataByID(ctx, crewID)
	crewCopy := crew

	if err != nil {
		c.logger.Error("SERVICE: GetCrewByID method failed", "id", crewID, "error", err)
		span.SetStatus(codes.Error, "GetCrewByID method failed")
		return crew, fmt.Errorf("SERVICE: GetCrewByID method failed")
	}

	c.logger.Info("SERVICE: Validating data")
	if !validClass(class) || !validSailNum(sailNum) {
		c.logger.Error("SERVICE: Invalid input", "class", class, "sail number", sailNum)
		span.SetStatus(codes.Error, "Invalid input")
		return crew, fmt.Errorf("SERVICE: Invalid input")
	}

	crew.Class = class
	crew.RatingID = ratingID
	crew.SailNum = sailNum

	crew, err = c.CrewRepository.Update(ctx, crew)
	if err != nil {
		crew = crewCopy
		c.logger.Error("SERVICE: UpdateCrew method failed", "error", err)
		span.SetStatus(codes.Error, "UpdateCrew method failed")
		return crew, err
	}

	c.logger.Info("SERVICE: Successfully updated crew", "crew", crew)
	span.SetStatus(codes.Ok, "Success")
	return crew, nil
}

func (c CrewService) GetCrewDataByID(id uuid.UUID) (*models.Crew, error) {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetCrewDataByID") // Спан для получения данных об экипаже
	defer span.End()

	crew, err := c.CrewRepository.GetCrewDataByID(ctx, id)
	if err != nil {
		c.logger.Error("SERVICE: GetCrewByID method failed", "id", id, "error", err)
		span.SetStatus(codes.Error, "GetCrewByID method failed")
		return nil, err
	}

	c.logger.Info("SERVICE: Successfully retrieved crew", "id", id)
	span.SetStatus(codes.Ok, "Success")
	return crew, nil
}

func (c CrewService) AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "AttachParticipantToCrew") // Спан для прикрепления участника к экипажу
	defer span.End()

	err := c.CrewRepository.AttachParticipantToCrew(ctx, participantID, crewID, helmsman)

	if err != nil {
		c.logger.Error("SERVICE: AttachParticipantToCrew method failed", "pid", participantID, "cid", crewID, "error", err)
		span.SetStatus(codes.Error, "AttachParticipantToCrew method failed")
		return err
	}

	c.logger.Info("SERVICE: Successfully attached participant to crew", "pid", participantID, "cid", crewID)
	span.SetStatus(codes.Ok, "Success")
	return nil
}

func (c CrewService) DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DetachParticipantFromCrew") // Спан для открепления участника от экипажа
	defer span.End()

	err := c.CrewRepository.DetachParticipantFromCrew(ctx, participantID, crewID)

	if err != nil {
		c.logger.Error("SERVICE: DetachParticipantFromCrew method failed", "pid", participantID, "cid", crewID, "error", err)
		span.SetStatus(codes.Error, "DetachParticipantFromCrew method failed")
		return err
	}

	c.logger.Info("SERVICE: Successfully detached participant from crew", "pid", participantID, "cid", crewID)
	span.SetStatus(codes.Ok, "Success")
	return nil
}

func (c CrewService) ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "ReplaceParticipantStatusInCrew") // Спан для замены статуса участника в экипаже
	defer span.End()

	if !validFlag(helmsman) || !validFlag(active) {
		c.logger.Error("SERVICE: Invalid input", "helmsman", helmsman, "active", active)
		span.SetStatus(codes.Error, "Invalid input")
		return fmt.Errorf("SERVICE: Invalid input")
	}

	err := c.CrewRepository.ReplaceParticipantStatusInCrew(ctx, participantID, crewID, helmsman, active)

	if err != nil {
		c.logger.Error("SERVICE: ReplaceParticipantStatusInCrew method failed", "pid", participantID, "cid", crewID, "error", err)
		span.SetStatus(codes.Error, "ReplaceParticipantStatusInCrew method failed")
		return err
	}

	c.logger.Info("SERVICE: Successfully replaced participant status in crew", "pid", participantID, "cid", crewID)
	span.SetStatus(codes.Ok, "Success")
	return nil
}

func (c CrewService) GetCrewsDataByRatingID(id uuid.UUID) ([]models.Crew, error) {
	// Извлекаем контекст из запроса
	ctx := context.Background() // Или c.Request.Context(), если у вас доступ к gin.Context

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetCrewsDataByRatingID") // Спан для получения данных об экипажах по рейтингу
	defer span.End()

	crews, err := c.CrewRepository.GetCrewsDataByRatingID(ctx, id)
	if err != nil {
		c.logger.Error("SERVICE: GetCrewsDataByRatingID method failed", "error", err)
		span.SetStatus(codes.Error, "GetCrewsDataByRatingID method failed")
		return nil, err
	}

	c.logger.Info("SERVICE: Successfully retrieved crews by rating id")
	span.SetStatus(codes.Ok, "Success")
	return crews, nil
}
