package main

import (
	"go-todo/config"
	"go-todo/models"
	"go-todo/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=changeme port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		config.ErrorLogger.Fatalf("Failed to connect to database: %v\n", err)
	}
	config.DB = db
	if err := config.DB.AutoMigrate(&models.Task{}); err != nil {
		config.ErrorLogger.Fatalf("Could not migrate task: %v\n", err)
	}

	router := routers.SetupRouter()
	if err := router.Run("localhost:8080"); err != nil {
		config.ErrorLogger.Fatalf("Couldn't setup router: %v\n", err)
	}
}
