.PHONY: help build up down restart logs test clean

## Show help
help:
	@echo "Available commands:"
	@echo "  make build     - Build all docker images"
	@echo "  make up        - Start all services"
	@echo "  make down      - Stop all services"
	@echo "  make restart   - Restart all services"
	@echo "  make logs      - Show logs from all services"
	@echo "  make test      - Run Go tests locally"
	@echo "  make clean     - Remove containers, volumes and images"

build:
	docker compose build

up:
	docker compose up

down:
	docker compose down

restart:
	docker compose down
	docker compose up --build

logs:
	docker compose logs -f

test:
	cd user-service && go test ./...
	cd order-service && go test ./...

clean:
	docker compose down -v
	docker system prune -f
