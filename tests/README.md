# Запуск тестов с последующей генерацией отчета в allure

`gotestsum --format=standard-quiet --junitfile="./tests/allure-results/unit_tests_repo.xml" -- ./tests/unit_tests/test_repository`

# Генерация отчета allure

`allure generate ./tests/allure-results --clean -o ./allure-report`

# Просмотр отчета

`allure serve ./tests/allure-report`

# Тестирование в случайном порядке

`go test -shuffle=on ./...`

# Тестирование в несколько потоков 

Если количество потоков должно быть равно количеству ядер, то

- узнаем, сколько ядер есть: ` [Environment]::ProcessorCount`
- `go test -parallel n ./...`, где n - кол-во ядер

# ЛР3

## Запуск контейнеров бд в докере 
```

docker run -d `
  --name postgres `
-e POSTGRES_USER=postgres `
  -e POSTGRES_PASSWORD=postgres `
-e POSTGRES_DB=postgres `
  -p 5433:5432 `
postgres


docker run -d `
  --name mongodb `
-p 27017:27017 `
mongo

```



## Запуск locust


## Запуск prometheus


## Запуск graphana

# ЛР 4

# ЛР 5

## Установка jaeger для экспорта метрик
docker run -d --name jaeger \
-p 16686:16686 \
-p 14268:14268 \
jaegertracing/all-in-one:1.41


# ЛР 6

Запуск проверок статического анализа локально
1. Проверка цикломатической сложности (gocyclo)

Команда:

`gocyclo -over 10 ./...`

Описание:

Эта команда анализирует все функции в вашем проекте и выводит те, у которых цикломатическая сложность превышает 10 (или другой указанный порог).

2. Проверка на мёртвый и неиспользуемый код (unused)
   
Команда:

`golangci-lint run --enable=unused ./...`

Описание:

Эта команда находит неиспользуемые переменные, функции и пакеты, а также недостижимый код в вашем проекте.

3. Проверка на сложность по Холстеду и уязвимости (gosec)

Команда:

`golangci-lint run --enable=gosec ./...`

Описание:

gosec анализирует код на наличие уязвимостей и потенциально опасных операций.

4. Полный запуск всех линтеров (golangci-lint)

   Команда:

`golangci-lint run ./...`
