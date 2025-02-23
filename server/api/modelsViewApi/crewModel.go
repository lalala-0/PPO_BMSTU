package modelsViewApi

import (
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
)

type CrewFormData struct {
	ID       uuid.UUID `json:"id"`
	RatingID uuid.UUID `json:"rating-id"`
	SailNum  int       `form:"name"`
	Class    string    `form:"class"`
}

func FromCrewModelToStringData(crew *models.Crew) (CrewFormData, error) {

	var class, err = modelTables.ClassToString(crew.Class)
	if err != nil {
		return CrewFormData{}, err
	}
	res := CrewFormData{crew.ID, crew.RatingID, crew.SailNum, class}
	return res, nil
}

func FromCrewModelsToStringData(crews []models.Crew) ([]CrewFormData, error) {
	var res []CrewFormData
	for _, crew := range crews {
		var el, err = FromCrewModelToStringData(&crew)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

//////////////

type ProtestCrewFormData struct {
	ID       uuid.UUID `json:"id"`
	RatingID uuid.UUID `json:"ratingID"`
	SailNum  int       `json:"sailNum"`
	Class    string    `json:"class"`
	Role     string    `json:"role"`
}

func FromProtestParticipantModelToStringData(crew *models.Crew, role int) (ProtestCrewFormData, error) {
	roleStr, err := modelTables.ProtestParticipantRoleToString(role)
	if err != nil {
		return ProtestCrewFormData{}, err
	}
	class, err := modelTables.ClassToString(crew.Class)
	if err != nil {
		return ProtestCrewFormData{}, err
	}
	res := ProtestCrewFormData{crew.ID, crew.RatingID, crew.SailNum, class, roleStr}
	return res, nil
}

func FromProtestParticipantModelsToStringData(crews []models.Crew, roles []int) ([]ProtestCrewFormData, error) {
	var res []ProtestCrewFormData
	for i, crew := range crews {
		var el, err = FromProtestParticipantModelToStringData(&crew, roles[i])
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

///////////////////////

type CrewInput struct {
	SailNum int `form:"sailNum" binding:"required"`
}

func FromCrewModelToInputData(crew *models.Crew) (CrewInput, error) {
	res := CrewInput{crew.SailNum}
	return res, nil
}

func FromProtestParticipantModelToInputData(crew *models.Crew, role int) (CrewInput, error) {
	res := CrewInput{crew.SailNum}
	return res, nil
}

type CrewParticipantDetachInput struct {
	ParticipantID string `form:"participantID" binding:"required"`
}

type CrewParticipantAttachInput struct {
	ParticipantID string `form:"participantID" binding:"required"`
	Helmsman      int    `form:"helmsman" binding:""`
}

var HelmsmanMap = map[int]string{
	0: "Не рулевой",
	1: "Рулевой",
}
