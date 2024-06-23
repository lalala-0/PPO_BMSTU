package mainMenus

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func RunMenu(a *registry.Services) error {
	fmt.Print("Кто вы?\n")
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "зритель",
				Handler: func() error {
					return ViewerMainMenu(*a)
				},
			},
			{
				Name: "судья",
				Handler: func() error {
					return JudgeLoginMenu(*a)
				},
			},
			{
				Name: "главный судья",
				Handler: func() error {
					return JudgeLoginMenu(*a)
				},
			},
		},
	)

	// Показать меню
	err := m.Menu()
	if err != nil {
		return err
	}

	return nil
}
