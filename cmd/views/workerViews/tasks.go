package workerViews

import (
	"PPO_BMSTU/cmd/menu"
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/cmd/views/taskViews"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func pickTaskForEditing(services registry.Services, tasks []models.Task) error {
	var err error
	var taskID int

	for {
		err = modelTables.Tasks(tasks)
		if err != nil {
			return err
		}

		fmt.Printf("Выберите услуги для заказа. Введите 0, чтобы вернуться обратно.\n")

		fmt.Scanf("%d", &taskID)
		if taskID == 0 {
			return nil
		}

		if taskID > 0 && taskID <= len(tasks) {
			updatedTask, updErr := taskViews.Update(services, tasks[taskID-1])
			if updErr != nil {
				fmt.Println(updErr.Error())
			} else {
				tasks[taskID-1] = *updatedTask
			}
		} else {
			fmt.Println("Неверный номер услуги")
		}
	}
}

func managerTasks(services registry.Services) error {
	var m menu.Menu
	m.CreateMenu(
		[]menu.Item{
			{
				Name: "Просмотреть все услуги",
				Handler: func() error {
					tasks, err := services.TaskService.GetAllTasks()
					if err != nil {
						fmt.Println(err.Error())
					}
					return pickTaskForEditing(services, tasks)
				},
			},
			{
				Name: "Просмотреть по категории",
				Handler: func() error {
					category := taskViews.ChooseTaskCategory()
					tasks, err := taskViews.TasksByCategory(services, category)
					if err != nil {
						fmt.Println(err.Error())
					}
					return pickTaskForEditing(services, tasks)
				},
			},
			{
				Name: "Создать новую услугу",
				Handler: func() error {
					return taskViews.Create(services)
				},
			},
		},
	)

	err := m.Menu()
	if err != nil {
		return err
	}

	return nil
}
