package metrics

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
	"os"
	"runtime"
	"time"
)

// createResourceMetrics создаёт метрики для CPU и памяти на основе serviceName и типа ресурса
func createResourceMetrics(serviceName, resourceType string) (cpuUsage, memoryUsage prometheus.Gauge) {
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        resourceType + "_cpu_usage_percent",
		Help:        "CPU usage percentage for " + resourceType,
		ConstLabels: prometheus.Labels{"service": serviceName},
	})

	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        resourceType + "_memory_usage_bytes",
		Help:        "Memory usage in bytes for " + resourceType,
		ConstLabels: prometheus.Labels{"service": serviceName},
	})

	// Регистрация метрик
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)

	return cpuUsage, memoryUsage
}

// trackResourceMetrics обновляет значения метрик CPU и памяти
func trackResourceMetrics(cpuUsage, memoryUsage prometheus.Gauge, updateInterval time.Duration, filePath string) {
	for {
		// Получение информации о CPU и памяти
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		// Установка значений метрик
		cpuPercent := getCPUUsage()
		memoryBytes := float64(memStats.Alloc)
		cpuUsage.Set(cpuPercent)
		memoryUsage.Set(memoryBytes)

		// Экспорт метрик в файл
		exportMetricsToFile(filePath, cpuPercent, memoryBytes)

		// Интервал обновления
		time.Sleep(updateInterval)
	}
}

// getCPUUsage возвращает использование CPU в процентах
func getCPUUsage() float64 {
	// Для реального вычисления можно использовать библиотеку `gopsutil`
	percentages, err := cpu.Percent(100*time.Millisecond, false)
	if err != nil || len(percentages) == 0 {
		return 0.0 // Если произошла ошибка, возвращаем 0
	}
	return percentages[0]
}

// exportMetricsToFile записывает метрики в файл в формате JSON
func exportMetricsToFile(filePath string, cpuPercent, memoryBytes float64) {
	data := map[string]interface{}{
		"timestamp":    time.Now().Format(time.RFC3339),
		"cpu_percent":  cpuPercent,
		"memory_bytes": memoryBytes,
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Логируем ошибку, но не прерываем выполнение
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		// Логируем ошибку, если запись не удалась
		return
	}
}

// TrackLogResources отслеживает ресурсы логгера
func TrackLogResources(serviceName string, filePath string) {
	cpuUsage, memoryUsage := createResourceMetrics(serviceName, "logger")
	go trackResourceMetrics(cpuUsage, memoryUsage, 5*time.Second, filePath)
}

// TrackTraceResources отслеживает ресурсы трассировщика
func TrackTraceResources(serviceName string, filePath string) {
	cpuUsage, memoryUsage := createResourceMetrics(serviceName, "tracer")
	go trackResourceMetrics(cpuUsage, memoryUsage, 5*time.Second, filePath)
}
