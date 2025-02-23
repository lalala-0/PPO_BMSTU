package modelsViewApi

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type RaceFormData struct {
	ID       uuid.UUID `json:"id"`
	RatingID uuid.UUID `json:"rating-id"`
	Date     string    `json:"date"`
	Number   int       `json:"number"`
	Class    string    `json:"class"`
}

func FromRaceModelToStringData(race *models.Race) (RaceFormData, error) {
	var class, err = modelTables.ClassToString(race.Class)
	if err != nil {
		return RaceFormData{}, err
	}
	res := RaceFormData{race.ID, race.RatingID, race.Date.Format("2006-01-02 15:04:05"), race.Number, class}
	return res, nil
}

func FromRaceModelsToStringData(races []models.Race) ([]RaceFormData, error) {
	var res []RaceFormData
	for _, race := range races {
		var el, err = FromRaceModelToStringData(&race)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

type RaceInput struct {
	Date   string `json:"date"   binding:"required"`
	Number int    `json:"number" binding:"required"`
	Class  int    `json:"class"  binding:""`
}

func FromRaceModelToInputData(race *models.Race) (RaceInput, error) {
	res := RaceInput{race.Date.String(), race.Number, race.Class}
	return res, nil
}

var SpecCircumstanceMap = map[int]string{
	0:  "-",   //
	1:  "DNS", // не стартовала (не подпадает под DNC и OCS)
	2:  "DNF", // не финишировала
	3:  "DNC", // не стартовала; не прибыла в район старта
	4:  "OCS", // не стартовала; находилась на стороне дистанции от стартовой линии в момент сигнала "Старт" для нее и не стартовала или нарушила правило 30.1
	5:  "ZFP", // 20% наказание по правилу 30.2
	6:  "UFD", // дисквалификация по правилу 30.3
	7:  "BFD", // дисквалификация по правилу 30.4
	8:  "SCP", // применено "Наказание штрафными очками"
	9:  "RET", // вышла из гонки
	10: "DSQ", // дисквалификация
	11: "DNE", // не исключаемая дисквалификация
	12: "RDG", // исправлен результат
	13: "DPI", // наказание по усмотрению протестового комитета
}

type StartInput struct {
	SpecCircumstance int   `json:"specCircumstance"` // Номер специального обстоятельства в случае фальш-старта
	FalseStartList   []int `json:"falseStartList"`   // Массив номеров фальш-стартовавших команд
}

// FromStartInputViewToStartInput преобразует список falseStartYacht в map[int]int, где
// map[falseStartYacht] = SpecCircumstance
func FromStartInputViewToStartInput(FalseStartList []int, SpecCircumstance int) map[int]int {
	res := make(map[int]int)
	for _, falseStartYacht := range FalseStartList {
		res[falseStartYacht] = SpecCircumstance
	}
	return res
}

type FinishInput struct {
	FinisherList []int `json:"finisherList"` // Массив номеров команд в порядке финиша
}

// FromFinishInputViewToFinishInput преобразует списки (FinisherList []int, AllCrewsList []int в два map[int]int, где
// finishersMap[FinisherListElem] = i - карта финишировавших, i - номер в массиве, начиная с 1
// nonFinishersMap[j] = 2 - карта не финишировавших, j - принадлежит множеству AllCrewsList \ FinisherList
func FromFinishInputViewToFinishInput(FinisherList []int, AllCrewsList []models.Crew) (finishersMap map[int]int, nonFinishersMap map[int]int) {
	// Карты для финишировавших и не финишировавших
	finishersMap = make(map[int]int, len(FinisherList))
	nonFinishersMap = make(map[int]int, len(AllCrewsList)-len(FinisherList))

	// Заполняем карту финишировавших
	for i, finisher := range FinisherList {
		finishersMap[finisher] = i + 1 // Нумерация начинается с 1
	}

	// Заполняем карту не финишировавших
	for _, crew := range AllCrewsList {
		if _, isFinisher := finishersMap[crew.SailNum]; !isFinisher {
			nonFinishersMap[crew.SailNum] = 2
		}
	}

	return finishersMap, nonFinishersMap
}
