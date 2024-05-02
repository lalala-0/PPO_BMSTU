package workerViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
)

func login(services registry.Services) (*models.Worker, error) {
	var email = utils.EndlessReadWord(stringConst.EmailRequest)
	var password = utils.EndlessReadWord(stringConst.PasswordRequest)

	worker, err := services.WorkerService.Login(email, password)
	if err != nil {
		return nil, err
	}

	return worker, nil
}
