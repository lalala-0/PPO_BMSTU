package judgeViews

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/cmd/views"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func judgeMenu(services registry.Services, judge *models.Judge) error {
	var m menu.Menu
	m.CreateMenu([]menu.Item{
		{
			Name: "Удалить профиль судьи",
			Handler: func() error {
				return DeleteJudge(services, judge)
			},
		},
		{
			Name: "Изменить профиль судьи",
			Handler: func() error {
				return UpdateJudge(services, judge)
			},
		},
		{
			Name: "Добавить судью к рейтингу",
			Handler: func() error {
				return AttachJudgeToRating(services, judge)
			},
		},
		{
			Name: "Удалить судью из рейтинга",
			Handler: func() error {
				return DetachJudgeFromRating(services, judge)
			},
		},
		{
			Name: "Добавить судью к рейтингу",
			Handler: func() error {
				return AttachJudgeToRating(services, judge)
			},
		},
		{
			Name: "Удалить судью из рейтинга",
			Handler: func() error {
				return DetachJudgeFromRating(services, judge)
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

func GetJudgeMenu(service registry.Services) error {
	judge := models.Judge{}
	err := views.GetJudge(service, &judge)
	if err != nil {
		fmt.Println(err)
	}
	err = judgeMenu(service, &judge)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
