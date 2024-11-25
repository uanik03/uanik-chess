package main

import (
	"chess-backend/config"
	"chess-backend/models"
	"fmt"
	"log"
)

func init() {
	dbConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	config.ConnectDB(&dbConfig)
}

func main() {
	err := config.DB.AutoMigrate(&models.User{}, &models.Game{}, &models.Move{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	} else {
		fmt.Println("✅ Migration complete")
	}
}

