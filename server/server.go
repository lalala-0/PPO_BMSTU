package server

import (
	_ "PPO_BMSTU/docs"
	"PPO_BMSTU/internal/registry"
	API "PPO_BMSTU/server/api/controllersApi"
	UI "PPO_BMSTU/server/ui/controllersUi"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
)

func RunServer(app *registry.App) error {

	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	funcMap := template.FuncMap{
		"add":      add,
		"mod":      mod,
		"contains": contains,
		"inArray":  inArray,
	}
	router.SetFuncMap(funcMap)

	UI.SetupRouter(app.Services, router)
	API.SetupRouter(app.Services, router)

	gin.SetMode(gin.DebugMode)

	port := app.Config.Port
	address := app.Config.Address
	err := router.Run(address + port)
	return err
}

func add(a, b int) int {
	return a + b
}

func mod(x, y int) int {
	return x % y
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func inArray(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
