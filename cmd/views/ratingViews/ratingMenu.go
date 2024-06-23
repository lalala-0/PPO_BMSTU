package ratingViews

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/cmd/views/crewViews"
	"PPO_BMSTU/cmd/views/raceViews"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
)

func ratingJudgeMenu(services registry.Services, rating *models.Rating, judge *models.Judge) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Просмотреть рейтинговую таблицу",
				Handler: func() error {
					return GetRatingTable(services, rating)
				},
			},
			{
				Name: "Удалить рейтинг",
				Handler: func() error {
					return DeleteRating(services, rating)
				},
			},
			{
				Name: "Изменить рейтинг",
				Handler: func() error {
					return UpdateRating(services, rating)
				},
			},
			{
				Name: "Список гонок",
				Handler: func() error {
					return raceViews.GetAllRaces(services, rating)
				},
			},
			{
				Name: "Создать гонку",
				Handler: func() error {
					return raceViews.CreateRace(services, rating)
				},
			},
			{
				Name: "Просмотреть конкретную гонку",
				Handler: func() error {
					return raceViews.GetRaceJudgeMenu(services, rating, judge)
				},
			},
			{
				Name: "Список команд",
				Handler: func() error {
					return crewViews.GetAllCrews(services, rating)
				},
			},
			{
				Name: "Добавить команду",
				Handler: func() error {
					return crewViews.CreateCrew(services, rating)
				},
			},
			{
				Name: "Рассмотреть конкретную команду",
				Handler: func() error {
					return crewViews.GetCrewMenu(services, rating)
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

func ratingViewerMenu(services registry.Services, rating *models.Rating) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Просмотреть рейтинговую таблицу",
				Handler: func() error {
					return GetRatingTable(services, rating)
				},
			},
			{
				Name: "Список гонок",
				Handler: func() error {
					return raceViews.GetAllRaces(services, rating)
				},
			},
			{
				Name: "Просмотреть конкретную гонку",
				Handler: func() error {
					return raceViews.GetRaceViewerMenu(services, rating)
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
