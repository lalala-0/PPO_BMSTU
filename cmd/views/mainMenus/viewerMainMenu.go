package mainMenus

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/cmd/views/ratingViews"
	"PPO_BMSTU/internal/registry"
)

func ViewerMainMenu(services registry.Services) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Список рейтингов",
				Handler: func() error {
					return ratingViews.GetAllRatings(services)
				},
			},
			{
				Name: "Просмотреть конкретный рейтинг",
				Handler: func() error {
					return ratingViews.GetRatingViewerMenu(services)
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
