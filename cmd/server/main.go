package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/kei3dev/todo-app-api-go/internal/handler"
	"github.com/kei3dev/todo-app-api-go/internal/repository"
	"github.com/kei3dev/todo-app-api-go/internal/usecase"
	"github.com/kei3dev/todo-app-api-go/pkg/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	if os.Getenv("APP_ENV") == "development" {
		db.MigrateDB()
	}

	userRepo := repository.NewUserRepository()
	todoRepo := repository.NewTodoRepository()

	userUsecase := usecase.NewUserUsecase(userRepo)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := chi.NewRouter()

	r.Post("/users", userHandler.RegisterUser)
	r.Get("/users/{id}", userHandler.GetUserByID)

	r.Post("/todos", todoHandler.CreateTodo)
	r.Get("/todos/{id}", todoHandler.GetTodoByID)

	port := "8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
