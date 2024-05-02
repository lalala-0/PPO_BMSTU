package taskViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func Update(services registry.Services, task models.Task) (*models.Task, error) {
	var name = utils.EndlessReadRow(stringConst.NameRequest)
	var price = utils.EndlessReadFloat64(stringConst.PriceRequest)
	var category = utils.EndlessReadInt(stringConst.CategoryRequest)

	updatedTask, err := services.TaskService.Update(task.ID, category, name, price)

	fmt.Println("Услуга успешно обновлена")
	return updatedTask, err
}
