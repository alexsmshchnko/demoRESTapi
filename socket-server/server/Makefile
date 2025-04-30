run_dev:
	go run main.go

stop_docker:
	docker-compose down

run_docker: stop_docker
	docker-compose up -d --no-deps --build

rebuild_app:
	docker-compose up -d --no-deps --build app