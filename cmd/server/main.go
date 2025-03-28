package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/kei3dev/todo-app-api-go/internal/handler"
	"github.com/kei3dev/todo-app-api-go/internal/repository"
	"github.com/kei3dev/todo-app-api-go/internal/usecase"
	"github.com/kei3dev/todo-app-api-go/pkg/db"
	"github.com/kei3dev/todo-app-api-go/pkg/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := db.NewDBConfig()
	_, err = db.InitDB(dbConfig)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	if os.Getenv("APP_ENV") == "development" {
		db.MigrateDB()
	}

	jwtConfig := middleware.NewJWTConfig(
		os.Getenv("JWT_SECRET"),
		24*time.Hour,
	)

	userRepo := repository.NewUserRepository()
	todoRepo := repository.NewTodoRepository()

	userUsecase := usecase.NewUserUsecase(userRepo)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	todoHandler := handler.NewTodoHandler(todoUsecase)
	authHandler := handler.NewAuthHandler(userUsecase, jwtConfig)

	r := chi.NewRouter()

	r.Use(middleware.CORS)

	r.Post("/users", userHandler.RegisterUser)
	r.Post("/login", authHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(jwtConfig.ValidateJWT)

		r.Post("/todos", todoHandler.CreateTodo)
		r.Get("/todos/{id}", todoHandler.GetTodoByID)
		r.Get("/todos", todoHandler.GetAllTodos)
		r.Put("/todos/{id}", todoHandler.UpdateTodo)
		r.Delete("/todos/{id}", todoHandler.DeleteTodo)
	})

	port := "8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
