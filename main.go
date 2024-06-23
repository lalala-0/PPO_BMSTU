package main

import (
	"PPO_BMSTU/cmd/views/mainMenus"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"github.com/charmbracelet/log"
)

func main() {
	app := registry.App{}

	err := app.Config.ParseConfig("config.json", "config")
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()

	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}

	if app.Config.Mode == "cmd" {
		//err = initAdmin(app.Services)
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}
		cmdErr := mainMenus.RunMenu(app.Services)
		if cmdErr != nil {
			log.Fatal(cmdErr)
			return
		}
	}
}

//func initAdmin(services *registry.Services) error {
//	admins, err := services.JudgeService.GetWorkersByRole(models.ManagerRole)
//	if err != nil {
//		return err
//	}
//
//	if len(admins) == 0 {
//		defaultAdmin := &models.Worker{
//			Email:       "default@admin.com",
//			Name:        "admin",
//			Surname:     "admin",
//			Role:        models.ManagerRole,
//			PhoneNumber: "+79999999999",
//			Address:     "admin address",
//		}
//		_, err = services.WorkerService.Create(defaultAdmin, "admin123")
//		if err != nil {
//			return err
//		}
//
//		log.Info("Default admin created")
//	}
//
//	return nil
//}
