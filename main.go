package main

import (
	"PPO_BMSTU/cmd/views/mainMenus"
	_ "PPO_BMSTU/docs"
	"PPO_BMSTU/internal/registry"
	"PPO_BMSTU/server"
	"fmt"
	"github.com/charmbracelet/log"
	"os"
)

func main() {
	app := registry.App{}

	// Чтение конфигурационного файла
	configFile := "config_test.json"
	if len(os.Args) > 1 { // Если переданы аргументы командной строки
		configFile = os.Args[1] // Использовать файл конфигурации, переданный в качестве аргумента
	}
	err := app.Config.ParseConfig(configFile, "config")
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}

	// Определяем режим работы приложения
	switch app.Config.Mode {
	case "cmd":
		log.Info("Start with command mode!")
		cmdErr := mainMenus.RunMenu(app.Services)
		if cmdErr != nil {
			log.Fatal(cmdErr)
		}

	case "server":
		log.Info("Start with server!")
		_, err = server.RunServer(&app)
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Error("Wrong app mode", "mode", app.Config.Mode)
	}
}
