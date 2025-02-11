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
	go build -o main ./cmd/server

run:
	go run ./cmd/server/main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...

migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down

migrate-create:
	migrate create -ext sql -dir ./migrations -seq $(name)

redis-clear:
	docker exec -it redis_cache redis-cli FLUSHALL
