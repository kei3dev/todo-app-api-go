package db

import (
	"log"

	"github.com/kei3dev/todo-app-api-go/internal/domain/model"
)

func MigrateDB() {
	err := DB.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}
	log.Println("✅ Database migrated successfully")
}
