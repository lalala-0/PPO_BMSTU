package crewViews

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func crewMenu(services registry.Services, crew *models.Crew) error {
	var m menu.Menu
	m.CreateMenu([]menu.Item{
		{
			Name: "Удалить команду",
			Handler: func() error {
				return DeleteCrew(services, crew)
			},
		},
		{
			Name: "Изменить команду",
			Handler: func() error {
				return UpdateCrew(services, crew)
			},
		},
	})

	// Показать меню
	err := m.Menu()
	if err != nil {
		return err
	}

	return nil
}

func GetCrewMenu(service registry.Services, rating *models.Rating) error {
	crew := models.Crew{}
	err := GetCrewInRating(service, &crew, rating.ID)
	if err != nil {
		fmt.Println(err)
	}
	err = crewMenu(service, &crew)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
