package raceViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/cmd/views"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"github.com/google/uuid"
)

func CreateRace(service registry.Services, rating *models.Rating) error {
	number := utils.EndlessReadInt(stringConst.RaceNumberRequest)
	date := utils.EndlessReadDateTime(stringConst.DateRequest)

	createdRace, err := service.RaceService.AddNewRace(uuid.New(), rating.ID, number, date, rating.Class)
	if err != nil {
		return err
	}

	classString, _ := modelTables.ClassToString(createdRace.Class)
	fmt.Printf("%d %s %s Гонка успешно создана\n\n\n", createdRace.Number, classString, createdRace.Date)

	return nil
}

func UpdateRace(service registry.Services, race *models.Race) error {
	rating := models.Rating{}
	err := views.GetRating(service, &rating)
	if err != nil {
		return err
	}

	number := utils.EndlessReadInt(stringConst.RaceNumberRequest)
	date := utils.EndlessReadDateTime(stringConst.DateRequest)
	class := utils.EndlessReadInt(stringConst.ClassRequest)

	updatedRace, err := service.RaceService.UpdateRaceByID(race.ID, rating.ID, number, date, class)
	if err != nil {
		return err
	}

	classString, _ := modelTables.ClassToString(updatedRace.Class)
	fmt.Printf("%d %s %s Гонка успешно обновлена\n\n\n", updatedRace.Number, classString, updatedRace.Date)

	return nil
}

func DeleteRace(service registry.Services, race *models.Race) error {
	err := service.RaceService.DeleteRaceByID(race.ID)

	if err != nil {
		return err
	}

	class, _ := modelTables.ClassToString(race.Class)
	fmt.Printf("%d %s %s Гонка успешно удалена\n\n\n", race.Number, class, race.Date)

	return nil
}

func MakeStartProcedure(services registry.Services, race *models.Race) error {
	// Ввод информации о фальш стартах
	falseStartYachtList := utils.EndlessReadIntIntMap(stringConst.FinishersWithSpecCircumstanceListRequest)

	err := services.RaceService.MakeStartProcedure(race.ID, falseStartYachtList)
	if err != nil {
		return err
	}
	return nil
}

func MakeFinishProcedure(services registry.Services, race *models.Race) error {
	// Ввод информации о порядке финиша
	finishersList := utils.EndlessReadIntSerialMap(stringConst.FinisherListRequest)
	finishersWithSpecCircumstanceList := utils.EndlessReadIntIntMap(stringConst.FinishersWithSpecCircumstanceListRequest)

	err := services.RaceService.MakeFinishProcedure(race.ID, finishersList, finishersWithSpecCircumstanceList)
	return err
}

func GetRaceJudgeMenu(service registry.Services, rating *models.Rating, judge *models.Judge) error {
	race := models.Race{}
	err := views.GetRaceInRating(service, rating.ID, &race)
	if err != nil {
		fmt.Println(err)
	}
	err = RaceJudgeMenu(service, &race, rating, judge)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GetAllRaces(services registry.Services, rating *models.Rating) error {
	races, err := services.RaceService.GetRacesDataByRatingID(rating.ID)
	if err != nil {
		return err
	}
	return modelTables.Races(races)
}

func GetRaceViewerMenu(service registry.Services, rating *models.Rating) error {
	race := models.Race{}
	err := views.GetRaceInRating(service, rating.ID, &race)
	if err != nil {
		fmt.Println(err)
	}
	err = RaceViewerMenu(service, &race, rating)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
