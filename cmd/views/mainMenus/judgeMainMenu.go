package mainMenus

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/cmd/views/judgeViews"
	"PPO_BMSTU/cmd/views/participantViews"
	"PPO_BMSTU/cmd/views/ratingViews"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func JudgeLoginMenu(services registry.Services) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "войти",
				Handler: func() error {
					judge, err := judgeViews.Login(services)
					if err == nil {
						if judge.Role == models.MainJudge {
							err = JudgesMainMenu(services, judge)
						} else if judge.Role == models.NotMainJudge {
							err = JudgesMainMenu(services, judge)
						} else {
							err = fmt.Errorf("Wrong judgeView role")
						}
					}
					return err
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

func JudgesMainMenu(services registry.Services, judge *models.Judge) error {
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
				Name: "Создать рейтинг",
				Handler: func() error {
					return ratingViews.CreateRating(services)
				},
			},
			{
				Name: "Просмотреть конкретный рейтинг",
				Handler: func() error {
					return ratingViews.GetRatingJudgeMenu(services, judge)
				},
			},
			{
				Name: "Список участников",
				Handler: func() error {
					return participantViews.GetAllParticipants(services)
				},
			},
			{
				Name: "Добавить участника",
				Handler: func() error {
					return participantViews.CreateParticipant(services)
				},
			},
			{
				Name: "Рассмотреть конкретного участника",
				Handler: func() error {
					return participantViews.GetParticipantMenu(services)
				},
			},
		})
	if judge.Role == models.MainJudge {
		m.AddItem(menu.Item{
			Name: "Список судей",
			Handler: func() error {
				return judgeViews.GetAllJudges(services)
			},
		})
		m.AddItem(menu.Item{
			Name: "Создать новый профиль судьи",
			Handler: func() error {
				return judgeViews.CreateJudge(services)
			},
		})
		m.AddItem(menu.Item{
			Name: "Рассмотреть конкретного судью",
			Handler: func() error {
				return judgeViews.GetJudgeMenu(services)
			},
		})
	}

	// Показать меню
	err := m.Menu()
	if err != nil {
		return err
	}

	return nil
}
