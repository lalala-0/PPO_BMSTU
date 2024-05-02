package userViews

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func Get(service registry.Services, user *models.User) error {
	_, err := service.UserService.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	fmt.Print("\nUser info:\n")
	fmt.Printf("Email: %s\nИмя: %s\nФамилия: %s\nТелефон: %s\nАдрес: %s\n", user.Email, user.Name, user.Surname, user.PhoneNumber, user.Address)
	fmt.Print("----------------\n")
	return nil
}
