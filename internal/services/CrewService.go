package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_interfaces"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type CrewService struct {
	CrewRepository repository_interfaces.ICrewRepository
	logger         *log.Logger
}

func NewCrewService(CrewRepository repository_interfaces.ICrewRepository, logger *log.Logger) service_interfaces.ICrewService {
	return &CrewService{
		CrewRepository: CrewRepository,
		logger:         logger,
	}
}

//fmt.Errof("")

func (c CrewService) AddNewCrew(crewID uuid.UUID, ratingID uuid.UUID, class int, sailNum int) (*models.Crew, error) {
	c.logger.Info("SERVICE: Validating data")
	if !validClass(class) || !(validSailNum(sailNum)) {
		c.logger.Error("SERVICE: Invalid input", "class", class, "sail number", sailNum)
		return nil, fmt.Errorf("SERVICE: Invalid input")
	}

	c.logger.Infof("SERVICE: Creating new crew %d", crewID)

	var crew = &models.Crew{
		ID:       crewID,
		RatingID: ratingID,
		Class:    class,
		SailNum:  sailNum,
	}

	createdCrew, err := c.CrewRepository.Create(crew)
	if err != nil {
		c.logger.Error("SERVICE: Create method failed", "error", err)
		return nil, fmt.Errorf("SERVICE: Create method failed")
	}

	c.logger.Info("SERVICE: Successfully created new crew with ", "id", createdCrew.ID)

	return createdCrew, nil
}

func (c CrewService) DeleteCrewByID(id uuid.UUID) error {
	_, err := c.CrewRepository.GetCrewDataByID(id)
	if err != nil {
		c.logger.Error("SERVICE: GetCrewDataByID method failed", "id", id, "error", err)
		return fmt.Errorf("SERVICE: GetCrewDataByID method failed")
	}

	err = c.CrewRepository.Delete(id)
	if err != nil {
		c.logger.Error("SERVICE: Delete method failed", "error", err)
		return fmt.Errorf("SERVICE: Delete method failed")
	}

	c.logger.Info("SERVICE: Successfully deleted crew", "id", id)
	return nil
}

func (c CrewService) UpdateCrewByID(crewID uuid.UUID, ratingID uuid.UUID, class int, sailNum int) (*models.Crew, error) {
	crew, err := c.CrewRepository.GetCrewDataByID(crewID)
	crewCopy := crew

	if err != nil {
		c.logger.Error("SERVICE: GetCrewByID method failed", "id", crewID, "error", err)
		return crew, fmt.Errorf("SERVICE: GetCrewByID method failed")
	}

	c.logger.Info("SERVICE: Validating data")
	if !validClass(class) || !validSailNum(sailNum) {
		c.logger.Error("SERVICE: Invalid input", "class", class, "sail number", sailNum)
		return crew, fmt.Errorf("SERVICE: Invalid input")
	}

	crew.Class = class
	crew.RatingID = ratingID
	crew.SailNum = sailNum

	crew, err = c.CrewRepository.Update(crew)
	if err != nil {
		crew = crewCopy
		c.logger.Error("SERVICE: UpdateCrew method failed", "error", err)
		return crew, err
	}

	c.logger.Info("SERVICE: Successfully updated crew", "crew", crew)
	return crew, nil
}

func (c CrewService) GetCrewDataByID(id uuid.UUID) (*models.Crew, error) {
	crew, err := c.CrewRepository.GetCrewDataByID(id)

	if err != nil {
		c.logger.Error("SERVICE: GetCrewByID method failed", "id", id, "error", err)
		return nil, err
	}

	c.logger.Info("SERVICE: Successfully got crew with GetCrewByID", "id", id)
	return crew, nil
}

func (c CrewService) AttachParticipantToCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int) error {
	err := c.CrewRepository.AttachParticipantToCrew(participantID, crewID, helmsman)

	if err != nil {
		c.logger.Error("SERVICE: AttachParticipantToCrew method failed", "pid", participantID, "cid", crewID, "error", err)
		return err
	}

	c.logger.Info("SERVICE: Successfully attach participant pid to crew cid", "pid", participantID, "cid", crewID)
	return nil
}

func (c CrewService) DetachParticipantFromCrew(participantID uuid.UUID, crewID uuid.UUID) error {
	err := c.CrewRepository.DetachParticipantFromCrew(participantID, crewID)

	if err != nil {
		c.logger.Error("SERVICE: DetachParticipantFromCrew method failed", "pid", participantID, "cid", crewID, "error", err)
		return err
	}

	c.logger.Info("SERVICE: Successfully detached participant pid from crew cid", "pid", participantID, "cid", crewID)
	return nil
}

func (c CrewService) ReplaceParticipantStatusInCrew(participantID uuid.UUID, crewID uuid.UUID, helmsman int, active int) error {
	if !validFlag(helmsman) || !(validFlag(active)) {
		c.logger.Error("SERVICE: Invalid input", "helmsman", helmsman, "active", active)
		return fmt.Errorf("SERVICE: Invalid input")
	}
	err := c.CrewRepository.ReplaceParticipantStatusInCrew(participantID, crewID, helmsman, active)

	if err != nil {
		c.logger.Error("SERVICE: ReplaceParticipantStatusInCrew method failed", "pid", participantID, "cid", crewID, "error", err)
		return err
	}

	c.logger.Info("SERVICE: Successfully replaced participant pid in crew cid", "pid", participantID, "cid", crewID)
	return nil
}
