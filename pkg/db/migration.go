package db

import (
	"log"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
)

func MigrateDB() {
	err := DB.AutoMigrate(&entity.User{}, &entity.Todo{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}
	log.Println("✅ Database migrated successfully")
}
