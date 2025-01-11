package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	"PPO_BMSTU/logger"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"time"
)

type ProtestService struct {
	ProtestRepository       repository_interfaces.IProtestRepository
	CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository
	CrewRepository          repository_interfaces.ICrewRepository
	logger                  *logger.CustomLogger
}

func NewProtestService(ProtestRepository repository_interfaces.IProtestRepository, CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository, CrewRepository repository_interfaces.ICrewRepository, logger *logger.CustomLogger) service_interfaces.IProtestService {
	return &ProtestService{
		ProtestRepository:       ProtestRepository,
		CrewResInRaceRepository: CrewResInRaceRepository,
		CrewRepository:          CrewRepository,
		logger:                  logger,
	}
}

func (p ProtestService) AddNewProtest(raceID uuid.UUID, ratingID uuid.UUID, judgeID uuid.UUID, ruleNum int, reviewDate time.Time, comment string, protesteeSailNum int, protestorSailNum int, witnessesSailNum []int) (*models.Protest, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "AddNewProtest")
	defer span.End()

	if !validRuleNum(ruleNum) {
		p.logger.Error("SERVICE: Invalid input data", "ruleNum", ruleNum)
		span.SetStatus(codes.Error, "Invalid ruleNum")
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}

	protestee, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(ctx, protesteeSailNum, ratingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "error", err)
		span.SetStatus(codes.Error, "GetCrewDataBySailNumAndRatingID failed")
		return nil, err
	}

	protest := &models.Protest{
		RaceID:     raceID,
		RatingID:   ratingID,
		JudgeID:    judgeID,
		RuleNum:    ruleNum,
		ReviewDate: reviewDate,
		Status:     models.PendingReview,
		Comment:    comment,
	}

	protest, err = p.ProtestRepository.Create(ctx, protest)
	if err != nil {
		p.logger.Error("SERVICE: CreateNewProtest method failed", "error", err)
		span.SetStatus(codes.Error, "CreateNewProtest failed")
		return nil, err
	}

	protestor, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(ctx, protestorSailNum, ratingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "error", err)
		span.SetStatus(codes.Error, "GetCrewDataBySailNumAndRatingID failed")
		return nil, err
	}

	err = p.ProtestRepository.AttachCrewToProtest(ctx, protestee.ID, protest.ID, models.Protestee)
	if err != nil {
		p.logger.Error("SERVICE: AttachCrewToProtest method failed", "error", err)
		span.SetStatus(codes.Error, "AttachCrewToProtest failed (protestee)")
		return nil, err
	}

	err = p.ProtestRepository.AttachCrewToProtest(ctx, protestor.ID, protest.ID, models.Protestor)
	if err != nil {
		_ = p.ProtestRepository.DetachCrewFromProtest(ctx, protestee.ID, protest.ID)
		p.logger.Error("SERVICE: AttachCrewToProtest method failed", "error", err)
		span.SetStatus(codes.Error, "AttachCrewToProtest failed (protestor)")
		return nil, err
	}

	for _, witnessSailNum := range witnessesSailNum {
		witness, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(ctx, witnessSailNum, ratingID)
		if err != nil {
			p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "error", err)
			span.SetStatus(codes.Error, "GetCrewDataBySailNumAndRatingID failed (witness)")
		}

		err = p.ProtestRepository.AttachCrewToProtest(ctx, witness.ID, protest.ID, models.Witness)
		if err != nil {
			p.logger.Error("SERVICE: AttachCrewToProtest method failed", "error", err)
			span.SetStatus(codes.Error, "AttachCrewToProtest failed (witness)")
		}
	}

	p.logger.Info("SERVICE: Successfully created new protest", "protest", protest)
	span.SetStatus(codes.Ok, "Protest created successfully")
	return protest, nil
}

func (p ProtestService) DeleteProtestByID(id uuid.UUID) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DeleteProtestByID")
	defer span.End()

	_, err := p.ProtestRepository.GetProtestDataByID(ctx, id)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", id, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	err = p.ProtestRepository.Delete(ctx, id)
	if err != nil {
		p.logger.Error("SERVICE: DeleteProtestByID method failed", "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	p.logger.Info("SERVICE: Successfully deleted protest", "protest", id)
	span.SetAttributes(attribute.String("protest_id", id.String())) // Добавляем атрибуты в span
	return nil
}

func (p ProtestService) UpdateProtestByID(protestID uuid.UUID, raceID uuid.UUID, judgeID uuid.UUID, ruleNum int, reviewDate time.Time, status int, comment string) (*models.Protest, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "UpdateProtestByID")
	defer span.End()

	protest, err := p.ProtestRepository.GetProtestDataByID(ctx, protestID)
	protestCopy := protest

	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return protest, err
	}

	if !validRuleNum(ruleNum) {
		p.logger.Error("SERVICE: Invalid ruleNum", "ruleNum", ruleNum)
		span.SetAttributes(attribute.Int("invalid_rule_num", ruleNum)) // Записываем атрибут
		return protest, service_errors.InvalidRuleNum
	}

	if !validStatus(status) {
		p.logger.Error("SERVICE: Invalid status", "status", status)
		span.SetAttributes(attribute.Int("invalid_status", status)) // Записываем атрибут
		return protest, service_errors.InvalidStatus
	}

	protest.RaceID = raceID
	protest.JudgeID = judgeID
	protest.RuleNum = ruleNum
	protest.Status = status
	protest.ReviewDate = reviewDate
	protest.Comment = comment

	protest, err = p.ProtestRepository.Update(ctx, protest)
	if err != nil {
		protest = protestCopy
		p.logger.Error("SERVICE: UpdateProtest method failed", "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return protest, err
	}

	p.logger.Info("SERVICE: Successfully updated protest", "protest", protest)
	span.SetAttributes(attribute.String("protest_id", protest.ID.String())) // Добавляем атрибуты в span
	return protest, nil
}

