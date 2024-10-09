package modelsViewApi

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type ParticipantFormData struct {
	ID       uuid.UUID `json:"id"`
	FIO      string    `form:"fio"`
	Category string    `form:"category"`
	Gender   string    `form:"gender"`
	Birthday string    `form:"birthday"`
	Coach    string    `form:"coach"`
}

func FromParticipantModelToStringData(participant *models.Participant) (ParticipantFormData, error) {
	category, err := modelTables.ParticipantCategoryToString(participant.Category)
	if err != nil {
		return ParticipantFormData{}, err
	}
	gender, err := modelTables.GenderToString(participant.Gender)
	if err != nil {
		return ParticipantFormData{}, err
	}
	res := ParticipantFormData{participant.ID, participant.FIO, category, gender, participant.Birthday.Format("2006-01-02"), participant.Coach}
	return res, nil
}

func FromParticipantModelsToStringData(participants []models.Participant) ([]ParticipantFormData, error) {
	var res []ParticipantFormData
	for _, participant := range participants {
		var el, err = FromParticipantModelToStringData(&participant)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

type ParticipantInput struct {
	ID       string `form:"id" binding:""`
	FIO      string `form:"fio" binding:"required"`
	Category int    `form:"category" binding:"required"`
	Gender   int    `form:"gender" binding:""`
	Birthday string `form:"birthday" binding:"required"`
	Coach    string `form:"coach" binding:"required"`
}

func FromParticipantModelToInputData(participant *models.Participant) (ParticipantInput, error) {
	res := ParticipantInput{participant.ID.String(), participant.FIO, participant.Category, participant.Gender, participant.Birthday.Format("2006-01-02T15:04"), participant.Coach}
	return res, nil
}

func FromParticipantModelsToInputData(participants []models.Participant) ([]ParticipantInput, error) {
	var res []ParticipantInput
	for _, participant := range participants {
		var el, err = FromParticipantModelToInputData(&participant)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

var CategoryMap = map[int]string{
	1: "Мастер спорта России междунородного класса",
	2: "Мастер спорта России",
	3: "Кандидат в мастера спорта",
	4: "1-ый спортивный разряд",
	5: "2-ой спортивный разряд",
	6: "3-ий спортивный разряд",
	7: "1-ый юношеский разряд",
	8: "2-ой юношеский разряд",
}

var GenderMap = map[int]string{
	1: "Муж.",
	2: "Жен.",
}
