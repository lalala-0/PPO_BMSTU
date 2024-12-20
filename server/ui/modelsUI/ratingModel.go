package modelsUI

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
	Name       string `form:"name" binding:"required"`
	Class      int    `form:"class" binding:"required"`
	BlowoutCnt int    `form:"blowout_cnt" binding:""`
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
