package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	"fmt"
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

func NewProtestService(ProtestRepository repository_interfaces.IProtestRepository, CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository, CrewRepository repository_interfaces.ICrewRepository, logger *log.Logger) service_interfaces.IProtestService {
	return &ProtestService{
		ProtestRepository:       ProtestRepository,
		CrewResInRaceRepository: CrewResInRaceRepository,
		CrewRepository:          CrewRepository,
		logger:                  logger,
	}
}

func (p ProtestService) AddNewProtest(raceID uuid.UUID, ratingID uuid.UUID, judgeID uuid.UUID, ruleNum int, reviewDate time.Time, comment string, protesteeSailNum int, protestorSailNum int, witnessesSailNum []int) (*models.Protest, error) {
	if !validRuleNum(ruleNum) {
		p.logger.Error("SERVICE: Invalid input data", "ruleNum", ruleNum)
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}

	protestee, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(protesteeSailNum, ratingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "error", err)
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

	err = p.ProtestRepository.AttachCrewToProtest(protestee.ID, protest.ID, models.Protestee)
	if err != nil {
		p.logger.Error("SERVICE: AttachCrewToProtest method failed", "error", err)
		return nil, err
	}

	err = p.ProtestRepository.AttachCrewToProtest(protestor.ID, protest.ID, models.Protestor)
	if err != nil {
		_ = p.ProtestRepository.DetachCrewFromProtest(protestee.ID, protest.ID)
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

		err = p.ProtestRepository.AttachCrewToProtest(witness.ID, protest.ID, models.Witness)
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

func (p ProtestService) UpdateProtestByID(protestID uuid.UUID, raceID uuid.UUID, judgeID uuid.UUID, ruleNum int, reviewDate time.Time, status int, comment string) (*models.Protest, error) {
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

func (p ProtestService) GetProtestParticipantsIDByID(id uuid.UUID) (map[uuid.UUID]int, error) {
	ids, err := p.ProtestRepository.GetProtestParticipantsIDByID(id)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestsDataByRaceID method failed", "error", err)
		return nil, err
	}

	p.logger.Info("SERVICE: Successfully got protest participants id by protest id", "ids", ids)
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
	protest, err := p.ProtestRepository.GetProtestDataByID(protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		return err
	}
	protest.Comment = comment
	protest.Status = models.Reviewed

	ids, err := p.ProtestRepository.GetProtestParticipantsIDByID(protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		return err
	}

	crewResInRace, err := p.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(protest.RaceID, mapkey(ids, models.Protestee))
	if err != nil {
		p.logger.Error("SERVICE: GetCrewResByRaceIDAndCrewID method failed", "id", protest.RaceID, "error", err)
		return err
	}
	crewResInRaceCopy := crewResInRace
	crewResInRace.Points = protesteePoints
	_, err = p.CrewResInRaceRepository.Update(crewResInRace)
	if err != nil {
		p.logger.Error("SERVICE: CompleteReview method failed", "error", err)
		return err
	}

	_, err = p.ProtestRepository.Update(protest)
	if err != nil {
		_, rc := p.CrewResInRaceRepository.Update(crewResInRaceCopy)
		if rc != nil {
			p.logger.Error("SERVICE: CompleteReview method failed", "error", err)
			return err
		}
		p.logger.Error("SERVICE: CompleteReview method failed", "id", protestID, "error", err)
		return err
	}

	p.logger.Info("SERVICE: Successfully reviewed protest id", "id", protestID)
	return nil
}

func (p ProtestService) AttachCrewToProtest(protestID uuid.UUID, sailNum int, role int) error {
	if !validProtestRole(role) {
		p.logger.Error("SERVICE: Invalid protest role", "protest role", role)
		return fmt.Errorf("SERVICE: Invalid protest role")
	}

	protest, err := p.ProtestRepository.GetProtestDataByID(protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		return err
	}

	crew, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, protest.RatingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "id", protestID, "error", err)
		return err
	}

	err = p.ProtestRepository.AttachCrewToProtest(crew.ID, protestID, role)

	if err != nil {
		p.logger.Error("SERVICE: AttachCrewToProtest method failed", "rid", protestID, "jid", crew.ID, "error", err)
		return err
	}

	p.logger.Info("SERVICE: Successfully attach crew jid to protest rid", "jid", crew.ID, "rid", protestID)
	return nil
}

func (p ProtestService) DetachCrewFromProtest(protestID uuid.UUID, sailNum int) error {
	protest, err := p.ProtestRepository.GetProtestDataByID(protestID)
	if err != nil {
		p.logger.Error("SERVICE: GetProtestDataByID method failed", "id", protestID, "error", err)
		return err
	}

	crew, err := p.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, protest.RatingID)
	if err != nil {
		p.logger.Error("SERVICE: GetCrewDataBySailNumAndRatingID method failed", "id", protest.RatingID, "error", err)
		return err
	}
	err = p.ProtestRepository.DetachCrewFromProtest(protestID, crew.ID)

	if err != nil {
		p.logger.Error("SERVICE: DetachCrewFromProtest method failed", "jid", crew.ID, "rid", protestID, "error", err)
		return err
	}

	p.logger.Info("SERVICE: Successfully detached crew jid from protest rid", "jid", crew.ID, "rid", protestID)
	return nil
}
