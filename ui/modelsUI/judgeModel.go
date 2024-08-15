package modelsUI

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type JudgeFormData struct {
	ID    uuid.UUID
	FIO   string
	Login string
	Role  string
	Post  string
}

func FromJudgeModelToStringData(judge *models.Judge) (JudgeFormData, error) {

	var role, err = modelTables.JudgeRoleToString(judge.Role)
	if err != nil {
		return JudgeFormData{}, err
	}
	res := JudgeFormData{judge.ID, judge.FIO, judge.Login, role, judge.Post}
	return res, nil
}

func FromJudgeModelsToStringData(judges []models.Judge) ([]JudgeFormData, error) {
	var res []JudgeFormData
	for _, judge := range judges {
		var el, err = FromJudgeModelToStringData(&judge)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

type PasswordInput struct {
	Password string `form:"password" binding:"required"`
}

type JudgeInput struct {
	ID       string `form:"id" binding:""`
	FIO      string `form:"fio" binding:"required"`
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:""`
	Role     int    `form:"role" binding:""`
	Post     string `form:"post" binding:""`
}

func FromJudgeModelToInputData(judge *models.Judge) (JudgeInput, error) {
	res := JudgeInput{judge.ID.String(), judge.FIO, judge.Login, judge.Password, judge.Role, judge.Post}
	return res, nil
}

func FromJudgeModelsToInputData(judges []models.Judge) ([]JudgeInput, error) {
	var res []JudgeInput
	for _, judge := range judges {
		var el, err = FromJudgeModelToInputData(&judge)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

var JudgeRoleMap = map[int]string{
	1: "Главный судья",
	2: "Судья",
}
