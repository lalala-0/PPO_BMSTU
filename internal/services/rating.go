package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type RatingService struct {
	RatingRepository repository_interfaces.IRatingRepository
	logger           *log.Logger
}

func NewRatingService(RatingRepository repository_interfaces.IRatingRepository, logger *log.Logger) service_interfaces.IRatingService {
	return &RatingService{
		RatingRepository: RatingRepository,
		logger:           logger,
	}
}

func (r RatingService) AddNewRating(ratingID uuid.UUID, class string, blowoutCnt int) (*models.Rating, error) {
	if !validClass(class) {
		r.logger.Error("SERVICE: Invalid class", "class", class)
		return nil, service_errors.InvalidClass
	}

	if !validBlowoutCnt(blowoutCnt) {
		r.logger.Error("SERVICE: Invalid BlowoutCnt", "BlowoutCnt", blowoutCnt)
		return nil, service_errors.InvalidBlowoutCnt
	}

	rating := &models.Rating{
		ID:         ratingID,
		Class:      class,
		BlowoutCnt: blowoutCnt,
	}

	rating, err := r.RatingRepository.Create(rating)
	if err != nil {
		r.logger.Error("SERVICE: CreateNewRating method failed", "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully created new rating", "rating", rating)
	return rating, nil
}

func (r RatingService) DeleteRatingByID(id uuid.UUID) error {
	_, err := r.RatingRepository.GetRatingDataByID(id)
	if err != nil {
		r.logger.Error("SERVICE: GetRatingDataByID method failed", "id", id, "error", err)
		return err
	}

	err = r.RatingRepository.Delete(id)
	if err != nil {
		r.logger.Error("SERVICE: DeleteRatingByID method failed", "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully deleted rating", "rating", id)
	return nil
}

func (r RatingService) UpdateRatingByID(ratingID uuid.UUID, class string, blowoutCnt int) (*models.Rating, error) {
	rating, err := r.RatingRepository.GetRatingDataByID(ratingID)
	ratingCopy := rating

	if err != nil {
		r.logger.Error("SERVICE: GetRatingByID method failed", "id", ratingID, "error", err)
		return rating, err
	}

	if !validClass(class) {
		r.logger.Error("SERVICE: Invalid class", "class", class)
		return rating, service_errors.InvalidClass
	}

	if !validBlowoutCnt(blowoutCnt) {
		r.logger.Error("SERVICE: Invalid BlowoutCnt", "BlowoutCnt", blowoutCnt)
		return rating, service_errors.InvalidBlowoutCnt
	}

	rating.Class = class
	rating.BlowoutCnt = blowoutCnt

	rating, err = r.RatingRepository.Update(rating)
	if err != nil {
		rating = ratingCopy
		r.logger.Error("SERVICE: UpdateRating method failed", "error", err)
		return rating, err
	}

	r.logger.Info("SERVICE: Successfully updated rating", "rating", rating)
	return rating, nil
}

func (r RatingService) GetRatingDataByID(id uuid.UUID) (*models.Rating, error) {
	rating, err := r.RatingRepository.GetRatingDataByID(id)

	if err != nil {
		r.logger.Error("SERVICE: GetRatingByID method failed", "id", id, "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully got rating with GetRatingByID", "id", id)
	return rating, nil
}

func (r RatingService) AttachJudgeToRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	err := r.RatingRepository.AttachJudgeToRating(ratingID, judgeID)

	if err != nil {
		r.logger.Error("SERVICE: AttachJudgeToRating method failed", "rid", ratingID, "jid", judgeID, "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully attach judge jid to rating rid", "jid", judgeID, "rid", ratingID)
	return nil
}

func (r RatingService) DetachJudgeFromRating(ratingID uuid.UUID, judgeID uuid.UUID) error {
	err := r.RatingRepository.DetachJudgeFromRating(ratingID, judgeID)

	if err != nil {
		r.logger.Error("SERVICE: DetachJudgeFromRating method failed", "jid", judgeID, "rid", ratingID, "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully detached judge jid from rating rid", "jid", judgeID, "rid", ratingID)
	return nil
}
