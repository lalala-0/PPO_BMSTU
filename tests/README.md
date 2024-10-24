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
