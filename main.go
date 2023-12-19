package main

import (
	"fmt"
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
		fmt.Printf("Could not open database: %v\n", err)
		return
	}
	config.DB = db
	config.DB.AutoMigrate(&models.Task{})

	router := routers.SetupRouter()
	if err := router.Run(); err != nil {
		fmt.Printf("Couldn't setup router: %v\n", err)
		return
	}
}
