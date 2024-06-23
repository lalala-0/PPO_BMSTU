package judgeViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
)

func Login(services registry.Services) (*models.Judge, error) {
	var login = utils.EndlessReadWord(stringConst.LoginRequest)
	var password = utils.EndlessReadWord(stringConst.PasswordRequest)

	judge, err := services.JudgeService.Login(login, password)
	if err != nil {
		return nil, err
	}

	return judge, nil
}