func (p ProtestService) GetProtestDataByID(id uuid.UUID) (*models.Protest, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetProtestDataByID")
	defer span.End()

	protest, err := p.ProtestRepository.GetProtestDataByID(ctx, id)

	if err != nil {
		p.logger.Error("SERVICE: GetProtestByID method failed", "id", id, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got protest with GetProtestDataByID", "id", id)
	span.SetAttributes(attribute.String("protest_id", id.String())) // Добавляем атрибуты в span
	return protest, nil
}

func (p ProtestService) GetProtestsDataByRaceID(raceID uuid.UUID) ([]models.Protest, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetProtestsDataByRaceID")
	defer span.End()

	protests, err := p.ProtestRepository.GetProtestsDataByRaceID(ctx, raceID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestsDataByRaceID method failed", "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got protests by race id", "protests", protests)
	span.SetAttributes(attribute.String("race_id", raceID.String())) // Добавляем атрибуты в span
	return protests, nil
}

func (p ProtestService) GetProtestParticipantsIDByID(id uuid.UUID) (map[uuid.UUID]int, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetProtestParticipantsIDByID")
	defer span.End()

	ids, err := p.ProtestRepository.GetProtestParticipantsIDByID(ctx, id)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestsDataByRaceID method failed", "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got protest participants id by protest id", "ids", ids)
	span.SetAttributes(attribute.String("protest_id", id.String())) // Добавляем атрибуты в span
	return ids, nil
}

func mapkey(m map[uuid.UUID]int, value int) (key uuid.UUID) {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return uuid.Nil
}

func (p ProtestService) CompleteReview(protestID uuid.UUID, protesteePoints int, comment string) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "CompleteReview")
	defer span.End()

	// Получаем данные о протесте
	protest, err := p.ProtestRepository.GetProtestDataByID(ctx, protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}
	protest.Comment = comment
	protest.Status = models.Reviewed

	// Получаем участников протеста
	ids, err := p.ProtestRepository.GetProtestParticipantsIDByID(ctx, protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	// Получаем данные о команде в гонке
	crewResInRace, err := p.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(ctx, protest.RaceID, mapkey(ids, models.Protestee))
	if err != nil {
		p.logger.Error("SERVICE: GetCrewResByRaceIDAndCrewID method failed", "id", protest.RaceID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}
	crewResInRaceCopy := crewResInRace
	crewResInRace.Points = protesteePoints

	// Обновляем данные о команде
	_, err = p.CrewResInRaceRepository.Update(ctx, crewResInRace)
	if err != nil {
		p.logger.Error("SERVICE: CompleteReview method failed", "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	// Обновляем данные о протесте
	_, err = p.ProtestRepository.Update(ctx, protest)
	if err != nil {
		// Если обновление не удалось, откатываем изменения в CrewResInRace
		_, rc := p.CrewResInRaceRepository.Update(ctx, crewResInRaceCopy)
		if rc != nil {
			p.logger.Error("SERVICE: CompleteReview method failed", "error", err)
			span.RecordError(err) // Записываем ошибку в span
			return err
		}
		p.logger.Error("SERVICE: CompleteReview method failed", "id", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	p.logger.Info("SERVICE: Successfully reviewed protest id", "id", protestID)
	span.SetAttributes(attribute.String("protest_id", protestID.String())) // Добавляем атрибуты в span
	return nil
}

func (p ProtestService) AttachCrewToProtest(protestID uuid.UUID, sailNum int, role int) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "AttachCrewToProtest")
	defer span.End()

	if !validProtestRole(role) {
		p.logger.Error("SERVICE: Invalid protest role", "protest role", role)
		span.RecordError(fmt.Errorf("SERVICE: Invalid protest role")) // Записываем ошибку в span
		return fmt.Errorf("SERVICE: Invalid protest role")
	}

	// Получаем данные о протесте
	protest, err := p.ProtestRepository.GetProtestDataByID(ctx, protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	// Получаем данные о команде
	crew, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(ctx, sailNum, protest.RatingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "id", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	// Прикрепляем команду к протесту
	err = p.ProtestRepository.AttachCrewToProtest(ctx, crew.ID, protestID, role)
	if err != nil {
		p.logger.Error("SERVICE: AttachCrewToProtest method failed", "rid", protestID, "jid", crew.ID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	p.logger.Info("SERVICE: Successfully attach crew jid to protest rid", "jid", crew.ID, "rid", protestID)
	span.SetAttributes(attribute.String("protest_id", protestID.String()), attribute.String("crew_id", crew.ID.String())) // Добавляем атрибуты в span
	return nil
}

func (p ProtestService) DetachCrewFromProtest(protestID uuid.UUID, sailNum int) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DetachCrewFromProtest")
	defer span.End()

	// Получаем данные о протесте
	protest, err := p.ProtestRepository.GetProtestDataByID(ctx, protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	// Получаем данные о команде
	crew, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(ctx, sailNum, protest.RatingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "id", protest.RatingID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	// Отключаем команду от протеста
	err = p.ProtestRepository.DetachCrewFromProtest(ctx, protestID, crew.ID)
	if err != nil {
		p.logger.Error("SERVICE: DetachCrewFromProtest method failed", "jid", crew.ID, "rid", protestID, "error", err)
		span.RecordError(err) // Записываем ошибку в span
		return err
	}

	p.logger.Info("SERVICE: Successfully detached crew jid from protest rid", "jid", crew.ID, "rid", protestID)
	span.SetAttributes(attribute.String("protest_id", protestID.String()), attribute.String("crew_id", crew.ID.String())) // Добавляем атрибуты в span
	return nil
}
