package modelsViewApi

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type RatingFormData struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `form:"name"`
	Class      string    `form:"class"`
	BlowoutCnt int       `form:"blowoutCnt"`
}

func FromRatingModelToStringData(rating *models.Rating) (RatingFormData, error) {
	var class, err = modelTables.ClassToString(rating.Class)
	if err != nil {
		return RatingFormData{}, err
	}
	res := RatingFormData{rating.ID, rating.Name, class, rating.BlowoutCnt}
	return res, nil
}

func FromRatingModelsToStringData(ratings []models.Rating) ([]RatingFormData, error) {
	var res []RatingFormData
	for _, rating := range ratings {
		var el, err = FromRatingModelToStringData(&rating)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

type RatingInput struct {
	Name       string `json:"name" binding:"required"`
	Class      int    `json:"class" binding:""`
	BlowoutCnt int    `json:"blowout_cnt" binding:""`
}

func FromRatingModelToInputData(rating *models.Rating) (RatingInput, error) {
	res := RatingInput{rating.Name, rating.Class, rating.BlowoutCnt}
	return res, nil
}

var ClassMap = map[int]string{
	1:  "Laser",
	2:  "Laser Radial",
	3:  "Optimist",
	4:  "Zoom8",
	5:  "Finn",
	6:  "SB20",
	7:  "J70",
	8:  "Nacra17",
	9:  "C49er",
	10: "RS:X",
	11: "Cadet",
}

type RatingTableLine struct {
	CrewID                uuid.UUID   `json:"crewID"`
	SailNum               int         `json:"SailNum"`
	ParticipantNames      []string    `json:"ParticipantNames"`
	ParticipantBirthDates []string    `json:"ParticipantBirthDates"`
	ParticipantCategories []int       `json:"ParticipantCategories"`
	ResInRace             map[int]int `json:"ResInRace"`
	PointsSum             int         `json:"PointsSum"`
	Rank                  int         `json:"Rank"`
	CoachNames            []string    `json:"CoachNames"`
}

type RaceInfo struct {
	RaceNum int       `json:"RaceNum"`
	RaceID  uuid.UUID `json:"RaceID"`
}

type RankingResponse struct {
	RankingTable []RatingTableLine `json:"RankingTable"`
	Races        []RaceInfo        `json:"Races"`
}

// Функция конвертации
func FromRatingTableLineModelTiStringData(original models.RatingTableLine, crew models.Crew) RatingTableLine {
	birthDates := make([]string, len(original.ParticipantBirthDates))
	for i, date := range original.ParticipantBirthDates {
		birthDates[i] = date.Format("2006-01-02") // Форматирование даты в строку
	}

	return RatingTableLine{
		CrewID:                crew.ID,
		SailNum:               original.SailNum,
		ParticipantNames:      original.ParticipantNames,
		ParticipantBirthDates: birthDates,
		ParticipantCategories: original.ParticipantCategories,
		ResInRace:             original.ResInRace,
		PointsSum:             original.PointsSum,
		Rank:                  original.Rank,
		CoachNames:            original.CoachNames,
	}
}

func FromRatingTableLinesModelTiStringData(original []models.RatingTableLine, crews []models.Crew) []RatingTableLine {
	var result []RatingTableLine
	for _, line := range original {
		var crew models.Crew
		for _, c := range crews {
			if c.SailNum == line.SailNum {
				crew = c
				break
			}
		}
		result = append(result, FromRatingTableLineModelTiStringData(line, crew))
	}
	return result
}
