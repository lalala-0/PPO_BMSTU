package server

import (
	"PPO_BMSTU/internal/registry"
	API "PPO_BMSTU/server/api/controllersApi"
	UI "PPO_BMSTU/server/ui/controllersUi"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"html/template"
)

// RunServer запускает сервер с трассировкой
func RunServer(app *registry.App) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(cors.Default()) // Разрешить все домены

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Добавление middleware для трассировки
	router.Use(TraceMiddleware())

	funcMap := template.FuncMap{
		"add":      add,
		"mod":      mod,
		"contains": contains,
		"inArray":  inArray,
	}
	router.SetFuncMap(funcMap)

	// Установка путей для ui и api
	UI.SetupRouter(app.Services, router)
	API.SetupRouter(app.Services, router)

	gin.SetMode(gin.DebugMode)

	port := app.Config.Port
	address := app.Config.Address
	err := router.Run(address + port)
	return router, err
}

// TraceMiddleware - middleware для трассировки
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Создание нового спана для запроса
		tracer := otel.Tracer("server")
		ctx, span := tracer.Start(c.Request.Context(), c.FullPath(),
			trace.WithAttributes(
				attribute.String("method", c.Request.Method),
				attribute.String("url", c.Request.URL.String()),
			),
		)
		defer span.End()

		// Добавление контекста с трассировкой в запрос
		c.Request = c.Request.WithContext(ctx)

		// Продолжение выполнения запроса
		c.Next()
	}
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
