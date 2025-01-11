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
)

type RatingService struct {
	RatingRepository repository_interfaces.IRatingRepository
	JudgeRepository  repository_interfaces.IJudgeRepository
	logger           *logger.CustomLogger
}

func NewRatingService(RatingRepository repository_interfaces.IRatingRepository, JudgeRepository repository_interfaces.IJudgeRepository, logger *logger.CustomLogger) service_interfaces.IRatingService {
	return &RatingService{
		RatingRepository: RatingRepository,
		JudgeRepository:  JudgeRepository,
		logger:           logger,
	}
}

func (r RatingService) AddNewRating(ratingID uuid.UUID, name string, class int, blowoutCnt int) (*models.Rating, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "AddNewRating")
	defer span.End()

	if !validClass(class) || !validBlowoutCnt(blowoutCnt) {
		r.logger.Error("SERVICE: Invalid input data", "class", class, "BlowoutCnt", blowoutCnt)
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}
	rating := &models.Rating{
		ID:         ratingID,
		Name:       name,
		Class:      class,
		BlowoutCnt: blowoutCnt,
	}

	rating, err := r.RatingRepository.Create(ctx, rating)
	if err != nil {
		r.logger.Error("SERVICE: CreateNewRating method failed", "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully created new rating", "rating", rating)
	return rating, nil
}

func (r RatingService) DeleteRatingByID(id uuid.UUID) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DeleteRatingByID")
	defer span.End()

	_, err := r.RatingRepository.GetRatingDataByID(ctx, id)
	if err != nil {
		r.logger.Error("SERVICE: GetRatingDataByID method failed", "id", id, "error", err)
		return err
	}

	judges, err := r.JudgeRepository.GetJudgesDataByRatingID(ctx, id)
	if err != nil {
		r.logger.Error("SERVICE: GetJudgesDataByRatingID method failed", "id", id, "error", err)
		return err
	}

	for _, judge := range judges {
		err := r.RatingRepository.DetachJudgeFromRating(ctx, id, judge.ID)
		if err != nil {
			r.logger.Error("SERVICE: DetachJudgeFromRating method failed", "id", id, "error", err)
			return err
		}
	}

	err = r.RatingRepository.Delete(ctx, id)
	if err != nil {
		for _, judge := range judges {
			err := r.RatingRepository.AttachJudgeToRating(ctx, id, judge.ID)
			if err != nil {
				r.logger.Error("SERVICE: AttachJudgeToRating method failed", "id", id, "error", err)
				return err
			}
		}
		r.logger.Error("SERVICE: DeleteRatingByID method failed", "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully deleted rating", "rating", id)
	return nil
}

func (r RatingService) UpdateRatingByID(ratingID uuid.UUID, name string, class int, blowoutCnt int) (*models.Rating, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "UpdateRatingByID")
	defer span.End()

	rating, err := r.RatingRepository.GetRatingDataByID(ctx, ratingID)
	ratingCopy := rating

	if err != nil {
		r.logger.Error("SERVICE: GetRatingByID method failed", "id", ratingID, "error", err)
		return rating, err
	}

	if !validClass(class) || !validBlowoutCnt(blowoutCnt) {
		r.logger.Error("SERVICE: Invalid input data", "class", class, "BlowoutCnt", blowoutCnt)
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}

	rating.Name = name
	rating.Class = class
	rating.BlowoutCnt = blowoutCnt

	rating, err = r.RatingRepository.Update(ctx, rating)
	if err != nil {
		rating = ratingCopy
		r.logger.Error("SERVICE: UpdateRating method failed", "error", err)
		return rating, err
	}

	r.logger.Info("SERVICE: Successfully updated rating", "rating", rating)
	return rating, nil
}

func (r RatingService) GetRatingDataByID(id uuid.UUID) (*models.Rating, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetRatingDataByID")
	defer span.End()

	rating, err := r.RatingRepository.GetRatingDataByID(ctx, id)

	if err != nil {
		r.logger.Error("SERVICE: GetRatingByID method failed", "id", id, "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully got rating with GetRatingByID", "id", id)
	return rating, nil
}

func (r RatingService) AttachJudgeToRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetAllCrewResInRace")
	defer span.End()

	err := r.RatingRepository.AttachJudgeToRating(ctx, ratingID, judgeID)

	if err != nil {
		r.logger.Error("SERVICE: AttachJudgeToRating method failed", "rid", ratingID, "jid", judgeID, "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully attach judgeView jid to rating rid", "jid", judgeID, "rid", ratingID)
	return nil
}

func (r RatingService) DetachJudgeFromRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DetachJudgeFromRating")
	defer span.End()

	err := r.RatingRepository.DetachJudgeFromRating(ctx, ratingID, judgeID)

	if err != nil {
		r.logger.Error("SERVICE: DetachJudgeFromRating method failed", "jid", judgeID, "rid", ratingID, "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully detached judgeView jid from rating rid", "jid", judgeID, "rid", ratingID)
	return nil
}

func (r RatingService) GetAllRatings() ([]models.Rating, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetAllRatings")
	defer span.End()

	rating, err := r.RatingRepository.GetAllRatings(ctx)

	if err != nil {
		r.logger.Error("SERVICE: GetAllRatings method failed", "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully got rating with GetAllRatings")
	return rating, nil
}

func (r RatingService) GetRatingTable(id uuid.UUID) ([]models.RatingTableLine, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetRatingTable")
	defer span.End()

	ratingTable, err := r.RatingRepository.GetRatingTable(ctx, id)

	if err != nil {
		r.logger.Error("SERVICE: GetRatingTable method failed", "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully got rating table with GetRatingTable")
	return ratingTable, nil
}
