include .env
export $(shell sed 's/=.*//' .env)

.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down --remove-orphans

.PHONY: stop
stop:
	docker compose stop

.PHONY: build
build:
	docker compose build

.PHONY: restart
restart:
	docker compose down --remove-orphans && docker compose up --build

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: run
run:
	docker compose exec api go run ./cmd/server/main.go

.PHONY: fmt
fmt:
	docker compose exec api go fmt ./...

.PHONY: lint
lint:
	docker compose exec api golangci-lint run

.PHONY: test
test:
	docker compose exec api go test ./...

.PHONY: tidy
tidy:
	docker compose exec api go mod tidy

.PHONY: mysql
mysql:
	docker compose exec db mysql -u$(DB_USER) -p$(DB_PASSWORD)

.PHONY: migrate-up
migrate-up:
	docker compose exec api migrate -path ./migrations -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down:
	docker compose exec api migrate -path ./migrations -database "$(DATABASE_URL)" down

.PHONY: migrate-create
migrate-create:
	docker compose exec api migrate create -ext sql -dir ./migrations -seq $(name)