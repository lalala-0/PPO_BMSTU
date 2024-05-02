package taskViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/registry"
)

func Create(services registry.Services) error {
	var name = utils.EndlessReadWord(stringConst.NameRequest)
	var price = utils.EndlessReadFloat64(stringConst.PriceRequest)
	var category = utils.EndlessReadInt(stringConst.CategoryRequest)

	_, err := services.TaskService.Create(name, price, category)
	if err != nil {
		println(err.Error())
	}

	println("Услуга успешно создана")
	return nil
}
