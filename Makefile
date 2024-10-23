test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="./tests" && \
	go test -tags=unit ./tests/unit_tests/test_service/... \
	./tests/unit_tests/test_repository/... --race --parallel 11
	cp environment.properties allure-results

allure:
	cp -R allure-reports/history allure-results
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
 	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
 	export DB_INIT_PATH="${GITHUB_WORKSPACE}/db/sql/init.sql" && \
 	go test -tags=unit ${GITHUB_WORKSPACE}/tests/unit_tests/test_service/ \
	${GITHUB_WORKSPACE}/tests/unit_tests/test_repository/ --race

local-unit:
	export ALLURE_OUTPUT_PATH="./tests" && \
 	export DB_INIT_PATH="./db/sql/init.sql" && \
 	go test -tags=unit ./tests/unit_tests/test_service/... \
	./tests/unit_tests/test_repository/... --race

ci-integration:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
 	export DB_INIT_PATH="${GITHUB_WORKSPACE}/db/sql/init.sql" && \
	go test -tags=integration ${GITHUB_WORKSPACE}/tests/integration_tests/test_repository \
	${GITHUB_WORKSPACE}/tests/integration_tests/test_service/ --race

local-integration:
	export ALLURE_OUTPUT_PATH="./tests" && \
 	export DB_INIT_PATH="./db/sql/init.sql" && \
	go test -tags=integration ./tests/integration/category_test.go --race \
	${GITHUB_WORKSPACE}/tests/integration_tests/test_service/ --race

ci-e2e:
	docker compose up -d
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="e2e-allure" && \
	go test -tags=e2e ${GITHUB_WORKSPACE}/tests/integration/e2e_test.go --race
	docker compose down
	docker image rm testing-backend:latest bitnami/postgresql:16 alpine:latest

local-e2e:
	docker compose up -d
	export ALLURE_OUTPUT_PATH="./tests" && \
	go test -tags=e2e ./tests/integration/e2e_test.go --race

rm-compose:
	docker compose down
	docker image rm testing-backend:latest

ci-concat-reports:
	mkdir allure-results
	cp unit-allure/* allure-results/
	cp integration-allure/* allure-results/
	#cp e2e-allure/* allure-results/
	cp environment.properties allure-results

.PHONY: test allure report ci-unit local-unit ci-integration local-integration ci-e2e local-e2e rm-compose ci-concat-reports
