package workerViews

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func Get(service registry.Services, worker *models.Worker) error {
	_, err := service.WorkerService.GetWorkerByID(worker.ID)
	if err != nil {
		return err
	}

	fmt.Print("\nWorker info:\n")
	fmt.Printf("Роль: %s\nEmail: %s\nИмя: %s\nФамилия: %s\nТелефон: %s\nАдрес: %s\n", models.WorkerRole[worker.Role], worker.Email, worker.Name, worker.Surname, worker.PhoneNumber, worker.Address)
	fmt.Print("----------------\n")
	return nil
}
