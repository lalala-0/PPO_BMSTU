package orderViews

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
)

func CancelOrder(services registry.Services, order *models.Order) error {
	_, err := services.OrderService.Update(order.ID, models.CancelledOrderStatus, order.Rate, order.WorkerID)
	if err != nil {
		return err
	}

	return nil
}
