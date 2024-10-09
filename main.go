package main

import (
	"PPO_BMSTU/cmd/views/mainMenus"
	_ "PPO_BMSTU/docs"
	"PPO_BMSTU/internal/registry"
	controllersUI "PPO_BMSTU/server"
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
		err = controllersUI.RunServer(&app)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		log.Error("Wrong app mode", "mode", app.Config.Mode)
	}
}

//else if app.Config.Mode == "api" {
//log.Info("Start with api!")
//err = controllersAPI.RunApi(&app)
//if err != nil {
//log.Fatal(err)
//return
//}
//}
