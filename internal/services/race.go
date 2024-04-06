package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
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

//func NewRaceService(RaceRepository repository_interfaces.IRaceRepository, CrewRepository repository_interfaces.ICrewRepository, CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository, logger *log.Logger) service_interfaces.IRaceService {
//	return &RaceService{
//		RaceRepository:          RaceRepository,
//		CrewRepository:          repository_interfaces.ICrewRepository,
//		CrewResInRaceRepository: repository_interfaces.ICrewResInRaceRepository,
//		logger:                  logger,
//	}
//}

func (r RaceService) AddNewRace(raceID uuid.UUID, ratingID uuid.UUID, number int, date time.Time, class string) (*models.Race, error) {
	if !validNumber(number) {
		r.logger.Error("SERVICE: Invalid number", "number", number)
		return nil, service_errors.InvalidNumber
	}

	if !validClass(class) {
		r.logger.Error("SERVICE: Invalid class", "class", class)
		return nil, service_errors.InvalidClass
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
	_, err := r.RaceRepository.GetRaceDataByID(id)
	if err != nil {
		r.logger.Error("SERVICE: GetRaceDataByID method failed", "id", id, "error", err)
		return err
	}

	err = r.RaceRepository.Delete(id)
	if err != nil {
		r.logger.Error("SERVICE: DeleteRaceByID method failed", "error", err)
		return err
	}

	r.logger.Info("SERVICE: Successfully deleted race", "race", id)
	return nil
}

func (r RaceService) UpdateRaceByID(raceID uuid.UUID, ratingID uuid.UUID, number int, date time.Time, class string) (*models.Race, error) {
	race, err := r.RaceRepository.GetRaceDataByID(raceID)
	raceCopy := race

	if err != nil {
		r.logger.Error("SERVICE: GetRaceByID method failed", "id", raceID, "error", err)
		return race, err
	}

	if !validNumber(number) {
		r.logger.Error("SERVICE: Invalid number", "number", number)
		return race, service_errors.InvalidNumber
	}

	if !validClass(class) {
		r.logger.Error("SERVICE: Invalid class", "class", class)
		return race, service_errors.InvalidClass
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

func (r RaceService) MakeStartProcedure(raceID uuid.UUID, falseStartYachtList map[int]string) error {
	race, rc := r.RaceRepository.GetRaceDataByID(raceID)

	if rc != nil {
		r.logger.Error("SERVICE: MakeStartProcedure method failed", "id", raceID, "error", rc)
		return rc
	}

	crews, rc := r.CrewRepository.GetCrewsDataByRatingID(race.RatingID)

	if rc != nil {
		r.logger.Error("SERVICE: GetCrewsDataByRatingID method failed", "id", raceID, "error", rc)
		return rc
	}
	for _, crew := range crews {
		crewResInRace := &models.CrewResInRace{
			CrewID: crew.ID,
			RaceID: raceID,
		}
		raceSailing, err := r.CrewResInRaceRepository.Create(crewResInRace)
		if err != nil {
			r.logger.Error("SERVICE: CreateNewRaceSailing method failed", "error", err)
			rc = err
		}

		r.logger.Info("SERVICE: Successfully created new raceSailing", "raceSailing", raceSailing)
	}

	for sailNum, specCircumstance := range falseStartYachtList {
		crew, err := r.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, race.RatingID)

		if err != nil {
			r.logger.Error("SERVICE: MakeStartProcedure method failed", "id", raceID, "error", err)
			rc = err
		} else {
			res, err := r.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(raceID, crew.ID)
			if err != nil {
				r.logger.Error("SERVICE: CreateNewRaceSailing method failed", "error", err)
				rc = err
			}
			res.SpecCircumstance = specCircumstance

			_, err = r.CrewResInRaceRepository.Update(res)
			if err != nil {
				r.logger.Error("SERVICE: CreateNewRaceSailing method failed", "error", err)
				rc = err
			}

			if rc == nil {
				r.logger.Info("SERVICE: Successfully created new raceSailing")
			}
		}
	}
	return rc
}

func (r RaceService) MakeFinishProcedure(raceID uuid.UUID, finishersList map[int]int, nonFinishersList map[int]string) error {
	race, rc := r.RaceRepository.GetRaceDataByID(raceID)

	if rc != nil {
		r.logger.Error("SERVICE: MakeStartProcedure method failed", "id", raceID, "error", rc)
		return rc
	}

	for sailNum, points := range finishersList {
		crew, err := r.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, race.RatingID)

		if err != nil {
			r.logger.Error("SERVICE: MakeStartProcedure method failed", "id", raceID, "error", err)
			rc = err
		} else {
			res, err := r.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(raceID, crew.ID)
			if err != nil {
				r.logger.Error("SERVICE: CreateNewRaceSailing method failed", "error", err)
				rc = err
			}
			res.Points = points

			_, err = r.CrewResInRaceRepository.Update(res)
			if err != nil {
				r.logger.Error("SERVICE: CreateNewRaceSailing method failed", "error", err)
				rc = err
			}

			if rc == nil {
				r.logger.Info("SERVICE: Successfully created new raceSailing")
			}
		}
	}

	for sailNum, specCircumstance := range nonFinishersList {
		crew, err := r.CrewRepository.GetCrewDataBySailNumAndRatingID(sailNum, race.RatingID)

		if err != nil {
			r.logger.Error("SERVICE: MakeStartProcedure method failed", "id", raceID, "error", err)
			rc = err
		} else {
			res, err := r.CrewResInRaceRepository.GetCrewResByRaceIDAndCrewID(raceID, crew.ID)
			if err != nil {
				r.logger.Error("SERVICE: CreateNewRaceSailing method failed", "error", err)
				rc = err
			}
			if len(res.SpecCircumstance) == 0 {
				res.SpecCircumstance = specCircumstance

				_, err = r.CrewResInRaceRepository.Update(res)
				if err != nil {
					r.logger.Error("SERVICE: CreateNewRaceSailing method failed", "error", err)
					rc = err
				}

				if rc == nil {
					r.logger.Info("SERVICE: Successfully created new raceSailing")
				}
			} else {
				r.logger.Info("SERVICE: the field SpecCircumstance was already filled in at the start")
			}
		}
	}
	return rc
}
