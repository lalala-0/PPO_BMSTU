package tests

import (
	"PPO_BMSTU/internal/registry"
	API "PPO_BMSTU/server/api/controllersApi"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"html/template"

	"github.com/stretchr/testify/suite"
	"runtime"
	"testing"
)

// e2eTestSuite описывает тестовый набор для Ie2e
type e2eTestSuite struct {
	suite.Suite
	app    registry.App
	router *gin.Engine
}

// TestIe2e запускает тесты в наборе
func TestIe2e(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Тестирование PostgreSQL
	suite.Run(t, new(e2eTestSuite))

	//// Тестирование MongoDB
	//os.Setenv("DB_TYPE", "mongo")
	//suite.Run(t, new(e2eTestSuite))
}

func (suite *e2eTestSuite) SetupSuite() {
	app := registry.App{}

	configFile := "../../../../../config/config_test.json"
	err := app.Config.ParseConfig(configFile, "config")
	assert.NoError(suite.T(), err)

	err = app.Run()
	assert.NoError(suite.T(), err)

	suite.router = runServer(&app)
}

// TearDownSuite выполняется один раз после завершения тестов
func (suite *e2eTestSuite) TearDownSuite() {

}

func runServer(app *registry.App) *gin.Engine {

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

	// Установка путей для ui и api
	API.SetupRouter(app.Services, router)

	gin.SetMode(gin.DebugMode)

	return router
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

func setupReadmeRote(router *gin.Engine) {

	// Настройка маршрута для статических файлов
	//router.StaticFile("/readme.md", "E:/PPO_BMSTU/README.md")              // Обслуживание файла README.md из корня проекта
	//router.StaticFile("/readme", "E:/PPO_BMSTU/server/static/readme.html") // Обслуживание файла README.md из корня проекта
	//
	//router.Static("/schemes", "E:/PPO_BMSTU/schemes")      // Путь к папке schemes
	//router.Static("/static", "E:/PPO_BMSTU/server/static") // Путь к папке static
}
