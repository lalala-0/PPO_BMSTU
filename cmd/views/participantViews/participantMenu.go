package participantViews

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/cmd/views"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func participantMenu(services registry.Services, participant *models.Participant) error {
	var m menu.Menu
	m.CreateMenu([]menu.Item{
		{
			Name: "Удалить профиль участника",
			Handler: func() error {
				return DeleteParticipant(services, participant)
			},
		},
		{
			Name: "Изменить профиль участника",
			Handler: func() error {
				return UpdateParticipant(services, participant)
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

func GetParticipantMenu(service registry.Services) error {
	participant := models.Participant{}
	err := views.GetParticipant(service, &participant)
	if err != nil {
		fmt.Println(err)
	}
	err = participantMenu(service, &participant)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
