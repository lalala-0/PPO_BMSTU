package protestViews

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
)

func protestJudgeMenu(services registry.Services, protest *models.Protest) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Список участников протеста",
				Handler: func() error {
					return GetProtestParticipants(services, protest)
				},
			},
			{
				Name: "Удалить протест",
				Handler: func() error {
					return DeleteProtest(services, protest)
				},
			},
			{
				Name: "Изменить протест",
				Handler: func() error {
					return UpdateProtest(services, protest)
				},
			},
			{
				Name: "Завершить рассмотрение протеста",
				Handler: func() error {
					return CompleteProtestReview(services, protest)
				},
			},
		},
	)

	err := m.Menu()
	if err != nil {
		return err
	}

	return nil
}

func protestViewerMenu(services registry.Services, protest *models.Protest) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Список участников протеста",
				Handler: func() error {
					return GetProtestParticipants(services, protest)
				},
			},
		},
	)

	err := m.Menu()
	if err != nil {
		return err
	}

	return nil
}
