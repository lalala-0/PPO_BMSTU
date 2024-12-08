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

type RaceService struct {
	RaceRepository          repository_interfaces.IRaceRepository
	CrewRepository          repository_interfaces.ICrewRepository
	CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository
	logger                  *log.Logger
}

func NewRaceService(RaceRepository repository_interfaces.IRaceRepository, CrewRepository repository_interfaces.ICrewRepository, CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository, logger *log.Logger) service_interfaces.IRaceService {
	return &RaceService{
		RaceRepository:          RaceRepository,
		CrewRepository:          CrewRepository,
		CrewResInRaceRepository: CrewResInRaceRepository,
		logger:                  logger,
	}
}

func (r RaceService) AddNewRace(raceID uuid.UUID, ratingID uuid.UUID, number int, date time.Time, class int) (*models.Race, error) {
	if !validNumber(number) || !validClass(class) {
		r.logger.Error("SERVICE: Invalid input data", "number", number, "class", class)
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}

	race := &models.Race{
		ID:       raceID,
		RatingID: ratingID,
		Date:     date,
		Number:   number,
		Class:    class,
	}

	race, err := r.RaceRepository.Create(race)
	if err != nil {
		r.logger.Error("SERVICE: CreateNewRace method failed", "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully created new race", "race", race)
	return race, nil
}

func (r RaceService) DeleteRaceByID(id uuid.UUID) error {
	err := r.RaceRepository.Delete(id)
	if err != nil {
		r.logger.Error("SERVICE: DeleteRaceByID method failed", "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully deleted race", "race", id)
	return nil
}

func (r RaceService) UpdateRaceByID(raceID uuid.UUID, ratingID uuid.UUID, number int, date time.Time, class int) (*models.Race, error) {
	race, err := r.RaceRepository.GetRaceDataByID(raceID)
	raceCopy := race

	if err != nil {
		r.logger.Error("SERVICE: GetRaceByID method failed", "id", raceID, "error", err)
		return race, err
	}

	if !validClass(class) || !validNumber(number) {
		r.logger.Error("SERVICE: Invalid input data", "class", class, "number", number)
		return race, fmt.Errorf("SERVICE: Invalid input data")
	}

	race.RatingID = ratingID
	race.Date = date
	race.Number = number
	race.Class = class

	race, err = r.RaceRepository.Update(race)
	if err != nil {
		race = raceCopy
		r.logger.Error("SERVICE: UpdateRace method failed", "error", err)
		return race, err
	}

	r.logger.Info("SERVICE: Successfully created new race", "race", race)
	return race, nil
}

func (r RaceService) GetRaceDataByID(id uuid.UUID) (*models.Race, error) {
	race, err := r.RaceRepository.GetRaceDataByID(id)

	if err != nil {
		r.logger.Error("SERVICE: GetRaceByID method failed", "id", id, "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully got race with GetRaceByID", "id", id)
	return race, nil
}

func (r RaceService) GetRacesDataByRatingID(ratingID uuid.UUID) ([]models.Race, error) {
	races, err := r.RaceRepository.GetRacesDataByRatingID(ratingID)
	if err != nil {
		r.logger.Error("SERVICE: GetRacesDataByRatingID method failed", "error", err)
		return nil, err
	}

	r.logger.Info("SERVICE: Successfully got races by rating id", "races", races)
	return races, nil
}

func (r RaceService) MakeStartProcedure(raceID uuid.UUID, falseStartYachtList map[int]int) error {
	race, rc := r.RaceRepository.GetRaceDataByID(raceID)
	if rc != nil {
		r.logError("SERVICE: MakeStartProcedure method failed", "id", raceID, rc)
		return rc
	}

	crews, rc := r.CrewRepository.GetCrewsDataByRatingID(race.RatingID)
	if rc != nil {
		r.logError("SERVICE: GetCrewsDataByRatingID method failed", "id", raceID, rc)
		return rc
	}

	rc = r.createRaceSailings(crews, raceID)
	if rc != nil {
		return rc
	}

	rc = r.processFalseStartYachts(falseStartYachtList, race, crews)

	r.logger.Info("SERVICE: Successfully start procedure")
	return rc
}

// createRaceSailings создает записи о старте для всех экипажей.
func (r RaceService) createRaceSailings(crews []models.Crew, raceID uuid.UUID) error {
	for _, crew := range crews {
		crewResInRace := &models.CrewResInRace{
			CrewID: crew.ID,
			RaceID: raceID,
		}
		if _, err := r.CrewResInRaceRepository.Create(crewResInRace); err != nil {
			r.logError("SERVICE: CreateNewRaceSailing method failed", err)
			return err
		}
		r.logger.Info("SERVICE: Successfully created new raceSailing", "crew", crew)
	}
	return nil
}

// processFalseStartYachts обрабатывает яхты с ложным стартом.
func (r RaceService) processFalseStartYachts(falseStartYachtList map[int]int, race *models.Race, crews []models.Crew) error {
	var rc error
	rc = nil
	for sailNum, specCircumstance := range falseStartYachtList {
		if !validSpecCircumstance(specCircumstance) {
			r.logger.Error("SERVICE: invalid input data", "SpecCircumstance", specCircumstance)
			return fmt.Errorf("SERVICE: invalid input data")
		}

		crew, err := r.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, race.RatingID)
		if err != nil {
			r.logError("SERVICE: GetCrewDataBySailNumAndRatingID method failed", err)
			rc = err
		} else {

			res, err := r.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(race.ID, crew.ID)
			if err != nil {
				r.logError("SERVICE: GetCrewResByRaceIDAndCrewID method failed", err)
				return err
			}

			res.SpecCircumstance = specCircumstance
			res.Points = len(crews) + 1
			if err := r.updateCrewRes(res); err != nil {
				return err
			}
		}
	}
	return rc
}

func (r RaceService) MakeFinishProcedure(raceID uuid.UUID, finishersList map[int]int, nonFinishersList map[int]int) error {
	var err error
	err = nil

	race, rc := r.RaceRepository.GetRaceDataByID(raceID)
	if rc != nil {
		r.logError("SERVICE: MakeStartProcedure method failed", raceID, rc)
		return rc
	}

	rc = r.processFinishers(finishersList, race)
	if rc != nil {
		err = rc
	}

	rc = r.processNonFinishers(nonFinishersList, race)
	if rc != nil {
		err = rc
	}

	rc = r.finalizeRaceResults(raceID)
	if rc != nil {
		return rc
	}

	r.logger.Info("SERVICE: Successfully finish procedure")
	return err
}

// processFinishers обрабатывает список финишировавших участников.
func (r RaceService) processFinishers(finishersList map[int]int, race *models.Race) error {
	var rc error
	rc = nil
	for sailNum, points := range finishersList {
		crew, err := r.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, race.RatingID)
		if err != nil {
			r.logError("SERVICE: MakeStartProcedure method failed", race.ID, err)
			rc = err
			continue
		}

		res, err := r.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(race.ID, crew.ID)
		if err != nil {
			r.logError("SERVICE: GetCrewResByRaceIDAndCrewID method failed", err)
			rc = err
			continue
		}

		res.Points = points
		if err := r.updateCrewRes(res); err != nil {
			return err
		}
	}
	return rc
}

// processNonFinishers обрабатывает список не финишировавших участников.
func (r RaceService) processNonFinishers(nonFinishersList map[int]int, race *models.Race) error {
	var rc error
	rc = nil
	for sailNum, specCircumstance := range nonFinishersList {
		if !validSpecCircumstance(specCircumstance) {
			r.logger.Error("SERVICE: invalid input data", "SpecCircumstance", specCircumstance)
			return fmt.Errorf("SERVICE: invalid input data")
		}

		crew, err := r.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, race.RatingID)
		if err != nil {
			r.logError("SERVICE: MakeStartProcedure method failed", race.ID, err)
			rc = err
			continue
		}

		res, err := r.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(race.ID, crew.ID)
		if err != nil {
			r.logError("SERVICE: CreateNewRaceSailing method failed", err)
			rc = err
			continue
		}

		if res.SpecCircumstance == 0 {
			res.SpecCircumstance = specCircumstance
			if err := r.updateCrewRes(res); err != nil {
				return err
			}
		} else {
			r.logger.Info("SERVICE: the field SpecCircumstance was already filled in at the start")
		}
	}
	return rc
}

