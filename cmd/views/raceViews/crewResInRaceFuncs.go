package raceViews

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
)

// GetAllCrewResInRace
func GetAllCrewResInRace(services registry.Services, rating *models.Rating, race *models.Race) error {
	allCrews, err := services.CrewService.GetCrewsDataByRatingID(rating.ID)
	if err != nil {
		return err
	}

	allCrewResInRace, err := services.RaceService.GetAllCrewResInRace(race)
	if err != nil {
		return err
	}
	return modelTables.AllCrewResInRace(allCrewResInRace, allCrews)
}
