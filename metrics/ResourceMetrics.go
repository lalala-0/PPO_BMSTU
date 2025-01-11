package metrics

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
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
func trackResourceMetrics(cpuUsage, memoryUsage prometheus.Gauge, updateInterval time.Duration) {
	for {
		// Получение информации о CPU и памяти
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		// Установка значений метрик
		cpuUsage.Set(getCPUUsage())
		memoryUsage.Set(float64(memStats.Alloc))

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

func trackLogResources(serviceName string) {
	cpuUsage, memoryUsage := createResourceMetrics(serviceName, "logger")
	go trackResourceMetrics(cpuUsage, memoryUsage, 5*time.Second)
}

func trackTraceResources(serviceName string) {
	cpuUsage, memoryUsage := createResourceMetrics(serviceName, "tracer")
	go trackResourceMetrics(cpuUsage, memoryUsage, 5*time.Second)
}
