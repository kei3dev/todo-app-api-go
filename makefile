# 環境変数を .env ファイルから読み込む
include .env
export $(shell sed 's/=.*//' .env)

# Docker コンテナ関連
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

# Golang のビルド & 実行
build:
	go build -o main ./cmd/server

run:
	go run ./cmd/server/main.go

fmt:
	go fmt ./...

# Golang の Linter & テスト
lint:
	golangci-lint run

test:
	go test ./...

# マイグレーション（DB 初期化）
migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down

migrate-create:
	migrate create -ext sql -dir ./migrations -seq $(name)

# Redis キャッシュクリア
redis-clear:
	docker exec -it redis_cache redis-cli FLUSHALL
