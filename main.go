package main

import (
	"PPO_BMSTU/cmd/views/mainMenus"
	"PPO_BMSTU/internal/registry"
	"PPO_BMSTU/ui/controllers"
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
		cmdErr := mainMenus.RunMenu(app.Services)
		if cmdErr != nil {
			log.Fatal(cmdErr)
			return
		}
	} else if app.Config.Mode == "server" {
		log.Info("Start with server!")
		err = controllers.RunServer(&app)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		log.Error("Wrong app mode", "mode", app.Config.Mode)
	}
}
