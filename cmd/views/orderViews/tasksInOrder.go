package orderViews

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func GetTasksInOrder(services registry.Services, order *models.Order) error {
	tasks, err := services.OrderService.GetTasksInOrder(order.ID)
	if err != nil {
		return err
	}

	fmt.Printf("\nУслуги в заказе:\n")
	for i, task := range tasks {
		taskAmount, _ := services.OrderService.GetTaskQuantity(order.ID, task.ID)
		fmt.Printf("%d.\t%s\t%d\n", i+1, task.Name, taskAmount)
	}

	return nil
}
