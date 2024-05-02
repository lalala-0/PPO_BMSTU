package userViews

import (
	"PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"strings"
)

func requestForChange(fieldName string, fieldValue string, word bool) string {
	fmt.Printf("Изменить %s (оставьте пустым, чтобы не менять): ", fieldName)

	input, err := cmdUtils.StringReader(word)
	strings.TrimSpace(input)

	if err != nil || len(input) == 0 {
		return fieldValue
	}

	return input
}

func Update(services registry.Services, user *models.User) error {
	var email = requestForChange("email", user.Email, true)
	var password = requestForChange("пароль", user.Password, true)
	var name = requestForChange("имя", user.Name, true)
	var surname = requestForChange("фамилию", user.Surname, true)
	var phoneNumber = requestForChange("номер телефона", user.PhoneNumber, true)
	var address = requestForChange("адрес", user.Address, false)

	_, err := services.UserService.Update(user.ID, name, surname, email, address, phoneNumber, password)

	if err != nil {
		return err
	}

	return nil
}
