package raceViews

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/cmd/views/crewViews"
	"PPO_BMSTU/cmd/views/protestViews"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
)

func RaceJudgeMenu(services registry.Services, race *models.Race, rating *models.Rating, judge *models.Judge) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Провести стартовую процедуру",
				Handler: func() error {
					return MakeStartProcedure(services, race)
				},
			},
			{
				Name: "Провести финишную процедуру",
				Handler: func() error {
					return MakeFinishProcedure(services, race)
				},
			},
			{
				Name: "Список команд",
				Handler: func() error {
					return crewViews.GetAllCrews(services, rating)
				},
			},
			{
				Name: "Список результатов экипажей в гонке",
				Handler: func() error {
					return GetAllCrewResInRace(services, rating, race)
				},
			},
			{
				Name: "Удалить гонку",
				Handler: func() error {
					return DeleteRace(services, race)
				},
			},
			{
				Name: "Изменить гонку",
				Handler: func() error {
					return UpdateRace(services, race)
				},
			},
			{
				Name: "Список протестов в гонке",
				Handler: func() error {
					return protestViews.GetAllProtestsInRace(services, race)
				},
			},
			{
				Name: "Просмотреть конкретный протест",
				Handler: func() error {
					return protestViews.GetProtestJudgeMenu(services, judge, race)
				},
			},
			{
				Name: "Создать протест",
				Handler: func() error {
					return protestViews.CreateProtest(services, race)
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

func RaceViewerMenu(services registry.Services, race *models.Race, rating *models.Rating) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Список результатов экипажей в гонке",
				Handler: func() error {
					return GetAllCrewResInRace(services, rating, race)
				},
			},
			{
				Name: "Список протестов в гонке",
				Handler: func() error {
					return protestViews.GetAllProtestsInRace(services, race)
				},
			},
			{
				Name: "Просмотреть конкретный протест",
				Handler: func() error {
					return protestViews.GetProtestViewerMenu(services, race)
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
