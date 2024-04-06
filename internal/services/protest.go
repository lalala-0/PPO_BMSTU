package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"time"
)

type ProtestService struct {
	ProtestRepository       repository_interfaces.IProtestRepository
	CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository
	CrewRepository          repository_interfaces.ICrewRepository
	logger                  *log.Logger
}

//func NewProtestService(ProtestRepository repository_interfaces.IProtestRepository, logger *log.Logger) service_interfaces.IProtestService {
//	return &ProtestService{
//		ProtestRepository: ProtestRepository,
//		logger:            logger,
//	}
//}

func (p ProtestService) AddNewProtest(protestID uuid.UUID, raceID uuid.UUID, ratingID uuid.UUID, ruleNum int, reviewDate time.Time, comment string, protesteeSailNum int, protestorSailNum int, witnessesSailNum []int) (*models.Protest, error) {
	if !validRuleNum(ruleNum) {
		p.logger.Error("SERVICE: Invalid ruleNum", "ruleNum", ruleNum)
		return nil, service_errors.InvalidRuleNum
	}

	protestee, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(protesteeSailNum, ratingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "error", err)
		return nil, err
	}

	protest := &models.Protest{
		ID:          protestID,
		RaceID:      raceID,
		RatingID:    ratingID,
		ProtesteeID: protestee.ID,
		RuleNum:     ruleNum,
		ReviewDate:  reviewDate,
		Status:      models.PendingReview,
		Comment:     comment,
	}

	protest, err = p.ProtestRepository.Create(protest)
	if err != nil {
		p.logger.Error("SERVICE: CreateNewProtest method failed", "error", err)
		return nil, err
	}

	protestor, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(protestorSailNum, ratingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "error", err)
		return nil, err
	}

	err = p.ProtestRepository.AttachCrewToProtest(protestee.ID, protestID, models.Protestee)
	if err != nil {
		p.logger.Error("SERVICE: AttachCrewToProtest method failed", "error", err)
		return nil, err
	}

	err = p.ProtestRepository.AttachCrewToProtest(protestor.ID, protestID, models.Protestor)
	if err != nil {
		err = p.ProtestRepository.DetachCrewFromProtest(protestee.ID, protestID)
		p.logger.Error("SERVICE: AttachCrewToProtest method failed", "error", err)
		return nil, err
	}

	rc := err
	for _, witnessSailNum := range witnessesSailNum {
		witness, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(witnessSailNum, ratingID)
		if err != nil {
			p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "error", err)
			rc = err
		}

		err = p.ProtestRepository.AttachCrewToProtest(witness.ID, protestID, models.Witness)
		if err != nil {
			p.logger.Error("SERVICE: AttachCrewToProtest method failed", "error", err)
			rc = err
		}
	}

	p.logger.Info("SERVICE: Successfully created new protest", "protest", protest)
	return protest, rc
}

func (p ProtestService) DeleteProtestByID(id uuid.UUID) error {
	_, err := p.ProtestRepository.GetProtestDataByID(id)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", id, "error", err)
		return err
	}

	err = p.ProtestRepository.Delete(id)
	if err != nil {
		p.logger.Error("SERVICE: DeleteProtestByID method failed", "error", err)
		return err
	}

	p.logger.Info("SERVICE: Successfully deleted protest", "protest", id)
	return nil
}

func (p ProtestService) UpdateProtestByID(protestID uuid.UUID, raceID uuid.UUID, ruleNum int, reviewDate time.Time, status int, comment string) (*models.Protest, error) {
	protest, err := p.ProtestRepository.GetProtestDataByID(protestID)
	protestCopy := protest

	if err != nil {
		p.logger.Error("SERVICE: GetParticipantByID method failed", "id", protestID, "error", err)
		return protest, err
	}

	if !validRuleNum(ruleNum) {
		p.logger.Error("SERVICE: Invalid ruleNum", "ruleNum", ruleNum)
		return protest, service_errors.InvalidRuleNum
	}

	if !validStatus(status) {
		p.logger.Error("SERVICE: Invalid status", "status", status)
		return protest, service_errors.InvalidStatus
	}

	protest.RaceID = raceID
	protest.RuleNum = ruleNum
	protest.Status = status
	protest.ReviewDate = reviewDate
	protest.Comment = comment

	protest, err = p.ProtestRepository.Update(protest)
	if err != nil {
		protest = protestCopy
		p.logger.Error("SERVICE: UpdateProtest method failed", "error", err)
		return protest, err
	}

	p.logger.Info("SERVICE: Successfully created new protest", "protest", protest)
	return protest, nil
}

func (p ProtestService) GetProtestDataByID(id uuid.UUID) (*models.Protest, error) {
	protest, err := p.ProtestRepository.GetProtestDataByID(id)

	if err != nil {
		p.logger.Error("SERVICE: GetProtestByID method failed", "id", id, "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got protest with GetProtestDataByID", "id", id)
	return protest, nil
}

func (p ProtestService) GetProtestsDataByRaceID(raceID uuid.UUID) ([]models.Protest, error) {
	protests, err := p.ProtestRepository.GetProtestsDataByRaceID(raceID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestsDataByRaceID method failed", "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got protests by protest id", "protests", protests)
	return protests, nil
}

func (p ProtestService) CompleteReview(protestID uuid.UUID, protesteePoints int, comment string) error {
	protest, err := p.ProtestRepository.GetProtestDataByID(protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		return err
	}
	protest.Comment = comment
	protest.Status = models.Reviewed

	protest, err = p.ProtestRepository.Update(protest)
	if err != nil {
		p.logger.Error("SERVICE: CompleteReview method failed", "id", protestID, "error", err)
		return err
	}

	crewResInRace, err := p.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(protest.RaceID, protest.ProtesteeID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewResByRaceIDAndCrewID method failed", "id", protest.RaceID, "error", err)
		return err
	}
	crewResInRace.Points = protesteePoints
	crewResInRace, err = p.CrewResInRaceRepository.Update(crewResInRace)
	if err != nil {
		p.logger.Error("SERVICE: CompleteReview method failed", "error", err)
		return err
	}

	p.logger.Info("SERVICE: Successfully reviewed protest id", "id", protestID)
	return nil
}
