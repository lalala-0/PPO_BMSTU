package logger

import (
	"PPO_BMSTU/metrics"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/prometheus/client_golang/prometheus"
	"os"
)

type CustomLogger struct {
	Logger         *log.Logger
	CPUUsageMetric prometheus.Gauge
	MemoryUsage    prometheus.Gauge
}

// NewCustomLogger создает новый CustomLogger с метриками
func NewCustomLogger(serviceName string) *CustomLogger {

	// Добавление метрик CPU и памяти
	go metrics.TrackLogResources("logger", "./metrics.log")
	// Создание и возврат логгера
	return &CustomLogger{
		Logger: log.New(os.Stdout),
	}
}

// Log метод для записи логов с метриками
func (cl *CustomLogger) Log(level string, message string, fields ...interface{}) {

	// Форматируем сообщение с полями
	formattedMessage := message
	if len(fields) > 0 {
		for i := 0; i < len(fields); i += 2 {
			if i+1 < len(fields) {
				formattedMessage += fmt.Sprintf(", %v=%v", fields[i], fields[i+1])
			} else {
				formattedMessage += fmt.Sprintf(", %v=<missing>", fields[i])
			}
		}
	}

	// Записываем лог через базовый логгер
	cl.Logger.Printf("[%s] %s", level, formattedMessage)

	// Если уровень "Fatal", вызываем os.Exit
	if level == "FATAL" {
		os.Exit(1)
	}
}

// Info метод для уровня INFO
func (cl *CustomLogger) Info(message string, fields ...interface{}) {
	cl.Log("INFO", message, fields...)
}

// Debug метод для уровня DEBUG
func (cl *CustomLogger) Debug(message string, fields ...interface{}) {
	cl.Log("DEBUG", message, fields...)
}

// Error метод для уровня ERROR
func (cl *CustomLogger) Error(message string, fields ...interface{}) {
	cl.Log("ERROR", message, fields...)
}

// Fatal метод для уровня FATAL
func (cl *CustomLogger) Fatal(message string, fields ...interface{}) {
	cl.Log("FATAL", message, fields...)
}
