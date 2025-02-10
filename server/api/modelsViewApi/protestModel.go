package modelsViewApi

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type ProtestFormData struct {
	ID         uuid.UUID `json:"id"`
	JudgeID    uuid.UUID `json:"judge-id"`
	RatingID   uuid.UUID `json:"rating-id"`
	RaceID     uuid.UUID `json:"race-id"`
	RuleNum    int       `json:"rule-num"`
	ReviewDate string    `json:"review-date"`
	Status     string    `json:"status"`
	Comment    string    `json:"comment"`
}

func FromProtestModelToStringData(protest *models.Protest) (ProtestFormData, error) {
	var status, err = modelTables.ProtestStatusToString(protest.Status)
	if err != nil {
		return ProtestFormData{}, err
	}
	res := ProtestFormData{protest.ID, protest.JudgeID, protest.RatingID, protest.RaceID, protest.RuleNum, protest.ReviewDate.Format("2006-01-02 15:04:05"), status, protest.Comment}
	return res, nil
}

func FromProtestModelsToStringData(protests []models.Protest) ([]ProtestFormData, error) {
	var res []ProtestFormData
	for _, protest := range protests {
		var el, err = FromProtestModelToStringData(&protest)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

type ProtestInput struct {
	JudgeID    uuid.UUID `json:"judge-id"`
	RuleNum    int       `json:"ruleNum" binding:"required"`
	ReviewDate string    `json:"reviewDate" binding:"required"`
	Status     int       `json:"status" binding:""`
	Comment    string    `json:"comment" binding:""`
}

func FromProtestModelToInputData(protest *models.Protest) (ProtestInput, error) {
	res := ProtestInput{protest.JudgeID, protest.RuleNum, protest.ReviewDate.String(), protest.Status, protest.Comment}
	return res, nil
}

//type ProtestUpdate struct {
//	RuleNum    int       `form:"ruleNum" binding:"required"`
//	ReviewDate time.Time `form:"reviewDate" binding:"required"`
//	Status     int       `form:"status" binding:"required"`
//	Comment    string    `form:"comment" binding:"required"`
//}

type ProtestCreate struct {
	JudgeID          uuid.UUID `json:"judge-id"`
	RuleNum          int       `form:"ruleNum" binding:"required"`
	ReviewDate       string    `form:"reviewDate" binding:"required"`
	Comment          string    `form:"comment" binding:""`
	ProtesteeSailNum int       `form:"protestee" binding:"required"`
	ProtestorSailNum int       `form:"protestor" binding:"required"`
	WitnessesSailNum []int     `form:"witnesses" binding:"required"`
}

type ProtestParticipantDetachInput struct {
	SailNum int `form:"sailNum" binding:"required"`
}

type ProtestParticipantAttachInput struct {
	SailNum int `form:"sailNum" binding:"required"`
	Role    int `form:"role" binding:"required"`
}

//// Protest participants role vars
//const Protestor = 1
//const Protestee = 2
//const Witness = 3

var RoleMap = map[int]string{
	1: "Протестующий",
	2: "Опротестованный",
	3: "Свидетель",
}

type ProtestComplete struct {
	ResPoints int    `form:"resPoints" binding:"required"`
	Comment   string `form:"comment" binding:""`
}