// finalizeRaceResults завершает обработку всех результатов.
func (r RaceService) finalizeRaceResults(raceID uuid.UUID) error {
	allCrewResInRace, err := r.CrewResInRaceRepository.GetAllCrewResInRace(raceID)
	if err != nil {
		return err
	}

	for _, crewRes := range allCrewResInRace {
		if crewRes.Points == 0 {
			crewRes.SpecCircumstance = models.DNF
			crewRes.Points = len(allCrewResInRace) + 1
		} else if crewRes.SpecCircumstance != 0 {
			crewRes.Points = len(allCrewResInRace) + 1
		}

		if err := r.updateCrewRes(&crewRes); err != nil {
			return err
		}
	}
	return nil
}

// updateCrewRes обновляет информацию о результатах экипажа в гонке.
func (r RaceService) updateCrewRes(res *models.CrewResInRace) error {
	_, err := r.CrewResInRaceRepository.Update(res)
	if err != nil {
		r.logError("SERVICE: UpdateRaceSailing method failed", err)
	}
	return err
}

func (r RaceService) GetAllCrewResInRace(race *models.Race) ([]models.CrewResInRace, error) {
	crews, err := r.CrewRepository.GetCrewsDataByRatingID(race.RatingID)
	if err != nil {
		r.logger.Error("SERVICE: GetAllJudges method failed", "error", err)
		return nil, err
	}

	var allResInRaces []models.CrewResInRace
	for _, crew := range crews {
		resInRace, err := r.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(race.ID, crew.ID)
		if err != nil {
			r.logger.Error("SERVICE: GetAllJudges method failed", "error", err)
			return nil, err
		}
		allResInRaces = append(allResInRaces, *resInRace)
	}

	r.logger.Info("SERVICE: Successfully got All Judges")
	return allResInRaces, nil
}

// logError логирует ошибку с дополнительными параметрами.
func (r RaceService) logError(message string, args ...interface{}) {
	r.logger.Error(message, args...)
}
