include .env
export $(shell sed 's/=.*//' .env)

up:
	docker compose up

down:
	docker compose down --remove-orphans

stop:
	docker compose stop

restart:
	docker compose down --remove-orphans && docker compose up -d

logs:
	docker compose logs -f

build:
	docker compose exec api go build -o main ./cmd/server

run:
	docker compose exec api go run ./cmd/server/main.go

fmt:
	docker compose exec api go fmt ./...

lint:
	docker compose exec api golangci-lint run

test:
	docker compose exec api go test ./...

tidy:
	docker compose exec api go mod tidy

mysql:
	docker compose exec db mysql -uuser -ppassword

migrate-up:
	docker compose exec api migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	docker compose exec api migrate -path ./migrations -database "$(DATABASE_URL)" down

migrate-create:
	docker compose exec api migrate create -ext sql -dir ./migrations -seq $(name)

redis-clear:
	docker compose exec redis redis-cli FLUSHALL

