test_race:
	go env -w CGO_ENABLED=1
	go test -count 1000 -race ./internal/adapters/repository

run_dev:
	go run main.go

stop_docker:
	docker-compose down

run_docker: stop_docker
	docker-compose up -d --no-deps --build

rebuild_app:
	docker-compose up -d --no-deps --build app