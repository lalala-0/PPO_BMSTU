# Отчет по оценке производительности

## 1. Введение
### Цель тестирования:

  Оценить производительность операций чтения/записи при одновременной работе множества клиентов.

### Объекты исследования:

  СУБД Mongo и Postgres.

- **Описание тестового окружения:**
    - Версия Docker: 3.8
    - Используемые инструменты: Locust + Prometheus + Grafana

---

## 2. Методология
- **Среда тестирования:**
    - Одинаковые Docker-образы для каждого прогона
- **Ограничения контейнеров:**
  - CPU: 1 
  - RAM: 2G
- **Оцениваемые параметры:**
    - Время выполнения операций
    - Утилизация ресурсов (CPU, RAM, Диск)
    - Задержки и перцентили

---

## 3. Описание тестов
### 3.1 Сценарии тестирования
- **Тест на конкурентные операции записи**

- **Тест на конкурентные операции чтения**
[mongo_benchmark_test.py](mongo_read_test.py)
[psql_benchmark_test.py](psql_read_test.py)
- **Тест на конкурентные операции записи + чтения**

### 3.2 Подготовка тестового окружения
- **Docker-конфигурация:**
  [docker-compose.yml](docker-compose.yml)
- **Настройка Prometheus:**
  [prometheus.yml](prometheus.yml)

---

## 4. Результаты тестов


Графики и запросы для Docker-контейнеров в Grafana:
Параметр	PromQL-запрос
- CPU Usage (%)	
  - sum(rate(container_cpu_usage_seconds_total{name="postgres"}[5m])) * 100
  - sum(rate(container_cpu_usage_seconds_total{name="mongo"}[5m])) * 100 
- RAM Usage (MB)	
  - container_memory_usage_bytes{name="postgres"} / 1024 / 1024
  - container_memory_usage_bytes{name="mongo"} / 1024 / 1024
- Network Inbound (MB)	
  - rate(container_network_receive_bytes_total{name="postgres"}[5m]) / 1024 / 1024
  - rate(container_network_receive_bytes_total{name="mongo"}[5m]) / 1024 / 1024
- Network Outbound (MB)	
  - rate(container_network_transmit_bytes_total{name="postgres"}[5m]) / 1024 / 1024
  - rate(container_network_transmit_bytes_total{name="mongo"}[5m]) / 1024 / 1024


### 4.1. Тест на конкурентные операции записи

#### 4.1.1. Время выполнения операций
| **MongoDB**                               | **PostgreSQL**                       |
|-------------------------------------------|--------------------------------------|
| ![write_mongo.png](img%2Fwrite_mongo.png) | ![write_psql.png](img%2Fwrite_psql.png) |


#### 4.1.2 Утилизация ресурсов
- **CPU Usage (%):**
![write_cpu_usage.png](img%2Fwrite_cpu_usage.png)
- **RAM Usage (MB):**
![write_ram_usage.png](img%2Fwrite_ram_usage.png)
- **Network:**
  - Inbound (MB):
![write_network_receive.png.png](img/write_network_receive.png)
  - Outbound (MB):
![write_network_transmit.png](img%2Fwrite_network_transmit.png)


### 4.2. Тест на конкурентные операции чтения


#### 4.2.1. Время выполнения операций
| **MongoDB**                          | **PostgreSQL**                       |
|--------------------------------------|--------------------------------------|
| ![read_mongo.png](img%2Fread_mongo.png) | ![read_postgres.png](img%2Fread_postgres.png) |


#### 4.2.2. Утилизация ресурсов
- **CPU Usage (%):**
  ![read_cpu_usage.png](img/read_cpu_usage.png)
- **RAM Usage (MB):**
  ![read_ram_usage.png](img%2Fread_ram_usage.png)
- **Network:**
  - Inbound (MB):
  ![read_network_receive.png](img%2Fread_network_receive.png)
  - Outbound (MB):
  ![read_network_transmit.png](img/read_network_transmit.png)

### 4.3.


#### 4.3.1. Время выполнения операций
| **MongoDB**                          | **PostgreSQL**                       |
|--------------------------------------|--------------------------------------|
| ![rw_mongo.png](img%2Frw_mongo.png) | ![rw_psql.png](img%2Frw_psql.png) |


#### 4.3.2. Утилизация ресурсов
- **CPU Usage (%):**
![rw_cpu_usage.png](img%2Frw_cpu_usage.png)
- **RAM Usage (MB):**
![rw_ram_usage.png](img%2Frw_ram_usage.png)
- **Network:**
  - Inbound (MB):
![rw_network_receive.png](img%2Frw_network_receive.png)
  - Outbound (MB):
![rw_network_transmit.png](img%2Frw_network_transmit.png)

---
