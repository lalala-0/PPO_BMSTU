package judgeViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/cmd/views"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"github.com/google/uuid"
)

func CreateJudge(service registry.Services) error {
	name := utils.EndlessReadRow(stringConst.NameRequest)
	login := utils.EndlessReadWord(stringConst.LoginRequest)
	password := utils.EndlessReadWord(stringConst.PasswordRequest)
	role := utils.EndlessReadInt(stringConst.JudgeRoleRequest)
	post := utils.EndlessReadRow(stringConst.JudgePostRequest)

	judge, err := service.JudgeService.CreateProfile(uuid.New(), name, login, password, role, post)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %d %s Профиль судьи успешно создан\n\n\n", judge.FIO, judge.Login, judge.Role, judge.Post)

	return nil
}

func DeleteJudge(service registry.Services, judge *models.Judge) error {
	err := service.JudgeService.DeleteProfile(judge.ID)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %d %s Профиль судьи успешно удалён\n\n\n", judge.FIO, judge.Login, judge.Role, judge.Post)

	return nil
}

func UpdateJudge(service registry.Services, judge *models.Judge) error {
	name := utils.EndlessReadRow(stringConst.NameRequest)
	login := utils.EndlessReadWord(stringConst.LoginRequest)
	password := utils.EndlessReadWord(stringConst.PasswordRequest)
	role := utils.EndlessReadInt(stringConst.JudgeRoleRequest)

	updatedJudge, err := service.JudgeService.UpdateProfile(judge.ID, name, login, password, role)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %d %s Профиль судьи успешно обновлён\n\n\n", updatedJudge.FIO, updatedJudge.Login, updatedJudge.Role, updatedJudge.Post)

	return nil
}

func GetAllJudges(services registry.Services) error {
	judges, err := services.JudgeService.GetAllJudges()
	if err != nil {
		return err
	}
	return modelTables.Judges(judges)
}

func AttachJudgeToRating(services registry.Services, judge *models.Judge) error {
	rating := models.Rating{}
	err := views.GetRating(services, &rating)
	if err != nil {
		return err
	}

	err = services.RatingService.AttachJudgeToRating(rating.ID, judge.ID)
	if err != nil {
		return err
	}
	return nil
}

func DetachJudgeFromRating(services registry.Services, judge *models.Judge) error {
	rating := models.Rating{}
	err := views.GetRating(services, &rating)
	if err != nil {
		return err
	}

	err = services.RatingService.DetachJudgeFromRating(rating.ID, judge.ID)
	if err != nil {
		return err
	}
	return nil
}
