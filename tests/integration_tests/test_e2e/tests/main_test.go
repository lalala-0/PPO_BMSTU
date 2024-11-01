package tests

import (
	"PPO_BMSTU/internal/registry"
	"PPO_BMSTU/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
	"os"
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
	os.Setenv("DB_TYPE", "postgres")
	suite.Run(t, new(e2eTestSuite))

	//// Тестирование MongoDB
	//os.Setenv("DB_TYPE", "mongo")
	//suite.Run(t, new(e2eTestSuite))
}

func (suite *e2eTestSuite) SetupSuite() {
	app := registry.App{}

	configFile := "../../../config/config_test.json"
	err := app.Config.ParseConfig(configFile, "config")
	assert.NoError(suite.T(), err)

	err = app.Run()
	assert.NoError(suite.T(), err)

	suite.router, err = server.RunServer(&app)
	assert.NoError(suite.T(), err)
}

// TearDownSuite выполняется один раз после завершения тестов
func (suite *e2eTestSuite) TearDownSuite() {

}
