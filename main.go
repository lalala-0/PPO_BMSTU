package main

import (
	"PPO_BMSTU/internal/registry"
	"PPO_BMSTU/server"
	"PPO_BMSTU/trace"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

// Создание метрик для мониторинга
var (
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "go_cpu_usage_seconds_total",
			Help: "CPU usage in seconds",
		},
		[]string{"operation"},
	)
	memoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "go_memory_usage_bytes",
			Help: "Memory usage in bytes",
		},
		[]string{"operation"},
	)
)

func init() {
	// Регистрируем метрики в Prometheus
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
}

func monitorResourceUsage() {
	// Здесь можно настроить мониторинг ресурсов (CPU, память)
	// Для примера просто записываем фиктивные данные
	cpuUsage.WithLabelValues("operation_name").Set(0.2)
	memoryUsage.WithLabelValues("operation_name").Set(1024)
}

func startMetricsServer() {
	// Экспонирование метрик через HTTP
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	app := registry.App{}

	// Чтение конфигурационного файла
	configFile := "config.json"
	if len(os.Args) > 1 { // Если переданы аргументы командной строки
		configFile = os.Args[1] // Использовать файл конфигурации, переданный в качестве аргумента
	}

	err := app.Config.ParseConfig(configFile, "config")
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err) // Более подробная ошибка
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Application failed to run: %v", err) // Более подробная ошибка
	}

	// Запуск сервера для сбора метрик
	go startMetricsServer()

	// Определяем режим работы приложения
	switch app.Config.Mode {
	case "server":
		log.Println("Start with server!")

		// Инициализация трассировки
		_, err := trace.InitTracer("http://localhost:14268/api/traces", "PPO")
		if err != nil {
			log.Fatal("init tracer", err)
		}

		// Запуск сервера с трассировкой
		server.RunServer(&app)

	default:
		log.Printf("Wrong app mode: %s", app.Config.Mode)
	}

	// Мониторинг ресурсов (CPU, память)
	monitorResourceUsage()

	// Формирование отчета (по желанию)
	log.Printf("CPU Usage (operation_name): %f", getMetricValue(cpuUsage, "operation_name"))
	log.Printf("Memory Usage (operation_name): %f", getMetricValue(memoryUsage, "operation_name"))
}

// Функция для получения значения метрики
func getMetricValue(gauge *prometheus.GaugeVec, label string) float64 {
	// Создаем срез метрик для сбора
	var value float64
	metricFamilies, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		log.Printf("Error gathering metrics: %v", err)
		return 0
	}

	// Поиск нужной метрики в собранных данных
	for _, mf := range metricFamilies {
		for _, metric := range mf.GetMetric() {
			// Проверка, что метка совпадает
			if metric.GetLabel()[0].GetValue() == label {
				value = metric.GetGauge().GetValue()
				return value
			}
		}
	}

	return value
}
