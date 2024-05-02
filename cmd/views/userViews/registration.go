package userViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func registration(services registry.Services) (*models.User, error) {
	var user *models.User
	var err error

	var email = utils.EndlessReadWord(stringConst.EmailRequest)
	var password = utils.EndlessReadWord(stringConst.PasswordRequest)
	var name = utils.EndlessReadWord(stringConst.NameRequest)
	var surname = utils.EndlessReadWord(stringConst.SurnameRequest)
	var phoneNumber = utils.EndlessReadWord(stringConst.PhoneRequest)
	var address = utils.EndlessReadRow(stringConst.AddressRequest)

	user, err = services.UserService.Register(&models.User{
		Email:       email,
		Name:        name,
		Surname:     surname,
		PhoneNumber: phoneNumber,
		Address:     address,
	}, password)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Пользователь %s %s успешно зарегистрирован\n\n\n", user.Name, user.Surname)

	return user, nil
}
