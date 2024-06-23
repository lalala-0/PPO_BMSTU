package participantViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"github.com/google/uuid"
)

func GetAllParticipants(services registry.Services) error {
	participants, err := services.ParticipantService.GetAllParticipants()
	if err != nil {
		return err
	}
	return modelTables.Participants(participants)
}

func DeleteParticipant(service registry.Services, participant *models.Participant) error {
	err := service.ParticipantService.DeleteParticipantByID(participant.ID)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %d %s %d Профиль участника успешно удалён\n\n\n", participant.FIO, participant.Birthday, participant.Category, participant.Coach, participant.Gender)

	return nil
}

func UpdateParticipant(service registry.Services, participant *models.Participant) error {
	fio := utils.EndlessReadRow(stringConst.NameRequest)
	category := utils.EndlessReadInt(stringConst.ParticipantCategoryRequest)
	birthDate := utils.EndlessReadDateTime(stringConst.BirthDateRequest)
	coach := utils.EndlessReadRow(stringConst.CoachNameRequest)

	updatedParticipant, err := service.ParticipantService.UpdateParticipantByID(participant.ID, fio, category, birthDate, coach)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %d %s %d Профиль участника успешно обновлён\n\n\n", updatedParticipant.FIO, updatedParticipant.Birthday, updatedParticipant.Category, updatedParticipant.Coach, updatedParticipant.Gender)

	return nil
}

func CreateParticipant(service registry.Services) error {
	fio := utils.EndlessReadRow(stringConst.NameRequest)
	gender := utils.EndlessReadInt(stringConst.GenderRequest)
	category := utils.EndlessReadInt(stringConst.ParticipantCategoryRequest)
	birthDate := utils.EndlessReadDateTime(stringConst.BirthDateRequest)
	coach := utils.EndlessReadRow(stringConst.CoachNameRequest)

	createdParticipant, err := service.ParticipantService.AddNewParticipant(uuid.New(), fio, category, gender, birthDate, coach)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %d %s %d Профиль участника успешно создан\n\n\n", createdParticipant.FIO, createdParticipant.Birthday, createdParticipant.Category, createdParticipant.Coach, createdParticipant.Gender)

	return nil
}
