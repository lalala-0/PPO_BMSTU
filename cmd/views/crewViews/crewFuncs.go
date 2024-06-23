package crewViews

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

func GetAllCrews(services registry.Services, rating *models.Rating) error {
	crews, err := services.CrewService.GetCrewsDataByRatingID(rating.ID)
	if err != nil {
		return err
	}
	return modelTables.Crews(services, crews)
}

func GetCrewInRating(service registry.Services, crew *models.Crew, ratingID uuid.UUID) error {

	crews, err := service.CrewService.GetCrewsDataByRatingID(ratingID)

	if err != nil {
		return err
	}

	err = modelTables.Crews(service, crews)
	if err != nil {
		return err
	}

	fmt.Printf("\n-----------\n" +
		"Введите номер команды\n" +
		"Введите 0, чтобы выйти\n\n")

	for {
		var crewNumber int
		_, err = fmt.Scanf("%d", &crewNumber)
		if err != nil {
			return err
		}

		if !utils.ValidateNumber(crewNumber, len(crews)) {
			fmt.Println("Неверный номер команды")
			continue
		}
		*crew = crews[crewNumber-1]
		return nil
	}
}

func DeleteCrew(service registry.Services, crew *models.Crew) error {
	err := service.CrewService.DeleteCrewByID(crew.ID)

	if err != nil {
		return err
	}

	fmt.Printf("%d %d Команда успешно удалена\n\n\n", crew.SailNum, crew.Class)

	return nil
}

func UpdateCrew(service registry.Services, crew *models.Crew) error {
	rating := models.Rating{}
	err := views.GetRating(service, &rating)
	if err != nil {
		return err
	}
	class := utils.EndlessReadInt(stringConst.ClassRequest)
	sailNum := utils.EndlessReadInt(stringConst.SailNumRequest)

	updatedCrew, err := service.CrewService.UpdateCrewByID(crew.ID, rating.ID, class, sailNum)

	if err != nil {
		return err
	}

	fmt.Printf("%d %d Команда успешно обновлена\n\n\n", updatedCrew.SailNum, updatedCrew.Class)

	return nil
}

func CreateCrew(service registry.Services, rating *models.Rating) error {
	class := utils.EndlessReadInt(stringConst.ClassRequest)
	sailNum := utils.EndlessReadInt(stringConst.SailNumRequest)

	createdCrew, err := service.CrewService.AddNewCrew(uuid.New(), rating.ID, class, sailNum)

	if err != nil {
		return err
	}

	fmt.Printf("%d %d Команда успешно создана\n\n\n", createdCrew.SailNum, createdCrew.Class)

	return nil
}
