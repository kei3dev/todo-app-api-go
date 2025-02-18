# Todo App API using Go

This is a simple Todo App API using Go.

## stack

- Go
- MySQL
- Redis
- Chi
- Gorm

## features

- user authentication
- CRUD operations for todos

## project structure

Directory structure based on Clean Architecture principles

```
todo-app-api-go/
├── cmd/
│   ├── server/
│   │   ├── main.go
├── config/
│   ├── config.go
├── internal/
│   ├── entity/
│   ├── usecase/
│   ├── repository/
│   ├── handler/
├── pkg/
│   ├── db/
│   │   ├── mysql.go
│   ├── middleware/
│   │   ├── auth.go
├── migrations/
│   ├── 000001_create_users_table.down.sql
│   ├── 000001_create_users_table.up.sql
│   ├── 000002_create_todos_table.down.sql
│   ├── 000002_create_todos_table.up.sql
├── .env
├── .gitignore
├── .air.toml
├── Dockerfile
├── Makefile
├── compose.yml
├── go.mod
├── go.sum
```

## Get started

### setup

```bash
make up
```

### stop

```bash
make stop
```

### delete

```bash
make down
```
